package services

import (
	"context"
	"hexagonal-example/domain/entities"
	"hexagonal-example/infrastructure/events"
)

// UserEventPublisher se encarga únicamente de publicar eventos relacionados con usuarios
// Esta separación permite que el procesamiento de la lógica de negocio no se vea
// afectado por problemas en la publicación de eventos
type UserEventPublisher struct {
	eventBus events.EventBus
}

// NewUserEventPublisher crea una nueva instancia del publicador de eventos de usuario
func NewUserEventPublisher(eventBus events.EventBus) *UserEventPublisher {
	return &UserEventPublisher{
		eventBus: eventBus,
	}
}

// PublishUserCreated publica un evento cuando se crea un usuario
func (p *UserEventPublisher) PublishUserCreated(ctx context.Context, user *entities.User) error {
	event := events.UserCreatedEvent{
		UserID:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return p.eventBus.Publish(ctx, "user.created", event)
}

// PublishUserUpdated publica un evento cuando se actualiza un usuario
func (p *UserEventPublisher) PublishUserUpdated(ctx context.Context, user *entities.User) error {
	event := events.UserUpdatedEvent{
		UserID:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "user.updated", event)
}

// PublishUserDeactivated publica un evento cuando se desactiva un usuario
func (p *UserEventPublisher) PublishUserDeactivated(ctx context.Context, user *entities.User) error {
	event := events.UserDeactivatedEvent{
		UserID:    user.ID,
		Email:     user.Email,
		DeactivatedAt: user.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "user.deactivated", event)
}

// PublishUserActivated publica un evento cuando se activa un usuario
func (p *UserEventPublisher) PublishUserActivated(ctx context.Context, user *entities.User) error {
	event := events.UserActivatedEvent{
		UserID:    user.ID,
		Email:     user.Email,
		ActivatedAt: user.UpdatedAt,
	}

	return p.eventBus.Publish(ctx, "user.activated", event)
}