# IoT Dashboard

A scalable IoT management platform built in Go, featuring real-time device monitoring and control capabilities.

## Features

- Real-time device monitoring using WebSocket connections
- Secure communication between IoT devices and centralized dashboard
- Responsive web interface for data visualization
- Support for multiple device types and sensors
- Real-time temperature, humidity, and battery monitoring
- Interactive charts for data visualization
- Automatic reconnection handling
- Graceful shutdown support

## Prerequisites

- Go 1.21 or higher
- Modern web browser with WebSocket support

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/iot-dashboard.git
cd iot-dashboard
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

1. Start the server:
```bash
go run cmd/server/main.go
```

2. Open your web browser and navigate to:
```
http://localhost:8080
```

## Project Structure

```
iot-dashboard/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── device/
│       └── device.go
├── static/
│   └── index.html
├── go.mod
├── go.sum
└── README.md
```

## WebSocket Protocol

The dashboard uses WebSocket for real-time communication. Devices can connect to the WebSocket endpoint at `/ws` with their device ID as a query parameter:

```
ws://localhost:8080/ws?device_id=<device_id>
```

### Message Format

Devices should send JSON messages in the following format:

```json
{
    "temperature": 25.5,
    "humidity": 60.0,
    "battery": 85.0,
    "timestamp": "2024-01-15T12:00:00Z"
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 