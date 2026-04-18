<script setup lang="ts">
    import { onMounted, ref, watch } from "vue"
    import type { ExpeditionStepResolveInfo, User } from "../model/types.ts"
    import type { Hero } from "../model/types.ts"
    import Team from "./Team.vue"
    import { global, fetchData, RequestType, postRequest, launchExpedition, formatTextTimeFromTimeMS } from "../model/utils.ts"

    // const currentHomeStatus = ref(HomeStatus.Noting);

    const props = defineProps<{
        user: User
    }>();

    let expeditions = ref<Record<string,number>|undefined>(undefined)

    onMounted(async () => {
        await fetchData<Record<string,number>>(expeditions, RequestType.AvailableExpeditions) 
        console.log(expeditions)
    })

    let selectedExp = ref("") 
    let selectionB = ref<Record<string,string>>({})
    function fillSelectionB(e: Record<string,number>) {
        for (const k in e) {
            if (k === selectedExp.value) {
                selectionB.value[k] = "eselected"
            } else {
                selectionB.value[k] = "enotselected"
            }
        }
    }

    watch(expeditions, (newExp) => {
        if (newExp) {
            fillSelectionB(newExp)
        }
    })

    function clickOnExpedition(e: string) {
        if (expeditions.value) {
            if (selectedExp.value !== e) {
                selectedExp.value = e
                fillSelectionB(expeditions.value)
            } else {
                selectedExp.value = ""
                fillSelectionB(expeditions.value)
            }
        }
    }

    let expStepInfo = ref<ExpeditionStepResolveInfo|undefined>(undefined)
    async function onclick() {
        let pathparms = [
            {id:"usr", value: ""+props.user.id},
            {id:"expId", value:selectedExp.value},
        ]
        launchExpedition(expStepInfo, props.user, selectedExp.value)
    }
    watch(expStepInfo, (newExpInfo) => {
        if (newExpInfo) {
            props.user.state.state = newExpInfo.stepState
        }
    })

</script>

<template>
    <div v-if="expeditions">
        <h1>Select an Expedition</h1>
        <div class="column">
            <div class="row"> 
                <div v-for="(value, key) in expeditions" v-on:click="clickOnExpedition(key)" :class="selectionB[key]">
                    <p style="text-align: center;">{{ key }}</p>
                    <p style="text-align: center;">time : {{ formatTextTimeFromTimeMS(value/1000000) }}</p>
                    <img :src="global.DOMAIN_NAME + global.EXPEDITION" width="150px">
                </div>
            </div>
            <div style="display: flex; align-items: center; justify-content: center;">
                <div>
                    <button v-on:click="onclick()">launch expedition</button>
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        <h1>Chargement...</h1>
    </div>
</template>


