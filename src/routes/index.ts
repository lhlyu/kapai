import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'
import useUserStore from '../stores/user'

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to, from) => {
    const store = useUserStore()
    if (!store.isLogin && to.name !== 'Login') {
        return { name: 'Login' }
    }
    if (store.isLogin && store.inRoom && to.name !== 'Game') {
        return { name: 'Game', params: { id: store.roomId } }
    }
    if (store.isLogin && !store.inRoom && to.name !== 'Room') {
        return { name: 'Room' }
    }
})

export default router
