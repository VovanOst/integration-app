<template>
  <div class="dashboard">
    <h1>Дашборд</h1>
    <div class="stats">
      <div class="stat-card">
        <h3>Активные подключения</h3>
        <p>{{ activeConnections }}</p>
      </div>
      <div class="stat-card">
        <h3>Успешных синхронизаций</h3>
        <p>{{ successfulSyncs }}</p>
      </div>
      <div class="stat-card">
        <h3>Ошибок</h3>
        <p>{{ failedSyncs }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const activeConnections = ref(0)
const successfulSyncs = ref(0)
const failedSyncs = ref(0)
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  margin-top: 20px;
}

.stat-card {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
</style>
