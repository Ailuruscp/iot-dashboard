package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"iot-dashboard/internal/device"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Handler handles HTTP requests
type Handler struct {
	hub          *Hub
	deviceManager *device.DeviceManager
	logger        *log.Logger
	upgrader      websocket.Upgrader
}

// NewHandler creates a new API handler
func NewHandler(hub *Hub, deviceManager *device.DeviceManager, logger *log.Logger) *Handler {
	return &Handler{
		hub:          hub,
		deviceManager: deviceManager,
		logger:        logger,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // In production, implement proper origin checking
			},
		},
	}
}

// RegisterRoutes registers all API routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// API routes
	router.HandleFunc("/api/devices", h.GetDevices).Methods("GET")
	router.HandleFunc("/api/devices/{id}", h.GetDevice).Methods("GET")
	router.HandleFunc("/api/devices", h.RegisterDevice).Methods("POST")
	router.HandleFunc("/api/devices/{id}", h.UnregisterDevice).Methods("DELETE")
	router.HandleFunc("/api/devices/{id}/data", h.UpdateDeviceData).Methods("POST")

	// WebSocket endpoint
	router.HandleFunc("/ws", h.ServeWS)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/dist")))
}

// ServeWS handles WebSocket connections
func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("device_id")
	if deviceID == "" {
		http.Error(w, "device_id is required", http.StatusBadRequest)
		return
	}

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Printf("Error upgrading connection: %v", err)
		return
	}

	client := NewClient(conn, h.hub, deviceID)
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

// GetDevices returns all devices
func (h *Handler) GetDevices(w http.ResponseWriter, r *http.Request) {
	devices := h.deviceManager.GetAllDevices()
	json.NewEncoder(w).Encode(devices)
}

// GetDevice returns a device by ID
func (h *Handler) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	device, exists := h.deviceManager.GetDevice(id)
	if !exists {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(device)
}

// RegisterDeviceRequest represents a request to register a device
type RegisterDeviceRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// RegisterDevice registers a new device
func (h *Handler) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	var req RegisterDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.ID == "" || req.Name == "" || req.Type == "" {
		http.Error(w, "ID, name, and type are required", http.StatusBadRequest)
		return
	}

	device := h.deviceManager.RegisterDevice(req.ID, req.Name, req.Type)
	json.NewEncoder(w).Encode(device)
}

// UnregisterDevice unregisters a device
func (h *Handler) UnregisterDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.deviceManager.UnregisterDevice(id)
	w.WriteHeader(http.StatusNoContent)
}

// UpdateDeviceDataRequest represents a request to update device data
type UpdateDeviceDataRequest struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Battery     float64 `json:"battery"`
}

// UpdateDeviceData updates device data
func (h *Handler) UpdateDeviceData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req UpdateDeviceDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := device.DeviceData{
		Temperature: req.Temperature,
		Humidity:    req.Humidity,
		Battery:     req.Battery,
		Timestamp:   time.Now(),
	}

	if err := h.deviceManager.UpdateDeviceData(id, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}