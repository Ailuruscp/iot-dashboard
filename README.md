# IoT Dashboard

A modern IoT management platform built with Go and Vue.js, featuring real-time device monitoring, data visualization, and control capabilities.

## Overview

The IoT Dashboard provides a centralized platform for managing and monitoring IoT devices. It features a responsive web interface for data visualization and real-time device status updates via WebSocket connections.

## Features

- **Real-time Monitoring**
  - Live device status updates via WebSocket
  - Temperature, humidity, and battery level tracking
  - Automatic reconnection handling
  - Graceful shutdown support

- **Device Management**
  - Register and unregister devices
  - View device details and status
  - Monitor device data in real-time
  - Support for multiple device types

- **Data Visualization**
  - Interactive charts for sensor data
  - Historical data tracking
  - Customizable dashboards
  - Responsive design for all devices

## Architecture

The project consists of two main components:

1. **Backend (Go)**
   - RESTful API for device management
   - WebSocket server for real-time updates
   - Device status monitoring
   - Data persistence

2. **Frontend (Vue.js)**
   - Responsive web interface
   - Real-time data visualization
   - Device management controls
   - WebSocket client for live updates

## Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- Modern web browser with WebSocket support

## Installation

### Backend Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/iot-dashboard.git
   cd iot-dashboard
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install frontend dependencies:
   ```bash
   npm install
   ```

## Running the Application

### Starting the Backend

1. From the project root, start the backend server:
   ```bash
   go run cmd/server/main.go
   ```

   The backend server will start on port 8080 by default.

### Starting the Frontend

1. In a new terminal, navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Start the development server:
   ```bash
   npm run serve
   ```

   The frontend will be available at http://localhost:8086

3. Access the dashboard in your web browser:
   ```
   http://localhost:8086
   ```

## Project Structure

```
iot-dashboard/
├── cmd/
│   └── server/          # Server entry point
│       └── main.go
├── internal/
│   ├── api/             # API handlers
│   ├── device/          # Device management
│   └── websocket/       # WebSocket implementation
├── frontend/            # Vue.js frontend
│   ├── public/          # Static assets
│   ├── src/             # Source code
│   │   ├── api/         # API client
│   │   ├── components/  # Vue components
│   │   ├── router/      # Vue router
│   │   ├── store/       # Vuex store
│   │   └── views/       # Vue views
│   ├── package.json     # Frontend dependencies
│   └── vue.config.js    # Vue configuration
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── README.md            # Project documentation
```

## API Documentation

### REST API Endpoints

- `GET /api/devices` - Get all registered devices
- `GET /api/devices/{id}` - Get a specific device by ID
- `POST /api/devices` - Register a new device
- `DELETE /api/devices/{id}` - Unregister a device
- `POST /api/devices/{id}/data` - Update device data

### WebSocket Protocol

Devices can connect to the WebSocket endpoint at `/ws` with their device ID as a query parameter:

```
ws://localhost:8080/ws?device_id=<device_id>
```

#### Message Format

Devices should send JSON messages in the following format:

```json
{
    "temperature": 25.5,
    "humidity": 60.0,
    "battery": 85.0,
    "timestamp": "2024-01-15T12:00:00Z"
}
```

## Development

### Building the Backend

```bash
go build -o iot-dashboard cmd/server/main.go
```

### Building the Frontend

```bash
cd frontend
npm run build
```

### Running Tests

```bash
# Backend tests
go test ./...

# Frontend tests
cd frontend
npm run test
```

## Troubleshooting

### WebSocket Connection Issues

If you encounter WebSocket connection issues:

1. Ensure the backend server is running on port 8080
2. Check that the frontend is configured to connect to the correct WebSocket URL
3. Verify that your browser supports WebSocket connections
4. Check the browser console for any error messages

### API Connection Issues

If you encounter API connection issues:

1. Ensure the backend server is running
2. Check that the frontend is configured to use the correct API URL
3. Verify that CORS is properly configured on the backend

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 