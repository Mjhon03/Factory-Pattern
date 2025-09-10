package events

import (
	"context"
	"sync"
)

// EventBus define la interfaz para el bus de eventos
// Este patrón permite desacoplar la publicación de eventos de su procesamiento
type EventBus interface {
	// Publish publica un evento en el bus
	Publish(ctx context.Context, eventType string, event interface{}) error

	// Subscribe suscribe un handler a un tipo de evento
	Subscribe(eventType string, handler EventHandler)

	// Unsubscribe desuscribe un handler de un tipo de evento
	Unsubscribe(eventType string, handler EventHandler)
}

// EventHandler define la interfaz para manejar eventos
type EventHandler interface {
	Handle(ctx context.Context, event interface{}) error
}

// EventHandlerFunc es una función que implementa EventHandler
type EventHandlerFunc func(ctx context.Context, event interface{}) error

// Handle implementa EventHandler para EventHandlerFunc
func (f EventHandlerFunc) Handle(ctx context.Context, event interface{}) error {
	return f(ctx, event)
}

// InMemoryEventBus implementa EventBus usando memoria
// Esta es una implementación simple para propósitos de demostración
type InMemoryEventBus struct {
	handlers map[string][]EventHandler
	mutex    sync.RWMutex
}

// NewInMemoryEventBus crea una nueva instancia del event bus en memoria
func NewInMemoryEventBus() EventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]EventHandler),
	}
}

// Publish publica un evento en el bus
func (b *InMemoryEventBus) Publish(ctx context.Context, eventType string, event interface{}) error {
	b.mutex.RLock()
	handlers := b.handlers[eventType]
	b.mutex.RUnlock()

	// Ejecutar todos los handlers para este tipo de evento
	for _, handler := range handlers {
		if err := handler.Handle(ctx, event); err != nil {
			// En un sistema real, podrías querer loggear el error
			// pero no fallar la publicación del evento
			// log.Printf("Error handling event %s: %v", eventType, err)
		}
	}

	return nil
}

// Subscribe suscribe un handler a un tipo de evento
func (b *InMemoryEventBus) Subscribe(eventType string, handler EventHandler) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

// Unsubscribe desuscribe un handler de un tipo de evento
func (b *InMemoryEventBus) Unsubscribe(eventType string, handler EventHandler) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	handlers := b.handlers[eventType]
	for i, h := range handlers {
		if h == handler {
			// Remover el handler de la lista
			b.handlers[eventType] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
}