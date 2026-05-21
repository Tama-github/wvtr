<script setup lang="ts">
    import { inject, onMounted, ref, watch } from "vue"
    import { type ExpeditionStepResolveInfo, type ExpToLaunch } from "../tools/types.ts"
    import { global, formatTextTimeFromTimeMS, getTypeFromName } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts"
import type { ExpeditionCategory } from "@/tools/expeditions.ts"

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
        
    const expeditionsCats = navigationHandler.getAvailableExpedition()
    const user = navigationHandler.getUser()

    let selectedExp = ref<ExpToLaunch|undefined>(undefined)
    //let selectedCat = ref("")
    const errorMsg = ref("") 
    let selectionB = ref<Record<string,string>>({})
    onMounted(()=>{
        fillSelectionB(expeditionsCats.value!)
    })

    function fillSelectionB(e: ExpeditionCategory[]) {
        for (let i = 0; i < e.length; i++) {
            for(let j = 0; j< e[i]!.expeditions.length; j++) {
                if (e[i]!.expeditions[j]!.key === selectedExp.value?.key) {
                    selectionB.value[e[i]!.expeditions[j]!.key] = "eselected"
                } else if (!e[i]!.expeditions[j]!.canBeLaunched) {
                    selectionB.value[e[i]!.expeditions[j]!.key] = "cantbeselected"
                } else {
                    selectionB.value[e[i]!.expeditions[j]!.key] = "enotselected"
                }
            }
        }
    }

    function clickOnExpedition(e: string, cat: string, cost: string, costNumber: number) {
        if (expeditionsCats.value) {
            if (!selectedExp.value || selectedExp.value!.key !== e) {
                selectedExp.value = {
                    cat: cat,
                    key: e,
                    cost: cost,
                    costNumber: costNumber,
                }
                fillSelectionB(expeditionsCats.value)
            } else {
                selectedExp.value = {
                    cat: "",
                    key: "",
                    cost: "",
                    costNumber: 0,
                }
                fillSelectionB(expeditionsCats.value)
            }
        }
    }

    let expStepInfo = ref<ExpeditionStepResolveInfo|undefined>(undefined)
    async function onclick() {
        if (user.value!.currentTeam!.heroes.length > 0) {
            console.log("cat : "+ selectedExp.value?.cat)
            console.log("id : "+ selectedExp.value?.key)
            let eType = getTypeFromName(selectedExp.value?.cost!) 
            if (eType != undefined) { // if the cost is an equipment we need the user to select wich one to spend.
                navigationHandler.setInventoryForExpeditionLaunch(eType, selectedExp.value!)
            } else {
                await navigationHandler.launchExpedition(expStepInfo, selectedExp.value?.cat!, selectedExp.value?.key!, undefined)
            }
        } else {
            errorMsg.value = "You should have at least one hero in your team before starting an expedition."
        }
    }
    watch(expStepInfo, (newExpInfo) => {
        if (newExpInfo) {
            user.value!.state!.state = newExpInfo.stepState
        }
    })

</script>

<template>
    <div v-if="expeditionsCats">
        <h1>Select an Expedition</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="onclick()">launch expedition</button>
                <p>{{ errorMsg }}</p>
            </div>
        </div>
        <div v-for="cat in expeditionsCats">
            <div class="column">
                <div>
                    <h1>{{cat.name}}</h1>
                </div>
                <div class="row"> 
                    <div v-for="e in cat.expeditions" :class="selectionB[e.key]">
                        <div v-if="e.canBeLaunched" v-on:click="clickOnExpedition(e.key, cat.name, e.costName, e.costNumber)">
                            <p style="text-align: center;">{{ e.key }}</p>
                            <p style="text-align: center;">time : {{ formatTextTimeFromTimeMS(e.duration/1000000) }}</p>
                            <p v-if="e.costName != ''" style="text-align: center;">Cost : {{ e.costName }} / {{ e.costNumber}} </p>
                            <img :src="e.imgURL" width="150px">
                        </div>
                        <div v-else :class="selectionB[e.key]" >
                            <p style="text-align: center;">{{ e.key }}</p>
                            <p style="text-align: center;">time : {{ formatTextTimeFromTimeMS(e.duration/1000000) }}</p>
                            <p v-if="e.costName != ''" style="text-align: center;">Cost : {{ e.costName }} / {{ e.costNumber}} </p>
                            <img :src="e.imgURL" width="150px">
                        </div>    
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        <h1>Loading...</h1>
    </div>
</template>


