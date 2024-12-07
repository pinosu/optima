package keeper

import (
	"context"

	"optima/x/optima/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EvaluateInvocable(goCtx context.Context, req *types.QueryEvaluateInvocableRequest) (*types.QueryEvaluateInvocableResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	_ = sdk.UnwrapSDKContext(goCtx)
	return &types.QueryEvaluateInvocableResponse{EvaluationResult: k.evaluate(req.InvocableName, req.InputData)}, nil
}
