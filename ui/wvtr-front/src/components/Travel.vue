<script setup lang="ts">
    import { onMounted, ref, watch } from "vue";
    import { EncounterState, type CurrentStepRequestMessage, type ExpeditionStepResolveInfo, type User } from "../model/types.ts"
    import { getCurrentExpeditionStepResolveInfo, postRequest, RequestType } from "../model/utils.ts"

    const props = defineProps<{
        user: User;
    }>();

    const timerd = ref("0")
    const timerh = ref("0")
    const timerm = ref("0")
    const timers = ref("0")
    const timertxt = ref("")

    const answer = ref<ExpeditionStepResolveInfo|undefined>(undefined)
    onMounted(async () => {
        await getCurrentExpeditionStepResolveInfo(answer, props.user.id) 
    })

    async function tick () {
        console.log("tick")
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
        var days = Math.floor(distance / (1000 * 60 * 60 * 24));
        var hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
        var seconds = Math.floor((distance % (1000 * 60)) / 1000);

        if (seconds>0) {
            timertxt.value = seconds+"s"
        }
        if (minutes>0) {
            timertxt.value = minutes+"m " + timertxt.value
        }
        if (hours>0) {
            timertxt.value = hours+"h " + timertxt.value
        }
        if (days>0) {
            timertxt.value = days+"d " + timertxt.value
        }

        // If the count down is finished, write some text
        if (distance < 0) {
            timertxt.value = "finished"
            await getCurrentExpeditionStepResolveInfo(answer, props.user.id) 
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
        if (newtarget) {
            props.user.state.state = newtarget.stepState
            timer = launchTimer()
        } else {
            props.user.state.state = EncounterState.Home
            clearInterval(timer);
        }
    })
</script>

<template>
    <div v-if="answer != undefined">
        <h1>Traveling</h1>
        <p>time before travel end {{ timertxt }}</p>
    </div>
    <div v-else>
        <h1>Trying to know where the party is</h1>
    </div>
</template>
