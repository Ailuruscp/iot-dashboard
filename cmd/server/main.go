package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"iot-dashboard/internal/api"
	"iot-dashboard/internal/device"
	"iot-dashboard/internal/websocket"

	"github.com/gorilla/mux"
)

func main() {
	// Parse command-line flags
	addr := flag.String("addr", ":8080", "HTTP service address")
	flag.Parse()

	// Set up logger
	logger := log.New(os.Stdout, "IOT-DASHBOARD: ", log.LstdFlags)

	// Initialize device manager
	deviceManager := device.NewDeviceManager(logger)

	// Initialize WebSocket hub
	hub := websocket.NewHub(logger, deviceManager)
	go hub.Run()

	// Initialize router
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// Initialize API handler
	handler := api.NewHandler(hub, deviceManager, logger)
	handler.RegisterRoutes(router)

	// Initialize HTTP server
	server := &http.Server{
		Addr:    *addr,
		Handler: router,
	}

	// Channel to listen for interrupt signal to gracefully terminate the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		logger.Printf("Starting server on %s", *addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	logger.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server gracefully stopped")
}
