package services

import (
	"context"
	"hexagonal-example/domain/entities"
)

// ProductService es el servicio principal que orquesta los servicios granulares de productos
type ProductService struct {
	validator  *ProductValidator
	processor  *ProductProcessor
	publisher  *ProductEventPublisher
}

// NewProductService crea una nueva instancia del servicio de producto
func NewProductService(validator *ProductValidator, processor *ProductProcessor, publisher *ProductEventPublisher) *ProductService {
	return &ProductService{
		validator: validator,
		processor: processor,
		publisher: publisher,
	}
}

// CreateProduct crea un nuevo producto con validación, procesamiento y publicación de eventos
func (s *ProductService) CreateProduct(ctx context.Context, id, name, description, category string, price float64, stock int) (*entities.Product, error) {
	// 1. Validar los datos de entrada
	if err := s.validator.ValidateCreateProduct(id, name, description, category, price, stock); err != nil {
		return nil, err
	}

	// 2. Procesar la creación del producto
	product, err := s.processor.CreateProduct(ctx, id, name, description, category, price, stock)
	if err != nil {
		return nil, err
	}

	// 3. Publicar evento de producto creado
	if err := s.publisher.PublishProductCreated(ctx, product); err != nil {
		// log.Printf("Failed to publish product created event: %v", err)
	}

	return product, nil
}

// UpdateProduct actualiza un producto existente
func (s *ProductService) UpdateProduct(ctx context.Context, id string, name, description, category *string, price *float64, stock *int) (*entities.Product, error) {
	// 1. Validar los datos de entrada
	if err := s.validator.ValidateUpdateProduct(id, name, description, category, price, stock); err != nil {
		return nil, err
	}

	// 2. Procesar la actualización del producto
	product, err := s.processor.UpdateProduct(ctx, id, name, description, category, price, stock)
	if err != nil {
		return nil, err
	}

	// 3. Publicar evento de producto actualizado
	if err := s.publisher.PublishProductUpdated(ctx, product); err != nil {
		// log.Printf("Failed to publish product updated event: %v", err)
	}

	return product, nil
}

// UpdateStock actualiza el stock de un producto
func (s *ProductService) UpdateStock(ctx context.Context, id string, newStock int) (*entities.Product, error) {
	// 1. Validar el nuevo stock
	if err := s.validator.ValidateUpdateProduct(id, nil, nil, nil, nil, &newStock); err != nil {
		return nil, err
	}

	// 2. Obtener el stock anterior para el evento
	oldProduct, err := s.processor.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	oldStock := oldProduct.Stock

	// 3. Procesar la actualización del stock
	product, err := s.processor.UpdateStock(ctx, id, newStock)
	if err != nil {
		return nil, err
	}

	// 4. Publicar evento de stock actualizado
	if err := s.publisher.PublishStockUpdated(ctx, product, oldStock); err != nil {
		// log.Printf("Failed to publish stock updated event: %v", err)
	}

	return product, nil
}

// AddStock añade stock a un producto
func (s *ProductService) AddStock(ctx context.Context, id string, quantity int) (*entities.Product, error) {
	// 1. Obtener el producto actual
	product, err := s.processor.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	oldStock := product.Stock

	// 2. Procesar la adición de stock
	product, err = s.processor.AddStock(ctx, id, quantity)
	if err != nil {
		return nil, err
	}

	// 3. Publicar evento de stock actualizado
	if err := s.publisher.PublishStockUpdated(ctx, product, oldStock); err != nil {
		// log.Printf("Failed to publish stock updated event: %v", err)
	}

	return product, nil
}

// RemoveStock reduce el stock de un producto
func (s *ProductService) RemoveStock(ctx context.Context, id string, quantity int) (*entities.Product, error) {
	// 1. Obtener el producto actual
	product, err := s.processor.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	oldStock := product.Stock

	// 2. Procesar la reducción de stock
	product, err = s.processor.RemoveStock(ctx, id, quantity)
	if err != nil {
		return nil, err
	}

	// 3. Publicar evento de stock actualizado
	if err := s.publisher.PublishStockUpdated(ctx, product, oldStock); err != nil {
		// log.Printf("Failed to publish stock updated event: %v", err)
	}

	return product, nil
}

// DeactivateProduct desactiva un producto
func (s *ProductService) DeactivateProduct(ctx context.Context, id string) (*entities.Product, error) {
	// 1. Procesar la desactivación del producto
	product, err := s.processor.DeactivateProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. Publicar evento de producto desactivado
	if err := s.publisher.PublishProductDeactivated(ctx, product); err != nil {
		// log.Printf("Failed to publish product deactivated event: %v", err)
	}

	return product, nil
}

// ActivateProduct activa un producto
func (s *ProductService) ActivateProduct(ctx context.Context, id string) (*entities.Product, error) {
	// 1. Procesar la activación del producto
	product, err := s.processor.ActivateProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. Publicar evento de producto activado
	if err := s.publisher.PublishProductActivated(ctx, product); err != nil {
		// log.Printf("Failed to publish product activated event: %v", err)
	}

	return product, nil
}

// GetProduct obtiene un producto por ID
func (s *ProductService) GetProduct(ctx context.Context, id string) (*entities.Product, error) {
	return s.processor.GetProduct(ctx, id)
}

// ListProducts obtiene una lista de productos
func (s *ProductService) ListProducts(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	return s.processor.ListProducts(ctx, limit, offset)
}

// ListAvailableProducts obtiene una lista de productos disponibles
func (s *ProductService) ListAvailableProducts(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	return s.processor.ListAvailableProducts(ctx, limit, offset)
}

// ListProductsByCategory obtiene productos por categoría
func (s *ProductService) ListProductsByCategory(ctx context.Context, category string, limit, offset int) ([]*entities.Product, error) {
	return s.processor.ListProductsByCategory(ctx, category, limit, offset)
}

// ListProductsByPriceRange obtiene productos en un rango de precios
func (s *ProductService) ListProductsByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]*entities.Product, error) {
	return s.processor.ListProductsByPriceRange(ctx, minPrice, maxPrice, limit, offset)
}