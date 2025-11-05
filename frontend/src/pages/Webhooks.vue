<template>
  <div class="webhooks">
    <h1>Управление вебхуками</h1>
    <button @click="showAddWebhook" class="btn btn-primary">+ Добавить вебхук</button>
    
    <div class="webhooks-list">
      <div v-for="webhook in webhooks" :key="webhook.id" class="webhook-item">
        <h3>{{ webhook.event_type }}</h3>
        <p>URL: {{ webhook.callback_url }}</p>
        <span :class="{ active: webhook.is_active }">{{ webhook.is_active ? 'Активен' : 'Неактивен' }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const webhooks = ref([])

const showAddWebhook = () => {
  alert('Откройте форму добавления вебхука')
}
</script>

<style scoped>
.webhooks {
  padding: 20px;
}

.webhooks-list {
  margin-top: 20px;
}

.webhook-item {
  background: white;
  border-left: 4px solid #0066cc;
  padding: 15px;
  margin-bottom: 10px;
  border-radius: 4px;
}

.active {
  color: #4caf50;
  font-weight: bold;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  background: #0066cc;
  color: white;
}
</style>
