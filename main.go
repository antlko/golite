package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/antlko/golite/internal"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

//go:embed ui/dist
var app embed.FS

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		slog.Warn("Failed to load .env file", "error", err)
	}

	// Load application configuration
	var cfg internal.AppConfig
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		slog.Error("Failed to process configuration", "error", err)
		return
	}

	// Start the UI server in a separate goroutine
	go func() {
		startUI(cfg)
	}()

	// Initialize the backend service
	if err := internal.InitService(cfg); err != nil {
		slog.Error("Failed to initialize the service", "error", err)
		return
	}
}

// startUI serves the Vue.js app
func startUI(cfg internal.AppConfig) {
	os.Setenv("VITE_SERVER_HOST", cfg.Ui.ServerPort)

	// Prepare the embedded Vue app files
	dist, err := fs.Sub(app, "ui/dist")
	if err != nil {
		slog.Error("Failed to prepare embedded UI files", "error", err)
		return
	}

	// Create a router and add logging middleware
	router := http.NewServeMux()
	router.Handle("/", withLogging(http.StripPrefix("/", http.FileServer(http.FS(dist)))))

	// Start the HTTP server
	addr := fmt.Sprintf(":%s", cfg.Ui.VuePort)
	slog.Info("Starting UI server", "address", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		slog.Error("Failed to start UI server", "error", err)
	}
}

// withLogging is a middleware for logging HTTP requests
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Incoming request",
			"method", r.Method,
			"url", r.URL.String(),
			"remoteAddr", r.RemoteAddr,
		)

		// Serve the static files or fallback to index.html for SPA routes
		// If the requested path is not a static asset, serve index.html
		if !isStaticAsset(r.URL.Path) {
			if r.URL.Path != "/" {
				// Rewrite the URL to `/` for SPA routing
				r.URL.Path = "/"
			}
			// Serve index.html for all routes
			http.ServeFile(w, r, "ui/dist/index.html")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// isStaticAsset checks if the requested path is a static asset
func isStaticAsset(path string) bool {
	// Define the list of static file extensions
	staticExtensions := []string{".js", ".css", ".ico", ".png", ".jpg", ".jpeg", ".gif", ".woff", ".woff2", ".ttf", ".svg"}

	// Check if the path ends with any of the static file extensions
	for _, ext := range staticExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
