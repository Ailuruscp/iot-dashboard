<template>
  <div class="device-card" @click="$emit('click')">
    <div class="device-header">
      <h3>{{ device.name }}</h3>
      <span :class="['status-badge', device.online ? 'success' : 'danger']">
        {{ device.online ? 'Online' : 'Offline' }}
      </span>
    </div>
    
    <div class="device-info">
      <p class="device-id">ID: {{ device.id }}</p>
      <p class="device-type">Type: {{ device.type }}</p>
      <p class="device-last-seen" v-if="!device.online">
        Last seen: {{ formatLastSeen(device.last_seen) }}
      </p>
    </div>
    
    <div class="device-data" v-if="device.data">
      <div v-for="(value, key) in device.data" :key="key" class="data-item">
        <span class="data-label">{{ formatLabel(key) }}:</span>
        <span class="data-value">{{ formatValue(value) }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DeviceCard',
  props: {
    device: {
      type: Object,
      required: true
    }
  },
  methods: {
    formatLastSeen(timestamp) {
      if (!timestamp) return 'Never'
      const date = new Date(timestamp)
      return date.toLocaleString()
    },
    formatLabel(key) {
      return key.split('_').map(word => 
        word.charAt(0).toUpperCase() + word.slice(1)
      ).join(' ')
    },
    formatValue(value) {
      if (typeof value === 'number') {
        return value.toFixed(2)
      }
      return value
    }
  }
}
</script>

<style lang="scss" scoped>
.device-card {
  background-color: white;
  border-radius: 0.5rem;
  padding: 1.5rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  
  .device-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
    
    h3 {
      font-size: 1.125rem;
      font-weight: 600;
      color: #111827;
    }
  }
  
  .device-info {
    margin-bottom: 1rem;
    
    p {
      margin: 0.25rem 0;
      color: #6b7280;
      font-size: 0.875rem;
    }
  }
  
  .device-data {
    border-top: 1px solid #e5e7eb;
    padding-top: 1rem;
    
    .data-item {
      display: flex;
      justify-content: space-between;
      margin-bottom: 0.5rem;
      
      .data-label {
        color: #6b7280;
        font-size: 0.875rem;
      }
      
      .data-value {
        font-weight: 500;
        color: #111827;
      }
    }
  }
}
</style> 