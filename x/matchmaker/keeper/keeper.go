package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/arcane-systems/rook/x/matchmaker/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Codec
		params   paramtypes.Subspace
		storeKey sdk.StoreKey
		router   *baseapp.MsgServiceRouter
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	router *baseapp.MsgServiceRouter,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		params:   paramSpace,
		router:   router,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateGame(ctx sdk.Context, room types.Room) {
	// NOTE: we use the block time as a deterministic source of randomness for generating
	// game maps. We may in the future want a better source of randomness but this suffices
	// for now.
	msg := room.MsgCreate(ctx.BlockTime().Unix())

	handler := k.router.Handler(msg)
	if handler == nil {
		panic(fmt.Sprintf("unrecognized message route to create game: %s", sdk.MsgTypeURL(msg)))
	}

	res, err := handler(ctx, msg)
	if err != nil {
		panic(err)
	}

	// emit the events from creating the game
	events := res.Events
	sdkEvents := make([]sdk.Event, 0, len(events))
	for i := 0; i < len(events); i++ {
		sdkEvents = append(sdkEvents, sdk.Event(events[i]))
	}
	ctx.EventManager().EmitEvents(sdkEvents)
}

// UpdateRooms loops through all rooms and performs the following:
// - Starts games that have full capacity
// - Starts games for rooms who have reached quorum and have exceeded the
// prestart_wait_period
// - Removes rooms that never reached the quorum within the room_lifespan
func (k Keeper) UpdateRooms(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	now := ctx.BlockTime()
	params := k.GetParams(ctx)

	roomIter := store.Iterator(
		types.RoomKey(0),
		types.RoomKey(1<<63-1),
	)
	for ; roomIter.Valid(); roomIter.Next() {
		var room types.Room
		k.cdc.MustUnmarshal(roomIter.Value(), &room)

		roomID := types.ParseRoomKey(roomIter.Key())

		switch t := room.Time.(type) {
		case *types.Room_Created:
			if room.IsFull() {
				k.CreateGame(ctx, room)
				k.DeleteRoom(ctx, roomID)
				k.RemoveRoomFromModePool(ctx, room.ModeId, roomID)
			} else if room.HasQuorum() {
				room.ReadyUp(now)
			} else if room.HasExpired(now, params.RoomLifespan) {
				k.DeleteRoom(ctx, roomID)
				k.RemoveRoomFromModePool(ctx, room.ModeId, roomID)
			}
		case *types.Room_Ready:
			if now.After(t.Ready.Add(params.PrestartWaitPeriod)) {
				k.CreateGame(ctx, room)
				k.DeleteRoom(ctx, roomID)
				k.RemoveRoomFromModePool(ctx, room.ModeId, roomID)
			}
		case *types.Room_Scheduled:
			if now.After(*t.Scheduled) {
				if room.HasQuorum() {
					k.CreateGame(ctx, room)
				}
				k.DeleteRoom(ctx, roomID)
				k.RemoveRoomFromModePool(ctx, room.ModeId, roomID)
			}
		}

	}
}

// FindPlayer performs a range scan returning the room that the player is currently in,
// if any. CONTRACT: a player can never be in more than one room.
// TODO: This process is performed often and could be made a lot more efficient if we used
// a secondary key.
func (k Keeper) FindPlayer(ctx sdk.Context, player string) (bool, types.IndexedRoom) {
	store := ctx.KVStore(k.storeKey)

	roomIter := store.Iterator(
		types.RoomKey(0),
		types.RoomKey(1<<63-1),
	)
	for ; roomIter.Valid(); roomIter.Next() {
		var room types.Room
		k.cdc.MustUnmarshal(roomIter.Value(), &room)

		for _, roomPlayer := range room.Players {
			if roomPlayer == player {
				return true, types.IndexedRoom{
					Id:   types.ParseRoomID(roomIter.Key()),
					Room: room,
				}
			}
		}
	}

	return false, types.IndexedRoom{}
}

func (k Keeper) GetRoom(ctx sdk.Context, roomID uint64) (types.Room, bool) {
	room := types.Room{}
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.RoomKey(roomID))
	if bz == nil {
		return room, false
	}

	k.cdc.MustUnmarshal(bz, &room)
	return room, true
}

func (k Keeper) SetRoom(ctx sdk.Context, roomID uint64, room types.Room) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&room)

	store.Set(types.RoomKey(roomID), bz)
}

func (k Keeper) DeleteRoom(ctx sdk.Context, roomID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RoomKey(roomID))
}

func (k Keeper) GetNextRoomID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	return types.ParseRoomID(store.Get(types.RoomIDKey))
}

func (k Keeper) IncrementNextRoomID(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	currentRoomID := types.ParseRoomID(store.Get(types.RoomIDKey))

	nextRoomIDBytes := types.RoomIDBytes(currentRoomID + 1)

	store.Set(types.RoomIDKey, nextRoomIDBytes)
}

func (k Keeper) GetNextModeID(ctx sdk.Context) uint32 {
	store := ctx.KVStore(k.storeKey)

	return types.ParseModeID(store.Get(types.ModeIDKey))
}

func (k Keeper) GetAllModes(ctx sdk.Context) []types.Mode {
	modes := make([]types.Mode, 0)
	store := ctx.KVStore(k.storeKey)

	modeIter := store.Iterator(
		types.ModeKey(0),
		types.ModeKey(1<<31-1),
	)
	for ; modeIter.Valid(); modeIter.Next() {
		var mode types.Mode
		k.cdc.MustUnmarshal(modeIter.Value(), &mode)
		modes = append(modes, mode)
	}

	return modes
}

func (k Keeper) IncrementNextModeID(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	currentModeID := types.ParseModeID(store.Get(types.ModeIDKey))

	nextModeIDBytes := types.ModeIDBytes(currentModeID + 1)

	store.Set(types.ModeIDKey, nextModeIDBytes)
}

func (k Keeper) SetMode(ctx sdk.Context, modeID uint32, mode types.Mode) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&mode)

	store.Set(types.ModeKey(modeID), bz)
}

func (k Keeper) GetMode(ctx sdk.Context, modeID uint32) (types.Mode, bool) {
	var mode types.Mode
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.ModeKey(modeID))
	if bz == nil {
		return mode, false
	}

	k.cdc.MustUnmarshal(bz, &mode)
	return mode, true
}

func (k Keeper) DeleteMode(ctx sdk.Context, modeID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ModeKey(modeID))
	store.Delete(types.RoomsByModeKey(modeID))
}

// NOTE: We return multiple rooms but I think it's not possible for there to be more than one room
// for a mode at a time.
func (k Keeper) GetRoomsByMode(ctx sdk.Context, modeID uint32) (types.Rooms, bool) {
	store := ctx.KVStore(k.storeKey)
	var rooms types.Rooms
	bz := store.Get(types.RoomsByModeKey(modeID))
	if bz == nil {
		return types.Rooms{}, false
	}
	k.cdc.MustUnmarshal(bz, &rooms)
	return rooms, true
}

func (k Keeper) SetRooms(ctx sdk.Context, modeID uint32, rooms types.Rooms) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RoomsByModeKey(modeID), k.cdc.MustMarshal(&rooms))
}

func (k Keeper) RemoveRoomFromModePool(ctx sdk.Context, modeID uint32, roomID uint64) {
	if modeID > 0 {
		rooms, exists := k.GetRoomsByMode(ctx, modeID)
		if !exists {
			panic("tried to remove a room from the mode pool that doesn't exist")
		}
		rooms.Remove(roomID)
		k.SetRooms(ctx, modeID, rooms)
	}
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.params.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	store := ctx.KVStore(k.storeKey)
	k.SetParams(ctx, genState.Params)
	for _, mode := range genState.InitialModes {
		k.SetMode(ctx, genState.NextModeId, mode)
		k.SetRooms(ctx, genState.NextModeId, types.Rooms{})
		genState.NextModeId++
	}
	store.Set(types.ModeIDKey, types.ModeIDBytes(genState.NextModeId))
	store.Set(types.RoomIDKey, types.RoomIDBytes(genState.NextRoomId))
}

func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	return types.GenesisState{
		Params:       k.GetParams(ctx),
		InitialModes: k.GetAllModes(ctx),
		NextRoomId:   k.GetNextRoomID(ctx),
		NextModeId:   k.GetNextModeID(ctx),
	}
}
