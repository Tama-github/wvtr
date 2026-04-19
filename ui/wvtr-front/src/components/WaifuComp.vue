<script setup lang="ts">
    import { ref, watch } from 'vue'
    import type { Hero, Waifu } from "../model/types.ts"
    import { global } from "../model/utils.ts"

    const props = defineProps<{
        waifu: Waifu;
    }>();

    let rankTextClass = "waifu-rank-e"

    switch (props.waifu.rank) {
        case "S":
            rankTextClass = "waifu-rank-s"
            break
        case "A":
            rankTextClass = "waifu-rank-a"
            break
        case "B":
            rankTextClass = "waifu-rank-b"
            break
        case "C":
            rankTextClass = "waifu-rank-c"
            break
        case "D":
            rankTextClass = "waifu-rank-d"
            break
        case "E":
            rankTextClass = "waifu-rank-e"
            break
    }
    
</script>

<template>
    <div class="column">
        <div>{{ props.waifu.name_user_preferred }}</div>
        <div class="waifu-image-container">
            <img :src="props.waifu.image_large" width="150"/>
            <div :class="rankTextClass" text>{{ props.waifu.rank }}</div>
        </div>
    </div>
</template>