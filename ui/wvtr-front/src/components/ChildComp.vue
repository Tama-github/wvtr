<script setup lang="ts">
// déclare les événements émis
const emit = defineEmits(['response'])

// emit avec un argument
emit('response', "hello à partir de l'enfant")
</script>
<template>
    <h1>Enfant</h1>
</template>