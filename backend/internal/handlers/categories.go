package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pharmacare-backend/internal/database"
	"pharmacare-backend/internal/models"

	"github.com/gorilla/mux"
)

// GetAllCategories returns all categories with drug count
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT c.id, c.name, c.slug, c.description, c.icon, c.created_at, c.updated_at,
		       COUNT(d.id) as drug_count
		FROM categories c
		LEFT JOIN drugs d ON c.id = d.category_id
		GROUP BY c.id, c.name, c.slug, c.description, c.icon, c.created_at, c.updated_at
		ORDER BY c.name
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.ID, &category.Name, &category.Slug, &category.Description,
			&category.Icon, &category.CreatedAt, &category.UpdatedAt, &category.Count,
		)
		if err != nil {
			http.Error(w, "Failed to scan category", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	response := models.APIResponse{
		Success: true,
		Data:    categories,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCategoryBySlug returns a category by slug
func GetCategoryBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	query := `
		SELECT c.id, c.name, c.slug, c.description, c.icon, c.created_at, c.updated_at,
		       COUNT(d.id) as drug_count
		FROM categories c
		LEFT JOIN drugs d ON c.id = d.category_id
		WHERE c.slug = $1
		GROUP BY c.id, c.name, c.slug, c.description, c.icon, c.created_at, c.updated_at
	`

	var category models.Category
	err := database.DB.QueryRow(query, slug).Scan(
		&category.ID, &category.Name, &category.Slug, &category.Description,
		&category.Icon, &category.CreatedAt, &category.UpdatedAt, &category.Count,
	)

	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	response := models.APIResponse{
		Success: true,
		Data:    category,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetDrugsByCategory returns all drugs in a category
func GetDrugsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	query := `
		SELECT d.id, d.name, d.description, d.composition, d.price, d.stock,
		       d.category_id, c.name as category_name, d.manufacturer, d.dosage,
		       d.side_effects, d.contraindications, d.image_url, d.requires_prescription,
		       d.created_at, d.updated_at
		FROM drugs d
		LEFT JOIN categories c ON d.category_id = c.id
		WHERE d.category_id = $1
		ORDER BY d.name
	`

	rows, err := database.DB.Query(query, categoryID)
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
			&drug.Manufacturer, &drug.Dosage, &drug.SideEffects,
			&drug.Contraindications, &drug.ImageURL,
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
