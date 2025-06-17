# Public Transportation E-Ticketing System

A comprehensive public transportation e-ticketing system built with Go, featuring offline capability and JWT authentication.

## Features

- 🔐 JWT-based authentication
- 🚇 Terminal management
- 💳 Prepaid card system
- 📊 Fare calculation matrix
- 🔄 Offline transaction sync
- 🛡️ Role-based access control
- 🐳 Docker support
- 📝 Comprehensive API documentation

## Architecture

The system follows clean architecture principles with clear separation of concerns:

```
root
├── Dockerfile
├── cmd
│   └── server
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── connection.go
│   ├── handlers
│   │   ├── auth.go
│   │   └── terminal.go
│   ├── middleware
│   │   └── auth.go
│   ├── models
│   │   ├── card.go
│   │   ├── fare_matrix.go
│   │   ├── offline_queue.go
│   │   ├── response.go
│   │   ├── terminal.go
│   │   ├── transaction.go
│   │   ├── user.go
│   │   └── validation_gate.go
│   ├── repositories
│   │   ├── terminal.go
│   │   └── user.go
│   ├── services
│   │   ├── auth.go
│   │   └── terminal.go
│   └── utils
│       ├── jwt.go
│       └── validator.go
└── pkg
    └── errors
        └── custom_error.go
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 12+
- Docker & Docker Compose (optional)

### Local Development

1. Clone the repository:
```bash
git clone <repository-url>
cd transportation-api
```

2. Install dependencies:
```bash
make deps
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Start PostgreSQL and run migrations:
```bash
make docker-up
make migrate
```

5. Run the application:
```bash
make run
```

### Docker Development

```bash
# Start all services
make docker-up

# View logs
make logs

# Stop services
make docker-down
```

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - User login

### Terminals (Protected)
- `POST /api/v1/terminals` - Create terminal

## Database Schema

The system uses PostgreSQL with the following main entities:
- Users (system administrators)
- Terminals (physical locations)
- Cards (prepaid cards)
- Transactions (journey records)
- Fare Matrix (pricing)
- Validation Gates (hardware devices)
- Offline Queue (sync mechanism)

## Testing

```bash
# Run all tests
make test

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `make fmt` and `make lint`
6. Submit a pull request
