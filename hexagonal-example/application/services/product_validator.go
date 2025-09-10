package services

import (
	"errors"
	"strings"
)

// ProductValidator se encarga únicamente de la validación de datos de producto
type ProductValidator struct{}

// NewProductValidator crea una nueva instancia del validador de productos
func NewProductValidator() *ProductValidator {
	return &ProductValidator{}
}

// ValidateCreateProduct valida los datos para crear un nuevo producto
func (v *ProductValidator) ValidateCreateProduct(id, name, description, category string, price float64, stock int) error {
	// Validar ID
	if err := v.validateID(id); err != nil {
		return err
	}

	// Validar nombre
	if err := v.validateName(name); err != nil {
		return err
	}

	// Validar descripción
	if err := v.validateDescription(description); err != nil {
		return err
	}

	// Validar categoría
	if err := v.validateCategory(category); err != nil {
		return err
	}

	// Validar precio
	if err := v.validatePrice(price); err != nil {
		return err
	}

	// Validar stock
	if err := v.validateStock(stock); err != nil {
		return err
	}

	return nil
}

// ValidateUpdateProduct valida los datos para actualizar un producto
func (v *ProductValidator) ValidateUpdateProduct(id string, name, description, category *string, price *float64, stock *int) error {
	// El ID siempre debe ser válido
	if err := v.validateID(id); err != nil {
		return err
	}

	// Validar campos opcionales si se proporcionan
	if name != nil {
		if err := v.validateName(*name); err != nil {
			return err
		}
	}

	if description != nil {
		if err := v.validateDescription(*description); err != nil {
			return err
		}
	}

	if category != nil {
		if err := v.validateCategory(*category); err != nil {
			return err
		}
	}

	if price != nil {
		if err := v.validatePrice(*price); err != nil {
			return err
		}
	}

	if stock != nil {
		if err := v.validateStock(*stock); err != nil {
			return err
		}
	}

	return nil
}

// validateID valida el ID del producto
func (v *ProductValidator) validateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("product ID cannot be empty")
	}
	if len(id) < 3 {
		return errors.New("product ID must be at least 3 characters long")
	}
	if len(id) > 50 {
		return errors.New("product ID cannot exceed 50 characters")
	}
	return nil
}

// validateName valida el nombre del producto
func (v *ProductValidator) validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("product name cannot be empty")
	}
	if len(name) < 2 {
		return errors.New("product name must be at least 2 characters long")
	}
	if len(name) > 200 {
		return errors.New("product name cannot exceed 200 characters")
	}
	return nil
}

// validateDescription valida la descripción del producto
func (v *ProductValidator) validateDescription(description string) error {
	if len(description) > 1000 {
		return errors.New("product description cannot exceed 1000 characters")
	}
	return nil
}

// validateCategory valida la categoría del producto
func (v *ProductValidator) validateCategory(category string) error {
	if strings.TrimSpace(category) == "" {
		return errors.New("product category cannot be empty")
	}
	if len(category) < 2 {
		return errors.New("product category must be at least 2 characters long")
	}
	if len(category) > 100 {
		return errors.New("product category cannot exceed 100 characters")
	}
	return nil
}

// validatePrice valida el precio del producto
func (v *ProductValidator) validatePrice(price float64) error {
	if price < 0 {
		return errors.New("product price cannot be negative")
	}
	if price > 1000000 {
		return errors.New("product price cannot exceed 1,000,000")
	}
	return nil
}

// validateStock valida el stock del producto
func (v *ProductValidator) validateStock(stock int) error {
	if stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	if stock > 1000000 {
		return errors.New("product stock cannot exceed 1,000,000")
	}
	return nil
}