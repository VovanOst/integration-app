import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import Connections from '../pages/Connections.vue'
import FieldMapping from '../pages/FieldMapping.vue'
import Webhooks from '../pages/Webhooks.vue'
import Logs from '../pages/Logs.vue'
import Settings from '../pages/Settings.vue'

const routes = [
  { path: '/', component: Dashboard },
  { path: '/connections', component: Connections },
  { path: '/mappings', component: FieldMapping },
  { path: '/webhooks', component: Webhooks },
  { path: '/logs', component: Logs },
  { path: '/settings', component: Settings },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
