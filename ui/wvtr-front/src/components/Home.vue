<script setup lang="ts">
    import { inject, onMounted, ref, watch } from "vue"
    import type { Hero, User } from "../model/types.ts"
    import Team from "./Team.vue"
    import TeamManagement from "./TeamManagement.vue"
    import ExpeditionsList from "./ExpeditionsList.vue"
import Waifus from "./Waifus.vue"
import InspectHero from "./InspectHero.vue"
import { HomeStatus, NavigationHandler } from "@/model/utils.ts"

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!

    const props = defineProps<{
        user: User;
    }>();

    
    function setHomeStatus (newStatus: HomeStatus) {
        navigationHandler!.setHomeStatus(newStatus)
    }
</script>

<template>
    <!-- <div class="home"> -->
        <Team :team="user!.currentTeam"/>
        <div class="column">
            <button v-on:click="setHomeStatus(HomeStatus.TeamManagement)">
            manage Team
            </button>
            <button v-on:click="setHomeStatus(HomeStatus.ExpeditionManagement)">
            launch expedition
            </button>
            <button v-on:click="setHomeStatus(HomeStatus.HeroGetter)">
            Check available waifus
            </button>
        </div>
        <TeamManagement v-if="navigationHandler.getHomeStatus().value == HomeStatus.TeamManagement" 
                :user="user" 
                :ownedHeroes="user.ownedHeroes"/>
        
        <ExpeditionsList v-else-if="navigationHandler.getHomeStatus().value == HomeStatus.ExpeditionManagement" :user="user"/>
        <Waifus v-else-if="navigationHandler.getHomeStatus().value == HomeStatus.HeroGetter" :user="user"/>
        <InspectHero v-else-if="navigationHandler.getHomeStatus().value == HomeStatus.InspectHero"/>
    <!-- </div> -->
</template>


