<script setup lang="ts">
    import { inject, onMounted, ref, watch, type Ref } from "vue";
    import { EncounterState, type CurrentStepRequestMessage, type ExpeditionDB, type ExpeditionStepResolveInfo, type Team, type User } from "../tools/types.ts"
    import { formatTextTimeFromTimeMS, getCurrentExpeditionStepResolveInfo, postRequest, RequestType } from "../tools/utils.ts"
    import TeamP from "./Team.vue";
    import Travel from "./Travel.vue";
    import Neutral from "./Neutral.vue";
    import Fight from "./Fight.vue";
    import type { NavigationHandler } from "@/tools/navigationHandler.ts";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()

    const timertxt = ref("");
    const answer = navigationHandler.getCurrentExpeditionStepResolveInfo()
    const eteam = ref<Team|undefined>(undefined)

    async function manageTimer(newtarget: Ref<ExpeditionStepResolveInfo | undefined>) {
        if (newtarget.value && newtarget.value.stepState) {
            eteam.value = newtarget.value.eTeam ? newtarget.value.eTeam : undefined
            console.log(newtarget.value.stepState)
            user.value!.state!.state = newtarget.value.stepState
            
            timer = launchTimer()
        } else {
            stopTimer();
            await navigationHandler.fetchExpeditionReport()
            user.value!.state!.state = EncounterState.Report
        }
    }
    
    onMounted(async () => {
        await navigationHandler.fetchCurrentExpeditionStepResolveInfo(user.value!.id)
        manageTimer(answer)
    })

    async function tick () {
        if (!answer.value || user.value!.state!.state == EncounterState.Report || !answer.value.timeline || !answer.value.timeline[answer.value.timeline.length-1]) {
            stopTimer()
            return
        }
        var countDownDate = Date.parse(answer.value.timeline[answer.value.timeline.length-1]!.when);
        // Get today's date and time
        var now = new Date().getTime();

        // Find the distance between now and the count down date
        var distance = countDownDate! - now;
        navigationHandler.applyTimelineEventToTeam(user, answer, now)
        
        // Time calculations for days, hours, minutes and seconds
        timertxt.value = formatTextTimeFromTimeMS(distance)

        // If the count down is finished, write some text
        if (distance < 0) {
            timertxt.value = "finished"
            await navigationHandler.fetchCurrentExpeditionStepResolveInfo(user.value!.id)
            stopTimer()
            manageTimer(answer)
        }
    }
    
    let timer: number | undefined

    function launchTimer() {
        let endAt = answer.value!.timeline[answer.value!.timeline!.length-1]!.when
        if (endAt) {
            console.log(answer.value)
            console.log(endAt)
            
            //console.log(props.user)
            // Update the count down every 1 second
            return setInterval(tick, 1000);
        }
    }

    function stopTimer() {
        clearInterval(timer);
    }


</script>

<template>
    <div v-if="answer != undefined">
        <p>time before encounter end {{ timertxt }}</p>
        <Travel v-if="user!.state!.state == EncounterState.Travel"/>
        <Neutral v-else-if="user!.state!.state == EncounterState.Neutral"/>
        <Fight v-if="user!.state!.state == EncounterState.Fight" :eteam="eteam"/>
    </div>
    <div v-else>
        <h1>Trying to know where the party is</h1>
    </div>
</template>
