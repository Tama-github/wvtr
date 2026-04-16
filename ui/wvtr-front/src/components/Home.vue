<script setup lang="ts">
    import { onMounted, ref, watch } from "vue"
    import type { User } from "../model/types.ts"
    import Team from "./Team.vue"
    import TeamManagement from "./TeamManagement.vue"
    import ExpeditionsList from "./ExpeditionsList.vue"

    enum HomeStatus {
        Noting = 1,
        ExpeditionManagement,
        TeamManagement,
    }

    const currentHomeStatus = ref(HomeStatus.Noting);

    const props = defineProps<{
        user: User;
    }>();

    
    function setHomeStatus (newStatus: HomeStatus) {
        currentHomeStatus.value = newStatus
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
        </div>
        <div>
            <TeamManagement v-if="currentHomeStatus == HomeStatus.TeamManagement" 
                    :user="user" 
                    :ownedHeroes="user.ownedHeroes"/>
            
            <ExpeditionsList v-else-if="currentHomeStatus == HomeStatus.ExpeditionManagement" :user="user"/>
        </div>
    <!-- </div> -->
</template>


