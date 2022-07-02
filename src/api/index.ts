import { request } from './request'
import {Encode} from '../util'

export const ApiAuthUrl = '/api/auth'


export interface LoginParam {
    account: string
    password: string
}

export interface UserModel {
    account: string
    name: string
    avatar: string
    token?: string
    decks?: number[]
    spells?: number[]
    isNew?: boolean
}

export interface Room {
    id: string
    status: number
    black?: UserModel
    red?: UserModel
}

// 登陆/注册
export const ApiLogin = async (param: LoginParam) => {
    return await request<UserModel>('login', 'k', Encode(JSON.stringify(param)))
}

// 创建房间
export const ApiCreateRoom = async () => {
    return await request<string>('create.room')
}


export interface RoomParam {
    room_id: string
}

// 加入房间
export const ApiJoineRoom = async (param: RoomParam) => {
    return await request<string>('join.room')
}

// 离开房间
export const ApiLeaveRoom = async (param: RoomParam) => {
    return await request<boolean>('leave.room')
}


// 获取房间信息
export const ApiGetRoom = async (roomId: string) => {
    return await request<Room>('get.room', 'r', roomId)
}

// 获取所有房间
export const ApiGetRooms = async () => {
    return await request<Room[]>('get.rooms')
}

export interface UpdateUserParam {
    name?: string
    avatar?: string
    password?: string
    decks?: number[]
    spells?: number[]
}

// 更新用户信息
export const ApiUpdateUser = async (param: UpdateUserParam) => {
    return await request<boolean>('update.user','k', Encode(JSON.stringify(param)))
}

