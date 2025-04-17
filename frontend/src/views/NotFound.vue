<template>
  <div class="not-found">
    <h1>404</h1>
    <h2>Page Not Found</h2>
    <p>The page you are looking for does not exist.</p>
    <router-link to="/" class="home-link">Return to Home</router-link>
  </div>
</template>

<script>
export default {
  name: 'NotFound'
}
</script>

<style lang="scss" scoped>
.not-found {
  text-align: center;
  padding: 40px 20px;
  
  h1 {
    font-size: 6rem;
    color: #2c3e50;
    margin: 0;
  }
  
  h2 {
    font-size: 2rem;
    color: #34495e;
    margin: 20px 0;
  }
  
  p {
    color: #666;
    margin-bottom: 30px;
  }
  
  .home-link {
    display: inline-block;
    padding: 10px 20px;
    background-color: #42b983;
    color: white;
    text-decoration: none;
    border-radius: 4px;
    transition: background-color 0.3s;
    
    &:hover {
      background-color: #3aa876;
    }
  }
}
</style> 