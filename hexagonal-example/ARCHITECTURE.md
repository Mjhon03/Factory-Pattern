# Diagrama de Arquitectura Hexagonal

## Arquitectura del Sistema

```mermaid
graph TB
    subgraph "Interfaces (Puertos)"
        HTTP[HTTP Controller]
        CLI[CLI Interface]
    end
    
    subgraph "Application Layer"
        US[User Service]
        PS[Product Service]
        UMS[User Management Service]
        PMS[Product Management Service]
        SF[Service Factory]
        DI[Container/DI]
    end
    
    subgraph "Domain Layer"
        UE[User Entity]
        PE[Product Entity]
        UR[User Repository Interface]
        PR[Product Repository Interface]
    end
    
    subgraph "Infrastructure Layer (Adaptadores)"
        URM[InMemory User Repository]
        PRM[InMemory Product Repository]
        EB[Event Bus]
        EH[Event Handlers]
    end
    
    subgraph "External Systems"
        DB[(Database)]
        MQ[Message Queue]
        LOG[Logging]
    end
    
    %% Connections
    HTTP --> US
    HTTP --> PS
    CLI --> UMS
    CLI --> PMS
    
    US --> UE
    PS --> PE
    UMS --> US
    PMS --> PS
    
    US --> UR
    PS --> PR
    UMS --> UR
    PMS --> PR
    
    SF --> US
    SF --> PS
    SF --> UMS
    SF --> PMS
    DI --> SF
    
    UR -.-> URM
    PR -.-> PRM
    
    URM --> DB
    PRM --> DB
    
    US --> EB
    PS --> EB
    EB --> EH
    EH --> MQ
    EH --> LOG
    
    %% Styling
    classDef domain fill:#e1f5fe
    classDef application fill:#f3e5f5
    classDef infrastructure fill:#e8f5e8
    classDef external fill:#fff3e0
    
    class UE,PE,UR,PR domain
    class US,PS,UMS,PMS,SF,DI application
    class URM,PRM,EB,EH infrastructure
    class DB,MQ,LOG external
```

## Flujo de Datos

```mermaid
sequenceDiagram
    participant Client
    participant Service
    participant Validator
    participant Processor
    participant Repository
    participant EventBus
    participant EventHandler
    
    Client->>Service: CreateUser(id, email, name)
    Service->>Validator: ValidateCreateUser()
    Validator-->>Service: Validation OK
    
    Service->>Processor: CreateUser()
    Processor->>Repository: Save(user)
    Repository-->>Processor: User saved
    Processor-->>Service: User created
    
    Service->>EventBus: PublishUserCreated()
    EventBus->>EventHandler: Handle event
    EventHandler-->>EventBus: Event processed
    
    Service-->>Client: User created successfully
```

## Patrones Implementados

### 1. Repository Pattern
- **Interfaz**: Define el contrato para acceso a datos
- **Implementación**: Encapsula la lógica de persistencia
- **Beneficio**: Desacopla el dominio de la infraestructura

### 2. Factory Pattern
- **Factory**: Crea servicios con dependencias inyectadas
- **Container**: Gestiona el ciclo de vida de las dependencias
- **Beneficio**: Centraliza la creación de objetos complejos

### 3. Service Layer Granular
- **Validator**: Validación de datos de entrada
- **Processor**: Lógica de negocio y persistencia
- **Publisher**: Publicación de eventos
- **Service**: Orquestación de los componentes

### 4. Dependency Injection
- **Container**: Contenedor de dependencias
- **Lazy Loading**: Servicios creados bajo demanda
- **Beneficio**: Inversión de control y testabilidad

## Principios SOLID Aplicados

- **S** - Single Responsibility: Cada clase tiene una responsabilidad
- **O** - Open/Closed: Abierto para extensión, cerrado para modificación
- **L** - Liskov Substitution: Las implementaciones son intercambiables
- **I** - Interface Segregation: Interfaces específicas y cohesivas
- **D** - Dependency Inversion: Dependencias hacia abstracciones