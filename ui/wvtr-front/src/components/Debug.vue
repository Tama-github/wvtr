<script setup lang="ts">
    import { ref, watch, reactive } from 'vue'
    import type { Waifu } from "../model/types.ts"
    import { global } from "../model/utils.ts"
    
    // Get Waifu by id
    // const waifu1: Waifu = JSON.parse(await fetch("/waifus/1"))
    const BASE_URL = 'https://tama.rhiobet.sh/';
    const waifuId = ref(1)
    //let waifu : Waifu = {}
    const jsonres = ref<Waifu>({})
    const azert = ref(false)

    // async function api<T>(path: string): Promise<T> {
    //     const response = await fetch(`${BASE_URL}/${path}`);

    //     if (!response.ok) {
    //     throw new Error(response.statusText);
    //     }

    //     //    And can also be used here ↴
    //     return await response.json() as T;
    // }

    async function fetchData() {
        azert.value = false
        jsonres.value = {}
        const res = await fetch(
            `https://tama.rhiobet.sh/waifus/${waifuId.value}`
        )
        azert.value = true
        jsonres.value = await res.json() as Waifu
    }

    fetchData()
</script>

<template>
    <p v-if="JSON.stringify(jsonres)==='{}'">Chargement...</p>
    <div v-else>
        <p> {{ JSON.stringify(jsonres) }} </p>
        <p> {{ jsonres.id }} </p>
        <p> {{ jsonres.imageUrl }} </p>

        <img :src="'https://tama.rhiobet.sh'+jsonres.imageUrl" width="150"/>
    </div>

</template>