<script setup lang="ts">
    import type { User } from "../model/types.ts"
    import { EncounterState } from "../model/types.ts"
    import Home from "./Home.vue"
    import Expedition from "./Expedition.vue"

    const props = defineProps<{
        user: User | undefined;
    }>();
</script>

<template>
    <div v-if="user" class="body">
        <Home v-if="user.state.state == EncounterState.Home" :user="user" />
        <div v-else-if="user.state.state == EncounterState.Error">
            <h1> There is a problem </h1>
        </div>
        <Expedition v-else :user="user"/>
    </div>
    <div v-else class="body">
        <h1>Chargement ...</h1>
    </div>
</template>