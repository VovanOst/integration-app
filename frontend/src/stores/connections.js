import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useConnectionsStore = defineStore('connections', () => {
  const connections = ref([])

  const addConnection = (connection) => {
    connections.value.push(connection)
  }

  const removeConnection = (id) => {
    connections.value = connections.value.filter(c => c.id !== id)
  }

  return { connections, addConnection, removeConnection }
})
