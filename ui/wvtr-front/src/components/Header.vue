<script setup lang="ts">
    import { inject } from "vue";
import type { NavigationHandler } from "@/tools/navigationHandler.ts";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser() 

</script>

<template>
    <div v-if="user" class="header row">
        <div>
            <h1>{{ user.name }}</h1>
        </div>
        <div class="curencies" style="padding-left: 300px;">
            <div v-if="user.inventory && user.inventory.currencies" v-for="c in user.inventory.currencies" class="row" style="padding-right: 15px;">
                <div><img :src="c.currency.iconURL" width="16"/></div>
                <div>{{ c.numberOwned }}</div>
            </div>
        </div>
    </div>
    <div v-else>
        <h1>Loading...</h1>
    </div>
</template>