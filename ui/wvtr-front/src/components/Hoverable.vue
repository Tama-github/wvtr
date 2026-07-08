<script setup lang="ts">
import { ref } from 'vue';

    const isHover = ref(false)
</script>

<template>
    <div class="base-content" v-on:mouseenter="isHover = true" v-on:mouseleave="isHover = false">
        <slot>

        </slot>
        <div v-if="isHover" class="hover-content">
            <slot name="hover">

            </slot>
        </div>
    </div>
</template>

<style scoped>
    .hover-content {
        position: absolute;
        background-color: aliceblue;
        top: calc(100%);
        overflow: visible;
        width: 450px;
        left: 50%;
        transform: translateX(-50%);
    }

    .base-content {
        position: relative;
    }
</style>