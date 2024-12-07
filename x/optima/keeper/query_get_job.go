package keeper

import (
	"context"

	"optima/x/optima/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetJob(goCtx context.Context, req *types.QueryGetJobRequest) (*types.QueryGetJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	job, err := k.getJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetJobResponse{EvaluationResult: job.Result}, nil
}
