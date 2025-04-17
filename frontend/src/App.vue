<template>
  <div class="app">
    <header class="header">
      <div class="container">
        <h1>IoT Dashboard</h1>
        <p>Real-time device monitoring and control</p>
      </div>
    </header>
    
    <main class="main">
      <div class="container">
        <router-view />
      </div>
    </main>
    
    <footer class="footer">
      <div class="container">
        <p>&copy; 2024 IoT Dashboard. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">
.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: #4f46e5;
  color: white;
  padding: 2rem 0;
  text-align: center;
  
  h1 {
    font-size: 2rem;
    font-weight: 700;
    margin-bottom: 0.5rem;
  }
  
  p {
    opacity: 0.8;
  }
}

.main {
  flex: 1;
  padding: 2rem 0;
}

.footer {
  background-color: #1f2937;
  color: white;
  padding: 1.5rem 0;
  text-align: center;
  font-size: 0.875rem;
  opacity: 0.7;
}
</style> 