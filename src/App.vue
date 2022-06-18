<template>

</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import * as PIXI from 'pixi.js'
import { gsap } from "gsap";
import { PixiPlugin } from "gsap/PixiPlugin"

gsap.registerPlugin(PixiPlugin);
PixiPlugin.registerPIXI(PIXI)

const app = new PIXI.Application({
    backgroundColor: 0x000000,
    backgroundAlpha: 0.5,
    resizeTo: document.body,
    antialias: true
})

let container = new PIXI.Container();



const sprite = PIXI.Sprite.from("https://s3-us-west-2.amazonaws.com/s.cdpn.io/693612/IaUrttj.png");

sprite.width = 64;
sprite.height = 64;


container.addChild(sprite);

window.addEventListener('mousemove', (ev) => {
    gsap.to(sprite, {
        pixi: {
            x: ev.x,
            y: ev.y,
        },
        duration: 0.4
    })
})

app.stage.addChild(container)



onMounted(() => {
    document.querySelector('#app')?.appendChild(app.view)
})


</script>
