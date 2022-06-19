import { defineStore } from 'pinia'

export interface IUser {
    id: string
    name: string
    roomId: string
    role?: 'Player' | 'Audience' | 'None'
}

// 用户
const useUserStore = defineStore({
    id: 'user',
    state: () =>
        ({
            id: '',
            name: '',
            roomId: ''
        } as IUser),
    getters: {
        isLogin: state => state.id.length > 0,
        inRoom: state => state.roomId.length > 0,
        getPlayer: state => ({
            id: state.id,
            name: state.name,
            roomId: state.roomId
        })
    },
    actions: {
        Login(name: string) {
            this.id = (+new Date()).toString(16)
            this.name = name
        },
        JoinRoom(roomId: string, role: 'Player' | 'Audience' = 'Player') {
            this.roomId = roomId
            this.role = role
        },
        LeaveRoom() {
            this.roomId = ''
            this.role = 'None'
        }
    },
    persist: {
        enabled: true,
        storage: window.sessionStorage
    }
})

export default useUserStore
