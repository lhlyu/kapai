import type { RouteRecordRaw } from 'vue-router'

import LoginView from '/src/views/login/index.vue'
import RoomView from '/src/views/room/index.vue'
import GameView from '/src/views/game/index.vue'

export const routes: RouteRecordRaw[] = [
    {
        name: 'Login',
        path: '/',
        component: LoginView
    },
    {
        name: 'Room',
        path: '/r',
        component: RoomView
    },
    {
        name: 'Game',
        path: '/g',
        component: GameView
    },
]
