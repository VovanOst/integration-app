import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMappingsStore = defineStore('mappings', () => {
  const mappings = ref({})

  const saveMappings = (newMappings) => {
    mappings.value = newMappings
  }

  return { mappings, saveMappings }
})
