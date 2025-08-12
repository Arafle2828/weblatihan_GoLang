package main

import (
	"log"
	"net/http"
	"os"

	"pharmacare-backend/internal/database"
	"pharmacare-backend/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize router
	r := mux.NewRouter()

	// API prefix
	api := r.PathPrefix("/api/v1").Subrouter()

	// Drug routes
	api.HandleFunc("/drugs", handlers.GetAllDrugs).Methods("GET")
	api.HandleFunc("/drugs/{id:[0-9]+}", handlers.GetDrugByID).Methods("GET")
	api.HandleFunc("/drugs/search", handlers.SearchDrugs).Methods("GET")

	// Category routes
	api.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")
	api.HandleFunc("/categories/{slug}", handlers.GetCategoryBySlug).Methods("GET")
	api.HandleFunc("/categories/{id:[0-9]+}/drugs", handlers.GetDrugsByCategory).Methods("GET")

	// Health check
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "service": "pharmacare-backend"}`))
	}).Methods("GET")

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Wrap router with CORS
	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 PharmaCare Backend API Server starting on port %s", port)
	log.Printf("📡 API endpoints available at http://localhost:%s/api/v1", port)
	log.Printf("🏥 Health check: http://localhost:%s/api/v1/health", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
