<template>
  <div class="modal-overlay" v-if="show" @click="$emit('close')">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>{{ device ? 'Device Details' : 'Add New Device' }}</h2>
        <button class="close-button" @click="$emit('close')">&times;</button>
      </div>
      
      <div class="modal-body">
        <template v-if="device">
          <div class="device-details">
            <div class="detail-group">
              <label>Device ID</label>
              <p>{{ device.id }}</p>
            </div>
            
            <div class="detail-group">
              <label>Name</label>
              <p>{{ device.name }}</p>
            </div>
            
            <div class="detail-group">
              <label>Type</label>
              <p>{{ device.type }}</p>
            </div>
            
            <div class="detail-group">
              <label>Status</label>
              <span :class="['status-badge', device.online ? 'success' : 'danger']">
                {{ device.online ? 'Online' : 'Offline' }}
              </span>
            </div>
            
            <div class="detail-group" v-if="!device.online">
              <label>Last Seen</label>
              <p>{{ formatLastSeen(device.last_seen) }}</p>
            </div>
            
            <div class="detail-group" v-if="device.data">
              <label>Current Data</label>
              <div class="data-grid">
                <div v-for="(value, key) in device.data" :key="key" class="data-item">
                  <span class="data-label">{{ formatLabel(key) }}</span>
                  <span class="data-value">{{ formatValue(value) }}</span>
                </div>
              </div>
            </div>
          </div>
        </template>
        
        <template v-else>
          <form @submit.prevent="handleSubmit" class="add-device-form">
            <div class="form-group">
              <label for="deviceId">Device ID</label>
              <input
                id="deviceId"
                v-model="form.deviceId"
                type="text"
                required
                placeholder="Enter device ID"
              />
            </div>
            
            <div class="form-group">
              <label for="name">Name</label>
              <input
                id="name"
                v-model="form.name"
                type="text"
                required
                placeholder="Enter device name"
              />
            </div>
            
            <div class="form-group">
              <label for="type">Type</label>
              <select id="type" v-model="form.type" required>
                <option value="">Select device type</option>
                <option value="sensor">Sensor</option>
                <option value="actuator">Actuator</option>
                <option value="controller">Controller</option>
              </select>
            </div>
            
            <div class="form-actions">
              <button type="button" class="cancel-button" @click="$emit('close')">
                Cancel
              </button>
              <button type="submit" class="submit-button" :disabled="loading">
                {{ loading ? 'Adding...' : 'Add Device' }}
              </button>
            </div>
          </form>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DeviceModal',
  props: {
    show: {
      type: Boolean,
      required: true
    },
    device: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      form: {
        deviceId: '',
        name: '',
        type: ''
      },
      loading: false
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
    },
    async handleSubmit() {
      this.loading = true
      try {
        await this.$store.dispatch('registerDevice', this.form)
        this.$emit('close')
      } catch (error) {
        console.error('Failed to add device:', error)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  border-radius: 0.5rem;
  width: 90%;
  max-width: 32rem;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 1px solid #e5e7eb;
  
  h2 {
    font-size: 1.25rem;
    font-weight: 600;
    color: #111827;
  }
  
  .close-button {
    background: none;
    border: none;
    font-size: 1.5rem;
    color: #6b7280;
    cursor: pointer;
    
    &:hover {
      color: #111827;
    }
  }
}

.modal-body {
  padding: 1.5rem;
}

.device-details {
  .detail-group {
    margin-bottom: 1.5rem;
    
    label {
      display: block;
      font-size: 0.875rem;
      font-weight: 500;
      color: #6b7280;
      margin-bottom: 0.5rem;
    }
    
    p {
      color: #111827;
      font-size: 1rem;
    }
  }
  
  .data-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 1rem;
    margin-top: 0.5rem;
    
    .data-item {
      background-color: #f3f4f6;
      padding: 0.75rem;
      border-radius: 0.375rem;
      
      .data-label {
        display: block;
        font-size: 0.75rem;
        color: #6b7280;
        margin-bottom: 0.25rem;
      }
      
      .data-value {
        font-size: 1rem;
        font-weight: 500;
        color: #111827;
      }
    }
  }
}

.add-device-form {
  .form-group {
    margin-bottom: 1.5rem;
    
    label {
      display: block;
      font-size: 0.875rem;
      font-weight: 500;
      color: #374151;
      margin-bottom: 0.5rem;
    }
    
    input,
    select {
      width: 100%;
      padding: 0.5rem;
      border: 1px solid #d1d5db;
      border-radius: 0.375rem;
      font-size: 1rem;
      
      &:focus {
        outline: none;
        border-color: #3b82f6;
        box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
      }
    }
  }
  
  .form-actions {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
    margin-top: 2rem;
    
    button {
      padding: 0.5rem 1rem;
      border-radius: 0.375rem;
      font-size: 0.875rem;
      font-weight: 500;
      cursor: pointer;
      
      &.cancel-button {
        background-color: white;
        border: 1px solid #d1d5db;
        color: #374151;
        
        &:hover {
          background-color: #f9fafb;
        }
      }
      
      &.submit-button {
        background-color: #3b82f6;
        border: none;
        color: white;
        
        &:hover {
          background-color: #2563eb;
        }
        
        &:disabled {
          background-color: #93c5fd;
          cursor: not-allowed;
        }
      }
    }
  }
}
</style> 