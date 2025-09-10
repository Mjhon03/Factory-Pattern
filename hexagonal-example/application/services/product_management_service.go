package services

import (
	"context"
	"hexagonal-example/domain/entities"
	"hexagonal-example/domain/repositories"
)

// ProductManagementService proporciona operaciones complejas de gestión de productos
type ProductManagementService struct {
	productService *ProductService
	productRepo    repositories.ProductRepository
}

// NewProductManagementService crea una nueva instancia del servicio de gestión de productos
func NewProductManagementService(productService *ProductService, productRepo repositories.ProductRepository) *ProductManagementService {
	return &ProductManagementService{
		productService: productService,
		productRepo:    productRepo,
	}
}

// BulkCreateProducts crea múltiples productos en una operación
func (s *ProductManagementService) BulkCreateProducts(ctx context.Context, products []CreateProductRequest) ([]*entities.Product, []error) {
	var createdProducts []*entities.Product
	var errors []error

	for i, req := range products {
		product, err := s.productService.CreateProduct(ctx, req.ID, req.Name, req.Description, req.Category, req.Price, req.Stock)
		if err != nil {
			errors = append(errors, &BulkOperationError{
				Index:   i,
				Message: err.Error(),
			})
		} else {
			createdProducts = append(createdProducts, product)
		}
	}

	return createdProducts, errors
}

// BulkUpdateStock actualiza el stock de múltiples productos
func (s *ProductManagementService) BulkUpdateStock(ctx context.Context, stockUpdates []StockUpdateRequest) ([]*entities.Product, []error) {
	var updatedProducts []*entities.Product
	var errors []error

	for i, req := range stockUpdates {
		product, err := s.productService.UpdateStock(ctx, req.ProductID, req.NewStock)
		if err != nil {
			errors = append(errors, &BulkOperationError{
				Index:   i,
				Message: err.Error(),
			})
		} else {
			updatedProducts = append(updatedProducts, product)
		}
	}

	return updatedProducts, errors
}

// GetProductStatistics obtiene estadísticas de productos
func (s *ProductManagementService) GetProductStatistics(ctx context.Context) (*ProductStatistics, error) {
	// Obtener el total de productos
	totalProducts, err := s.productRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// Obtener productos disponibles
	availableProducts, err := s.productRepo.FindAvailable(ctx, 1000, 0)
	if err != nil {
		return nil, err
	}

	// Calcular productos no disponibles
	unavailableProducts := totalProducts - len(availableProducts)

	return &ProductStatistics{
		TotalProducts:      totalProducts,
		AvailableProducts:  len(availableProducts),
		UnavailableProducts: unavailableProducts,
	}, nil
}

// GetCategoryStatistics obtiene estadísticas por categoría
func (s *ProductManagementService) GetCategoryStatistics(ctx context.Context) (map[string]int, error) {
	// En una implementación real, esto requeriría una consulta más compleja
	// Por simplicidad, retornamos un mapa vacío
	return make(map[string]int), nil
}

// SearchProducts busca productos por diferentes criterios
func (s *ProductManagementService) SearchProducts(ctx context.Context, criteria ProductSearchCriteria) ([]*entities.Product, error) {
	// Implementar lógica de búsqueda más compleja
	if criteria.Category != "" {
		return s.productService.ListProductsByCategory(ctx, criteria.Category, criteria.Limit, criteria.Offset)
	}

	if criteria.MinPrice > 0 || criteria.MaxPrice > 0 {
		minPrice := criteria.MinPrice
		maxPrice := criteria.MaxPrice
		if maxPrice == 0 {
			maxPrice = 1000000 // Valor alto por defecto
		}
		return s.productService.ListProductsByPriceRange(ctx, minPrice, maxPrice, criteria.Limit, criteria.Offset)
	}

	// Si no hay criterios específicos, retornar todos los productos
	return s.productService.ListProducts(ctx, criteria.Limit, criteria.Offset)
}

// CreateProductRequest representa una solicitud para crear un producto
type CreateProductRequest struct {
	ID          string
	Name        string
	Description string
	Category    string
	Price       float64
	Stock       int
}

// StockUpdateRequest representa una solicitud para actualizar stock
type StockUpdateRequest struct {
	ProductID string
	NewStock  int
}

// ProductStatistics contiene estadísticas de productos
type ProductStatistics struct {
	TotalProducts       int
	AvailableProducts   int
	UnavailableProducts int
}

// ProductSearchCriteria define criterios de búsqueda para productos
type ProductSearchCriteria struct {
	Category string
	MinPrice float64
	MaxPrice float64
	Limit    int
	Offset   int
}