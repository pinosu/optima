package keeper

/*
#cgo LDFLAGS:  -L/usr/local/lib -ldmnengine
#include <stdlib.h>

size_t load_models(const char* c_dir);
const char* evaluate_invocable(const char* c_invocable_name, const char* c_input_data, size_t* output_data_len);
void free_memory(void* ptr, size_t len);
*/
import "C"
import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"unsafe"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"optima/x/optima/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}

	Job struct {
		Id     uint64
		Result string
	}
)

var keyPrefixJob = []byte("job:")

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) LoadModels(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("invalid path for models")
	}
	cDir := C.CString(path)
	defer C.free(unsafe.Pointer(cDir))
	if C.load_models(cDir) == 0 {
		return fmt.Errorf("error loading models")
	}
	return nil
}

func (k Keeper) Evaluate(ctx sdk.Context, job_id uint64, invocableName string, inputData string) error {
	res := k.evaluate(invocableName, inputData)
	job := Job{
		Id:     job_id,
		Result: res,
	}
	return k.setJob(ctx, job)
}

func (k Keeper) setJob(ctx sdk.Context, job Job) error {
	key := jobKey(job.Id)
	keyExists, err := k.storeService.OpenKVStore(ctx).Has(key)
	if err != nil {
		return err
	}
	if keyExists {
		return fmt.Errorf("job id already exists")
	}
	bz, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return k.storeService.OpenKVStore(ctx).Set(jobKey(job.Id), bz)
}

func (k Keeper) getJob(ctx sdk.Context, jobID uint64) (Job, error) {
	key := jobKey(jobID)
	bz, err := k.storeService.OpenKVStore(ctx).Get(key)
	if err != nil {
		return Job{}, err
	}
	var job Job
	err = json.Unmarshal(bz, &job)
	if err != nil {
		return Job{}, err
	}
	return job, nil
}

func (k Keeper) evaluate(invocableName string, inputData string) string {
	cInvocableName := C.CString(invocableName)
	defer C.free(unsafe.Pointer(cInvocableName))
	cInputData := C.CString(inputData)
	defer C.free(unsafe.Pointer(cInputData))
	var outputDataLen C.size_t
	outputData := C.evaluate_invocable(cInvocableName, cInputData, &outputDataLen)
	defer C.free_memory(unsafe.Pointer(outputData), outputDataLen)
	return string(C.GoBytes(unsafe.Pointer(outputData), C.int(outputDataLen)))
}

func jobKey(jobID uint64) []byte {
	jobIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(jobIDBytes, jobID)
	return append([]byte(keyPrefixJob), jobIDBytes...)
}
