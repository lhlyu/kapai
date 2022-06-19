<template>
    <section>
        <h1>房间</h1>
        <small>
            <p><strong>绿点房间</strong>表示空闲，可以加入房间</p>
            <p><strong>红点房间</strong>表示正在游戏，点击玩家可以观战</p>
        </small>
        <ul>
            <li v-if="!store.inRoom">
                <button class="cursor" @click="createRoom">创建房间</button>
            </li>
            <li v-for="v in rooms" :key="v.id">
                <h3><small>房间号：</small>{{ v.id }}</h3>
                <span :class="statuses[v.status]"></span>
                <button @click="() => joinRoom(v.id, 1)" class="player" :class="v?.player1 && v.status === 1 ? '' : 'cursor'">
                    {{ v?.player1 ? v.player1.name : '+加入房间' }}
                </button>
                <button @click="() => joinRoom(v.id, 2)" class="player" :class="v?.player2 && v.status === 1 ? '' : 'cursor'">
                    {{ v?.player2 ? v.player2.name : '+加入房间' }}
                </button>
            </li>
        </ul>
    </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ApiGetRoom, ApiGetRooms, ApiCreateRoom, ApiJoinRoom } from '../../api'
import type { Room } from '../../api'
import useUserStore from '../../stores/user'
import { useRouter } from 'vue-router'

const store = useUserStore()
const router = useRouter()

const rooms = ref<Room[]>([])

const statuses: Record<number, string> = {
    1: 'status-wait',
    2: 'status-full'
}

const loadRooms = async () => {
    const result = await ApiGetRooms()
    if (result.code) {
        return
    }
    rooms.value = result.data
}

// 加入房间
const joinRoom = async (roomId: string, index: number) => {
    const result = await ApiGetRoom(roomId)
    if (result.code) {
        return
    }
    const room = result.data
    if (room.status === 1) {
        switch (index) {
            case 1:
                if (room.player1) {
                    alert('加入失败：位置已被其他用户占据')
                    await loadRooms()
                    return
                }
            case 2:
                if (room.player2) {
                    alert('加入失败：位置已被其他用户占据')
                    await loadRooms()
                    return
                }
        }

        const result1 = await ApiJoinRoom({
            ...store.getPlayer,
            roomId: roomId
        })
        if (result1.code) {
            return
        }
        store.JoinRoom(roomId, 'Player')
        router.replace({ path: '/g', params: { id: roomId } })

        return
    }

    if (room.status === 2) {
        store.JoinRoom(roomId, 'Audience')
        router.replace({ path: '/g', params: { id: roomId } })
        return
    }
}

// 创建房间
const createRoom = async () => {
    const result = await ApiCreateRoom(store.getPlayer)
    if (result.code) {
        return
    }
    store.JoinRoom(result.data.roomId)
    router.replace({ path: '/g', params: { id: result.data.roomId } })
}

onMounted(async () => {
    await loadRooms()
})
</script>

<style lang="scss" scoped>
section {
    padding: 20px 1.33333vw;
    box-sizing: border-box;
    text-align: center;
    h1 {
        display: block;
    }
    small {
        margin-top: 20px;
        display: block;
        line-height: 2;
    }
    ul {
        list-style-type: none;
        padding: 0;
        display: flex;
        flex-direction: row;
        justify-content: center;
        flex-wrap: wrap;
        width: 100%;
        margin: 20px 0;
        li {
            position: relative;
            border: 1px dashed #666;
            border-radius: 10px;
            max-width: 300px;
            width: 40%;
            margin: 10px;
            padding: 3.13333vw;
            display: inline-block;
            box-sizing: border-box;
            transition: all 0.2s linear;
            &:hover {
                border: 1px solid #666;
            }
            .status-wait {
                position: absolute;
                right: 10px;
                top: 10px;
                width: 8px;
                height: 8px;
                border-radius: 50%;
                background: #2ed573;
            }
            .status-full {
                position: absolute;
                right: 10px;
                top: 10px;
                width: 8px;
                height: 8px;
                border-radius: 50%;
                background: red;
            }
            h3 {
                margin: 0 10px 10px;
                padding-bottom: 10px;
                border-bottom: 1px dashed #666;
                small {
                    display: inline-block;
                }
            }
            button {
                padding: 5px 8px;
                border: 1px dashed #666;
                border-radius: 5px;
                margin: 5px;
                outline: none;
                background: white;
                transition: all 0.2s linear;
                &:hover {
                    border: 1px solid #666;
                }
            }
            .cursor {
                cursor: pointer;
            }
        }
    }
}
</style>
