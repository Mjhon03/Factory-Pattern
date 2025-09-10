package services

import (
	"context"
	"errors"
	"hexagonal-example/domain/entities"
	"hexagonal-example/domain/repositories"
)

// UserProcessor se encarga del procesamiento de la lógica de negocio de usuarios
// Esta clase maneja la creación, actualización y manipulación de entidades de usuario
// sin preocuparse por la validación (que hace UserValidator) ni la publicación de eventos
type UserProcessor struct {
	userRepo repositories.UserRepository
}

// NewUserProcessor crea una nueva instancia del procesador de usuarios
func NewUserProcessor(userRepo repositories.UserRepository) *UserProcessor {
	return &UserProcessor{
		userRepo: userRepo,
	}
}

// CreateUser crea un nuevo usuario en el sistema
func (p *UserProcessor) CreateUser(ctx context.Context, id, email, name string) (*entities.User, error) {
	// Verificar si el usuario ya existe
	exists, err := p.userRepo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	// Verificar si el email ya está en uso
	existingUser, err := p.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Crear la entidad de usuario
	user, err := entities.NewUser(id, email, name)
	if err != nil {
		return nil, err
	}

	// Guardar en el repositorio
	if err := p.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser actualiza un usuario existente
func (p *UserProcessor) UpdateUser(ctx context.Context, id string, email, name *string) (*entities.User, error) {
	// Buscar el usuario existente
	user, err := p.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Actualizar email si se proporciona
	if email != nil {
		// Verificar si el nuevo email ya está en uso por otro usuario
		existingUser, err := p.userRepo.FindByEmail(ctx, *email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already in use by another user")
		}

		if err := user.UpdateEmail(*email); err != nil {
			return nil, err
		}
	}

	// Actualizar nombre si se proporciona
	if name != nil {
		if err := user.UpdateName(*name); err != nil {
			return nil, err
		}
	}

	// Guardar los cambios
	if err := p.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeactivateUser desactiva un usuario
func (p *UserProcessor) DeactivateUser(ctx context.Context, id string) (*entities.User, error) {
	user, err := p.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Deactivate()
	if err := p.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// ActivateUser activa un usuario
func (p *UserProcessor) ActivateUser(ctx context.Context, id string) (*entities.User, error) {
	user, err := p.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Activate()
	if err := p.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser obtiene un usuario por ID
func (p *UserProcessor) GetUser(ctx context.Context, id string) (*entities.User, error) {
	user, err := p.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetUserByEmail obtiene un usuario por email
func (p *UserProcessor) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := p.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// ListUsers obtiene una lista de usuarios
func (p *UserProcessor) ListUsers(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	return p.userRepo.FindAll(ctx, limit, offset)
}

// ListActiveUsers obtiene una lista de usuarios activos
func (p *UserProcessor) ListActiveUsers(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	return p.userRepo.FindActive(ctx, limit, offset)
}