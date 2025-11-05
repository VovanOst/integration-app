<template>
  <div class="connections">
    <h1>Управление подключениями</h1>
    <button @click="showAddConnection" class="btn btn-primary">+ Добавить подключение</button>
    
    <div class="connections-list">
      <div v-for="connection in connections" :key="connection.id" class="connection-card">
        <h3>{{ connection.name }}</h3>
        <p>Система: {{ connection.system_type }}</p>
        <button @click="deleteConnection(connection.id)" class="btn btn-danger">Удалить</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const connections = ref([])

const showAddConnection = () => {
  alert('Откройте форму добавления подключения')
}

const deleteConnection = (id) => {
  connections.value = connections.value.filter(c => c.id !== id)
}
</script>

<style scoped>
.connections {
  padding: 20px;
}

.connections-list {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
  margin-top: 20px;
}

.connection-card {
  background: white;
  border: 1px solid #ddd;
  padding: 20px;
  border-radius: 8px;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 10px;
}

.btn-primary {
  background: #0066cc;
  color: white;
}

.btn-danger {
  background: #ff6b6b;
  color: white;
}
</style>
