import { defineStore } from 'pinia'

export const useLiveReadoutsStore = defineStore('liveReadouts', {
  state: () => ({
    readout: 0
  }),
  actions: {
    startListening() {
      const ws = new WebSocket('ws://localhost:3000/api/temperature-readouts/live') // TODO: change
      ws.onmessage = (event) => {
        const temp = parseFloat(event.data)
        this.readout = temp
        console.log({ temp })
      }
    }
  }
})
