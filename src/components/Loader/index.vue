<template>
    <section>
        <div class="loader">
            <div class="face">
                <div class="circle"></div>
            </div>
            <div class="face">
                <div class="circle"></div>
            </div>
        </div>
    </section>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
section {

    position: absolute;
    width: 100%;
    height: 100%;
    z-index: 9999;
    display: flex;
    justify-content: center;
    align-items: center;
    background: #212121;

    transition: opacity 0.4s linear;

    .loader {
        width: 6em;
        height: 6em;
        font-size: 10px;
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;

        .face {
            position: absolute;
            border-radius: 50%;
            border-style: solid;
            animation: animate 3s linear infinite;
            &:nth-child(1) {
                width: 100%;
                height: 100%;
                color: gold;
                border-color: currentColor transparent transparent currentColor;
                border-width: 0.2em 0.2em 0em 0em;
                --deg: -45deg;
                animation-direction: normal;
            }
            &:nth-child(2) {
                width: 70%;
                height: 70%;
                color: lime;
                border-color: currentColor currentColor transparent transparent;
                border-width: 0.2em 0em 0em 0.2em;
                --deg: -135deg;
                animation-direction: reverse;
            }

            .circle {
                position: absolute;
                width: 50%;
                height: 0.1em;
                top: 50%;
                left: 50%;
                background-color: transparent;
                transform: rotate(var(--deg));
                transform-origin: left;
                &::before {
                    position: absolute;
                    top: -0.5em;
                    right: -0.5em;
                    content: '';
                    width: 1em;
                    height: 1em;
                    background-color: currentColor;
                    border-radius: 50%;
                    box-shadow: 0 0 2em,
                    0 0 4em,
                    0 0 6em,
                    0 0 8em,
                    0 0 10em,
                    0 0 0 0.5em rgba(255, 255, 0, 0.1);
                }
            }
        }

    }
}

@keyframes animate {
    to {
        transform: rotate(1turn);
    }
}
</style>