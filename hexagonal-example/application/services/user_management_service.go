package services

import (
	"context"
	"errors"
	"hexagonal-example/domain/entities"
	"hexagonal-example/domain/repositories"
)

// UserManagementService proporciona operaciones complejas de gestión de usuarios
// Este servicio combina múltiples servicios para operaciones que requieren
// coordinación entre diferentes partes del sistema
type UserManagementService struct {
	userService *UserService
	userRepo    repositories.UserRepository
}

// NewUserManagementService crea una nueva instancia del servicio de gestión de usuarios
func NewUserManagementService(userService *UserService, userRepo repositories.UserRepository) *UserManagementService {
	return &UserManagementService{
		userService: userService,
		userRepo:    userRepo,
	}
}

// BulkCreateUsers crea múltiples usuarios en una operación
func (s *UserManagementService) BulkCreateUsers(ctx context.Context, users []CreateUserRequest) ([]*entities.User, []error) {
	var createdUsers []*entities.User
	var errors []error

	for i, req := range users {
		user, err := s.userService.CreateUser(ctx, req.ID, req.Email, req.Name)
		if err != nil {
			errors = append(errors, &BulkOperationError{
				Index:   i,
				Message: err.Error(),
			})
		} else {
			createdUsers = append(createdUsers, user)
		}
	}

	return createdUsers, errors
}

// BulkDeactivateUsers desactiva múltiples usuarios
func (s *UserManagementService) BulkDeactivateUsers(ctx context.Context, userIDs []string) ([]*entities.User, []error) {
	var deactivatedUsers []*entities.User
	var errors []error

	for i, userID := range userIDs {
		user, err := s.userService.DeactivateUser(ctx, userID)
		if err != nil {
			errors = append(errors, &BulkOperationError{
				Index:   i,
				Message: err.Error(),
			})
		} else {
			deactivatedUsers = append(deactivatedUsers, user)
		}
	}

	return deactivatedUsers, errors
}

// GetUserStatistics obtiene estadísticas de usuarios
func (s *UserManagementService) GetUserStatistics(ctx context.Context) (*UserStatistics, error) {
	// Obtener el total de usuarios
	totalUsers, err := s.userRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// Obtener usuarios activos
	activeUsers, err := s.userRepo.FindActive(ctx, 1000, 0) // Obtener hasta 1000 usuarios activos
	if err != nil {
		return nil, err
	}

	// Calcular usuarios inactivos
	inactiveUsers := totalUsers - len(activeUsers)

	return &UserStatistics{
		TotalUsers:    totalUsers,
		ActiveUsers:   len(activeUsers),
		InactiveUsers: inactiveUsers,
	}, nil
}

// SearchUsers busca usuarios por diferentes criterios
func (s *UserManagementService) SearchUsers(ctx context.Context, criteria SearchCriteria) ([]*entities.User, error) {
	// Implementar lógica de búsqueda más compleja
	// Por simplicidad, aquí solo implementamos búsqueda por email
	if criteria.Email != "" {
		user, err := s.userService.GetUserByEmail(ctx, criteria.Email)
		if err != nil {
			return nil, err
		}
		return []*entities.User{user}, nil
	}

	// Si no hay criterios específicos, retornar todos los usuarios
	return s.userService.ListUsers(ctx, criteria.Limit, criteria.Offset)
}

// CreateUserRequest representa una solicitud para crear un usuario
type CreateUserRequest struct {
	ID    string
	Email string
	Name  string
}

// UserStatistics contiene estadísticas de usuarios
type UserStatistics struct {
	TotalUsers    int
	ActiveUsers   int
	InactiveUsers int
}

// SearchCriteria define criterios de búsqueda para usuarios
type SearchCriteria struct {
	Email  string
	Limit  int
	Offset int
}

// BulkOperationError representa un error en una operación en lote
type BulkOperationError struct {
	Index   int
	Message string
}

func (e *BulkOperationError) Error() string {
	return errors.New(e.Message).Error()
}