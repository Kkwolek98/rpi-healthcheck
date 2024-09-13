<script setup lang="ts">
import HistoricalReadout from "@/components/HistoricalReadout.vue";
import LiveReadout from "@/components/LiveReadout.vue";
</script>

<template>
<div class="flex flex-col gap-4">
  <h1 class="text-xl">Dashboard</h1>

  <LiveReadout />
  <HistoricalReadout />
</div>
</template>