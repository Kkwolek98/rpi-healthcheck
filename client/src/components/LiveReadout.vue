<script setup lang="ts">
import { useLiveReadoutsStore } from '@/store/liveReadouts';
import { onBeforeMount } from 'vue';

const readouts = useLiveReadoutsStore()

onBeforeMount(() => {
  readouts.startListening()
})
</script>

<template>
  <div class="rounded-xl w-32 h-24 p-4 bg-red-500 flex flex-col gap-4 text-white">
    <h2 class=" text-gray-300">
      GPU:
    </h2>
    <div class="flex justify-end text-xl">
      {{ readouts.temperature }}°C
    </div>
  </div>
</template>