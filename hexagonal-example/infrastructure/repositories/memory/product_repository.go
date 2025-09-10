package memory

import (
	"context"
	"errors"
	"hexagonal-example/domain/entities"
	"hexagonal-example/domain/repositories"
	"sync"
)

// InMemoryProductRepository implementa ProductRepository usando memoria
type InMemoryProductRepository struct {
	products map[string]*entities.Product
	mutex    sync.RWMutex
}

// NewProductRepository crea una nueva instancia del repositorio de productos en memoria
func NewProductRepository() repositories.ProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*entities.Product),
	}
}

// Save guarda un producto en el repositorio
func (r *InMemoryProductRepository) Save(ctx context.Context, product *entities.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Crear una copia del producto para evitar modificaciones externas
	productCopy := *product
	r.products[product.ID] = &productCopy
	return nil
}

// FindByID busca un producto por su ID
func (r *InMemoryProductRepository) FindByID(ctx context.Context, id string) (*entities.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, nil
	}

	// Retornar una copia para evitar modificaciones externas
	productCopy := *product
	return &productCopy, nil
}

// FindByName busca productos por nombre
func (r *InMemoryProductRepository) FindByName(ctx context.Context, name string) ([]*entities.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var result []*entities.Product
	for _, product := range r.products {
		if product.Name == name {
			// Retornar una copia para evitar modificaciones externas
			productCopy := *product
			result = append(result, &productCopy)
		}
	}

	return result, nil
}

// FindByCategory busca productos por categoría
func (r *InMemoryProductRepository) FindByCategory(ctx context.Context, category string, limit, offset int) ([]*entities.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var products []*entities.Product
	for _, product := range r.products {
		if product.Category == category {
			products = append(products, product)
		}
	}

	// Aplicar paginación
	start := offset
	end := offset + limit
	if start >= len(products) {
		return []*entities.Product{}, nil
	}
	if end > len(products) {
		end = len(products)
	}

	// Retornar copias para evitar modificaciones externas
	result := make([]*entities.Product, 0, end-start)
	for i := start; i < end; i++ {
		productCopy := *products[i]
		result = append(result, &productCopy)
	}

	return result, nil
}

// FindAll retorna todos los productos
func (r *InMemoryProductRepository) FindAll(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	products := make([]*entities.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}

	// Aplicar paginación
	start := offset
	end := offset + limit
	if start >= len(products) {
		return []*entities.Product{}, nil
	}
	if end > len(products) {
		end = len(products)
	}

	// Retornar copias para evitar modificaciones externas
	result := make([]*entities.Product, 0, end-start)
	for i := start; i < end; i++ {
		productCopy := *products[i]
		result = append(result, &productCopy)
	}

	return result, nil
}

// FindAvailable retorna productos disponibles (activos y con stock)
func (r *InMemoryProductRepository) FindAvailable(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var availableProducts []*entities.Product
	for _, product := range r.products {
		if product.IsAvailable() {
			availableProducts = append(availableProducts, product)
		}
	}

	// Aplicar paginación
	start := offset
	end := offset + limit
	if start >= len(availableProducts) {
		return []*entities.Product{}, nil
	}
	if end > len(availableProducts) {
		end = len(availableProducts)
	}

	// Retornar copias para evitar modificaciones externas
	result := make([]*entities.Product, 0, end-start)
	for i := start; i < end; i++ {
		productCopy := *availableProducts[i]
		result = append(result, &productCopy)
	}

	return result, nil
}

// FindByPriceRange busca productos en un rango de precios
func (r *InMemoryProductRepository) FindByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]*entities.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var products []*entities.Product
	for _, product := range r.products {
		if product.Price >= minPrice && product.Price <= maxPrice {
			products = append(products, product)
		}
	}

	// Aplicar paginación
	start := offset
	end := offset + limit
	if start >= len(products) {
		return []*entities.Product{}, nil
	}
	if end > len(products) {
		end = len(products)
	}

	// Retornar copias para evitar modificaciones externas
	result := make([]*entities.Product, 0, end-start)
	for i := start; i < end; i++ {
		productCopy := *products[i]
		result = append(result, &productCopy)
	}

	return result, nil
}

// Delete elimina un producto del repositorio
func (r *InMemoryProductRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(r.products, id)
	return nil
}

// Exists verifica si un producto existe por ID
func (r *InMemoryProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, exists := r.products[id]
	return exists, nil
}

// Count retorna el número total de productos
func (r *InMemoryProductRepository) Count(ctx context.Context) (int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return len(r.products), nil
}

// CountByCategory retorna el número de productos en una categoría
func (r *InMemoryProductRepository) CountByCategory(ctx context.Context, category string) (int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	count := 0
	for _, product := range r.products {
		if product.Category == category {
			count++
		}
	}

	return count, nil
}