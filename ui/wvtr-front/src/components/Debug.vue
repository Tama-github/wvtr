<script setup lang="ts">
    import { ref, watch, reactive } from 'vue'
    import type { Hero } from "../model/types.ts"
    import { global } from "../model/utils.ts"
    
    const heroId = ref(1)
    //const jsonres = ref<Hero>({})
    const azert = ref(false)

    // async function fetchData() {
    //     azert.value = false
    //     jsonres.value = {}
    //     const res = await fetch(
    //         `https://tama.rhiobet.sh/hero/${heroId.value}`
    //     )
    //     azert.value = true
    //     jsonres.value = await res.json() as Hero
    // }

    // fetchData()
</script>

<template>
    <!-- <p v-if="JSON.stringify(jsonres)==='{}'">Chargement...</p>
    <div v-else>
        <p> {{ JSON.stringify(jsonres) }} </p>
        <p> {{ jsonres.id }} </p>
        <p> {{ jsonres.imageUrl }} </p>

        <img :src="'https://tama.rhiobet.sh'+jsonres.imageUrl" width="150"/>
    </div> -->

</template>