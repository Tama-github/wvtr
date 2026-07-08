<script setup lang="ts">
    import type { Skill } from '@/tools/types';
import { global } from '@/tools/utils';
import { ref, watch } from 'vue'
import Hoverable from '../Hoverable.vue';

    const props = defineProps<{
        skill: Skill;
    }>();
    
</script>

<template>
    <Hoverable>
        <template #default>
            <slot></slot>
        </template>
        <template #hover>
            <div class="column" v-if="skill">
                <h2>{{ skill.name }}</h2>
                <img width="75" :src="skill.image_url"/>
                <div> {{ skill.description }} </div>
            </div>
        </template>
    </Hoverable>
</template>