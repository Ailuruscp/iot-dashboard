<template>
  <div class="device-detail">
    <h1>Device Details</h1>
    <div v-if="device" class="device-info">
      <h2>{{ device.name }}</h2>
      <p>ID: {{ device.id }}</p>
      <p>Type: {{ device.type }}</p>
      <p>Status: {{ device.online ? 'Online' : 'Offline' }}</p>
      <p>Last Seen: {{ formatDate(device.last_seen) }}</p>
    </div>
    <div v-else class="loading">
      Loading device information...
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import api from '@/api/axios'

export default {
  name: 'DeviceDetail',
  data() {
    return {
      device: null
    }
  },
  methods: {
    ...mapActions(['fetchDeviceById']),
    formatDate(date) {
      return new Date(date).toLocaleString()
    },
    async fetchDeviceDetails() {
      try {
        const response = await api.get(`/api/devices/${this.$route.params.id}`)
        this.device = response.data
      } catch (error) {
        console.error('Error fetching device details:', error)
      }
    }
  },
  created() {
    this.fetchDeviceDetails()
  }
}
</script>

<style lang="scss" scoped>
.device-detail {
  padding: 20px;
  
  .device-info {
    margin-top: 20px;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    
    h2 {
      margin-bottom: 15px;
      color: #2c3e50;
    }
    
    p {
      margin: 10px 0;
      color: #666;
    }
  }
  
  .loading {
    text-align: center;
    color: #666;
    margin-top: 20px;
  }
}
</style> 