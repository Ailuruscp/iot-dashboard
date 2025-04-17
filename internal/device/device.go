package device

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// DeviceStatus represents the connection status of a device
type DeviceStatus string

// Device status constants
const (
	StatusOnline  DeviceStatus = "online"
	StatusOffline DeviceStatus = "offline"
	StatusError   DeviceStatus = "error"
)

// DeviceData represents sensor data from an IoT device
type DeviceData struct {
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Battery     float64   `json:"battery"`
	Timestamp   time.Time `json:"timestamp"`
}

// Device represents an IoT device in the system
type Device struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Status     DeviceStatus           `json:"status"`
	Data       map[string]interface{} `json:"data"`
	LastSeen   time.Time              `json:"last_seen"`
	CreatedAt  time.Time              `json:"created_at"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	LastData   DeviceData             `json:"last_data"`
	Connected  bool                   `json:"connected"`
	mu         sync.RWMutex
}

// SensorReading represents a single sensor reading
type SensorReading struct {
	DeviceID  string      `json:"device_id"`
	SensorID  string      `json:"sensor_id"`
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
	Unit      string      `json:"unit,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// DeviceManager manages all IoT devices
type DeviceManager struct {
	devices map[string]*Device
	mu      sync.RWMutex
	logger  *log.Logger
}

// NewDeviceManager creates a new device manager
func NewDeviceManager(logger *log.Logger) *DeviceManager {
	return &DeviceManager{
		devices: make(map[string]*Device),
		logger:  logger,
	}
}

// RegisterDevice registers a new device
func (dm *DeviceManager) RegisterDevice(id, name, deviceType string) *Device {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	now := time.Now()
	device := &Device{
		ID:         id,
		Name:       name,
		Type:       deviceType,
		Status:     StatusOnline,
		Data:       make(map[string]interface{}),
		LastSeen:   now,
		CreatedAt:  now,
		Connected:  true,
		Attributes: map[string]interface{}{
			"firmware_version": "1.0.0",
		},
	}

	dm.devices[id] = device
	dm.logger.Printf("Device registered: %s", device.String())
	return device
}

// UnregisterDevice unregisters a device
func (dm *DeviceManager) UnregisterDevice(id string) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if device, exists := dm.devices[id]; exists {
		device.Connected = false
		device.Status = StatusOffline
		dm.logger.Printf("Device unregistered: %s", device.String())
		delete(dm.devices, id)
	}
}

// UpdateDeviceData updates the device data
func (dm *DeviceManager) UpdateDeviceData(id string, data DeviceData) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	device, exists := dm.devices[id]
	if !exists {
		return fmt.Errorf("device not found: %s", id)
	}

	device.mu.Lock()
	device.LastData = data
	device.LastSeen = time.Now()
	device.Status = StatusOnline
	device.Connected = true
	device.mu.Unlock()

	dm.logger.Printf("Device data updated: %s", device.String())
	return nil
}

// GetDevice returns a device by ID
func (dm *DeviceManager) GetDevice(id string) (*Device, bool) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	device, exists := dm.devices[id]
	return device, exists
}

// GetAllDevices returns all registered devices
func (dm *DeviceManager) GetAllDevices() []*Device {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	devices := make([]*Device, 0, len(dm.devices))
	for _, device := range dm.devices {
		devices = append(devices, device)
	}
	return devices
}

// BroadcastUpdate returns a JSON representation of all devices
func (dm *DeviceManager) BroadcastUpdate() ([]byte, error) {
	devices := dm.GetAllDevices()
	return json.Marshal(devices)
}

// String returns a string representation of the device
func (d *Device) String() string {
	return fmt.Sprintf("%s (%s) - %s", d.Name, d.ID, d.Status)
}

// NewDevice creates a new device
func NewDevice(id, name, deviceType string) *Device {
	now := time.Now()
	return &Device{
		ID:         id,
		Name:       name,
		Type:       deviceType,
		Status:     StatusOffline,
		Data:       make(map[string]interface{}),
		LastSeen:   now,
		CreatedAt:  now,
		Connected:  false,
		Attributes: make(map[string]interface{}),
	}
}

// SetStatus sets the device status
func (d *Device) SetStatus(status DeviceStatus) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Status = status
}

// UpdateData updates the device data
func (d *Device) UpdateData(data map[string]interface{}) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Data = data
	d.LastSeen = time.Now()
}
