package services

import (
	"context"
	"hexagonal-example/domain/entities"
)

// UserService es el servicio principal que orquesta los servicios granulares
// Este servicio combina la validación, procesamiento y publicación de eventos
// para proporcionar una interfaz unificada para las operaciones de usuario
type UserService struct {
	validator  *UserValidator
	processor  *UserProcessor
	publisher  *UserEventPublisher
}

// NewUserService crea una nueva instancia del servicio de usuario
// Recibe las dependencias de los servicios granulares
func NewUserService(validator *UserValidator, processor *UserProcessor, publisher *UserEventPublisher) *UserService {
	return &UserService{
		validator: validator,
		processor: processor,
		publisher: publisher,
	}
}

// CreateUser crea un nuevo usuario con validación, procesamiento y publicación de eventos
func (s *UserService) CreateUser(ctx context.Context, id, email, name string) (*entities.User, error) {
	// 1. Validar los datos de entrada
	if err := s.validator.ValidateCreateUser(id, email, name); err != nil {
		return nil, err
	}

	// 2. Procesar la creación del usuario
	user, err := s.processor.CreateUser(ctx, id, email, name)
	if err != nil {
		return nil, err
	}

	// 3. Publicar evento de usuario creado
	if err := s.publisher.PublishUserCreated(ctx, user); err != nil {
		// Log del error pero no fallar la operación
		// En un sistema real, podrías querer implementar un mecanismo de retry
		// o guardar el evento en una cola para procesamiento posterior
		// log.Printf("Failed to publish user created event: %v", err)
	}

	return user, nil
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(ctx context.Context, id string, email, name *string) (*entities.User, error) {
	// 1. Validar los datos de entrada
	if err := s.validator.ValidateUpdateUser(id, email, name); err != nil {
		return nil, err
	}

	// 2. Procesar la actualización del usuario
	user, err := s.processor.UpdateUser(ctx, id, email, name)
	if err != nil {
		return nil, err
	}

	// 3. Publicar evento de usuario actualizado
	if err := s.publisher.PublishUserUpdated(ctx, user); err != nil {
		// Log del error pero no fallar la operación
		// log.Printf("Failed to publish user updated event: %v", err)
	}

	return user, nil
}

// DeactivateUser desactiva un usuario
func (s *UserService) DeactivateUser(ctx context.Context, id string) (*entities.User, error) {
	// 1. Procesar la desactivación del usuario
	user, err := s.processor.DeactivateUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. Publicar evento de usuario desactivado
	if err := s.publisher.PublishUserDeactivated(ctx, user); err != nil {
		// log.Printf("Failed to publish user deactivated event: %v", err)
	}

	return user, nil
}

// ActivateUser activa un usuario
func (s *UserService) ActivateUser(ctx context.Context, id string) (*entities.User, error) {
	// 1. Procesar la activación del usuario
	user, err := s.processor.ActivateUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. Publicar evento de usuario activado
	if err := s.publisher.PublishUserActivated(ctx, user); err != nil {
		// log.Printf("Failed to publish user activated event: %v", err)
	}

	return user, nil
}

// GetUser obtiene un usuario por ID
func (s *UserService) GetUser(ctx context.Context, id string) (*entities.User, error) {
	return s.processor.GetUser(ctx, id)
}

// GetUserByEmail obtiene un usuario por email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return s.processor.GetUserByEmail(ctx, email)
}

// ListUsers obtiene una lista de usuarios
func (s *UserService) ListUsers(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	return s.processor.ListUsers(ctx, limit, offset)
}

// ListActiveUsers obtiene una lista de usuarios activos
func (s *UserService) ListActiveUsers(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	return s.processor.ListActiveUsers(ctx, limit, offset)
}