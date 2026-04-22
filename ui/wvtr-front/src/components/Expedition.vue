<script setup lang="ts">
    import { onMounted, ref, watch } from "vue";
    import { EncounterState, type CurrentStepRequestMessage, type ExpeditionStepResolveInfo, type User } from "../model/types.ts"
    import { formatTextTimeFromTimeMS, getCurrentExpeditionStepResolveInfo, postRequest, RequestType } from "../model/utils.ts"
    import Team from "./Team.vue";
    import Travel from "./Travel.vue";
    import Neutral from "./Neutral.vue";
    import Fight from "./Fight.vue";

    const props = defineProps<{
        user: User;
    }>();

    const timertxt = ref("")

    const answer = ref<ExpeditionStepResolveInfo|undefined>(undefined)
    const tmp = ref<ExpeditionStepResolveInfo|undefined>(undefined)
    onMounted(async () => {
        await getCurrentExpeditionStepResolveInfo(answer, props.user.id) 
    })

    async function tick () {
        if (!answer.value) {
            return
        }
        var countDownDate = Date.parse(answer.value.stepEndAt);
        // Get today's date and time
        var now = new Date().getTime();

        // Find the distance between now and the count down date
        var distance = countDownDate! - now;
        //console.log(distance)
        // Time calculations for days, hours, minutes and seconds
        timertxt.value = formatTextTimeFromTimeMS(distance)

        // If the count down is finished, write some text
        if (distance < 0) {
            timertxt.value = "finished"
            await getCurrentExpeditionStepResolveInfo(tmp, props.user.id) 
            answer.value = tmp.value
        }
    }
    
    let timer: number | undefined

    function launchTimer() {
        console.log(props.user)
        if (answer.value && answer.value.stepEndAt) {
            console.log(answer.value)
            console.log(answer.value.stepEndAt)
            
            //console.log(props.user)
            // Update the count down every 1 second
            return setInterval(tick, 1000);
        }
    }

    watch(answer, (newtarget)=>{
        if (newtarget && newtarget.stepState) {
            props.user.state.state = newtarget.stepState
            timer = launchTimer()
        } else {
            props.user.state.state = EncounterState.Home
            clearInterval(timer);
        }
    })
</script>

<template>
    <Team :team="props.user.currentTeam"/>
    <div v-if="answer != undefined">
        <Travel v-if="props.user.state.state == EncounterState.Travel"/>
        <Neutral v-else-if="props.user.state.state == EncounterState.Neutral"/>
        <Fight v-if="props.user.state.state == EncounterState.Fight"/>
        <p>time before encounter end {{ timertxt }}</p>
    </div>
    <div v-else>
        <h1>Trying to know where the party is</h1>
    </div>
</template>
