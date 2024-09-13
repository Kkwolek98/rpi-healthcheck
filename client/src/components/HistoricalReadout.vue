<script setup lang="ts">
import { useHistoricalReadoutsStore } from '@/store/historicalReadouts';
import { computed, onBeforeMount } from 'vue';
import VueApexCharts from 'vue3-apexcharts';

const readouts = useHistoricalReadoutsStore();

onBeforeMount(() => {
  readouts.getLastWeek();
});

const chartData = computed(() => {
  const timestamps = readouts.lastWeek.map((el) => {
    const date = new Date(el.ts);
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const year = date.getFullYear();
    
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');

    return `${day}.${month}.${year} ${hours}:${minutes}:${seconds}`;
  });
  const temperatures = readouts.lastWeek.map((el) => el.temp);
  return {
    options: {
      chart: {
        id: 'temperature-last-week'
      },
      xaxis: {
        categories: timestamps
      },
    },
    series: [{
      name: 'Temperature (°C)',
      data: temperatures
    }]
  };
})
</script>

<template>
    <VueApexCharts type="bar" :options="chartData.options" :series="chartData.series" height="400px"/>
</template>