package services

import (
	"context"
	"hexagonal-example/domain/entities"
	"hexagonal-example/infrastructure/events"
)

// ProductEventPublisher se encarga únicamente de publicar eventos relacionados con productos
type ProductEventPublisher struct {
	eventBus events.EventBus
}

// NewProductEventPublisher crea una nueva instancia del publicador de eventos de producto
func NewProductEventPublisher(eventBus events.EventBus) *ProductEventPublisher {
	return &ProductEventPublisher{
		eventBus: eventBus,
	}
}

// PublishProductCreated publica un evento cuando se crea un producto
func (p *ProductEventPublisher) PublishProductCreated(ctx context.Context, product *entities.Product) error {
	event := events.ProductCreatedEvent{
		ProductID:  product.ID,
		Name:       product.Name,
		Category:   product.Category,
		Price:      product.Price,
		Stock:      product.Stock,
		CreatedAt:  product.CreatedAt,
	}

	return p.eventBus.Publish(ctx, "product.created", event)
}

// PublishProductUpdated publica un evento cuando se actualiza un producto
func (p *ProductEventPublisher) PublishProductUpdated(ctx context.Context, product *entities.Product) error {
	event := events.ProductUpdatedEvent{
		ProductID:  product.ID,
		Name:       product.Name,
		Category:   product.Category,
		Price:      product.Price,
		Stock:      product.Stock,
		UpdatedAt:  product.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "product.updated", event)
}

// PublishStockUpdated publica un evento cuando se actualiza el stock de un producto
func (p *ProductEventPublisher) PublishStockUpdated(ctx context.Context, product *entities.Product, oldStock int) error {
	event := events.StockUpdatedEvent{
		ProductID:  product.ID,
		Name:       product.Name,
		OldStock:   oldStock,
		NewStock:   product.Stock,
		UpdatedAt:  product.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "product.stock.updated", event)
}

// PublishProductDeactivated publica un evento cuando se desactiva un producto
func (p *ProductEventPublisher) PublishProductDeactivated(ctx context.Context, product *entities.Product) error {
	event := events.ProductDeactivatedEvent{
		ProductID:      product.ID,
		Name:           product.Name,
		DeactivatedAt:  product.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "product.deactivated", event)
}

// PublishProductActivated publica un evento cuando se activa un producto
func (p *ProductEventPublisher) PublishProductActivated(ctx context.Context, product *entities.Product) error {
	event := events.ProductActivatedEvent{
		ProductID:    product.ID,
		Name:         product.Name,
		ActivatedAt:  product.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "product.activated", event)
}