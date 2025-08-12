package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"pharmacare-backend/internal/database"
	"pharmacare-backend/internal/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

// sendErrorResponse sends a consistent JSON error response
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.APIResponse{
		Success: false,
		Error:   message,
	}

	json.NewEncoder(w).Encode(response)
}

// GetAllDrugs returns all drugs
func GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	log.Println("🔍 GetAllDrugs called")

	query := `
		SELECT d.id, d.name, d.description, d.composition, d.price, d.stock,
		       d.category_id, c.name as category_name, d.manufacturer, d.dosage,
		       d.side_effects, d.contraindications, d.image_url, d.requires_prescription,
		       d.created_at, d.updated_at
		FROM drugs d
		LEFT JOIN categories c ON d.category_id = c.id
		ORDER BY d.name
	`

	log.Println("📊 Executing query...")
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("❌ Query failed: %v", err)
		sendErrorResponse(w, "Failed to fetch drugs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Println("✅ Query executed successfully")

	var drugs []models.Drug
	rowCount := 0
	for rows.Next() {
		rowCount++
		log.Printf("📄 Processing row %d", rowCount)

		var drug models.Drug
		var imageURL sql.NullString
		var categoryName sql.NullString
		var description sql.NullString
		var composition sql.NullString
		var manufacturer sql.NullString
		var dosage sql.NullString

		err := rows.Scan(
			&drug.ID, &drug.Name, &description, &composition,
			&drug.Price, &drug.Stock, &drug.CategoryID, &categoryName,
			&manufacturer, &dosage, pq.Array(&drug.SideEffects),
			pq.Array(&drug.Contraindications), &imageURL,
			&drug.RequiresPrescription, &drug.CreatedAt, &drug.UpdatedAt,
		)
		if err != nil {
			log.Printf("❌ Failed to scan row %d: %v", rowCount, err)
			sendErrorResponse(w, "Failed to scan drug", http.StatusInternalServerError)
			return
		}

		// Convert sql.NullString to string
		if description.Valid {
			drug.Description = description.String
		}
		if composition.Valid {
			drug.Composition = composition.String
		}
		if categoryName.Valid {
			drug.CategoryName = categoryName.String
		}
		if manufacturer.Valid {
			drug.Manufacturer = manufacturer.String
		}
		if dosage.Valid {
			drug.Dosage = dosage.String
		}
		if imageURL.Valid {
			drug.ImageURL = imageURL.String
		}

		drugs = append(drugs, drug)
	}

	log.Printf("📊 Total rows processed: %d", rowCount)

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
		sendErrorResponse(w, "Invalid drug ID", http.StatusBadRequest)
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
		sendErrorResponse(w, "Drug not found", http.StatusNotFound)
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
		sendErrorResponse(w, "Failed to search drugs", http.StatusInternalServerError)
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
			sendErrorResponse(w, "Failed to scan drug", http.StatusInternalServerError)
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
