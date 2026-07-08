<script setup lang="ts">
    import type { Weapon } from '@/tools/types';
import { global } from '@/tools/utils';
import { ref, watch } from 'vue'
import Hoverable from '../Hoverable.vue';

    const props = defineProps<{
        weapon: Weapon;
    }>();
    
</script>

<template>


    <Hoverable>
        <template #default>
            <slot></slot>
        </template>
        <template #hover>
            <div class="column" v-if="weapon">
                <h2>{{ weapon.name }}</h2>
                <img width="75" :src="weapon.iconURL"/>
                <div v-if="weapon.baseDamage">
                    <div>Damage:</div>
                    <div class="row">
                        <div class="column">
                            <div v-if="weapon.baseDamage.slashDmg > 0">Slash:</div>
                            <div v-if="weapon.baseDamage.bluntDmg > 0">Blunt:</div>
                            <div v-if="weapon.baseDamage.pierceDmg > 0">Pierce:</div>
                            <div v-if="weapon.baseDamage.fireDmg > 0">Fire:</div>
                            <div v-if="weapon.baseDamage.frostDmg > 0">Frost:</div> 
                            <div v-if="weapon.baseDamage.lightningDmg > 0">Lightning:</div>
                        </div>
                        <div class="column">
                            <div v-if="weapon.baseDamage.slashDmg > 0">{{ weapon.baseDamage.slashDmg.toFixed(2) }}</div>
                            <div v-if="weapon.baseDamage.bluntDmg > 0">{{ weapon.baseDamage.bluntDmg.toFixed(2) }}</div>
                            <div v-if="weapon.baseDamage.pierceDmg > 0">{{ weapon.baseDamage.pierceDmg.toFixed(2) }}</div>
                            <div v-if="weapon.baseDamage.fireDmg > 0">{{ weapon.baseDamage.fireDmg.toFixed(2) }}</div>
                            <div v-if="weapon.baseDamage.frostDmg > 0">{{ weapon.baseDamage.frostDmg.toFixed(2) }}</div>
                            <div v-if="weapon.baseDamage.lightningDmg > 0">{{ weapon.baseDamage.lightningDmg.toFixed(2) }}</div>
                        </div>
                    </div>    
                </div>
                <div>Crit rate: {{ weapon.baseCritRate!.value.toFixed(2) }}</div>
                <div>Attack speed: {{ weapon.baseAttackSpeed!.value.toFixed(2) }}</div>
                <div class="column" v-if="weapon.affixes && weapon.affixes.length>0">
                    <div v-for="a in weapon.affixes">
                        {{ a.name }}
                    </div>
                </div>
            </div>
        </template>
    </Hoverable>
</template>