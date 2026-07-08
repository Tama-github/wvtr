<script setup lang="ts">
    import type { Omamori } from '@/tools/types';
import { global } from '@/tools/utils';
import { ref, watch } from 'vue'
import Hoverable from '../Hoverable.vue';

    const props = defineProps<{
        omamori: Omamori;
    }>();
    
</script>

<template>
    <Hoverable>
        <template #default>
            <slot></slot>
        </template>
        <template #hover>
            <div class="column" v-if="omamori">
                <h2>{{ omamori.name }}</h2>
                <img width="75" :src="omamori.iconURL"/>
                <div class="column" v-if="omamori.affixes && omamori.affixes.length>0">
                    <div v-for="a in omamori.affixes">
                        {{ a.name }}
                    </div>
                </div>
            </div>
        </template>
    </Hoverable>
</template>