<template>
    <section>
        <div class="login">
            <h1>登陆</h1>
            <input type="text" placeholder="输入昵称" @keydown.enter="loginFn" v-model="name" maxlength="6" />
            <button @click="loginFn">登陆</button>
        </div>
    </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import useUserStore from '../../stores/user'

const router = useRouter()
const store = useUserStore()

const name = ref('')

const loginFn = () => {
    name.value = name.value.trim()
    if (!name.value.length) {
        alert('昵称不能为空')
        return
    }
    store.Login(name.value)
    router.replace({
        path: '/r'
    })
}
</script>

<style lang="scss" scoped>
section {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    min-height: 100vh;
    .login {
        width: 90%;
        max-width: 460px;
        box-sizing: border-box;
        padding: 36px 15px;
        margin: 0 auto;
        border: 1px dashed #666;
        border-radius: 10px;
        text-align: center;
        input {
            height: 36px;
            font-size: 20px;
            outline: none;
            display: block;
            margin: 26px auto;
            border: 1px dashed #666;
            border-radius: 5px;
            text-align: center;
        }
        button {
            display: block;
            margin: 0 auto;
            padding: 6px 40px;
            border: 1px dashed #666;
            border-radius: 5px;
            outline: none;
            background: white;
            &:hover {
                background: whitesmoke;
            }
        }
    }
}
</style>
