# user-api
User API wrote in GO and gRPC

# Project Setup
## Prerequisites
- Install GO
- Install Protocoll Buffer: https://github.com/protocolbuffers/protobuf/releases
- Install go plugins for protocol compilers: https://grpc.io/docs/languages/go/quickstart/

# Project organization

    project-root/
    ├── cmd/                # Main application entry points (binaries)
    ├── internal/           # Core domain logic (optional, restricts access to internal packages)
    │   ├── domain/         # Domain layer (Entities, Value Objects, Aggregates, Interfaces)
    │   ├── application/    # Application layer (Use cases, orchestrates domain logic)
    │   ├── infrastructure/ # Infrastructure layer (Database, gRPC, HTTP, external APIs)
    │   ├── interfaces/     # Interface adapters (Controllers, Repositories, Gateways)
    │   └── config/         # Application configuration and settings
    ├── pkg/                # Shared libraries (reusable, domain-independent)
    ├── api/                # Protobuf/gRPC or OpenAPI definitions
    └── docs/               # Documentation (e.g., architecture, API specs)
