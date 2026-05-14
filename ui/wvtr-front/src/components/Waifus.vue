<script setup lang="ts">
    import { inject, onMounted, ref, watch } from "vue"
    import type { User, Waifu } from "../tools/types.ts"
    import type { Hero } from "../tools/types.ts"
    import { global, fetchData, RequestType, createAHeroFromAWaifu } from "../tools/utils.ts"
    import WaifuComp from "./WaifuComp.vue"
    import { NavigationStatus, NavigationHandler } from "../tools/navigationHandler.ts"

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const userWaifus = navigationHandler.getUserWaifus()

    let selectedWaifu = ref<Waifu|undefined>(undefined) 
    let selectionB = ref<Record<string,string>>({})

    onMounted(()=>{
        if (userWaifus.value) {
            fillSelectionB(userWaifus.value)
        }
    })
    function fillSelectionB(waifus: Waifu[]) {
        for (let i = 0; i < waifus.length; i++) {
            if (waifus[i] ==  selectedWaifu.value) {
                selectionB.value[waifus[i]!.id] = "eselected"
            } else {
                selectionB.value[waifus[i]!.id] = "enotselected"
            }
        }
    }

    function clickOnWaifu(waifu: Waifu) {
        if (userWaifus.value) {
            if (selectedWaifu.value != waifu) {
                selectedWaifu.value = waifu
                fillSelectionB(userWaifus.value!)
            } else {
                selectedWaifu.value = undefined
                fillSelectionB(userWaifus.value!)
            }
        }
    }

    async function onclick() {
        if (selectedWaifu.value) {
            navigationHandler.createAHeroFromAWaifu(selectedWaifu.value)
        }
    }

    watch(userWaifus, (nuw)=>{
        if (nuw) {
            fillSelectionB(nuw)
        }
    })

</script>

<template>
    <div v-if="userWaifus">
        <h1>Select a Waifu</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="onclick()">Make a hero</button>
            </div>
        </div>
        <div class="column">
            <div class="row" style="display: flex;flex-wrap: wrap;"> 
                <WaifuComp v-for="w in userWaifus" v-on:click="clickOnWaifu(w)" :waifu="w" :class="selectionB[w.id]"/>
            </div>
        </div>
    </div>
    <div v-else>
        <h1>Loading...</h1>
    </div>
</template>


