import { createStore } from 'vuex'
import api from '@/api/axios'

export default createStore({
  state: {
    devices: [],
    selectedDevice: null,
    loading: false,
    error: null,
    wsConnected: false
  },
  getters: {
    getDeviceById: (state) => (id) => {
      return state.devices.find(device => device.id === id)
    },
    onlineDevices: (state) => {
      return state.devices.filter(device => device.connected)
    },
    offlineDevices: (state) => {
      return state.devices.filter(device => !device.connected)
    }
  },
  mutations: {
    SET_DEVICES(state, devices) {
      state.devices = devices
    },
    SET_SELECTED_DEVICE(state, device) {
      state.selectedDevice = device
    },
    SET_LOADING(state, loading) {
      state.loading = loading
    },
    SET_ERROR(state, error) {
      state.error = error
    },
    SET_WS_CONNECTED(state, connected) {
      state.wsConnected = connected
    },
    UPDATE_DEVICE(state, updatedDevice) {
      const index = state.devices.findIndex(device => device.id === updatedDevice.id)
      if (index !== -1) {
        state.devices.splice(index, 1, updatedDevice)
      }
    }
  },
  actions: {
    async fetchDevices({ commit }) {
      commit('SET_LOADING', true)
      try {
        const response = await api.get('/api/devices')
        commit('SET_DEVICES', response.data)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', 'Failed to fetch devices')
        console.error('Error fetching devices:', error)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async fetchDeviceById({ commit }, id) {
      commit('SET_LOADING', true)
      try {
        const response = await api.get(`/api/devices/${id}`)
        commit('SET_SELECTED_DEVICE', response.data)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', 'Failed to fetch device details')
        console.error('Error fetching device details:', error)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async registerDevice({ commit, dispatch }, deviceData) {
      commit('SET_LOADING', true)
      try {
        await api.post('/api/devices', deviceData)
        commit('SET_ERROR', null)
        dispatch('fetchDevices')
      } catch (error) {
        commit('SET_ERROR', 'Failed to register device')
        console.error('Error registering device:', error)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async unregisterDevice({ commit, dispatch }, id) {
      commit('SET_LOADING', true)
      try {
        await api.delete(`/api/devices/${id}`)
        commit('SET_ERROR', null)
        dispatch('fetchDevices')
      } catch (error) {
        commit('SET_ERROR', 'Failed to unregister device')
        console.error('Error unregistering device:', error)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async updateDeviceData({ commit }, { id, data }) {
      try {
        await api.post(`/api/devices/${id}/data`, data)
        commit('SET_ERROR', null)
      } catch (error) {
        commit('SET_ERROR', 'Failed to update device data')
        console.error('Error updating device data:', error)
      }
    },
    connectWebSocket({ commit, dispatch }) {
      const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
      const ws = new WebSocket(`${protocol}//localhost:8080/ws?device_id=dashboard`)
      
      ws.onopen = () => {
        console.log('WebSocket connected')
        commit('SET_WS_CONNECTED', true)
        commit('SET_ERROR', null)
      }
      
      ws.onmessage = (event) => {
        try {
          const message = JSON.parse(event.data)
          if (message.type === 'device_list') {
            commit('SET_DEVICES', message.devices)
          }
        } catch (error) {
          console.error('Error parsing WebSocket message:', error)
        }
      }
      
      ws.onclose = () => {
        console.log('WebSocket disconnected')
        commit('SET_WS_CONNECTED', false)
        // Attempt to reconnect after 5 seconds
        setTimeout(() => {
          dispatch('connectWebSocket')
        }, 5000)
      }
      
      ws.onerror = (error) => {
        console.error('WebSocket error:', error)
        commit('SET_ERROR', 'WebSocket connection error')
      }
    }
  }
}) 