<script setup lang="ts">
import LayoutDrawer from "./components/LayoutDrawer.vue"
import { RouterView } from 'vue-router'
</script>

<template>
<main>
  <LayoutDrawer>
    <RouterView />
  </LayoutDrawer>
</main>
</template>

<style scoped>
</style>
