<script setup lang="ts">
    import type { User } from "../model/types.ts"
    import { global } from "../model/utils.ts"

    const props = defineProps<{
        user: User | undefined;
    }>();
</script>

<template>
    <div class="header">
        <div v-if="user">
            <h1>{{ user.name }}</h1>
        </div>
        <div v-else>
            <h1>Chargement ...</h1>
        </div>
    </div>
</template>