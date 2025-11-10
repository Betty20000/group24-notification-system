package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BerylCAtieno/group24-notification-system/services/orchestrator/internal/config"
	"github.com/BerylCAtieno/group24-notification-system/services/orchestrator/internal/handlers"
	"github.com/BerylCAtieno/group24-notification-system/services/orchestrator/internal/middleware"
	"github.com/BerylCAtieno/group24-notification-system/services/orchestrator/internal/mocks"
	"github.com/BerylCAtieno/group24-notification-system/services/orchestrator/internal/services"
	"github.com/BerylCAtieno/group24-notification-system/services/orchestrator/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	if err := logger.Initialize(cfg.Logging.Level, cfg.Logging.Format); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync()

	logger.Log.Info("Starting Orchestrator Service",
		zap.String("port", cfg.Server.Port),
		zap.Bool("mock_services", cfg.Services.UseMockServices),
	)

	// Initialize clients (using mocks for now)
	var userClient = mocks.NewUserServiceMock()
	var templateClient = mocks.NewTemplateServiceMock()

	logger.Log.Info("Using mock services for development")

	// Initialize services
	orchestrationService := services.NewOrchestrationService(userClient, templateClient)

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()
	notificationHandler := handlers.NewNotificationHandler(orchestrationService)

	// Setup Gin router
	if cfg.Logging.Format == "json" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Logging())

	// Health check endpoints
	router.GET("/health/live", healthHandler.Live)
	router.GET("/health/ready", healthHandler.Ready)
	router.GET("/health", healthHandler.Live) // Simple health check

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Notification endpoints
		v1.POST("/notifications", notificationHandler.Create)
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start server in a goroutine
	go func() {
		logger.Log.Info("Server started", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Log.Info("Server exited")
}
