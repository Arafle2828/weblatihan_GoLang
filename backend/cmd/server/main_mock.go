package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Mock data structures
type Drug struct {
	ID                   int      `json:"id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Composition          string   `json:"composition"`
	Price                int      `json:"price"`
	Stock                int      `json:"stock"`
	CategoryID           int      `json:"category_id"`
	CategoryName         string   `json:"category_name,omitempty"`
	Manufacturer         string   `json:"manufacturer"`
	Dosage               string   `json:"dosage"`
	SideEffects          []string `json:"side_effects"`
	Contraindications    []string `json:"contraindications"`
	ImageURL             string   `json:"image_url"`
	RequiresPrescription bool     `json:"requires_prescription"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Count       int    `json:"count,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Mock data
var mockCategories = []Category{
	{ID: 1, Name: "Obat Demam", Slug: "obat-demam", Description: "Obat untuk menurunkan demam", Icon: "🌡️", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: 2, Name: "Obat Batuk", Slug: "obat-batuk", Description: "Obat untuk mengatasi batuk", Icon: "🫁", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: 3, Name: "Obat Sakit Kepala", Slug: "obat-sakit-kepala", Description: "Obat untuk mengatasi sakit kepala", Icon: "🧠", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: 4, Name: "Vitamin", Slug: "vitamin", Description: "Suplemen vitamin dan mineral", Icon: "💊", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: 5, Name: "Obat Luar", Slug: "obat-luar", Description: "Obat untuk penggunaan luar", Icon: "🧴", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: 6, Name: "Obat Lambung", Slug: "obat-lambung", Description: "Obat untuk masalah lambung", Icon: "🫄", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
}

var mockDrugs = []Drug{
	{
		ID: 1, Name: "Paracetamol 500mg", Description: "Obat penurun demam dan pereda nyeri",
		Composition: "Paracetamol 500mg", Price: 15000, Stock: 100, CategoryID: 1,
		Manufacturer: "Kimia Farma", Dosage: "3x1 tablet per hari",
		SideEffects: []string{"Mual", "Muntah", "Ruam kulit"}, Contraindications: []string{"Gangguan hati berat"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 2, Name: "OBH Combi", Description: "Obat batuk berdahak dewasa",
		Composition: "Sukralfat 500mg", Price: 25000, Stock: 50, CategoryID: 2,
		Manufacturer: "Dexa Medica", Dosage: "3x1 sendok teh per hari",
		SideEffects: []string{"Mengantuk", "Mual"}, Contraindications: []string{"Anak dibawah 6 tahun"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 3, Name: "Bodrex", Description: "Obat sakit kepala dan pusing",
		Composition: "Paracetamol 500mg, Kafein 50mg", Price: 12000, Stock: 75, CategoryID: 3,
		Manufacturer: "Tempo Scan Pacific", Dosage: "2x1 tablet per hari",
		SideEffects: []string{"Jantung berdebar", "Sulit tidur"}, Contraindications: []string{"Hipertensi", "Gangguan jantung"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 4, Name: "Vitamin C 1000mg", Description: "Suplemen vitamin C untuk daya tahan tubuh",
		Composition: "Ascorbic Acid 1000mg", Price: 45000, Stock: 200, CategoryID: 4,
		Manufacturer: "Blackmores", Dosage: "1x1 tablet per hari",
		SideEffects: []string{"Diare", "Mual"}, Contraindications: []string{"Batu ginjal"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 5, Name: "Betadine", Description: "Antiseptik untuk luka luar",
		Composition: "Povidone Iodine 10%", Price: 35000, Stock: 30, CategoryID: 5,
		Manufacturer: "Mundipharma", Dosage: "Oleskan pada luka 2-3x sehari",
		SideEffects: []string{"Iritasi kulit", "Reaksi alergi"}, Contraindications: []string{"Hipersensitif iodine", "Gangguan tiroid"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 6, Name: "Antasida Doen", Description: "Obat untuk mengatasi asam lambung",
		Composition: "Aluminium Hydroxide, Magnesium Hydroxide", Price: 18000, Stock: 60, CategoryID: 6,
		Manufacturer: "Dankos Farma", Dosage: "2x1 tablet sesudah makan",
		SideEffects: []string{"Konstipasi", "Diare"}, Contraindications: []string{"Gangguan ginjal berat"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 7, Name: "Amoxicillin 500mg", Description: "Antibiotik untuk infeksi bakteri",
		Composition: "Amoxicillin 500mg", Price: 55000, Stock: 25, CategoryID: 1,
		Manufacturer: "Sanbe Farma", Dosage: "3x1 kapsul per hari",
		SideEffects: []string{"Diare", "Mual", "Ruam kulit"}, Contraindications: []string{"Alergi penisilin"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: true,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
	{
		ID: 8, Name: "Ibuprofen 400mg", Description: "Anti-inflamasi dan pereda nyeri",
		Composition: "Ibuprofen 400mg", Price: 22000, Stock: 80, CategoryID: 3,
		Manufacturer: "Kalbe Farma", Dosage: "3x1 tablet per hari sesudah makan",
		SideEffects: []string{"Nyeri perut", "Mual", "Pusing"}, Contraindications: []string{"Tukak lambung", "Gangguan ginjal"},
		ImageURL: "/api/placeholder/300/300", RequiresPrescription: false,
		CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z",
	},
}

// Handler functions
func getAllDrugs(w http.ResponseWriter, r *http.Request) {
	// Add category names to drugs
	for i := range mockDrugs {
		for _, cat := range mockCategories {
			if cat.ID == mockDrugs[i].CategoryID {
				mockDrugs[i].CategoryName = cat.Name
				break
			}
		}
	}

	response := APIResponse{
		Success: true,
		Data:    mockDrugs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getDrugByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid drug ID",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, drug := range mockDrugs {
		if drug.ID == id {
			// Add category name
			for _, cat := range mockCategories {
				if cat.ID == drug.CategoryID {
					drug.CategoryName = cat.Name
					break
				}
			}

			response := APIResponse{
				Success: true,
				Data:    drug,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response := APIResponse{
		Success: false,
		Error:   "Drug not found",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func searchDrugs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		getAllDrugs(w, r)
		return
	}

	var results []Drug
	query = strings.ToLower(query)

	for _, drug := range mockDrugs {
		if strings.Contains(strings.ToLower(drug.Name), query) ||
			strings.Contains(strings.ToLower(drug.Description), query) ||
			strings.Contains(strings.ToLower(drug.Composition), query) {

			// Add category name
			for _, cat := range mockCategories {
				if cat.ID == drug.CategoryID {
					drug.CategoryName = cat.Name
					break
				}
			}
			results = append(results, drug)
		}
	}

	response := APIResponse{
		Success: true,
		Data:    results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	// Add drug count to categories
	for i := range mockCategories {
		count := 0
		for _, drug := range mockDrugs {
			if drug.CategoryID == mockCategories[i].ID {
				count++
			}
		}
		mockCategories[i].Count = count
	}

	response := APIResponse{
		Success: true,
		Data:    mockCategories,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getCategoryBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, category := range mockCategories {
		if category.Slug == slug {
			// Add drug count
			count := 0
			for _, drug := range mockDrugs {
				if drug.CategoryID == category.ID {
					count++
				}
			}
			category.Count = count

			response := APIResponse{
				Success: true,
				Data:    category,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response := APIResponse{
		Success: false,
		Error:   "Category not found",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func getDrugsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid category ID",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	var results []Drug
	for _, drug := range mockDrugs {
		if drug.CategoryID == categoryID {
			// Add category name
			for _, cat := range mockCategories {
				if cat.ID == drug.CategoryID {
					drug.CategoryName = cat.Name
					break
				}
			}
			results = append(results, drug)
		}
	}

	response := APIResponse{
		Success: true,
		Data:    results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "healthy",
		"service": "PharmaCare API (Mock)",
		"time":    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Health check
	api.HandleFunc("/health", healthCheck).Methods("GET")

	// Drug routes
	api.HandleFunc("/drugs", getAllDrugs).Methods("GET")
	api.HandleFunc("/drugs/{id:[0-9]+}", getDrugByID).Methods("GET")
	api.HandleFunc("/drugs/search", searchDrugs).Methods("GET")

	// Category routes
	api.HandleFunc("/categories", getAllCategories).Methods("GET")
	api.HandleFunc("/categories/{slug:[a-z-]+}", getCategoryBySlug).Methods("GET")
	api.HandleFunc("/categories/{id:[0-9]+}/drugs", getDrugsByCategory).Methods("GET")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 PharmaCare API Server (Mock) starting on port %s", port)
	log.Printf("📍 Health check: http://localhost:%s/api/v1/health", port)
	log.Printf("📊 Drugs API: http://localhost:%s/api/v1/drugs", port)
	log.Printf("📂 Categories API: http://localhost:%s/api/v1/categories", port)

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
