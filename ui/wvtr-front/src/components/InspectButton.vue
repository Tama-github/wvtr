<script setup lang="ts">
    import { inject } from "vue";
import type { Hero } from "../model/types.ts"
    import { global, HomeStatus, NavigationHandler } from "../model/utils.ts"

    const props = defineProps<{
        hero: Hero;
    }>();

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    function onClickInspectButton() {
        navigationHandler.setHomeStatus(HomeStatus.InspectHero)
        navigationHandler.setHeroToInspect(props.hero)
    }
</script>

<template>
    <button class="inspection-button" v-on:click="onClickInspectButton()">
        <img :src="global.DOMAIN_NAME + '/imgs/magnifyingglassicon.png'" width="25px"/>
    </button>
</template>