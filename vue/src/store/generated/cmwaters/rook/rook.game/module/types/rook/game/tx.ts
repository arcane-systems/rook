/* eslint-disable */
import { Direction, Settlement, Position, directionFromJSON, directionToJSON, settlementFromJSON, settlementToJSON } from '../../rook/game/types'
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { GameConfig } from '../../rook/game/config'

export const protobufPackage = 'rook.game'

export interface MsgMove {
  creator: string
  gameId: number
  position: Position | undefined
  direction: Direction
  population: number
}

export interface MsgMoveResponse {}

export interface MsgBuild {
  creator: string
  gameId: number
  settlement: Settlement
  position: Position | undefined
}

export interface MsgBuildResponse {}

export interface MsgCreate {
  /** all players must be signers */
  players: string[]
  config: GameConfig | undefined
}

export interface MsgCreateResponse {
  gameId: number
}

const baseMsgMove: object = { creator: '', gameId: 0, direction: 0, population: 0 }

export const MsgMove = {
  encode(message: MsgMove, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint64(message.gameId)
    }
    if (message.position !== undefined) {
      Position.encode(message.position, writer.uint32(26).fork()).ldelim()
    }
    if (message.direction !== 0) {
      writer.uint32(32).int32(message.direction)
    }
    if (message.population !== 0) {
      writer.uint32(40).uint32(message.population)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMove {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgMove } as MsgMove
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.gameId = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.position = Position.decode(reader, reader.uint32())
          break
        case 4:
          message.direction = reader.int32() as any
          break
        case 5:
          message.population = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgMove {
    const message = { ...baseMsgMove } as MsgMove
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId)
    } else {
      message.gameId = 0
    }
    if (object.position !== undefined && object.position !== null) {
      message.position = Position.fromJSON(object.position)
    } else {
      message.position = undefined
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = directionFromJSON(object.direction)
    } else {
      message.direction = 0
    }
    if (object.population !== undefined && object.population !== null) {
      message.population = Number(object.population)
    } else {
      message.population = 0
    }
    return message
  },

  toJSON(message: MsgMove): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.gameId !== undefined && (obj.gameId = message.gameId)
    message.position !== undefined && (obj.position = message.position ? Position.toJSON(message.position) : undefined)
    message.direction !== undefined && (obj.direction = directionToJSON(message.direction))
    message.population !== undefined && (obj.population = message.population)
    return obj
  },

  fromPartial(object: DeepPartial<MsgMove>): MsgMove {
    const message = { ...baseMsgMove } as MsgMove
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId
    } else {
      message.gameId = 0
    }
    if (object.position !== undefined && object.position !== null) {
      message.position = Position.fromPartial(object.position)
    } else {
      message.position = undefined
    }
    if (object.direction !== undefined && object.direction !== null) {
      message.direction = object.direction
    } else {
      message.direction = 0
    }
    if (object.population !== undefined && object.population !== null) {
      message.population = object.population
    } else {
      message.population = 0
    }
    return message
  }
}

const baseMsgMoveResponse: object = {}

export const MsgMoveResponse = {
  encode(_: MsgMoveResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgMoveResponse } as MsgMoveResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgMoveResponse {
    const message = { ...baseMsgMoveResponse } as MsgMoveResponse
    return message
  },

  toJSON(_: MsgMoveResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgMoveResponse>): MsgMoveResponse {
    const message = { ...baseMsgMoveResponse } as MsgMoveResponse
    return message
  }
}

const baseMsgBuild: object = { creator: '', gameId: 0, settlement: 0 }

export const MsgBuild = {
  encode(message: MsgBuild, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint64(message.gameId)
    }
    if (message.settlement !== 0) {
      writer.uint32(24).int32(message.settlement)
    }
    if (message.position !== undefined) {
      Position.encode(message.position, writer.uint32(34).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuild {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgBuild } as MsgBuild
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.gameId = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.settlement = reader.int32() as any
          break
        case 4:
          message.position = Position.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgBuild {
    const message = { ...baseMsgBuild } as MsgBuild
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId)
    } else {
      message.gameId = 0
    }
    if (object.settlement !== undefined && object.settlement !== null) {
      message.settlement = settlementFromJSON(object.settlement)
    } else {
      message.settlement = 0
    }
    if (object.position !== undefined && object.position !== null) {
      message.position = Position.fromJSON(object.position)
    } else {
      message.position = undefined
    }
    return message
  },

  toJSON(message: MsgBuild): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.gameId !== undefined && (obj.gameId = message.gameId)
    message.settlement !== undefined && (obj.settlement = settlementToJSON(message.settlement))
    message.position !== undefined && (obj.position = message.position ? Position.toJSON(message.position) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgBuild>): MsgBuild {
    const message = { ...baseMsgBuild } as MsgBuild
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId
    } else {
      message.gameId = 0
    }
    if (object.settlement !== undefined && object.settlement !== null) {
      message.settlement = object.settlement
    } else {
      message.settlement = 0
    }
    if (object.position !== undefined && object.position !== null) {
      message.position = Position.fromPartial(object.position)
    } else {
      message.position = undefined
    }
    return message
  }
}

const baseMsgBuildResponse: object = {}

export const MsgBuildResponse = {
  encode(_: MsgBuildResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuildResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgBuildResponse } as MsgBuildResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgBuildResponse {
    const message = { ...baseMsgBuildResponse } as MsgBuildResponse
    return message
  },

  toJSON(_: MsgBuildResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgBuildResponse>): MsgBuildResponse {
    const message = { ...baseMsgBuildResponse } as MsgBuildResponse
    return message
  }
}

const baseMsgCreate: object = { players: '' }

export const MsgCreate = {
  encode(message: MsgCreate, writer: Writer = Writer.create()): Writer {
    for (const v of message.players) {
      writer.uint32(18).string(v!)
    }
    if (message.config !== undefined) {
      GameConfig.encode(message.config, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreate } as MsgCreate
    message.players = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 2:
          message.players.push(reader.string())
          break
        case 3:
          message.config = GameConfig.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreate {
    const message = { ...baseMsgCreate } as MsgCreate
    message.players = []
    if (object.players !== undefined && object.players !== null) {
      for (const e of object.players) {
        message.players.push(String(e))
      }
    }
    if (object.config !== undefined && object.config !== null) {
      message.config = GameConfig.fromJSON(object.config)
    } else {
      message.config = undefined
    }
    return message
  },

  toJSON(message: MsgCreate): unknown {
    const obj: any = {}
    if (message.players) {
      obj.players = message.players.map((e) => e)
    } else {
      obj.players = []
    }
    message.config !== undefined && (obj.config = message.config ? GameConfig.toJSON(message.config) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreate>): MsgCreate {
    const message = { ...baseMsgCreate } as MsgCreate
    message.players = []
    if (object.players !== undefined && object.players !== null) {
      for (const e of object.players) {
        message.players.push(e)
      }
    }
    if (object.config !== undefined && object.config !== null) {
      message.config = GameConfig.fromPartial(object.config)
    } else {
      message.config = undefined
    }
    return message
  }
}

const baseMsgCreateResponse: object = { gameId: 0 }

export const MsgCreateResponse = {
  encode(message: MsgCreateResponse, writer: Writer = Writer.create()): Writer {
    if (message.gameId !== 0) {
      writer.uint32(8).uint64(message.gameId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateResponse } as MsgCreateResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.gameId = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateResponse {
    const message = { ...baseMsgCreateResponse } as MsgCreateResponse
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId)
    } else {
      message.gameId = 0
    }
    return message
  },

  toJSON(message: MsgCreateResponse): unknown {
    const obj: any = {}
    message.gameId !== undefined && (obj.gameId = message.gameId)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateResponse>): MsgCreateResponse {
    const message = { ...baseMsgCreateResponse } as MsgCreateResponse
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId
    } else {
      message.gameId = 0
    }
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Move(request: MsgMove): Promise<MsgMoveResponse>
  Build(request: MsgBuild): Promise<MsgBuildResponse>
  Create(request: MsgCreate): Promise<MsgCreateResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Move(request: MsgMove): Promise<MsgMoveResponse> {
    const data = MsgMove.encode(request).finish()
    const promise = this.rpc.request('rook.game.Msg', 'Move', data)
    return promise.then((data) => MsgMoveResponse.decode(new Reader(data)))
  }

  Build(request: MsgBuild): Promise<MsgBuildResponse> {
    const data = MsgBuild.encode(request).finish()
    const promise = this.rpc.request('rook.game.Msg', 'Build', data)
    return promise.then((data) => MsgBuildResponse.decode(new Reader(data)))
  }

  Create(request: MsgCreate): Promise<MsgCreateResponse> {
    const data = MsgCreate.encode(request).finish()
    const promise = this.rpc.request('rook.game.Msg', 'Create', data)
    return promise.then((data) => MsgCreateResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
