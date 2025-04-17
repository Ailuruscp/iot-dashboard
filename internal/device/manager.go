package device

import (
	"errors"
	"log"
	"sync"
	"time"
)

// Manager handles device registration, status, and data
type Manager struct {
	devices map[string]*Device
	mu      sync.RWMutex
	logger  *log.Logger
}

// NewManager creates a new device manager
func NewManager(logger *log.Logger) *Manager {
	manager := &Manager{
		devices: make(map[string]*Device),
		logger:  logger,
	}
	
	// Start device status checker
	go manager.statusChecker()
	
	return manager
}

// RegisterDevice registers a new device or updates an existing one
func (m *Manager) RegisterDevice(id, name, deviceType string) *Device {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// Check if device already exists
	if device, exists := m.devices[id]; exists {
		device.Name = name
		device.Type = deviceType
		device.SetStatus(StatusOnline)
		m.logger.Printf("Updated existing device: %s", device)
		return device
	}
	
	// Create new device
	device := NewDevice(id, name, deviceType)
	device.SetStatus(StatusOnline)
	m.devices[id] = device
	
	m.logger.Printf("Registered new device: %s", device)
	return device
}

// GetDevice retrieves a device by ID
func (m *Manager) GetDevice(id string) *Device {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	return m.devices[id]
}

// UpdateDeviceData updates device data with new readings
func (m *Manager) UpdateDeviceData(id string, data map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	device, exists := m.devices[id]
	if !exists {
		return errors.New("device not found")
	}
	
	device.UpdateData(data)
	return nil
}

// GetAllDevices returns a slice of all registered devices
func (m *Manager) GetAllDevices() []*Device {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	devices := make([]*Device, 0, len(m.devices))
	for _, device := range m.devices {
		devices = append(devices, device)
	}
	
	return devices
}

// RemoveDevice removes a device from the manager
func (m *Manager) RemoveDevice(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if device, exists := m.devices[id]; exists {
		m.logger.Printf("Removing device: %s", device)
		delete(m.devices, id)
	}
}

// SetDeviceStatus sets the status of a device
func (m *Manager) SetDeviceStatus(id string, status DeviceStatus) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	device, exists := m.devices[id]
	if !exists {
		return errors.New("device not found")
	}
	
	device.SetStatus(status)
	return nil
}

// statusChecker periodically checks for offline devices
func (m *Manager) statusChecker() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		m.checkDeviceStatuses()
	}
}

// checkDeviceStatuses marks devices as offline if they haven't sent data recently
func (m *Manager) checkDeviceStatuses() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	now := time.Now()
	offlineThreshold := 2 * time.Minute
	
	for _, device := range m.devices {
		if device.Status == StatusOnline && now.Sub(device.LastSeen) > offlineThreshold {
			device.Status = StatusOffline
			m.logger.Printf("Device went offline: %s", device)
		}
	}
}