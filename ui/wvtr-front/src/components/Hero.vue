<script setup lang="ts">
    import { ref, watch } from 'vue'
    import type { Hero } from "../model/types.ts"
    import { global } from "../model/utils.ts"

    const props = defineProps<{
        hero: Hero | undefined;
    }>();
    
</script>

<template>
    <div v-if="hero" class="hero">
        <img :src="global.DOMAIN_NAME + hero.imageUrl" width="150"/>
        <div class="column">
            <h1>{{ hero.name }}<label style="font-size: medium;"> lvl.{{ hero.level }}</label></h1>
            <label>hp : {{ hero.currentHP }}/{{ hero.maxHP }} </label>
            <progress :max="hero.maxHP" :value="hero.currentHP"/>
            <label>xp : {{ hero.currentXP }}/{{ hero.xpBeforLvlUp }} </label>
            <progress :max="hero.xpBeforLvlUp" :value="hero.currentXP"/>
        </div>
    </div>
    <div v-else class="hero">
        <img :src="global.DOMAIN_NAME + global.NO_IMAGE" width="150"/>
        <div class="column">
            <h1>{{  }}<label style="font-size: medium;"> lvl.{{ 0 }}</label></h1>
            <label>hp : {{ 0 }}/{{ 0 }} </label>
            <progress :max="0" :value="0"/>
            <label>xp : {{ 0 }}/{{ 0 }} </label>
            <progress :max="0" :value="0"/>
        </div>
    </div>
</template>