<template>
  <div class="field-mapping-container">
    <h3>Сопоставление полей Bitrix24 и Facebook</h3>

    <div class="mapping-grid">
      <!-- Левая колонка: Bitrix24 поля (draggable) -->
      <div class="column bitrix-column">
        <h4>Bitrix24 поля</h4>
        <div class="fields-list">
          <div
            v-for="(field, index) in bitrixFields"
            :key="`bitrix-${index}`"
            class="field-item draggable"
            :class="{ 'is-mapped': isBitrixFieldMapped(field) }"
            draggable="true"
            @dragstart="dragStart($event, field, 'bitrix')"
            @dragend="dragEnd"
          >
            {{ field }}
            <span v-if="isBitrixFieldMapped(field)" class="mapped-badge">✓</span>
          </div>
        </div>
      </div>

      <!-- Правая колонка: Facebook поля (drop zone) -->
      <div class="column facebook-column">
        <h4>Facebook поля</h4>
        <div class="fields-list">
          <div
            v-for="(field, index) in facebookFields"
            :key="`facebook-${index}`"
            class="field-item droppable"
            :class="{ 'drag-over': dragOverField === field }"
            @dragover.prevent="dragOver($event, field)"
            @dragleave="dragLeave"
            @drop="drop($event, field)"
          >
            <div class="field-name">{{ field }}</div>
            <div v-if="getMappedBitrixField(field)" class="mapped-from">
              ← {{ getMappedBitrixField(field) }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Отображение текущих сопоставлений -->
    <div class="mappings-summary">
      <h4>Текущие сопоставления:</h4>
      <div v-if="Object.keys(mappings).length > 0" class="mappings-list">
        <div v-for="(facebookField, bitrixField) in mappings" :key="bitrixField" class="mapping-item">
          <span class="bitrix-field">{{ bitrixField }}</span>
          <span class="arrow">→</span>
          <span class="facebook-field">{{ facebookField }}</span>
          <button @click="removeMapping(bitrixField)" class="remove-btn">✕</button>
        </div>
      </div>
      <div v-else class="no-mappings">Сопоставления отсутствуют</div>
    </div>

    <!-- Кнопки действий -->
    <div class="actions">
      <button @click="saveMappings" class="btn btn-primary">Сохранить</button>
      <button @click="clearMappings" class="btn btn-secondary">Очистить</button>
      <button @click="exportMappings" class="btn btn-info">Экспортировать JSON</button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const bitrixFields = ref([
  'Lead.NAME',
  'Lead.PHONE',
  'Lead.EMAIL',
  'Contact.FULL_NAME',
  'Contact.PHONE_MOBILE',
  'Contact.EMAIL',
  'Deal.TITLE',
  'Deal.AMOUNT',
  'Deal.STAGE_ID',
])

const facebookFields = ref([
  'full_name',
  'phone_number',
  'email_address',
  'first_name',
  'last_name',
  'mobile_phone',
  'business_email',
  'campaign_name',
  'purchase_amount',
])

const mappings = ref({})
const draggedField = ref(null)
const draggedSource = ref(null)
const dragOverField = ref(null)

// Проверить, сопоставлено ли поле Bitrix24
const isBitrixFieldMapped = (field) => {
  return Object.keys(mappings.value).includes(field)
}

// Получить Bitrix24 поле, которое сопоставлено с Facebook полем
const getMappedBitrixField = (facebookField) => {
  for (const [bitrixField, fbField] of Object.entries(mappings.value)) {
    if (fbField === facebookField) {
      return bitrixField
    }
  }
  return null
}

// Начало перетаскивания
const dragStart = (event, field, source) => {
  draggedField.value = field
  draggedSource.value = source
  event.dataTransfer.effectAllowed = 'move'
  event.dataTransfer.setData('text/plain', field)
}

// Конец перетаскивания
const dragEnd = () => {
  draggedField.value = null
  draggedSource.value = null
  dragOverField.value = null
}

// Перемещение над drop zone
const dragOver = (event, field) => {
  event.preventDefault()
  event.dataTransfer.dropEffect = 'move'
  dragOverField.value = field
}

// Покидание drop zone
const dragLeave = () => {
  dragOverField.value = null
}

// Drop действие
const drop = (event, facebookField) => {
  event.preventDefault()
  dragOverField.value = null

  if (draggedSource.value === 'bitrix' && draggedField.value) {
    // Если это поле уже сопоставлено с другим Facebook полем, удалить старое
    if (mappings.value[draggedField.value]) {
      delete mappings.value[draggedField.value]
    }

    // Добавить новое сопоставление
    mappings.value[draggedField.value] = facebookField
  }

  draggedField.value = null
  draggedSource.value = null
}

// Удалить сопоставление
const removeMapping = (bitrixField) => {
  delete mappings.value[bitrixField]
}

// Очистить все сопоставления
const clearMappings = () => {
  mappings.value = {}
}

// Сохранить сопоставления (например, отправить на сервер)
const saveMappings = () => {
  console.log('Сохраненные сопоставления:', mappings.value)
  alert('Сопоставления сохранены! Проверьте консоль браузера.')
}

// Экспортировать в JSON
const exportMappings = () => {
  const dataStr = JSON.stringify(mappings.value, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(dataBlob)
  const link = document.createElement('a')
  link.href = url
  link.download = 'field-mappings.json'
  link.click()
  URL.revokeObjectURL(url)
}
</script>

<style scoped>
.field-mapping-container {
  padding: 20px;
  font-family: Arial, sans-serif;
  max-width: 1200px;
  margin: 0 auto;
}

h3 {
  text-align: center;
  color: #333;
  margin-bottom: 30px;
}

h4 {
  margin-bottom: 15px;
  color: #555;
  font-size: 16px;
}

.mapping-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30px;
  margin-bottom: 30px;
}

.column {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
  background-color: #f9f9f9;
}

.bitrix-column {
  background-color: #f0f5ff;
  border-color: #0066cc;
}

.facebook-column {
  background-color: #fff0f5;
  border-color: #cc0066;
}

.fields-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.field-item {
  padding: 12px;
  background-color: white;
  border: 2px solid #ddd;
  border-radius: 4px;
  cursor: move;
  transition: all 0.2s ease;
  position: relative;
}

.draggable {
  background-color: #e3f2fd;
  border-color: #0066cc;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.draggable:hover {
  background-color: #bbdefb;
  box-shadow: 0 2px 8px rgba(0, 102, 204, 0.2);
  transform: translateY(-2px);
}

.draggable.is-mapped {
  background-color: #c8e6c9;
  border-color: #4caf50;
}

.mapped-badge {
  color: #4caf50;
  font-weight: bold;
  font-size: 14px;
}

.droppable {
  background-color: #fce4ec;
  border-color: #cc0066;
  /* КЛЮЧЕВОЕ ИСПРАВЛЕНИЕ: Точно такая же высота как у левых элементов */
  min-height: 50px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  /* Добавляем box-sizing для точного расчета размеров */
  box-sizing: border-box;
}

.droppable.drag-over {
  background-color: #f8bbd0;
  border-color: #880e4f;
  border-width: 2px;
  box-shadow: 0 0 8px rgba(204, 0, 102, 0.3);
}

.field-name {
  font-weight: 500;
  color: #333;
}

.mapped-from {
  font-size: 12px;
  color: #666;
  margin-top: 5px;
  font-style: italic;
}

.mappings-summary {
  background-color: #f5f5f5;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
}

.mappings-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.mapping-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px;
  background-color: white;
  border-radius: 4px;
  border-left: 4px solid #0066cc;
}

.bitrix-field {
  font-weight: bold;
  color: #0066cc;
  flex: 1;
}

.arrow {
  color: #999;
}

.facebook-field {
  font-weight: bold;
  color: #cc0066;
  flex: 1;
}

.remove-btn {
  background-color: #ff6b6b;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 5px 10px;
  cursor: pointer;
  font-size: 12px;
  transition: background-color 0.2s;
}

.remove-btn:hover {
  background-color: #ff5252;
}

.no-mappings {
  color: #999;
  font-style: italic;
  text-align: center;
  padding: 20px;
}

.actions {
  display: flex;
  gap: 10px;
  justify-content: center;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.2s;
}

.btn-primary {
  background-color: #0066cc;
  color: white;
}

.btn-primary:hover {
  background-color: #0052a3;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 102, 204, 0.3);
}

.btn-secondary {
  background-color: #999;
  color: white;
}

.btn-secondary:hover {
  background-color: #777;
}

.btn-info {
  background-color: #17a2b8;
  color: white;
}

.btn-info:hover {
  background-color: #138496;
}
</style>