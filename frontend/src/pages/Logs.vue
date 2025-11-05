<template>
  <div class="logs">
    <h1>Логи синхронизации</h1>
    
    <table class="logs-table">
      <thead>
        <tr>
          <th>Дата</th>
          <th>Тип события</th>
          <th>Статус</th>
          <th>Источник</th>
          <th>Цель</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="log in logs" :key="log.id">
          <td>{{ log.created_at }}</td>
          <td>{{ log.event_type }}</td>
          <td :class="log.status">{{ log.status }}</td>
          <td>{{ log.source }}</td>
          <td>{{ log.target }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const logs = ref([])
</script>

<style scoped>
.logs {
  padding: 20px;
}

.logs-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.logs-table th,
.logs-table td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: left;
}

.logs-table th {
  background: #f5f5f5;
}

.success {
  color: #4caf50;
  font-weight: bold;
}

.error {
  color: #ff6b6b;
  font-weight: bold;
}
</style>
