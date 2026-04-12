<script setup lang="ts">
    import { ref, watch, onMounted } from 'vue'
    import type { Team } from "./model/types.ts"
    import type { GameState } from "./model/types.ts"
    import type { User } from "./model/types.ts"
    import { EncounterState } from "./model/types.ts"
    import { RequestType } from "./model/utils.ts"
    import VTeam from "./components/Team.vue"
    import { global, fetchData } from "./model/utils.ts"

    const user = ref<User|undefined>(undefined)

    // async function fetchData() {
    //     user.value = {}
    //     const res = await fetch(request)
    //     user.value = await res.json() as User
    // }
    // onMounted(async () => {
    //     user.value = await fetchData<User>(RequestType.User, [{id: "id", value: "1"}]) 
    // })
    onMounted(async () => {
        await fetchData<User>(user, RequestType.User, [{id: "id", value: "2"}]) 
    })
</script>

<template>
<!-- <Debug/> -->
<p v-if="!user">Chargement...</p>
<div v-else-if="user.state.state == EncounterState.Home" class="wrapper">
  <p>user home</p>
</div>
<div v-else-if="user.state.state == EncounterState.Travel" class="wrapper">
  <p>travel</p>
</div>
<div v-else-if="user.state.state == EncounterState.Fight" class="wrapper">
  <p>fight</p>
</div>
<div v-else-if="user.state.state == EncounterState.Neutral" class="wrapper">
  <p>neutral</p>
</div>
<div v-else-if="user.state.state == EncounterState.Error" class="wrapper">
  <p>error</p>
</div>
</template>

<style>
.wrapper {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(5, 200px);
  background-color: bisque;
}

.box1 {
  grid-column-start: 1;
  grid-column-end: 1;
  grid-row-start: 2;
  grid-row-end: 5;
  border-style: solid;
}

.box2 {
  grid-column-start: 3;
  grid-column-end: 3;
  grid-row-start: 2;
  grid-row-end: 5;
  border-style: solid;
}
</style>
