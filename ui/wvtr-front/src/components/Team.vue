<script setup lang="ts">
    import { ref, watch } from 'vue'
    import type { Team } from "../model/types.ts"
    import Waifu from './Waifu.vue'
    import { global } from "../model/utils.ts"

    const {team} = defineProps<{
        team: Team;
    }>();

    // const {teamId} = defineProps<{
    //     teamId: number;
    // }>();

    // //const teamId = ref(1)
    // const team = ref<Team>({})
    // const request = global.REQ_TEAM + teamId

    // async function fetchData() {
    //     team.value = {}
    //     const res = await fetch(request)
    //     team.value = await res.json() as Team
    // }

    // fetchData()
</script>

<template>
    <!-- <p>{{ JSON.stringify(team) }}</p> -->
    <!-- <p v-if="JSON.stringify(team)==='{}'">Chargement...</p> -->
    <div>
        <Waifu v-for="waifu in team.waifus" :waifu="waifu"/>
    </div>
</template>