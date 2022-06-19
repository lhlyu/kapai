<template>
    <button @click="leaveRoom">离开游戏</button>
</template>

<script setup lang="ts">
import useUserStore from '../../stores/user'
import { useRouter } from 'vue-router'
import { ApiLeaveRoom } from '../../api'

const store = useUserStore()
const router = useRouter()

const leaveRoom = async () => {
    if (store.role === 'Audience') {
        store.LeaveRoom()
        router.replace({ path: '/r' })
        return
    }
    if (store.role === 'Player') {
        store.LeaveRoom()
        await ApiLeaveRoom(store.getPlayer)
    }
}
</script>

<style scoped></style>
