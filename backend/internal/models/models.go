package models

import (
	"time"
)

type Drug struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Composition          string    `json:"composition"`
	Price                float64   `json:"price"`
	Stock                int       `json:"stock"`
	CategoryID           int       `json:"category_id"`
	CategoryName         string    `json:"category_name,omitempty"`
	Manufacturer         string    `json:"manufacturer"`
	Dosage               string    `json:"dosage"`
	SideEffects          []string  `json:"side_effects"`
	Contraindications    []string  `json:"contraindications"`
	ImageURL             string    `json:"image_url"`
	RequiresPrescription bool      `json:"requires_prescription"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Count       int       `json:"count,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID         int                    `json:"id"`
	Email      string                 `json:"email"`
	Name       string                 `json:"name"`
	Phone      string                 `json:"phone"`
	Address    map[string]interface{} `json:"address,omitempty"`
	IsVerified bool                   `json:"is_verified"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

type CartItem struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Drug      Drug      `json:"drug"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID              int                    `json:"id"`
	UserID          int                    `json:"user_id"`
	Total           float64                `json:"total"`
	Status          string                 `json:"status"`
	ShippingAddress map[string]interface{} `json:"shipping_address"`
	PaymentMethod   string                 `json:"payment_method"`
	PaymentStatus   string                 `json:"payment_status"`
	Items           []OrderItem            `json:"items,omitempty"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	Drug      Drug      `json:"drug"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
