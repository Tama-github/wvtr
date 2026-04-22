<script setup lang="ts">
    import { ref, watch } from "vue"
    import type { Team, User } from "../model/types.ts"
    import type { Hero } from "../model/types.ts"
    import { global, postRequest, RequestType } from "../model/utils.ts"

    // const currentHomeStatus = ref(HomeStatus.Noting);

    const props = defineProps<{
        user: User
        ownedHeroes: Hero[];
    }>();

    let selectedH = ref<Hero[]>([])
    let selectionB = ref(new Array(props.ownedHeroes.length).fill(false))
    
    function clickOnHero(h: Hero) {
        console.log("clicked!")
        console.log(h)
        if (selectedH.value.includes(h)) {
            console.log("already selected")
            let idx = selectedH.value.indexOf(h)
            if (idx > -1) {
                console.log("remove from selected")
                selectedH.value.splice(idx, 1);
                selectionB.value[props.ownedHeroes.indexOf(h)] = false
            }
        } else if (selectedH.value.length < 3) {
            console.log("Add to selected")
            selectedH.value.push(h)
            selectionB.value[props.ownedHeroes.indexOf(h)] = true
        }
        console.log("selected :" + selectedH)
        console.log("selection filter :" +selectionB)
    }

    async function saveTeam() {
        // send request to modify current team de user
        props.user.currentTeam.heroes = selectedH.value
        selectedH = ref<Hero[]>([])
        selectionB.value.fill(false)
        let tmpTeam = ref<Team|undefined>(undefined)
        await postRequest<Team, User>(tmpTeam, props.user, RequestType.UpdateTeam)
        console.log(tmpTeam)
        if (tmpTeam.value) {
            props.user.currentTeam = tmpTeam.value
        }
    }

</script>

<template>
    <div>
        <h1>Select You team</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="saveTeam()">Save</button>
            </div>
        </div>
        <div class="row"> 
            <div v-for="h in ownedHeroes">
                <img v-if="!selectionB[ownedHeroes.indexOf(h)]" class="hnotselected" :src="h.imageUrl" v-on:click="clickOnHero(h)">
                <img v-else class="hselected" :src="h.imageUrl" v-on:click="clickOnHero(h)">
            </div>
        </div>
    </div>
</template>
