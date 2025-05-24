# TODO List REST API (Go)

A clean architecture REST API for managing TODO lists, built with Go and PostgreSQL. Features JWT authentication, Docker deployment, and database migrations.

## Features

- **RESTful API Design** - Follows REST conventions with proper HTTP methods and status codes
- **Clean Architecture** - Separation of concerns with clear layers (delivery, service, repository)
- **JWT Authentication** - Secure user registration/login with JWT tokens
- **PostgreSQL Database** - Relational data storage with proper schema migrations
- **Docker Support** - Easy containerization for development and deployment
- **Configuration Management** - Environment variables and config files via Viper
- **Graceful Shutdown** - Proper handling of server termination

## Technologies

- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL (+ [sqlx](https://github.com/jmoiron/sqlx) for queries)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Auth**: JWT (JSON Web Tokens)
- **Containerization**: Docker

## API Endpoints

### Authentication
- `POST /auth/sign-up` - Register new user
- `POST /auth/sign-in` - Login user (returns JWT)

### TODO Lists
- `POST /api/lists` - Create new list
- `GET /api/lists` - Get all lists
- `GET /api/lists/:id` - Get list by ID
- `PUT /api/lists/:id` - Update list
- `DELETE /api/lists/:id` - Delete list

### TODO Items
- `POST /api/lists/:id/items` - Create item in list
- `GET /api/lists/:id/items` - Get all items in list
- `GET /api/items/:id` - Get item by ID
- `PUT /api/items/:id` - Update item
- `DELETE /api/items/:id` - Delete item

## Creating migrations 
`migrate create -ext sql -dir ./migrations -seq migration_name`
