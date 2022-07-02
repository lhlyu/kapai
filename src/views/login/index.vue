<template>
    <section>
        <div class="login">
            <input v-model="loginParam.account" spellcheck="false" type="text" size="30" maxlength="12" placeholder="登陆or注册"  autocomplete="off">
            <!--      设置成password, 浏览器存在自动填充的问题，解决不了，改成text      -->
            <input v-model="loginParam.password" spellcheck="false" type="text" size="30" maxlength="16" placeholder="密_码" @keydown.enter="loginFn" autocomplete="off">
        </div>
    </section>
</template>

<script setup lang="ts">
import {ref} from "vue";
import useSiteStore from "../../stores/site";
import type { LoginParam } from "../../api";
import { ApiLogin } from "../../api";

const site = useSiteStore()

const loginParam = ref<LoginParam>({
    account: '',
    password: ''
})

const loginFn = async () => {
    site.loading = true
    const result = await ApiLogin(loginParam.value)

    setTimeout(() => {
        site.loading = false
        console.log(result)
    }, 1000)
}

</script>

<style lang="scss" scoped>
section {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    .login {
        max-width: 400px;
        width: 100%;
        text-align: center;
        input {
            box-sizing: border-box;
            outline: 0;
            border: 1px solid #fff;
            width: 90%;
            height: 46px;
            border-radius: 20px;
            padding: 5px 20px;
            color: inherit;
            background-color: #fff;
            font-size: 16px;
            text-align: center;
            margin: 8px 0;
        }
    }
}
</style>
