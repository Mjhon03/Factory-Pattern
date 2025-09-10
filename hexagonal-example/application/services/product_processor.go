package services

import (
	"context"
	"errors"
	"hexagonal-example/domain/entities"
	"hexagonal-example/domain/repositories"
)

// ProductProcessor se encarga del procesamiento de la lógica de negocio de productos
type ProductProcessor struct {
	productRepo repositories.ProductRepository
}

// NewProductProcessor crea una nueva instancia del procesador de productos
func NewProductProcessor(productRepo repositories.ProductRepository) *ProductProcessor {
	return &ProductProcessor{
		productRepo: productRepo,
	}
}

// CreateProduct crea un nuevo producto en el sistema
func (p *ProductProcessor) CreateProduct(ctx context.Context, id, name, description, category string, price float64, stock int) (*entities.Product, error) {
	// Verificar si el producto ya existe
	exists, err := p.productRepo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("product already exists")
	}

	// Crear la entidad de producto
	product, err := entities.NewProduct(id, name, description, category, price, stock)
	if err != nil {
		return nil, err
	}

	// Guardar en el repositorio
	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct actualiza un producto existente
func (p *ProductProcessor) UpdateProduct(ctx context.Context, id string, name, description, category *string, price *float64, stock *int) (*entities.Product, error) {
	// Buscar el producto existente
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	// Actualizar campos si se proporcionan
	if name != nil {
		product.Name = *name
	}
	if description != nil {
		product.Description = *description
	}
	if category != nil {
		product.Category = *category
	}
	if price != nil {
		if err := product.UpdatePrice(*price); err != nil {
			return nil, err
		}
	}
	if stock != nil {
		if err := product.UpdateStock(*stock); err != nil {
			return nil, err
		}
	}

	// Guardar los cambios
	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateStock actualiza el stock de un producto
func (p *ProductProcessor) UpdateStock(ctx context.Context, id string, newStock int) (*entities.Product, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if err := product.UpdateStock(newStock); err != nil {
		return nil, err
	}

	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// AddStock añade stock a un producto
func (p *ProductProcessor) AddStock(ctx context.Context, id string, quantity int) (*entities.Product, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if err := product.AddStock(quantity); err != nil {
		return nil, err
	}

	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// RemoveStock reduce el stock de un producto
func (p *ProductProcessor) RemoveStock(ctx context.Context, id string, quantity int) (*entities.Product, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if err := product.RemoveStock(quantity); err != nil {
		return nil, err
	}

	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// DeactivateProduct desactiva un producto
func (p *ProductProcessor) DeactivateProduct(ctx context.Context, id string) (*entities.Product, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	product.Deactivate()
	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// ActivateProduct activa un producto
func (p *ProductProcessor) ActivateProduct(ctx context.Context, id string) (*entities.Product, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	product.Activate()
	if err := p.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// GetProduct obtiene un producto por ID
func (p *ProductProcessor) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// ListProducts obtiene una lista de productos
func (p *ProductProcessor) ListProducts(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	return p.productRepo.FindAll(ctx, limit, offset)
}

// ListAvailableProducts obtiene una lista de productos disponibles
func (p *ProductProcessor) ListAvailableProducts(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	return p.productRepo.FindAvailable(ctx, limit, offset)
}

// ListProductsByCategory obtiene productos por categoría
func (p *ProductProcessor) ListProductsByCategory(ctx context.Context, category string, limit, offset int) ([]*entities.Product, error) {
	return p.productRepo.FindByCategory(ctx, category, limit, offset)
}

// ListProductsByPriceRange obtiene productos en un rango de precios
func (p *ProductProcessor) ListProductsByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]*entities.Product, error) {
	return p.productRepo.FindByPriceRange(ctx, minPrice, maxPrice, limit, offset)
}