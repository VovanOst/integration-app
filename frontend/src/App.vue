<template>
  <div class="app">
    <nav class="navbar">
      <div class="nav-brand">Integration App</div>
      <div class="nav-links">
        <router-link to="/">Dashboard</router-link>
        <router-link to="/connections">Connections</router-link>
        <router-link to="/mappings">Field Mapping</router-link>
        <router-link to="/webhooks">Webhooks</router-link>
        <router-link to="/logs">Logs</router-link>
        <router-link to="/settings">Settings</router-link>
      </div>
    </nav>

    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<style scoped>
.app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.navbar {
  background: #0066cc;
  color: white;
  padding: 16px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.nav-brand {
  font-size: 20px;
  font-weight: bold;
}

.nav-links {
  display: flex;
  gap: 20px;
}

.nav-links a {
  color: white;
  text-decoration: none;
  padding: 8px 12px;
  border-radius: 4px;
  transition: background 0.2s;
}

.nav-links a:hover {
  background: rgba(255,255,255,0.2);
}

.nav-links a.router-link-active {
  background: rgba(255,255,255,0.3);
  font-weight: bold;
}

.main-content {
  flex: 1;
  padding: 20px;
}

@media (max-width: 768px) {
  .navbar {
    flex-direction: column;
    gap: 10px;
  }

  .nav-links {
    flex-direction: column;
    gap: 10px;
    width: 100%;
  }

  .nav-links a {
    display: block;
    text-align: center;
  }
}
</style>
