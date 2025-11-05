<template>
  <div class="settings">
    <h1>Настройки</h1>
    <div class="settings-group">
      <h2>Параметры приложения</h2>
      <div class="setting-item">
        <label>Интервал синхронизации (минуты):</label>
        <input v-model="syncInterval" type="number" />
      </div>
      <div class="setting-item">
        <label>Шифрование токенов:</label>
        <input v-model="encryptionEnabled" type="checkbox" />
      </div>
      <button @click="saveSettings" class="btn btn-primary">Сохранить</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const syncInterval = ref(15)
const encryptionEnabled = ref(true)

const saveSettings = () => {
  alert('Настройки сохранены')
}
</script>

<style scoped>
.settings {
  padding: 20px;
}

.settings-group {
  background: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 600px;
}

.setting-item {
  margin-bottom: 20px;
}

.setting-item label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
}

.setting-item input {
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  width: 100%;
}

.btn {
  padding: 10px 20px;
  background: #0066cc;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>
