<template>

</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {
    Application,
    Container,
    Sprite,
    filters,
    Graphics
} from 'pixi.js'

const app = new Application({
    backgroundColor: 0x000000,
    backgroundAlpha: 0.5,
    resizeTo: document.body,
    antialias: true
})

let container = new Container();

app.stage.addChild(container)

let sprite = Sprite.from("https://s3-us-west-2.amazonaws.com/s.cdpn.io/693612/IaUrttj.png");

sprite.width = 128;
sprite.height = 128;

container.addChild(sprite);



onMounted(() => {
    document.querySelector('#app')?.appendChild(app.view)
})


</script>
