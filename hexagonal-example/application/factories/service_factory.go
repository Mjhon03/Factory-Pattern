package factories

import (
	"hexagonal-example/application/services"
	"hexagonal-example/domain/repositories"
	"hexagonal-example/infrastructure/events"
)

// ServiceFactory implementa el patrón Factory para crear servicios
// Este patrón encapsula la lógica de creación de objetos complejos
// y proporciona una interfaz unificada para crear diferentes tipos de servicios
type ServiceFactory struct {
	userRepo    repositories.UserRepository
	productRepo repositories.ProductRepository
	eventBus    events.EventBus
}

// NewServiceFactory crea una nueva instancia del factory de servicios
// Recibe todas las dependencias necesarias para crear los servicios
func NewServiceFactory(
	userRepo repositories.UserRepository,
	productRepo repositories.ProductRepository,
	eventBus events.EventBus,
) *ServiceFactory {
	return &ServiceFactory{
		userRepo:    userRepo,
		productRepo: productRepo,
		eventBus:    eventBus,
	}
}

// CreateUserService crea un servicio de usuario con todas sus dependencias
// El factory se encarga de inyectar las dependencias correctas
func (f *ServiceFactory) CreateUserService() *services.UserService {
	// Crear los servicios granulares
	validator := services.NewUserValidator()
	processor := services.NewUserProcessor(f.userRepo)
	publisher := services.NewUserEventPublisher(f.eventBus)

	// Crear el servicio principal que orquesta los servicios granulares
	return services.NewUserService(validator, processor, publisher)
}

// CreateProductService crea un servicio de producto con todas sus dependencias
func (f *ServiceFactory) CreateProductService() *services.ProductService {
	// Crear los servicios granulares
	validator := services.NewProductValidator()
	processor := services.NewProductProcessor(f.productRepo)
	publisher := services.NewProductEventPublisher(f.eventBus)

	// Crear el servicio principal que orquesta los servicios granulares
	return services.NewProductService(validator, processor, publisher)
}

// CreateUserManagementService crea un servicio de gestión de usuarios
// Este servicio combina múltiples servicios para operaciones complejas
func (f *ServiceFactory) CreateUserManagementService() *services.UserManagementService {
	userService := f.CreateUserService()
	return services.NewUserManagementService(userService, f.userRepo)
}

// CreateProductManagementService crea un servicio de gestión de productos
func (f *ServiceFactory) CreateProductManagementService() *services.ProductManagementService {
	productService := f.CreateProductService()
	return services.NewProductManagementService(productService, f.productRepo)
}

// CreateAllServices crea todos los servicios disponibles
// Útil para inicializar toda la aplicación de una vez
func (f *ServiceFactory) CreateAllServices() *AllServices {
	return &AllServices{
		UserService:            f.CreateUserService(),
		ProductService:         f.CreateProductService(),
		UserManagementService:  f.CreateUserManagementService(),
		ProductManagementService: f.CreateProductManagementService(),
	}
}

// AllServices contiene todas las instancias de servicios creadas
// Facilita el acceso a todos los servicios desde un solo lugar
type AllServices struct {
	UserService            *services.UserService
	ProductService         *services.ProductService
	UserManagementService  *services.UserManagementService
	ProductManagementService *services.ProductManagementService
}