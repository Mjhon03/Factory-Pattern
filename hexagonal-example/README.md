# Arquitectura Hexagonal con Patrones Repository y Factory

Este proyecto demuestra cómo implementar una **Arquitectura Hexagonal** (también conocida como **Ports and Adapters**) en Go, utilizando los patrones **Repository** y **Factory** junto con **Dependency Injection** explícita.

## 🏗️ Arquitectura del Proyecto

```
hexagonal-example/
├── domain/                    # Capa de Dominio (Núcleo del negocio)
│   ├── entities/             # Entidades de dominio
│   │   ├── user.go          # Entidad Usuario
│   │   └── product.go       # Entidad Producto
│   └── repositories/        # Interfaces Repository (Puertos)
│       ├── user_repository.go
│       └── product_repository.go
├── application/              # Capa de Aplicación
│   ├── services/            # Servicios de aplicación
│   │   ├── user_*.go       # Servicios granulares de usuario
│   │   └── product_*.go    # Servicios granulares de producto
│   └── factories/           # Factory Pattern
│       └── service_factory.go
├── infrastructure/          # Capa de Infraestructura (Adaptadores)
│   ├── repositories/        # Implementaciones concretas
│   │   └── memory/         # Repositorios en memoria
│   ├── events/             # Sistema de eventos
│   └── config/             # Configuración y DI
│       └── container.go    # Contenedor de dependencias
└── main.go                 # Punto de entrada
```

## 🎯 Patrones Implementados

### 1. **Patrón Repository**

El patrón Repository encapsula la lógica de acceso a datos y proporciona una interfaz más orientada a objetos para acceder a la capa de persistencia.

**Interfaz (Puerto):**
```go
type UserRepository interface {
    Save(ctx context.Context, user *entities.User) error
    FindByID(ctx context.Context, id string) (*entities.User, error)
    FindByEmail(ctx context.Context, email string) (*entities.User, error)
    // ... más métodos
}
```

**Implementación (Adaptador):**
```go
type InMemoryUserRepository struct {
    users map[string]*entities.User
    mutex sync.RWMutex
}
```

**Beneficios:**
- ✅ Desacopla la lógica de negocio del acceso a datos
- ✅ Facilita el testing con implementaciones mock
- ✅ Permite cambiar la implementación sin afectar el dominio
- ✅ Centraliza la lógica de acceso a datos

### 2. **Patrón Factory**

El patrón Factory encapsula la lógica de creación de objetos complejos y proporciona una interfaz unificada para crear diferentes tipos de servicios.

```go
type ServiceFactory struct {
    userRepo    repositories.UserRepository
    productRepo repositories.ProductRepository
    eventBus    events.EventBus
}

func (f *ServiceFactory) CreateUserService() *services.UserService {
    validator := services.NewUserValidator()
    processor := services.NewUserProcessor(f.userRepo)
    publisher := services.NewUserEventPublisher(f.eventBus)
    
    return services.NewUserService(validator, processor, publisher)
}
```

**Beneficios:**
- ✅ Centraliza la creación de objetos complejos
- ✅ Encapsula la lógica de inyección de dependencias
- ✅ Facilita la configuración y el mantenimiento
- ✅ Permite crear variaciones de servicios fácilmente

### 3. **Service Layer Granular**

Hemos separado las responsabilidades en servicios específicos:

- **Validator**: Se encarga únicamente de la validación de datos
- **Processor**: Maneja la lógica de negocio y persistencia
- **EventPublisher**: Publica eventos del sistema
- **Service**: Orquesta los servicios granulares

```go
type UserService struct {
    validator  *UserValidator
    processor  *UserProcessor
    publisher  *UserEventPublisher
}

func (s *UserService) CreateUser(ctx context.Context, id, email, name string) (*entities.User, error) {
    // 1. Validar
    if err := s.validator.ValidateCreateUser(id, email, name); err != nil {
        return nil, err
    }
    
    // 2. Procesar
    user, err := s.processor.CreateUser(ctx, id, email, name)
    if err != nil {
        return nil, err
    }
    
    // 3. Publicar evento
    s.publisher.PublishUserCreated(ctx, user)
    
    return user, nil
}
```

**Beneficios:**
- ✅ Separación clara de responsabilidades
- ✅ Facilita el testing individual de cada componente
- ✅ Permite reutilizar componentes en diferentes contextos
- ✅ Hace el código más mantenible y legible

### 4. **Dependency Injection Explícita**

El contenedor de dependencias se encarga de crear y configurar todas las dependencias de manera centralizada.

```go
type Container struct {
    userRepo    repositories.UserRepository
    productRepo repositories.ProductRepository
    eventBus    events.EventBus
    serviceFactory *factories.ServiceFactory
    // ... servicios lazy-loaded
}

func NewContainer() *Container {
    // Crear implementaciones concretas
    userRepo := memory.NewUserRepository()
    productRepo := memory.NewProductRepository()
    eventBus := events.NewInMemoryEventBus()
    
    // Crear factory
    serviceFactory := factories.NewServiceFactory(userRepo, productRepo, eventBus)
    
    return &Container{
        userRepo:       userRepo,
        productRepo:    productRepo,
        eventBus:       eventBus,
        serviceFactory: serviceFactory,
    }
}
```

**Beneficios:**
- ✅ Configuración centralizada de dependencias
- ✅ Lazy loading de servicios
- ✅ Facilita el testing con mocks
- ✅ Inversión de control clara

## 🚀 Cómo Ejecutar el Proyecto

1. **Clonar y navegar al proyecto:**
```bash
cd /workspace/hexagonal-example
```

2. **Ejecutar el ejemplo:**
```bash
go run main.go
```

3. **Ver la salida:**
```
=== Ejemplo de Arquitectura Hexagonal con Patrones Repository y Factory ===

=== EJEMPLOS DE USUARIOS ===

1. Creando usuarios...
📧 Evento: Usuario creado - ID: user1, Email: juan@example.com, Nombre: Juan Pérez
✅ Usuario creado: Juan Pérez (juan@example.com)
📧 Evento: Usuario creado - ID: user2, Email: maria@example.com, Nombre: María García
✅ Usuario creado: María García (maria@example.com)

2. Actualizando usuario...
📧 Evento: Usuario actualizado - ID: user1, Email: juan@example.com, Nombre: Juan Carlos Pérez
✅ Usuario actualizado: Juan Carlos Pérez

3. Obteniendo usuario...
✅ Usuario obtenido: Juan Carlos Pérez (juan@example.com)

4. Listando usuarios...
✅ Total de usuarios: 2
   - Juan Carlos Pérez (juan@example.com) - Activo: true
   - María García (maria@example.com) - Activo: true

5. Desactivando usuario...
📧 Evento: Usuario desactivado - ID: user2, Email: maria@example.com
✅ Usuario desactivado: María García

=== EJEMPLOS DE PRODUCTOS ===

1. Creando productos...
📦 Evento: Producto creado - ID: prod1, Nombre: Laptop Gaming, Precio: $1299.99
✅ Producto creado: Laptop Gaming - $1299.99 (Stock: 10)
📦 Evento: Producto creado - ID: prod2, Nombre: Mouse Inalámbrico, Precio: $29.99
✅ Producto creado: Mouse Inalámbrico - $29.99 (Stock: 50)

2. Actualizando precio...
📦 Evento: Producto actualizado - ID: prod1, Nombre: Laptop Gaming, Precio: $1199.99
✅ Precio actualizado: Laptop Gaming - $1199.99

3. Actualizando stock...
📦 Evento: Stock actualizado - Producto: Laptop Gaming, Stock anterior: 10, Stock nuevo: 15
✅ Stock actualizado: Laptop Gaming - Stock: 15

4. Listando productos disponibles...
✅ Productos disponibles: 2
   - Laptop Gaming - $1199.99 (Stock: 15) - Disponible: true
   - Mouse Inalámbrico - $29.99 (Stock: 50) - Disponible: true

5. Buscando productos por categoría...
✅ Productos en Electrónicos: 1
   - Laptop Gaming - $1199.99

=== EJEMPLOS DE SERVICIOS DE GESTIÓN ===

1. Estadísticas de usuarios...
✅ Estadísticas de usuarios:
   - Total: 2
   - Activos: 1
   - Inactivos: 1

2. Estadísticas de productos...
✅ Estadísticas de productos:
   - Total: 2
   - Disponibles: 2
   - No disponibles: 0

3. Creación en lote de usuarios...
📧 Evento: Usuario creado - ID: user3, Email: ana@example.com, Nombre: Ana López
✅ Usuario creado en lote: Ana López
📧 Evento: Usuario creado - ID: user4, Email: carlos@example.com, Nombre: Carlos Rodríguez
✅ Usuario creado en lote: Carlos Rodríguez
📧 Evento: Usuario creado - ID: user5, Email: laura@example.com, Nombre: Laura Martínez
✅ Usuario creado en lote: Laura Martínez
```

## 🔍 Conceptos Clave Explicados

### **Arquitectura Hexagonal**

La arquitectura hexagonal (Ports and Adapters) separa la lógica de negocio de los detalles de implementación:

- **Puertos (Ports)**: Interfaces que definen qué puede hacer el sistema
- **Adaptadores (Adapters)**: Implementaciones concretas que conectan con sistemas externos
- **Dominio**: Contiene la lógica de negocio pura, sin dependencias externas

### **Flujo de Datos**

1. **Entrada**: Los controladores HTTP/CLI reciben las peticiones
2. **Aplicación**: Los servicios de aplicación procesan la lógica de negocio
3. **Dominio**: Las entidades y reglas de negocio se ejecutan
4. **Infraestructura**: Los repositorios persisten los datos y se publican eventos
5. **Salida**: Se retorna la respuesta al cliente

### **Ventajas de esta Arquitectura**

- ✅ **Testabilidad**: Cada capa se puede probar independientemente
- ✅ **Mantenibilidad**: Cambios en una capa no afectan a las otras
- ✅ **Flexibilidad**: Se pueden cambiar implementaciones sin afectar el dominio
- ✅ **Escalabilidad**: Fácil agregar nuevas funcionalidades
- ✅ **Desacoplamiento**: Las dependencias apuntan hacia el centro (dominio)

## 🧪 Testing

Para probar esta arquitectura, puedes:

1. **Crear mocks** de los repositorios para testing unitario
2. **Usar el contenedor** para inyectar dependencias de prueba
3. **Probar cada capa** independientemente
4. **Usar tests de integración** con implementaciones reales

## 🔧 Extensiones Posibles

- **Base de datos real**: Reemplazar los repositorios en memoria con PostgreSQL/MySQL
- **API REST**: Agregar controladores HTTP con Gin o Echo
- **Eventos asíncronos**: Implementar colas de mensajes con RabbitMQ/Kafka
- **Caché**: Agregar Redis para mejorar el rendimiento
- **Logging**: Implementar logging estructurado
- **Métricas**: Agregar monitoreo y métricas

## 📚 Referencias

- [Arquitectura Hexagonal - Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Clean Architecture - Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Repository Pattern - Martin Fowler](https://martinfowler.com/eaaCatalog/repository.html)
- [Factory Pattern - Gang of Four](https://en.wikipedia.org/wiki/Factory_method_pattern)

---

Este proyecto demuestra cómo implementar una arquitectura limpia y mantenible en Go, utilizando patrones de diseño probados y principios SOLID. ¡Espero que te sea útil para entender estos conceptos! 🚀