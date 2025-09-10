package entities

import (
	"errors"
	"time"
)

// Product representa la entidad de producto en el dominio
// Contiene toda la lógica de negocio relacionada con productos
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
}

// NewProduct crea una nueva instancia de Product con validaciones de dominio
func NewProduct(id, name, description, category string, price float64, stock int) (*Product, error) {
	// Validaciones de dominio
	if id == "" {
		return nil, errors.New("product ID cannot be empty")
	}
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if price < 0 {
		return nil, errors.New("product price cannot be negative")
	}
	if stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	// Crear el producto con valores por defecto
	now := time.Now()
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Category:    category,
		CreatedAt:   now,
		UpdatedAt:   now,
		IsActive:    true, // Los productos se crean activos por defecto
	}, nil
}

// UpdatePrice actualiza el precio del producto
func (p *Product) UpdatePrice(newPrice float64) error {
	if newPrice < 0 {
		return errors.New("price cannot be negative")
	}
	
	p.Price = newPrice
	p.UpdatedAt = time.Now()
	return nil
}

// UpdateStock actualiza el stock del producto
func (p *Product) UpdateStock(newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	
	p.Stock = newStock
	p.UpdatedAt = time.Now()
	return nil
}

// AddStock añade stock al producto
func (p *Product) AddStock(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	
	p.Stock += quantity
	p.UpdatedAt = time.Now()
	return nil
}

// RemoveStock reduce el stock del producto
func (p *Product) RemoveStock(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	if p.Stock < quantity {
		return errors.New("insufficient stock")
	}
	
	p.Stock -= quantity
	p.UpdatedAt = time.Now()
	return nil
}

// Deactivate desactiva un producto
func (p *Product) Deactivate() {
	p.IsActive = false
	p.UpdatedAt = time.Now()
}

// Activate activa un producto
func (p *Product) Activate() {
	p.IsActive = true
	p.UpdatedAt = time.Now()
}

// IsAvailable verifica si el producto está disponible para venta
func (p *Product) IsAvailable() bool {
	return p.IsActive && p.Stock > 0
}

// IsValid verifica si el producto es válido según las reglas de negocio
func (p *Product) IsValid() bool {
	return p.ID != "" && p.Name != "" && p.Price >= 0 && p.Stock >= 0
}