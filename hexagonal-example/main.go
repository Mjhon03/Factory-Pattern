package main

import (
	"context"
	"fmt"
	"log"
	"hexagonal-example/infrastructure/config"
	"hexagonal-example/infrastructure/events"
)

func main() {
	fmt.Println("=== Ejemplo de Arquitectura Hexagonal con Patrones Repository y Factory ===")
	fmt.Println()

	// Crear el contenedor de dependencias
	// Este es el punto de entrada principal donde se configuran todas las dependencias
	container := config.NewContainer()

	// Configurar event handlers para demostrar el sistema de eventos
	setupEventHandlers(container)

	// Ejecutar ejemplos
	runUserExamples(container)
	runProductExamples(container)
	runManagementExamples(container)
}

// setupEventHandlers configura los handlers de eventos para demostrar el sistema
func setupEventHandlers(container *config.Container) {
	eventBus := container.GetEventBus()

	// Handler para eventos de usuario
	eventBus.Subscribe("user.created", events.EventHandlerFunc(func(ctx context.Context, event interface{}) error {
		if userEvent, ok := event.(events.UserCreatedEvent); ok {
			fmt.Printf("📧 Evento: Usuario creado - ID: %s, Email: %s, Nombre: %s\n", 
				userEvent.UserID, userEvent.Email, userEvent.Name)
		}
		return nil
	}))

	eventBus.Subscribe("user.updated", events.EventHandlerFunc(func(ctx context.Context, event interface{}) error {
		if userEvent, ok := event.(events.UserUpdatedEvent); ok {
			fmt.Printf("📧 Evento: Usuario actualizado - ID: %s, Email: %s, Nombre: %s\n", 
				userEvent.UserID, userEvent.Email, userEvent.Name)
		}
		return nil
	}))

	// Handler para eventos de producto
	eventBus.Subscribe("product.created", events.EventHandlerFunc(func(ctx context.Context, event interface{}) error {
		if productEvent, ok := event.(events.ProductCreatedEvent); ok {
			fmt.Printf("📦 Evento: Producto creado - ID: %s, Nombre: %s, Precio: $%.2f\n", 
				productEvent.ProductID, productEvent.Name, productEvent.Price)
		}
		return nil
	}))

	eventBus.Subscribe("product.stock.updated", events.EventHandlerFunc(func(ctx context.Context, event interface{}) error {
		if stockEvent, ok := event.(events.StockUpdatedEvent); ok {
			fmt.Printf("📦 Evento: Stock actualizado - Producto: %s, Stock anterior: %d, Stock nuevo: %d\n", 
				stockEvent.Name, stockEvent.OldStock, stockEvent.NewStock)
		}
		return nil
	}))
}

// runUserExamples demuestra el uso del servicio de usuarios
func runUserExamples(container *config.Container) {
	fmt.Println("=== EJEMPLOS DE USUARIOS ===")
	
	ctx := context.Background()
	userService := container.GetUserService()

	// Crear usuarios
	fmt.Println("\n1. Creando usuarios...")
	user1, err := userService.CreateUser(ctx, "user1", "juan@example.com", "Juan Pérez")
	if err != nil {
		log.Printf("Error creando usuario 1: %v", err)
	} else {
		fmt.Printf("✅ Usuario creado: %s (%s)\n", user1.Name, user1.Email)
	}

	user2, err := userService.CreateUser(ctx, "user2", "maria@example.com", "María García")
	if err != nil {
		log.Printf("Error creando usuario 2: %v", err)
	} else {
		fmt.Printf("✅ Usuario creado: %s (%s)\n", user2.Name, user2.Email)
	}

	// Actualizar usuario
	fmt.Println("\n2. Actualizando usuario...")
	updatedUser, err := userService.UpdateUser(ctx, "user1", nil, stringPtr("Juan Carlos Pérez"))
	if err != nil {
		log.Printf("Error actualizando usuario: %v", err)
	} else {
		fmt.Printf("✅ Usuario actualizado: %s\n", updatedUser.Name)
	}

	// Obtener usuario
	fmt.Println("\n3. Obteniendo usuario...")
	retrievedUser, err := userService.GetUser(ctx, "user1")
	if err != nil {
		log.Printf("Error obteniendo usuario: %v", err)
	} else {
		fmt.Printf("✅ Usuario obtenido: %s (%s)\n", retrievedUser.Name, retrievedUser.Email)
	}

	// Listar usuarios
	fmt.Println("\n4. Listando usuarios...")
	users, err := userService.ListUsers(ctx, 10, 0)
	if err != nil {
		log.Printf("Error listando usuarios: %v", err)
	} else {
		fmt.Printf("✅ Total de usuarios: %d\n", len(users))
		for _, user := range users {
			fmt.Printf("   - %s (%s) - Activo: %t\n", user.Name, user.Email, user.IsActive)
		}
	}

	// Desactivar usuario
	fmt.Println("\n5. Desactivando usuario...")
	deactivatedUser, err := userService.DeactivateUser(ctx, "user2")
	if err != nil {
		log.Printf("Error desactivando usuario: %v", err)
	} else {
		fmt.Printf("✅ Usuario desactivado: %s\n", deactivatedUser.Name)
	}
}

// runProductExamples demuestra el uso del servicio de productos
func runProductExamples(container *config.Container) {
	fmt.Println("\n=== EJEMPLOS DE PRODUCTOS ===")
	
	ctx := context.Background()
	productService := container.GetProductService()

	// Crear productos
	fmt.Println("\n1. Creando productos...")
	product1, err := productService.CreateProduct(ctx, "prod1", "Laptop Gaming", "Laptop de alto rendimiento para gaming", "Electrónicos", 1299.99, 10)
	if err != nil {
		log.Printf("Error creando producto 1: %v", err)
	} else {
		fmt.Printf("✅ Producto creado: %s - $%.2f (Stock: %d)\n", product1.Name, product1.Price, product1.Stock)
	}

	product2, err := productService.CreateProduct(ctx, "prod2", "Mouse Inalámbrico", "Mouse inalámbrico ergonómico", "Accesorios", 29.99, 50)
	if err != nil {
		log.Printf("Error creando producto 2: %v", err)
	} else {
		fmt.Printf("✅ Producto creado: %s - $%.2f (Stock: %d)\n", product2.Name, product2.Price, product2.Stock)
	}

	// Actualizar precio
	fmt.Println("\n2. Actualizando precio...")
	updatedProduct, err := productService.UpdateProduct(ctx, "prod1", nil, nil, nil, float64Ptr(1199.99), nil)
	if err != nil {
		log.Printf("Error actualizando producto: %v", err)
	} else {
		fmt.Printf("✅ Precio actualizado: %s - $%.2f\n", updatedProduct.Name, updatedProduct.Price)
	}

	// Actualizar stock
	fmt.Println("\n3. Actualizando stock...")
	stockUpdatedProduct, err := productService.AddStock(ctx, "prod1", 5)
	if err != nil {
		log.Printf("Error actualizando stock: %v", err)
	} else {
		fmt.Printf("✅ Stock actualizado: %s - Stock: %d\n", stockUpdatedProduct.Name, stockUpdatedProduct.Stock)
	}

	// Listar productos disponibles
	fmt.Println("\n4. Listando productos disponibles...")
	availableProducts, err := productService.ListAvailableProducts(ctx, 10, 0)
	if err != nil {
		log.Printf("Error listando productos: %v", err)
	} else {
		fmt.Printf("✅ Productos disponibles: %d\n", len(availableProducts))
		for _, product := range availableProducts {
			fmt.Printf("   - %s - $%.2f (Stock: %d) - Disponible: %t\n", 
				product.Name, product.Price, product.Stock, product.IsAvailable())
		}
	}

	// Buscar por categoría
	fmt.Println("\n5. Buscando productos por categoría...")
	electronics, err := productService.ListProductsByCategory(ctx, "Electrónicos", 10, 0)
	if err != nil {
		log.Printf("Error buscando por categoría: %v", err)
	} else {
		fmt.Printf("✅ Productos en Electrónicos: %d\n", len(electronics))
		for _, product := range electronics {
			fmt.Printf("   - %s - $%.2f\n", product.Name, product.Price)
		}
	}
}

// runManagementExamples demuestra el uso de los servicios de gestión
func runManagementExamples(container *config.Container) {
	fmt.Println("\n=== EJEMPLOS DE SERVICIOS DE GESTIÓN ===")
	
	ctx := context.Background()
	userManagementService := container.GetUserManagementService()
	productManagementService := container.GetProductManagementService()

	// Estadísticas de usuarios
	fmt.Println("\n1. Estadísticas de usuarios...")
	userStats, err := userManagementService.GetUserStatistics(ctx)
	if err != nil {
		log.Printf("Error obteniendo estadísticas de usuarios: %v", err)
	} else {
		fmt.Printf("✅ Estadísticas de usuarios:\n")
		fmt.Printf("   - Total: %d\n", userStats.TotalUsers)
		fmt.Printf("   - Activos: %d\n", userStats.ActiveUsers)
		fmt.Printf("   - Inactivos: %d\n", userStats.InactiveUsers)
	}

	// Estadísticas de productos
	fmt.Println("\n2. Estadísticas de productos...")
	productStats, err := productManagementService.GetProductStatistics(ctx)
	if err != nil {
		log.Printf("Error obteniendo estadísticas de productos: %v", err)
	} else {
		fmt.Printf("✅ Estadísticas de productos:\n")
		fmt.Printf("   - Total: %d\n", productStats.TotalProducts)
		fmt.Printf("   - Disponibles: %d\n", productStats.AvailableProducts)
		fmt.Printf("   - No disponibles: %d\n", productStats.UnavailableProducts)
	}

	// Operaciones en lote
	fmt.Println("\n3. Creación en lote de usuarios...")
	bulkUsers := []struct {
		ID    string
		Email string
		Name  string
	}{
		{"user3", "ana@example.com", "Ana López"},
		{"user4", "carlos@example.com", "Carlos Rodríguez"},
		{"user5", "laura@example.com", "Laura Martínez"},
	}

	createRequests := make([]interface{}, len(bulkUsers))
	for i, user := range bulkUsers {
		createRequests[i] = struct {
			ID    string
			Email string
			Name  string
		}{user.ID, user.Email, user.Name}
	}

	// Nota: En una implementación real, necesitarías convertir esto correctamente
	// Por simplicidad, creamos usuarios individualmente
	userService := container.GetUserService()
	for _, user := range bulkUsers {
		_, err := userService.CreateUser(ctx, user.ID, user.Email, user.Name)
		if err != nil {
			log.Printf("Error creando usuario %s: %v", user.ID, err)
		} else {
			fmt.Printf("✅ Usuario creado en lote: %s\n", user.Name)
		}
	}
}

// Funciones auxiliares para crear punteros
func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}