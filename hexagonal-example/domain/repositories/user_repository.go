package repositories

import (
	"context"
	"hexagonal-example/domain/entities"
)

// UserRepository define la interfaz para el repositorio de usuarios
// Esta interfaz está en el dominio y define el contrato que deben cumplir
// todas las implementaciones concretas (infraestructura)
//
// El patrón Repository encapsula la lógica de acceso a datos y proporciona
// una interfaz más orientada a objetos para acceder a la capa de persistencia
type UserRepository interface {
	// Save guarda un usuario en el repositorio
	// Si el usuario ya existe, lo actualiza; si no, lo crea
	Save(ctx context.Context, user *entities.User) error

	// FindByID busca un usuario por su ID
	// Retorna nil si no se encuentra
	FindByID(ctx context.Context, id string) (*entities.User, error)

	// FindByEmail busca un usuario por su email
	// Retorna nil si no se encuentra
	FindByEmail(ctx context.Context, email string) (*entities.User, error)

	// FindAll retorna todos los usuarios
	// Puede incluir filtros opcionales
	FindAll(ctx context.Context, limit, offset int) ([]*entities.User, error)

	// FindActive retorna todos los usuarios activos
	FindActive(ctx context.Context, limit, offset int) ([]*entities.User, error)

	// Delete elimina un usuario del repositorio
	Delete(ctx context.Context, id string) error

	// Exists verifica si un usuario existe por ID
	Exists(ctx context.Context, id string) (bool, error)

	// Count retorna el número total de usuarios
	Count(ctx context.Context) (int, error)
}

// UserRepositoryError define errores específicos del repositorio de usuarios
type UserRepositoryError struct {
	Message string
	Err     error
}

func (e *UserRepositoryError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *UserRepositoryError) Unwrap() error {
	return e.Err
}

// Errores comunes del repositorio
var (
	ErrUserNotFound    = &UserRepositoryError{Message: "user not found"}
	ErrUserAlreadyExists = &UserRepositoryError{Message: "user already exists"}
	ErrInvalidUserData   = &UserRepositoryError{Message: "invalid user data"}
)