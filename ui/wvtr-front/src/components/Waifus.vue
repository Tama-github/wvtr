<script setup lang="ts">
    import { onMounted, ref, watch } from "vue"
    import type { ExpeditionStepResolveInfo, User, Waifu } from "../model/types.ts"
    import type { Hero } from "../model/types.ts"
    import Team from "./Team.vue"
    import { global, fetchData, RequestType, postRequest, launchExpedition, formatTextTimeFromTimeMS, createAnHeroFromAWaifu } from "../model/utils.ts"
    import WaifuComp from "./WaifuComp.vue"

    // const currentHomeStatus = ref(HomeStatus.Noting);

    const props = defineProps<{
        user: User
    }>();

    let userWaifus = ref<Waifu[]|undefined>(undefined)

    onMounted(async () => {
        await fetchData<Waifu[]>(userWaifus, RequestType.UserWaifus, [{id: "id", value: `${props.user.id}`}]) 
        console.log(userWaifus)
    })

    let selectedWaifu = ref<Waifu|undefined>(undefined) 
    let selectionB = ref<Record<string,string>>({})
    function fillSelectionB(waifus: Waifu[]) {
        for (let i = 0; i < waifus.length; i++) {
            if (waifus[i] ==  selectedWaifu.value) {
                selectionB.value[waifus[i]!.id] = "eselected"
            } else {
                selectionB.value[waifus[i]!.id] = "enotselected"
            }
        }
    }

    watch(userWaifus, (newExp) => {
        if (newExp) {
            fillSelectionB(newExp)
        }
    })

    function clickOnWaifu(waifu: Waifu) {
        if (userWaifus.value) {
            if (selectedWaifu.value != waifu) {
                selectedWaifu.value = waifu
                fillSelectionB(userWaifus.value)
            } else {
                selectedWaifu.value = undefined
                fillSelectionB(userWaifus.value)
            }
        }
    }

    async function onclick() {
        createAnHeroFromAWaifu(ref<Hero|undefined>(undefined), props.user)
    }
    // watch(expStepInfo, (newExpInfo) => {
    //     if (newExpInfo) {
    //         props.user.state.state = newExpInfo.stepState
    //     }
    // })

</script>

<template>
    <div v-if="userWaifus">
        <h1>Select a Waifu</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="onclick()">Get hero</button>
            </div>
        </div>
        <div class="column">
            <div class="row" style="display: flex;flex-wrap: wrap;"> 
                <WaifuComp v-for="w in userWaifus" v-on:click="clickOnWaifu(w)" :waifu="w" :class="selectionB[w.id]"/>
            </div>
            
        </div>
    </div>
    <div v-else>
        <h1>Chargement...</h1>
    </div>
</template>


