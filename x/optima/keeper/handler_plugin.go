package keeper

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// abstract keeper
type msKeeper interface {
	Evaluate(invocableName string, inputData string) string
}

type CustomMsgHandler struct {
	k *Keeper
}

func NewCustomMsgHandler(k *Keeper) *CustomMsgHandler {
	return &CustomMsgHandler{k: k}
}

type OptimaMsg struct {
	Invocable Invocable `json:"invocable"`
}

type Invocable struct {
	InvocableName string `json:"invocable_name"`
	InputData     string `json:"input_data"`
}

func (h CustomMsgHandler) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) (events []sdk.Event, data [][]byte, msgResponses [][]*codectypes.Any, err error) {
	if msg.Custom == nil {
		return nil, nil, nil, wasmtypes.ErrUnknownMsg
	}
	var customMsg OptimaMsg
	if err := json.Unmarshal(msg.Custom, &customMsg); err != nil {
		return nil, nil, nil, sdkerrors.ErrJSONUnmarshal.Wrap("custom message")
	}

	res := h.k.Evaluate(customMsg.Invocable.InvocableName, customMsg.Invocable.InputData)
	ctx.Logger().Info("=========================================================")
	ctx.Logger().Info(res)
	return []sdk.Event{}, nil, nil, nil
}
