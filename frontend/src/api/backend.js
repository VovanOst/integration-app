const API_URL = 'http://localhost:8080/api'

export const api = {
  // Connections
  getConnections: () => fetch(`${API_URL}/connections`).then(r => r.json()),
  createConnection: (data) => fetch(`${API_URL}/connections`, { method: 'POST', body: JSON.stringify(data) }).then(r => r.json()),
  
  // Mappings
  getMappings: () => fetch(`${API_URL}/mappings`).then(r => r.json()),
  saveMappings: (data) => fetch(`${API_URL}/mappings`, { method: 'POST', body: JSON.stringify(data) }).then(r => r.json()),
  
  // Webhooks
  getWebhooks: () => fetch(`${API_URL}/webhooks`).then(r => r.json()),
  createWebhook: (data) => fetch(`${API_URL}/webhooks`, { method: 'POST', body: JSON.stringify(data) }).then(r => r.json()),
  
  // Logs
  getLogs: (filters = {}) => fetch(`${API_URL}/sync/logs`).then(r => r.json()),
}
