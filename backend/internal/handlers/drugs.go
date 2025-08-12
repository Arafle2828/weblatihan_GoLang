package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pharmacare-backend/internal/database"
	"pharmacare-backend/internal/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

// GetAllDrugs returns all drugs
func GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT d.id, d.name, d.description, d.composition, d.price, d.stock,
		       d.category_id, c.name as category_name, d.manufacturer, d.dosage,
		       d.side_effects, d.contraindications, d.image_url, d.requires_prescription,
		       d.created_at, d.updated_at
		FROM drugs d
		LEFT JOIN categories c ON d.category_id = c.id
		ORDER BY d.name
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch drugs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var drugs []models.Drug
	for rows.Next() {
		var drug models.Drug
		err := rows.Scan(
			&drug.ID, &drug.Name, &drug.Description, &drug.Composition,
			&drug.Price, &drug.Stock, &drug.CategoryID, &drug.CategoryName,
			&drug.Manufacturer, &drug.Dosage, pq.Array(&drug.SideEffects),
			pq.Array(&drug.Contraindications), &drug.ImageURL,
			&drug.RequiresPrescription, &drug.CreatedAt, &drug.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to scan drug", http.StatusInternalServerError)
			return
		}
		drugs = append(drugs, drug)
	}

	response := models.APIResponse{
		Success: true,
		Data:    drugs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetDrugByID returns a single drug by ID
func GetDrugByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}

	query := `
		SELECT d.id, d.name, d.description, d.composition, d.price, d.stock,
		       d.category_id, c.name as category_name, d.manufacturer, d.dosage,
		       d.side_effects, d.contraindications, d.image_url, d.requires_prescription,
		       d.created_at, d.updated_at
		FROM drugs d
		LEFT JOIN categories c ON d.category_id = c.id
		WHERE d.id = $1
	`

	var drug models.Drug
	err = database.DB.QueryRow(query, id).Scan(
		&drug.ID, &drug.Name, &drug.Description, &drug.Composition,
		&drug.Price, &drug.Stock, &drug.CategoryID, &drug.CategoryName,
		&drug.Manufacturer, &drug.Dosage, pq.Array(&drug.SideEffects),
		pq.Array(&drug.Contraindications), &drug.ImageURL,
		&drug.RequiresPrescription, &drug.CreatedAt, &drug.UpdatedAt,
	)

	if err != nil {
		http.Error(w, "Drug not found", http.StatusNotFound)
		return
	}

	response := models.APIResponse{
		Success: true,
		Data:    drug,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SearchDrugs searches drugs by name, description, or composition
func SearchDrugs(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("q")
	if searchTerm == "" {
		GetAllDrugs(w, r)
		return
	}

	query := `
		SELECT d.id, d.name, d.description, d.composition, d.price, d.stock,
		       d.category_id, c.name as category_name, d.manufacturer, d.dosage,
		       d.side_effects, d.contraindications, d.image_url, d.requires_prescription,
		       d.created_at, d.updated_at
		FROM drugs d
		LEFT JOIN categories c ON d.category_id = c.id
		WHERE d.name ILIKE $1 OR d.description ILIKE $1 OR d.composition ILIKE $1
		ORDER BY d.name
	`

	rows, err := database.DB.Query(query, "%"+searchTerm+"%")
	if err != nil {
		http.Error(w, "Failed to search drugs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var drugs []models.Drug
	for rows.Next() {
		var drug models.Drug
		err := rows.Scan(
			&drug.ID, &drug.Name, &drug.Description, &drug.Composition,
			&drug.Price, &drug.Stock, &drug.CategoryID, &drug.CategoryName,
			&drug.Manufacturer, &drug.Dosage, pq.Array(&drug.SideEffects),
			pq.Array(&drug.Contraindications), &drug.ImageURL,
			&drug.RequiresPrescription, &drug.CreatedAt, &drug.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to scan drug", http.StatusInternalServerError)
			return
		}
		drugs = append(drugs, drug)
	}

	response := models.APIResponse{
		Success: true,
		Data:    drugs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
