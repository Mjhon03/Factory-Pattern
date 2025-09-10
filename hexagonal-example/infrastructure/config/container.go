package config

import (
	"hexagonal-example/application/factories"
	"hexagonal-example/application/services"
	"hexagonal-example/domain/repositories"
	"hexagonal-example/infrastructure/events"
	"hexagonal-example/infrastructure/repositories/memory"
)

// Container implementa el patrón de Dependency Injection
// Este contenedor se encarga de crear y configurar todas las dependencias
// de la aplicación de manera centralizada
type Container struct {
	// Repositorios
	userRepo    repositories.UserRepository
	productRepo repositories.ProductRepository

	// Event Bus
	eventBus events.EventBus

	// Factory
	serviceFactory *factories.ServiceFactory

	// Servicios (lazy-loaded)
	userService            *services.UserService
	productService         *services.ProductService
	userManagementService  *services.UserManagementService
	productManagementService *services.ProductManagementService
}

// NewContainer crea una nueva instancia del contenedor de dependencias
// Aquí es donde se configuran todas las implementaciones concretas
func NewContainer() *Container {
	// Crear implementaciones concretas de repositorios
	userRepo := memory.NewUserRepository()
	productRepo := memory.NewProductRepository()

	// Crear implementación concreta del event bus
	eventBus := events.NewInMemoryEventBus()

	// Crear el factory de servicios
	serviceFactory := factories.NewServiceFactory(userRepo, productRepo, eventBus)

	return &Container{
		userRepo:       userRepo,
		productRepo:    productRepo,
		eventBus:       eventBus,
		serviceFactory: serviceFactory,
	}
}

// GetUserService retorna la instancia del servicio de usuario
// Implementa lazy loading para crear el servicio solo cuando se necesita
func (c *Container) GetUserService() *services.UserService {
	if c.userService == nil {
		c.userService = c.serviceFactory.CreateUserService()
	}
	return c.userService
}

// GetProductService retorna la instancia del servicio de producto
func (c *Container) GetProductService() *services.ProductService {
	if c.productService == nil {
		c.productService = c.serviceFactory.CreateProductService()
	}
	return c.productService
}

// GetUserManagementService retorna la instancia del servicio de gestión de usuarios
func (c *Container) GetUserManagementService() *services.UserManagementService {
	if c.userManagementService == nil {
		c.userManagementService = c.serviceFactory.CreateUserManagementService()
	}
	return c.userManagementService
}

// GetProductManagementService retorna la instancia del servicio de gestión de productos
func (c *Container) GetProductManagementService() *services.ProductManagementService {
	if c.productManagementService == nil {
		c.productManagementService = c.serviceFactory.CreateProductManagementService()
	}
	return c.productManagementService
}

// GetAllServices retorna todas las instancias de servicios
func (c *Container) GetAllServices() *factories.AllServices {
	return c.serviceFactory.CreateAllServices()
}

// GetUserRepository retorna la instancia del repositorio de usuarios
func (c *Container) GetUserRepository() repositories.UserRepository {
	return c.userRepo
}

// GetProductRepository retorna la instancia del repositorio de productos
func (c *Container) GetProductRepository() repositories.ProductRepository {
	return c.productRepo
}

// GetEventBus retorna la instancia del event bus
func (c *Container) GetEventBus() events.EventBus {
	return c.eventBus
}

// GetServiceFactory retorna la instancia del factory de servicios
func (c *Container) GetServiceFactory() *factories.ServiceFactory {
	return c.serviceFactory
}