package services

import (
	"errors"
	"regexp"
	"strings"
)

// UserValidator se encarga únicamente de la validación de datos de usuario
// Esta es una responsabilidad específica que se separa del procesamiento
// y la persistencia de datos
type UserValidator struct {
	emailRegex *regexp.Regexp
}

// NewUserValidator crea una nueva instancia del validador de usuarios
func NewUserValidator() *UserValidator {
	// Compilar la expresión regular para validar emails una sola vez
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	return &UserValidator{
		emailRegex: emailRegex,
	}
}

// ValidateCreateUser valida los datos para crear un nuevo usuario
func (v *UserValidator) ValidateCreateUser(id, email, name string) error {
	// Validar ID
	if err := v.validateID(id); err != nil {
		return err
	}

	// Validar email
	if err := v.validateEmail(email); err != nil {
		return err
	}

	// Validar nombre
	if err := v.validateName(name); err != nil {
		return err
	}

	return nil
}

// ValidateUpdateUser valida los datos para actualizar un usuario
func (v *UserValidator) ValidateUpdateUser(id string, email, name *string) error {
	// El ID siempre debe ser válido
	if err := v.validateID(id); err != nil {
		return err
	}

	// Validar email si se proporciona
	if email != nil {
		if err := v.validateEmail(*email); err != nil {
			return err
		}
	}

	// Validar nombre si se proporciona
	if name != nil {
		if err := v.validateName(*name); err != nil {
			return err
		}
	}

	return nil
}

// validateID valida el ID del usuario
func (v *UserValidator) validateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("user ID cannot be empty")
	}
	if len(id) < 3 {
		return errors.New("user ID must be at least 3 characters long")
	}
	if len(id) > 50 {
		return errors.New("user ID cannot exceed 50 characters")
	}
	return nil
}

// validateEmail valida el email del usuario
func (v *UserValidator) validateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(email) > 255 {
		return errors.New("email cannot exceed 255 characters")
	}
	if !v.emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// validateName valida el nombre del usuario
func (v *UserValidator) validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	if len(name) > 100 {
		return errors.New("name cannot exceed 100 characters")
	}
	return nil
}