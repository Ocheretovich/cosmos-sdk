package keeper

import (
	"context"

	"cosmossdk.io/x/mint/types"
)

var _ types.QueryServer = queryServer{}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// Params returns params of the mint module.
func (q queryServer) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params, err := q.k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: params}, nil
}

// Deprecated: DO NOT USE
// This method uses deprecated query request.
// Inflation returns minter.Inflation of the mint module.
func (q queryServer) Inflation(ctx context.Context, _ *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	minter, err := q.k.Minter.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryInflationResponse{Inflation: minter.Inflation}, nil
}

// Deprecated: DO NOT USE
// This method uses deprecated query request.
// AnnualProvisions returns minter.AnnualProvisions of the mint module.
func (q queryServer) AnnualProvisions(ctx context.Context, _ *types.QueryAnnualProvisionsRequest) (*types.QueryAnnualProvisionsResponse, error) {
	minter, err := q.k.Minter.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryAnnualProvisionsResponse{AnnualProvisions: minter.AnnualProvisions}, nil
}

func (q queryServer) EpochProvisions(ctx context.Context, _ *types.QueryEpochProvisionsRequest) (*types.QueryEpochProvisionsResponse, error) {
	minter, err := q.k.Minter.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryEpochProvisionsResponse{EpochProvisions: minter.EpochProvisions}, nil
}
