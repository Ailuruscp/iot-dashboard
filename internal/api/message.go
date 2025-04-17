package api

import (
	"iot-dashboard/internal/device"
)

// Message types
const (
	MessageTypeDeviceData     = "device_data"     // Data from device sensors
	MessageTypeCommand        = "command"         // Command to a device
	MessageTypeDeviceList     = "device_list"     // List of all devices
	MessageTypeDeviceRegister = "device_register" // Device registration
)

// Message represents a message exchanged between clients and the server
type Message struct {
	// Message type
	Type string `json:"type"`

	// Device identification
	DeviceID   string `json:"device_id,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	DeviceType string `json:"device_type,omitempty"`

	// For device data messages
	Data map[string]interface{} `json:"data,omitempty"`

	// For command messages
	Command string                 `json:"command,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty"`

	// For device list messages
	Devices []*device.Device `json:"devices,omitempty"`

	// Timestamp (set by server)
	Timestamp int64 `json:"timestamp,omitempty"`

	// Error information (if applicable)
	Error string `json:"error,omitempty"`
}
