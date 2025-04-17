<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <h2>Device Dashboard</h2>
      <div class="connection-status">
        <span :class="['status-indicator', { 'connected': wsConnected }]"></span>
        {{ wsConnected ? 'Connected' : 'Disconnected' }}
      </div>
    </div>
    
    <div v-if="error" class="error-message">
      {{ error }}
    </div>
    
    <div v-if="loading" class="loading">
      Loading devices...
    </div>
    
    <div v-else>
      <div class="stats-cards">
        <div class="card stat-card">
          <h3>Total Devices</h3>
          <p class="stat-value">{{ devices.length }}</p>
        </div>
        <div class="card stat-card">
          <h3>Online Devices</h3>
          <p class="stat-value">{{ onlineDevices.length }}</p>
        </div>
        <div class="card stat-card">
          <h3>Offline Devices</h3>
          <p class="stat-value">{{ offlineDevices.length }}</p>
        </div>
      </div>
      
      <div class="device-actions">
        <button class="btn btn-primary" @click="showAddDeviceModal = true">
          Add Device
        </button>
      </div>
      
      <div v-if="devices.length === 0" class="no-devices">
        <p>No devices registered. Add a device to get started.</p>
      </div>
      
      <div v-else class="grid grid-cols-3">
        <device-card
          v-for="device in devices"
          :key="device.id"
          :device="device"
          @click="viewDeviceDetails(device.id)"
        />
      </div>
    </div>
    
    <!-- Add Device Modal -->
    <div v-if="showAddDeviceModal" class="modal">
      <div class="modal-content">
        <h3>Add New Device</h3>
        <form @submit.prevent="registerNewDevice">
          <div class="form-group">
            <label for="deviceId">Device ID</label>
            <input
              id="deviceId"
              v-model="newDevice.id"
              type="text"
              required
              placeholder="e.g., device-001"
            />
          </div>
          <div class="form-group">
            <label for="deviceName">Device Name</label>
            <input
              id="deviceName"
              v-model="newDevice.name"
              type="text"
              required
              placeholder="e.g., Living Room Sensor"
            />
          </div>
          <div class="form-group">
            <label for="deviceType">Device Type</label>
            <select id="deviceType" v-model="newDevice.type" required>
              <option value="">Select a type</option>
              <option value="sensor">Sensor</option>
              <option value="controller">Controller</option>
              <option value="camera">Camera</option>
              <option value="other">Other</option>
            </select>
          </div>
          <div class="form-actions">
            <button type="button" class="btn" @click="showAddDeviceModal = false">
              Cancel
            </button>
            <button type="submit" class="btn btn-primary">
              Add Device
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from 'vuex'
import DeviceCard from '@/components/DeviceCard.vue'

export default {
  name: 'Dashboard',
  components: {
    DeviceCard
  },
  data() {
    return {
      showAddDeviceModal: false,
      newDevice: {
        id: '',
        name: '',
        type: ''
      }
    }
  },
  computed: {
    ...mapState(['devices', 'loading', 'error', 'wsConnected']),
    ...mapGetters(['onlineDevices', 'offlineDevices'])
  },
  created() {
    this.fetchDevices()
    this.connectWebSocket()
  },
  methods: {
    ...mapActions(['fetchDevices', 'registerDevice', 'connectWebSocket']),
    viewDeviceDetails(id) {
      this.$router.push(`/device/${id}`)
    },
    async registerNewDevice() {
      await this.registerDevice(this.newDevice)
      this.showAddDeviceModal = false
      this.newDevice = {
        id: '',
        name: '',
        type: ''
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  .dashboard-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    
    h2 {
      font-size: 1.5rem;
      font-weight: 600;
    }
  }
  
  .connection-status {
    display: flex;
    align-items: center;
    font-size: 0.875rem;
    
    .status-indicator {
      width: 10px;
      height: 10px;
      border-radius: 50%;
      background-color: #ef4444;
      margin-right: 0.5rem;
      
      &.connected {
        background-color: #10b981;
      }
    }
  }
  
  .error-message {
    background-color: rgba(239, 68, 68, 0.1);
    color: #ef4444;
    padding: 1rem;
    border-radius: 0.375rem;
    margin-bottom: 1.5rem;
  }
  
  .loading {
    text-align: center;
    padding: 2rem;
    color: #6b7280;
  }
  
  .stats-cards {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1.5rem;
    margin-bottom: 2rem;
    
    @media (max-width: 768px) {
      grid-template-columns: 1fr;
    }
    
    .stat-card {
      text-align: center;
      
      h3 {
        font-size: 1rem;
        color: #6b7280;
        margin-bottom: 0.5rem;
      }
      
      .stat-value {
        font-size: 2rem;
        font-weight: 700;
        color: #4f46e5;
      }
    }
  }
  
  .device-actions {
    margin-bottom: 1.5rem;
  }
  
  .no-devices {
    text-align: center;
    padding: 3rem;
    background-color: white;
    border-radius: 0.5rem;
    color: #6b7280;
  }
  
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    
    .modal-content {
      background-color: white;
      border-radius: 0.5rem;
      padding: 2rem;
      width: 100%;
      max-width: 500px;
      
      h3 {
        margin-bottom: 1.5rem;
        font-size: 1.25rem;
        font-weight: 600;
      }
      
      .form-group {
        margin-bottom: 1rem;
        
        label {
          display: block;
          margin-bottom: 0.5rem;
          font-weight: 500;
        }
        
        input, select {
          width: 100%;
          padding: 0.5rem;
          border: 1px solid #e5e7eb;
          border-radius: 0.375rem;
          
          &:focus {
            outline: none;
            border-color: #4f46e5;
          }
        }
      }
      
      .form-actions {
        display: flex;
        justify-content: flex-end;
        gap: 1rem;
        margin-top: 1.5rem;
      }
    }
  }
}
</style> 