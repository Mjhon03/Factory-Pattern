package repositories

import (
	"context"
	"hexagonal-example/domain/entities"
)

// ProductRepository define la interfaz para el repositorio de productos
// Sigue el mismo patrón que UserRepository pero para productos
type ProductRepository interface {
	// Save guarda un producto en el repositorio
	Save(ctx context.Context, product *entities.Product) error

	// FindByID busca un producto por su ID
	FindByID(ctx context.Context, id string) (*entities.Product, error)

	// FindByName busca productos por nombre (puede retornar múltiples)
	FindByName(ctx context.Context, name string) ([]*entities.Product, error)

	// FindByCategory busca productos por categoría
	FindByCategory(ctx context.Context, category string, limit, offset int) ([]*entities.Product, error)

	// FindAll retorna todos los productos
	FindAll(ctx context.Context, limit, offset int) ([]*entities.Product, error)

	// FindAvailable retorna productos disponibles (activos y con stock)
	FindAvailable(ctx context.Context, limit, offset int) ([]*entities.Product, error)

	// FindByPriceRange busca productos en un rango de precios
	FindByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]*entities.Product, error)

	// Delete elimina un producto del repositorio
	Delete(ctx context.Context, id string) error

	// Exists verifica si un producto existe por ID
	Exists(ctx context.Context, id string) (bool, error)

	// Count retorna el número total de productos
	Count(ctx context.Context) (int, error)

	// CountByCategory retorna el número de productos en una categoría
	CountByCategory(ctx context.Context, category string) (int, error)
}

// ProductRepositoryError define errores específicos del repositorio de productos
type ProductRepositoryError struct {
	Message string
	Err     error
}

func (e *ProductRepositoryError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *ProductRepositoryError) Unwrap() error {
	return e.Err
}

// Errores comunes del repositorio de productos
var (
	ErrProductNotFound      = &ProductRepositoryError{Message: "product not found"}
	ErrProductAlreadyExists = &ProductRepositoryError{Message: "product already exists"}
	ErrInvalidProductData   = &ProductRepositoryError{Message: "invalid product data"}
	ErrInsufficientStock    = &ProductRepositoryError{Message: "insufficient stock"}
)