<script setup lang="ts">
    import type { User } from "../model/types.ts"
    import { EncounterState } from "../model/types.ts"
    import Home from "./Home.vue"
    import Travel from "./Travel.vue"
    import Fight from "./Fight.vue"
    import Neutral from "./Neutral.vue"

    const props = defineProps<{
        user: User | undefined;
    }>();
</script>

<template>
    <div v-if="user" class="body">
        <Home v-if="user.state.state == EncounterState.Home" :user="user" />
        <Travel v-else-if="user.state.state == EncounterState.Travel" :user="user" />
        <Fight v-else-if="user.state.state == EncounterState.Fight" :user="user" />
        <Neutral v-else-if="user.state.state == EncounterState.Neutral" :user="user" />
        <div v-else-if="user.state.state == EncounterState.Error">
            <h1> There is a problem </h1>
        </div>
    </div>
    <div v-else class="body">
        <h1>Chargement ...</h1>
    </div>
</template>