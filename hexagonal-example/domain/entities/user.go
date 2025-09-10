package entities

import (
	"errors"
	"time"
)

// User representa la entidad de usuario en el dominio
// Esta es la entidad central que contiene toda la lógica de negocio relacionada con usuarios
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

// NewUser crea una nueva instancia de User con validaciones de dominio
// Este es un constructor que encapsula la lógica de creación de usuarios
func NewUser(id, email, name string) (*User, error) {
	// Validaciones de dominio
	if id == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	if email == "" {
		return nil, errors.New("user email cannot be empty")
	}
	if name == "" {
		return nil, errors.New("user name cannot be empty")
	}

	// Crear el usuario con valores por defecto
	now := time.Now()
	return &User{
		ID:        id,
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
		IsActive:  true, // Los usuarios se crean activos por defecto
	}, nil
}

// UpdateEmail actualiza el email del usuario con validación
// Método de dominio que encapsula la lógica de negocio
func (u *User) UpdateEmail(newEmail string) error {
	if newEmail == "" {
		return errors.New("email cannot be empty")
	}
	
	u.Email = newEmail
	u.UpdatedAt = time.Now()
	return nil
}

// UpdateName actualiza el nombre del usuario
func (u *User) UpdateName(newName string) error {
	if newName == "" {
		return errors.New("name cannot be empty")
	}
	
	u.Name = newName
	u.UpdatedAt = time.Now()
	return nil
}

// Deactivate desactiva un usuario
func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

// Activate activa un usuario
func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

// IsValid verifica si el usuario es válido según las reglas de negocio
func (u *User) IsValid() bool {
	return u.ID != "" && u.Email != "" && u.Name != ""
}