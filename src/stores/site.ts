import { defineStore } from 'pinia'

// 网站
const useSiteStore = defineStore({
    id: 'site',
    state: () => ({
        loading: false
    }),
    persist: {
        enabled: true,
        storage: window.sessionStorage
    }
})

export default useSiteStore
