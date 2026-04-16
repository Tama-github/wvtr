<script setup lang="ts">
    import { ref, watch } from 'vue'
    import type { Team } from "../model/types.ts"
    import Hero from './Hero.vue'
    import { global } from "../model/utils.ts"

    const {team} = defineProps<{
        team: Team;
    }>();

</script>

<template>
    <div class="team">
        <Hero v-if="team.heroes" :hero="team.heroes[0]"/>
        <Hero v-else :hero="undefined"/>
        <Hero v-if="team.heroes" :hero="team.heroes[1]"/>
        <Hero v-else :hero="undefined"/>
        <Hero v-if="team.heroes" :hero="team.heroes[2]"/>
        <Hero v-else :hero="undefined"/>
    </div>
</template>
