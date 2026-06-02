package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"placepicker"
	"placepicker/internal/config"
	"placepicker/internal/connection"
	"placepicker/internal/db"
	"placepicker/internal/handlers"
)

func main() {
	// ENV
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to initialize configuration: %v", err)
	}

	// initialize the database pool
	conn, err := connection.Connect(cfg.DBURL, cfg.AuthToken)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to Turso!")

	queries := db.New(conn)

	// handlers
	locationHandler := handlers.NewHandler(queries)

	// routes
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", locationHandler.GetAvailableLocations)

	// static assets
	staticFS, err := fs.Sub(placepicker.EmbeddedAssets, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
