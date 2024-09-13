import { defineStore } from 'pinia'

export type TemperatureReadout = {
  temp: number
  ts: string
}

export const useHistoricalReadoutsStore = defineStore('historicalReadouts', {
  state: () => ({
    lastWeek: [] as TemperatureReadout[]
  }),
  actions: {
    getLastWeek() {
      fetch('http://localhost:3000/api/temperature-readouts/last-week')
        .then((res) => res.json())
        .then((res) => {
          this.lastWeek = res
        })
    }
  }
})
