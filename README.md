# MediaDex API

REST API backend for **MediaDex**, a personal media library management application. It allows users to track their media (movies, series, books, games, etc.), organize them with tags and custom fields, and create filtered collections.

---

## Features

- User registration and login with JWT (access token + refresh token)
- Full CRUD on **media** entries (movies, series, books, etc.)
- Management of colored **tags** associated with media
- Management of **custom fields** with per-media values (`MediaField`)
- Creation of **collections** with dynamic filters
- All data is scoped per user
- Interactive documentation via Swagger UI

## Tech Stack

| Component      | Technology                           |
|----------------|--------------------------------------|
| Language       | Go 1.26                              |
| Router         | [chi v5](https://github.com/go-chi/chi) |
| ORM            | [GORM](https://gorm.io)              |
| Database       | PostgreSQL                           |
| Authentication | JWT (HS256) + bcrypt                 |
| Documentation  | Swagger via [swaggo](https://github.com/swaggo/swag) |
| Config         | `.env` via [godotenv](https://github.com/joho/godotenv) |

## Prerequisites

- [Go](https://go.dev/dl/) >= 1.21
- [PostgreSQL](https://www.postgresql.org/) >= 14

## Installation

```bash
# Clone the repository
git clone <repo-url>
cd MediaDex-API

# Install dependencies
go mod download
```

## Configuration

Copy the `.env.example` file and rename it to `.env` at the root of the project:

```env
PORT=8080
JWT_SECRET_KEY=your_long_and_secure_jwt_secret
CONNECTION_STRING=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable&TimeZone=UTC
```

| Variable            | Description                          |
|---------------------|--------------------------------------|
| `PORT`              | HTTP server listening port           |
| `JWT_SECRET_KEY`    | JWT token signing secret key         |
| `CONNECTION_STRING` | PostgreSQL connection URL            |

## Running

```bash
go run main.go
```

The server starts on the configured port:

```
Server running on http://localhost:8080
Swagger UI available at http://localhost:8080/swagger/index.html
```

## Project Structure

```
MediaDex-API/
├── main.go                  # Entry point, router definition
├── go.mod
├── config/
│   └── config.go            # .env loading, repository initialization
├── database/
│   ├── database.go          # GORM AutoMigrate
│   └── dbmodel/             # GORM models + Repository interfaces
│       ├── user.go
│       ├── media.go
│       ├── collection.go
│       ├── field.go
│       ├── media_field.go
│       └── tag.go
├── pkg/
│   ├── model/               # Request/response structs (bind/render)
│   ├── authentication/      # Login, register, refresh, JWT, middleware
│   ├── collection/          # Handler + routes /collections
│   ├── field/               # Handler + routes /fields
│   ├── media/               # Handler + routes /media
│   ├── mediaField/          # Handler + routes /mediaFields
│   ├── tag/                 # Handler + routes /tags
│   └── user/                # Handler + routes /users
├── docs/                    # Auto-generated Swagger files (do not edit manually)
└── bruno_mediadex/          # Bruno collection for API testing
```

## API Endpoints

### Authentication — `/api/v1/auth`

| Method | Route            | Description                          |
|--------|------------------|--------------------------------------|
| POST   | `/auth/login`    | Login (email/username + password)    |
| POST   | `/auth/register` | Register a new account               |
| POST   | `/auth/refresh`  | Refresh the access token             |

> All routes below require an `Authorization: Bearer <access_token>` header.

### Media — `/api/v1/media`

| Method | Route                        | Description                    |
|--------|------------------------------|--------------------------------|
| GET    | `/media`                     | List all media                 |
| POST   | `/media`                     | Create a media entry           |
| GET    | `/media/{id}`                | Get a media entry by ID        |
| PATCH  | `/media/{id}`                | Update a media entry           |
| DELETE | `/media/{id}`                | Delete a media entry           |
| GET    | `/media/{id}/tags`           | List tags of a media entry     |
| POST   | `/media/{id}/tags/{tagId}`   | Add a tag to a media entry     |
| DELETE | `/media/{id}/tags/{tagId}`   | Remove a tag from a media entry|
| GET    | `/media/{id}/fields`         | List fields of a media entry   |

### Collections — `/api/v1/collections`

| Method | Route                  | Description                       |
|--------|------------------------|-----------------------------------|
| GET    | `/collections`         | List all collections              |
| POST   | `/collections`         | Create a collection               |
| GET    | `/collections/{id}`    | Get a collection by ID            |
| PUT    | `/collections/{id}`    | Update a collection               |
| DELETE | `/collections/{id}`    | Delete a collection               |

### Tags — `/api/v1/tags`

| Method | Route                  | Description                       |
|--------|------------------------|-----------------------------------|
| GET    | `/tags`                | List all tags                     |
| POST   | `/tags`                | Create a tag                      |
| GET    | `/tags/{id}`           | Get a tag by ID                   |
| PUT    | `/tags/{id}`           | Update a tag                      |
| DELETE | `/tags/{id}`           | Delete a tag                      |
| GET    | `/tags/{id}/media`     | List media associated with a tag  |

### Custom Fields — `/api/v1/fields`

| Method | Route                  | Description                       |
|--------|------------------------|-----------------------------------|
| GET    | `/fields`              | List all fields                   |
| POST   | `/fields`              | Create a field                    |
| GET    | `/fields/{id}`         | Get a field by ID                 |
| PUT    | `/fields/{id}`         | Update a field                    |
| DELETE | `/fields/{id}`         | Delete a field                    |
| GET    | `/fields/{id}/media`   | List media linked to a field      |

### Field Values — `/api/v1/mediaFields`

| Method | Route                                | Description                    |
|--------|--------------------------------------|--------------------------------|
| GET    | `/mediaFields`                       | List all field values          |
| POST   | `/mediaFields`                       | Create a field value           |
| GET    | `/mediaFields/{fieldId}/{mediaId}`   | Get a field value by IDs       |
| PUT    | `/mediaFields/{fieldId}/{mediaId}`   | Update a field value           |
| DELETE | `/mediaFields/{fieldId}/{mediaId}`   | Delete a field value           |

### User Account — `/api/v1/users`

| Method | Route        | Description                    |
|--------|--------------|--------------------------------|
| GET    | `/users/me`  | Get current user profile       |
| PUT    | `/users/me`  | Update current user profile    |
| DELETE | `/users/me`  | Delete current user account    |

## Authentication

The API uses **JWT HS256**:

- **Access token**: 2-hour lifetime, sent in the `Authorization: Bearer <token>` header
- **Refresh token**: 3-hour lifetime, used on `POST /api/v1/auth/refresh`

Passwords are hashed with **bcrypt** before storage.

## Swagger Documentation

Once the server is running, the interactive documentation is available at:

```
http://localhost:<PORT>/swagger/index.html
```

Click **Authorize** and enter `Bearer <access_token>` to test protected routes.

## Testing (Bruno)

The `bruno_mediadex/` folder contains a [Bruno](https://www.usebruno.com/) collection with all API requests.

1. Open Bruno and import the `bruno_mediadex/` folder
2. Configure the `bruno_mediadex` environment (base URL + tokens)
3. Run requests starting with **Authentication**
