package keeper

/*
#cgo LDFLAGS:  -L/usr/local/lib -ldmnengine
#include <stdlib.h>

size_t load_models(const char* c_dir);
*/
import "C"
import (
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
)

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
