package memory

import (
	"context"
	"errors"
	"hexagonal-example/domain/entities"
	"hexagonal-example/domain/repositories"
	"sync"
)

// InMemoryUserRepository implementa UserRepository usando memoria
// Esta es una implementación concreta del patrón Repository
// que almacena los datos en memoria para propósitos de demostración
type InMemoryUserRepository struct {
	users map[string]*entities.User
	mutex sync.RWMutex
}

// NewUserRepository crea una nueva instancia del repositorio en memoria
func NewUserRepository() repositories.UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entities.User),
	}
}

// Save guarda un usuario en el repositorio
func (r *InMemoryUserRepository) Save(ctx context.Context, user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Crear una copia del usuario para evitar modificaciones externas
	userCopy := *user
	r.users[user.ID] = &userCopy
	return nil
}

// FindByID busca un usuario por su ID
func (r *InMemoryUserRepository) FindByID(ctx context.Context, id string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}

	// Retornar una copia para evitar modificaciones externas
	userCopy := *user
	return &userCopy, nil
}

// FindByEmail busca un usuario por su email
func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			// Retornar una copia para evitar modificaciones externas
			userCopy := *user
			return &userCopy, nil
		}
	}

	return nil, nil
}

// FindAll retorna todos los usuarios
func (r *InMemoryUserRepository) FindAll(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	// Aplicar paginación
	start := offset
	end := offset + limit
	if start >= len(users) {
		return []*entities.User{}, nil
	}
	if end > len(users) {
		end = len(users)
	}

	// Retornar copias para evitar modificaciones externas
	result := make([]*entities.User, 0, end-start)
	for i := start; i < end; i++ {
		userCopy := *users[i]
		result = append(result, &userCopy)
	}

	return result, nil
}

// FindActive retorna todos los usuarios activos
func (r *InMemoryUserRepository) FindActive(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	activeUsers := make([]*entities.User, 0)
	for _, user := range r.users {
		if user.IsActive {
			activeUsers = append(activeUsers, user)
		}
	}

	// Aplicar paginación
	start := offset
	end := offset + limit
	if start >= len(activeUsers) {
		return []*entities.User{}, nil
	}
	if end > len(activeUsers) {
		end = len(activeUsers)
	}

	// Retornar copias para evitar modificaciones externas
	result := make([]*entities.User, 0, end-start)
	for i := start; i < end; i++ {
		userCopy := *activeUsers[i]
		result = append(result, &userCopy)
	}

	return result, nil
}

// Delete elimina un usuario del repositorio
func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

// Exists verifica si un usuario existe por ID
func (r *InMemoryUserRepository) Exists(ctx context.Context, id string) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, exists := r.users[id]
	return exists, nil
}

// Count retorna el número total de usuarios
func (r *InMemoryUserRepository) Count(ctx context.Context) (int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return len(r.users), nil
}