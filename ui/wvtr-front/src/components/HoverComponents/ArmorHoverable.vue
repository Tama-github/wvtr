<script setup lang="ts">
    import type { Armor } from '@/tools/types';
import { global } from '@/tools/utils';
import { ref, watch } from 'vue'
import Hoverable from '../Hoverable.vue';

    const props = defineProps<{
        armor: Armor;
    }>();
    
</script>

<template>


    <Hoverable>
        <template #default>
            <slot></slot>
        </template>
        <template #hover>
            <div class="column" v-if="armor">
                <h2>{{ armor.name }}</h2>
                <img width="75" :src="armor.iconURL"/>
                <div v-if="armor.baseResistancesRange">
                    <div>Damage:</div>
                    <div class="row">
                        <div class="column">
                            <div v-if="armor.baseResistancesRange.slashDmg > 0">Slash res:</div>
                            <div v-if="armor.baseResistancesRange.bluntDmg > 0">Blunt res:</div>
                            <div v-if="armor.baseResistancesRange.pierceDmg > 0">Pierce res:</div>
                            <div v-if="armor.baseResistancesRange.fireDmg > 0">Fire res:</div>
                            <div v-if="armor.baseResistancesRange.frostDmg > 0">Frost res:</div> 
                            <div v-if="armor.baseResistancesRange.lightningDmg > 0">Lightning res:</div>
                        </div>
                        <div class="column">
                            <div v-if="armor.baseResistancesRange.slashDmg > 0">{{ armor.baseResistancesRange.slashDmg.toFixed(2) }}</div>
                            <div v-if="armor.baseResistancesRange.bluntDmg > 0">{{ armor.baseResistancesRange.bluntDmg.toFixed(2) }}</div>
                            <div v-if="armor.baseResistancesRange.pierceDmg > 0">{{ armor.baseResistancesRange.pierceDmg.toFixed(2) }}</div>
                            <div v-if="armor.baseResistancesRange.fireDmg > 0">{{ armor.baseResistancesRange.fireDmg.toFixed(2) }}</div>
                            <div v-if="armor.baseResistancesRange.frostDmg > 0">{{ armor.baseResistancesRange.frostDmg.toFixed(2) }}</div>
                            <div v-if="armor.baseResistancesRange.lightningDmg > 0">{{ armor.baseResistancesRange.lightningDmg.toFixed(2) }}</div>
                        </div>
                    </div>    
                </div>
                <div>Block score: {{ armor.blockScore!.value.toFixed(2) }}</div>
                <div>Evade score: {{ armor.evadeScore!.value.toFixed(2) }}</div>
                <div class="column" v-if="armor.affixes && armor.affixes.length>0">
                    <div v-for="a in armor.affixes">
                        {{ a.name }}
                    </div>
                </div>
            </div>
        </template>
    </Hoverable>
</template>