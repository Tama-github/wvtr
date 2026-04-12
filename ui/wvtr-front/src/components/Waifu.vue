<script setup lang="ts">
    import { ref, watch } from 'vue'
    import type { Waifu } from "../model/types.ts"
    import { global } from "../model/utils.ts"

    const props = defineProps<{
        waifu: Waifu;
    }>();
</script>

<template>
    <img :src="global.DOMAIN_NAME + waifu.imageUrl" width="150"/>
</template>