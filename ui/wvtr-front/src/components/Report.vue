<script setup lang="ts">
    import type { NavigationHandler } from '@/tools/navigationHandler';
    import { EncounterState, HeroTakeDamageStatus, type ExpeditionDB, type User } from '@/tools/types';
    import { formatTextTimeFromTimeMS, getEncounterStateString, getStringFromFAD } from '@/tools/utils';
    import { inject, onMounted, ref } from 'vue';


    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    
    const report = navigationHandler.getReport()
    const reportJson = JSON.stringify(report)
    async function onclick() {
        await navigationHandler.closeReport()
    }
</script>

<template>
    <div>
        <h1>Report</h1>
        <div v-for="evR in report.whatHappened">
            <h2>{{ getEncounterStateString(evR.stepState!) }}</h2>
            <div v-for="truc in evR.timeline">
                <div class="row">
                    <div>
                        {{  new Date(truc.when).toLocaleString() }} :
                    </div>
                    <div class="column">
                        <div>
                        {{ truc.what }}
                        </div>
                        <div v-if="truc.whatAction">
                            <div v-for="txt in getStringFromFAD(truc.whatAction)">
                                {{ txt }}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div>
            <h1>Rewards</h1>
            <div v-if="report.expeditionRewards">
                <div v-if="report.expeditionRewards.xp > 0">
                    <p>The team got {{ report.expeditionRewards.xp }}xp points.</p>
                </div>
                <div class="row" v-if="report.expeditionRewards.loot!.currencies.length>0">
                    <div class="column" v-for="c in report.expeditionRewards.loot!.currencies">
                        <img :src="c.currency!.iconURL" width="16px">
                        <div> {{ c.numberOwned }} </div>
                    </div>
                </div>
                <div v-if="report.expeditionRewards.loot!.weapons.length>0">
                    <div class="column" v-for="w in report.expeditionRewards.loot!.weapons">
                        <img :src="w.iconURL" width="50px">
                        <div> Damage: {{ w.baseDamage }} </div>
                        <div> Attack speed: {{ w.baseAttackSpeed }} </div>
                        <div> Crit rate: {{ w.baseCritRate }} </div>
                        <div v-if="w.affixes.length>0">
                            <div>Affixes:</div>
                            <div v-for="a in w.affixes">
                                <div> {{ a.name }} </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div v-if="report.expeditionRewards.loot!.armors.length>0">
                    <div class="column" v-for="w in report.expeditionRewards.loot!.armors">
                        <img :src="w.iconURL" width="50px">
                        <div> Resistances: {{ w.baseResistancesRange }} </div>
                        <div> Block score: {{ w.blockScore?.value }} </div>
                        <div> Evade score: {{ w.evadeScore?.value }} </div>
                        <div v-if="w.affixes.length>0">
                            <div>Affixes:</div>
                            <div v-for="a in w.affixes">
                                <div> {{ a.name }} </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div v-if="report.expeditionRewards.loot!.omamoris.length>0">
                    <div class="column" v-for="w in report.expeditionRewards.loot!.omamoris">
                        <img :src="w.iconURL" width="50px">
                        <div v-if="w.affixes.length>0">
                            <div>Affixes:</div>
                            <div v-for="a in w.affixes">
                                <div> {{ a.name }} </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <button v-on:click="onclick()">ok</button>
    </div>
</template>