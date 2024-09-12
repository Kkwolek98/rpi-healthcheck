import { defineStore } from 'pinia'

let ws: WebSocket

export const useLiveReadoutsStore = defineStore('liveReadouts', {
  state: () => ({
    temperature: 0
  }),
  actions: {
    startListening() {
      ws = new WebSocket('ws://localhost:3000/api/temperature-readouts/live') // TODO: change
      ws.onmessage = (event) => {
        const temp = parseFloat(event.data)
        this.temperature = temp
      }
    }
  }
})
