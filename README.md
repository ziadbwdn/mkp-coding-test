# Public Transportation E-Ticketing System

A comprehensive public transportation e-ticketing system built with Go, featuring offline capability and JWT authentication.

## Features

- 🔐 JWT-based authentication
- 🚇 Terminal management
- 💳 Prepaid card system
- 📊 Fare calculation matrix
- 🔄 Offline transaction sync
- 🛡️ Role-based access control
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

## Database Schema

┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│     USERS       │     │    TERMINALS    │     │     CARDS       │
├─────────────────┤     ├─────────────────┤     ├─────────────────┤
│ id (PK)         │     │ id (PK)         │     │ id (PK)         │
│ username        │     │ name            │     │ card_number     │
│ email           │     │ code            │     │ balance         │
│ password_hash   │     │ location        │     │ user_id (FK)    │
│ role            │     │ is_active       │     │ is_active       │
│ created_at      │     │ created_at      │     │ created_at      │
│ updated_at      │     │ updated_at      │     │ updated_at      │
└─────────────────┘     └─────────────────┘     └─────────────────┘
         │                        │                        │
         │                        │                        │
         └────────────────────────┼────────────────────────┘
                                  │
┌─────────────────┐               │              ┌─────────────────┐
│  TRANSACTIONS   │               │              │   FARE_MATRIX   │
├─────────────────┤               │              ├─────────────────┤
│ id (PK)         │               │              │ id (PK)         │
│ card_id (FK)    │──────────────┘               │ from_terminal   │
│ entry_terminal  │                              │ to_terminal     │
│ exit_terminal   │                              │ fare_amount     │
│ entry_time      │                              │ created_at      │
│ exit_time       │                              │ updated_at      │
│ fare_amount     │                              └─────────────────┘
│ status          │
│ created_at      │     ┌─────────────────┐
│ updated_at      │     │ VALIDATION_GATES│
└─────────────────┘     ├─────────────────┤
         │              │ id (PK)         │
         │              │ terminal_id (FK)│
         │              │ gate_code       │
         │              │ is_active       │
         │              │ last_sync       │
         │              │ created_at      │
         │              │ updated_at      │
         │              └─────────────────┘
         │                       │
┌─────────────────┐              │
│ OFFLINE_QUEUE   │              │
├─────────────────┤              │
│ id (PK)         │              │
│ gate_id (FK)    │──────────────┘
│ transaction_data│
│ timestamp       │
│ is_synced       │
│ created_at      │
└─────────────────┘
# Business Logic and Flow

## Design Description (Online Mode)

1. User Journey:

- User taps prepaid card at validation gate (check-in)
- System records entry terminal, timestamp, card ID
- User travels to destination
- User taps card at exit validation gate (check-out)
- System calculates fare based on entry/exit terminals
- Fare is deducted from prepaid card balance


2. System Components:

- Validation Gates: Hardware devices at each terminal with card readers
- API Gateway: Handles all requests from validation gates
- Authentication Service: Manages JWT tokens for gate authentication
- Database Service: Stores transactions, terminals, card data
- Fare Calculation Engine: Determines fare based on terminal matrix


3. Data Flow (Online):
Gate → API Gateway → Auth Check → Business Logic → Database → Response


## Design Description (Offline Mode)

1. Offline Capability:

- Each validation gate has local storage for transactions
- Pre-loaded fare matrix and terminal data
- Cached card balances (last known state)
- Local transaction queue


2. Offline Process:

- Gate validates card locally using cached data
- Transactions stored in local queue with timestamps
- Fare calculated using local fare matrix
- Balance updated locally (optimistic approach)


3. Sync Process (When connection restored):

- Gate sends all queued transactions to server
- Server validates and processes transactions chronologically
- Conflicts resolved (insufficient balance, duplicate transactions)
- Updated card balances and fare matrices pushed to gates

  ------

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
