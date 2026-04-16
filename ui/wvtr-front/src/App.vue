<script setup lang="ts">
    import { ref, onMounted } from 'vue'
    import { fetchData, RequestType } from "./model/utils.ts"
    import Header from "./components/Header.vue"
    import Body from "./components/Body.vue"
    import type { User } from "./model/types.ts"

    const user = ref<User|undefined>(undefined)

    onMounted(async () => {
        await fetchData<User>(user, RequestType.User, [{id: "id", value: "1"}]) 
    })
</script>

<template>
<div class="page">
  <Header :user="user"/>
  <Body  :user="user"/>
</div>
</template>

