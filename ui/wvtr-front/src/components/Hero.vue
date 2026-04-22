<script setup lang="ts">
    import { ref, watch } from 'vue'
    import type { Hero } from "../model/types.ts"
    import { global } from "../model/utils.ts"
import InspectButton from './InspectButton.vue';

    const props = defineProps<{
        hero: Hero | undefined;
    }>();
    
</script>

<template>
    <div v-if="hero" class="hero">
        <div class="waifu-image-container">
            <img :src="hero.imageUrl" width="150"/>
            <InspectButton :hero="hero"/>
        </div>
        
        <div class="column">
            <h1>{{ hero.name }}<label style="font-size: medium;"> lvl.{{ hero.attributes.level }}</label></h1>    
            <div class="column" style="display: inline-flex;">
                <label>hp : {{ hero.attributes.currentHP.toFixed(0) }}/{{ hero.attributes.maxHP.toFixed(0) }} </label>
                <progress :max="hero.attributes.maxHP.toFixed(0)" :value="hero.attributes.currentHP.toFixed(0)" class="hero-progress"/>
                <label>xp : {{ hero.attributes.currentXP.toFixed(0) }}/{{ hero.attributes.xpBeforLvlUp.toFixed(0) }} </label>
                <progress :max="hero.attributes.xpBeforLvlUp.toFixed(0)" :value="hero.attributes.currentXP.toFixed(0)" class="hero-progress"/>
            </div>
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