package events

import "time"

// UserCreatedEvent representa el evento cuando se crea un usuario
type UserCreatedEvent struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// UserUpdatedEvent representa el evento cuando se actualiza un usuario
type UserUpdatedEvent struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserDeactivatedEvent representa el evento cuando se desactiva un usuario
type UserDeactivatedEvent struct {
	UserID         string    `json:"user_id"`
	Email          string    `json:"email"`
	DeactivatedAt  time.Time `json:"deactivated_at"`
}

// UserActivatedEvent representa el evento cuando se activa un usuario
type UserActivatedEvent struct {
	UserID       string    `json:"user_id"`
	Email        string    `json:"email"`
	ActivatedAt  time.Time `json:"activated_at"`
}