package main

import (
	"context"
	"fmt"
	"testing"
	"hexagonal-example/infrastructure/config"
)

// TestExample demuestra cómo probar la arquitectura hexagonal
// Este es un ejemplo básico de testing unitario
func TestExample(t *testing.T) {
	// Crear el contenedor de dependencias
	container := config.NewContainer()
	
	// Obtener el servicio de usuario
	userService := container.GetUserService()
	
	// Crear un contexto
	ctx := context.Background()
	
	// Crear un usuario
	user, err := userService.CreateUser(ctx, "test-user", "test@example.com", "Test User")
	if err != nil {
		t.Fatalf("Error creando usuario: %v", err)
	}
	
	// Verificar que el usuario se creó correctamente
	if user.ID != "test-user" {
		t.Errorf("Expected user ID 'test-user', got '%s'", user.ID)
	}
	
	if user.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", user.Email)
	}
	
	if user.Name != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", user.Name)
	}
	
	if !user.IsActive {
		t.Error("Expected user to be active")
	}
	
	// Obtener el usuario
	retrievedUser, err := userService.GetUser(ctx, "test-user")
	if err != nil {
		t.Fatalf("Error obteniendo usuario: %v", err)
	}
	
	// Verificar que es el mismo usuario
	if retrievedUser.ID != user.ID {
		t.Error("Retrieved user ID doesn't match")
	}
}

// TestProductService demuestra el testing del servicio de productos
func TestProductService(t *testing.T) {
	container := config.NewContainer()
	productService := container.GetProductService()
	ctx := context.Background()
	
	// Crear un producto
	product, err := productService.CreateProduct(ctx, "test-product", "Test Product", "Test Description", "Test Category", 99.99, 10)
	if err != nil {
		t.Fatalf("Error creando producto: %v", err)
	}
	
	// Verificar el producto
	if product.ID != "test-product" {
		t.Errorf("Expected product ID 'test-product', got '%s'", product.ID)
	}
	
	if product.Price != 99.99 {
		t.Errorf("Expected price 99.99, got %.2f", product.Price)
	}
	
	if product.Stock != 10 {
		t.Errorf("Expected stock 10, got %d", product.Stock)
	}
	
	// Actualizar el stock
	updatedProduct, err := productService.AddStock(ctx, "test-product", 5)
	if err != nil {
		t.Fatalf("Error actualizando stock: %v", err)
	}
	
	// Verificar que el stock se actualizó
	if updatedProduct.Stock != 15 {
		t.Errorf("Expected stock 15, got %d", updatedProduct.Stock)
	}
}

// TestValidation demuestra el testing de validación
func TestValidation(t *testing.T) {
	container := config.NewContainer()
	userService := container.GetUserService()
	ctx := context.Background()
	
	// Intentar crear un usuario con datos inválidos
	_, err := userService.CreateUser(ctx, "", "invalid-email", "")
	if err == nil {
		t.Error("Expected validation error for invalid user data")
	}
	
	// Intentar crear un usuario con email inválido
	_, err = userService.CreateUser(ctx, "valid-id", "invalid-email", "Valid Name")
	if err == nil {
		t.Error("Expected validation error for invalid email")
	}
}

// TestEventSystem demuestra el testing del sistema de eventos
func TestEventSystem(t *testing.T) {
	container := config.NewContainer()
	userService := container.GetUserService()
	ctx := context.Background()
	
	// Crear un usuario (esto debería disparar un evento)
	_, err := userService.CreateUser(ctx, "event-user", "event@example.com", "Event User")
	if err != nil {
		t.Fatalf("Error creando usuario: %v", err)
	}
	
	// En un test real, podrías verificar que el evento se publicó
	// Por ejemplo, usando un mock del event bus
}

// BenchmarkUserCreation mide el rendimiento de la creación de usuarios
func BenchmarkUserCreation(b *testing.B) {
	container := config.NewContainer()
	userService := container.GetUserService()
	ctx := context.Background()
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		userID := fmt.Sprintf("bench-user-%d", i)
		email := fmt.Sprintf("bench%d@example.com", i)
		name := fmt.Sprintf("Bench User %d", i)
		
		_, err := userService.CreateUser(ctx, userID, email, name)
		if err != nil {
			b.Fatalf("Error creando usuario: %v", err)
		}
	}
}

// TestConcurrentAccess demuestra el testing de acceso concurrente
func TestConcurrentAccess(t *testing.T) {
	container := config.NewContainer()
	userService := container.GetUserService()
	ctx := context.Background()
	
	// Crear múltiples usuarios concurrentemente
	done := make(chan bool, 10)
	
	for i := 0; i < 10; i++ {
		go func(i int) {
			userID := fmt.Sprintf("concurrent-user-%d", i)
			email := fmt.Sprintf("concurrent%d@example.com", i)
			name := fmt.Sprintf("Concurrent User %d", i)
			
			_, err := userService.CreateUser(ctx, userID, email, name)
			if err != nil {
				t.Errorf("Error creando usuario concurrente: %v", err)
			}
			
			done <- true
		}(i)
	}
	
	// Esperar a que terminen todas las goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
	
	// Verificar que se crearon todos los usuarios
	users, err := userService.ListUsers(ctx, 20, 0)
	if err != nil {
		t.Fatalf("Error listando usuarios: %v", err)
	}
	
	if len(users) < 10 {
		t.Errorf("Expected at least 10 users, got %d", len(users))
	}
}