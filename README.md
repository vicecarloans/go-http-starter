# go-http-starter

Starter code for a Go HTTP API: [Fiber](https://gofiber.io/) for routing, [GORM](https://gorm.io/) on PostgreSQL for persistence, and [Atlas](https://atlasgo.io/) for versioned schema migrations. It ships with a dev container, so the database and toolchain come up with one click.

The sample domain is a small books and authors API, laid out so each feature is easy to copy when you add your own.

## Stack

| Concern | Choice |
|---|---|
| Language | Go 1.22 |
| HTTP framework | Fiber v2 (CORS, compression, request logging, panic recovery, request IDs) |
| ORM | GORM with the Postgres driver |
| Migrations | Atlas, with the schema generated from the GORM models |
| Database | PostgreSQL 16 |
| Dev environment | VS Code Dev Container with Docker Compose |

## Project layout

```
.devcontainer/          Dev container: Go, Atlas, and a Postgres 16 service
server/
  cmd/
    go-http-server/     API entry point
    migrations/         Prints the GORM schema for Atlas to diff against
  internal/
    authors/            domain, handler, service, repository
    books/              domain, handler, service, repository
    infrastructure/     Fiber app wiring and the Postgres connection
  migrations/           Versioned SQL migrations managed by Atlas
  utils/                Environment config loading
  Dockerfile            Multi-stage build of the API image
```

Each feature package follows the same four layers: `domain.go` (models and DTOs), `repository.go` (database access), `service.go` (business logic), and `handler.go` (HTTP routes).

## Getting started

### With the dev container (recommended)

1. Open the repository in VS Code and choose **Reopen in Container**. This starts Postgres and installs Go and Atlas.
2. Create your environment file:

   ```bash
   cd server
   cp .env.example .env
   ```

3. Apply the migrations:

   ```bash
   atlas migrate apply \
     --dir "file://migrations" \
     --url "postgres://postgres:postgres@localhost:5432/go-project?sslmode=disable"
   ```

4. Run the server:

   ```bash
   go run ./cmd/go-http-server
   ```

The API listens on `http://localhost:8080`.

### Without the dev container

You need Go 1.22+, the Atlas CLI, and a PostgreSQL 16 instance. Put your connection details in `server/.env`, then follow steps 3 and 4 above.

## API

| Method | Path | Description |
|---|---|---|
| GET | `/ping` | Health check |
| GET | `/api/v1/books?offset=0&limit=10` | List books |
| GET | `/api/v1/books/:id` | Get a book |
| POST | `/api/v1/books` | Create a book |
| PATCH | `/api/v1/books/:id` | Update a book |
| DELETE | `/api/v1/books/:id` | Delete a book |
| GET | `/api/v1/authors` | List authors |
| GET | `/api/v1/authors/:id` | Get an author |
| POST | `/api/v1/authors` | Create an author |
| PATCH | `/api/v1/authors/:id` | Update an author |
| DELETE | `/api/v1/authors/:id` | Delete an author |

Responses use a consistent envelope:

```json
{ "status": "success", "data": {} }
```

```json
{ "status": "fail", "message": "Book not found" }
```

Try it:

```bash
curl http://localhost:8080/ping
curl "http://localhost:8080/api/v1/books?limit=5"
```

## Changing the schema

The GORM models are the source of truth. After editing a model, generate a new migration from the diff:

```bash
cd server
atlas migrate diff your_change_name --env gorm
```

Atlas writes a new SQL file to `server/migrations/`. Review it, then apply it with the `atlas migrate apply` command above.

## Configuration

Set in `server/.env`:

| Variable | Example |
|---|---|
| `DB_HOST` | `localhost` |
| `DB_PORT` | `5432` |
| `DB_USER` | `postgres` |
| `DB_PASSWORD` | `postgres` |
| `DB_NAME` | `go-project` |

## Docker

Build and run the API image:

```bash
cd server
docker build -t go-http-server .
docker run --env-file .env -p 8080:8080 go-http-server
```

## Adding a feature

1. Copy `internal/books` to a new package and rename the types.
2. Register the model in `cmd/migrations/main.go`.
3. Wire the repository, service, and handler in `internal/infrastructure/fiber.go`.
4. Generate and apply a migration.
