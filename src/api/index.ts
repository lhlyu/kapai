import instance from './request'
import {base64Encode} from '../util'

export const ApiAuthUrl = '/api/auth'

export interface Player {
    id: string
    name: string
    roomId: string
}

export interface Room {
    id: string
    status: number
    player1?: Player
    player2?: Player
}

// 创建房间
export const ApiCreateRoom = async (player: Player) => {
    return await instance.get<Player>(`/api/create_room?r=${base64Encode(JSON.stringify(player))}`)
}

// 获取房间列表
export const ApiGetRooms = async () => {
    return await instance.get<Room[]>('/api/get_rooms')
}

// 获取房间基本信息
export const ApiGetRoom = async (id: string) => {
    return await instance.get<Room>(`/api/get_room?id=${id}`)
}

// 加入房间
export const ApiJoinRoom = async (player: Player) => {
    return await instance.get(`/api/join_room?r=${base64Encode(JSON.stringify(player))}`)
}

// 离开房间
export const ApiLeaveRoom = async (player: Player) => {
    return await instance.get(`/api/leave_room?g=${base64Encode(JSON.stringify(player))}`)
}
