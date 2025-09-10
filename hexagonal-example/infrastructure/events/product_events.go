package events

import "time"

// ProductCreatedEvent representa el evento cuando se crea un producto
type ProductCreatedEvent struct {
	ProductID  string    `json:"product_id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
}

// ProductUpdatedEvent representa el evento cuando se actualiza un producto
type ProductUpdatedEvent struct {
	ProductID  string    `json:"product_id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// StockUpdatedEvent representa el evento cuando se actualiza el stock de un producto
type StockUpdatedEvent struct {
	ProductID  string    `json:"product_id"`
	Name       string    `json:"name"`
	OldStock   int       `json:"old_stock"`
	NewStock   int       `json:"new_stock"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ProductDeactivatedEvent representa el evento cuando se desactiva un producto
type ProductDeactivatedEvent struct {
	ProductID     string    `json:"product_id"`
	Name          string    `json:"name"`
	DeactivatedAt time.Time `json:"deactivated_at"`
}

// ProductActivatedEvent representa el evento cuando se activa un producto
type ProductActivatedEvent struct {
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	ActivatedAt time.Time `json:"activated_at"`
}