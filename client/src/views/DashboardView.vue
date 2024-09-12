<script lang="ts">
import { useLiveReadoutsStore } from '@/store/liveReadouts';

const readoutsStore = useLiveReadoutsStore()

readoutsStore.startListening()
</script>

<template>
  Dashboard
</template>