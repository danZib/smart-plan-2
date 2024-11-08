package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"smart-plan-2/internal/api"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Config struct {
	Port string
}

func main() {
	// Config
	cfg := Config{
		Port: "8080",
	}

	// Context
	ctx := context.Background()

	// Controller
	controller := api.NewController()

	// Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))

	router.Get("/health", controller.GetHealth)

	// Server
	server := http.Server{
		Addr:    net.JoinHostPort("", cfg.Port),
		Handler: router,
		// share context for all request to share logger, values and graceful shutdown
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	// Start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
