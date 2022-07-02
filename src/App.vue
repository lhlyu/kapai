<template>
    <Loader v-show="store.loading"></Loader>
    <router-view></router-view>
</template>

<script setup lang="ts">
import Loader from '/src/components/Loader/index.vue'
import useSiteStore from "./stores/site";

const store = useSiteStore()
</script>