package keeper

import (
	"context"
	"errors"

	"github.com/arcane-systems/rook/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Config.Map.Seed == 0 {
		return nil, errors.New("no map seed specified")
	}

	params := m.Keeper.GetLatestParamsVersion(ctx)

	overview, state, err := types.SetupGame(msg.Players, &msg.Config, params)
	if err != nil {
		return nil, err
	}

	gameID, err := m.Keeper.GetGameID(ctx)
	if err != nil {
		return nil, err
	}
	// save the next game ID
	m.Keeper.SetGameID(ctx, gameID+1)

	// persist the game overview. We will save game state in the end block
	m.Keeper.SetGameOverview(ctx, gameID, overview)
	m.Keeper.SetGameState(ctx, gameID, state)

	// emit the events for the new game
	ctx.EventManager().EmitTypedEvent(&types.EventNewGame{
		GameId:        gameID,
		Players:       msg.Players,
		Config:        &msg.Config,
		ParamsVersion: params,
	})

	return &types.MsgCreateResponse{GameId: gameID}, nil
}

func (m msgServer) Build(goCtx context.Context, msg *types.MsgBuild) (*types.MsgBuildResponse, error) {
	resp := &types.MsgBuildResponse{}
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, err := m.Keeper.GetGame(ctx, msg.GameId)
	if err != nil {
		return resp, err
	}

	err = game.Build(msg.Creator, msg.Populace, msg.Settlement)
	if err != nil {
		return resp, err
	}

	m.Keeper.SetGameState(ctx, msg.GameId, game.State())
	return resp, nil
}

func (m msgServer) Move(goCtx context.Context, msg *types.MsgMove) (*types.MsgMoveResponse, error) {
	resp := &types.MsgMoveResponse{}
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, err := m.Keeper.GetGame(ctx, msg.GameId)
	if err != nil {
		return resp, err
	}

	err = game.Move(msg.Creator, msg.Populace, msg.Direction, msg.Population)
	if err != nil {
		return resp, err
	}

	m.Keeper.SetGameState(ctx, msg.GameId, game.State())
	return resp, err
}

func (m msgServer) ChangeParams(goCtx context.Context, msg *types.MsgChangeParams) (*types.MsgChangeParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.Params.ValidateBasic()
	if err != nil {
		return &types.MsgChangeParamsResponse{}, err
	}

	version := m.Keeper.SetParams(ctx, msg.Params)

	return &types.MsgChangeParamsResponse{Version: version}, nil
}
