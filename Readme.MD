<p align="center"><img height="100px" width="100px" src="https://crying-cats.org/logo.png" alt="Crying Cats Logo"> </p>

<p align="center"><h1>Runway</h1></p>

<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/cryingcatscloud/runway" alt="Go version">
<img src="https://img.shields.io/github/v/release/cryingcatscloud/runway" alt="Latest Stable Version">
<img src="https://img.shields.io/github/license/cryingcatscloud/runway" alt="License">
</p>

## Contents

- [About](#about)
  - [Why use Runway?](#why-use-runway)
- [Installation](#installation)
- [Creating a New Project](#creating-a-new-project)
  - [Interactive mode (default)](#interactive-mode-default)
  - [Non-interactive mode](#non-interactive-mode)
- [Scaffolding with crew make](#scaffolding-with-crew-make)
  - [Creating a feature module](#creating-a-feature-module)
  - [Creating an Ent model (schema)](#creating-an-ent-model-schema)
  - [Creating a config module](#creating-a-config-module)
  - [Why use crew make?](#why-use-crew-make)
- [Code Generation with crew gen](#code-generation-with-crew-gen)
  - [crew gen ent](#crew-gen-ent)
  - [crew gen openapi](#crew-gen-openapi)
  - [crew gen server](#crew-gen-server)
  - [crew gen all](#crew-gen-all)
  - [Typical workflow](#typical-workflow)
- [Defining Routes](#defining-routes)
- [Architecture & Modularity](#architecture--modularity)
## About

**Runway** is a lightweight Go framework that lets you spin up CRUD‑style REST APIs with medium‑complex business logic in minutes.  

It **isolates the transport layer**, generates all boilerplate (including an OpenAPI spec), and gives you a clean, modular project layout out of the box.

### Why use Runway?

- **Fast scaffolding** – one command creates a fully working service skeleton.  
- **OpenAPI generation** – your API docs are generated automatically.  
- **Transport‑layer isolation** – runway will take care of the transport layer and provide convenient DTOs, so you only need to focus on business logic


## Installation

Install the `crew` CLI:

```bash
go install github.com/cryingcatscloud/runway/cmd/crew@latest
```

> Make sure `$GOPATH/bin` (or `$HOME/go/bin`) is in your `PATH` so the `crew` binary is available from any terminal.

Verify installation:

```bash
crew --help
```

---

## Creating a New Project

Runway projects are bootstrapped using the `crew new` command.

### Interactive mode (default)

```bash
crew new my-awesome-service
```

You will be prompted to choose:

1. **Project name** – defaults to the argument you passed
2. **Required infrastructure** – Postgres, Redis, or both
3. **ORM for Postgres** – Ent (recommended) or raw SQL

### Non-interactive mode

```bash
crew new my-awesome-service \
  --req-infra=pg,redis \
  --req-orm=ent \
  --no-interactive
```

| Flag               | Meaning                                  |
| ------------------ | ---------------------------------------- |
| `--no-interactive` | Skip prompts and use flags only          |
| `--req-infra`      | Comma-separated list: `pg`, `redis`      |
| `--req-orm`        | `ent` (requires Postgres) or leave empty |
| `--skip-tidy`      | Skip `go mod tidy` after scaffolding     |

---

## Scaffolding with `crew make`

After a project is created, you will spend most of your time using `crew make`.
This command generates new application components inside your existing service and keeps the structure consistent across the codebase.

```bash
crew make
```

```text
Scaffold helpers (module, config, model)

Usage:
  crew make [command]

Available Commands:
  config      Create a config in app/config
  model       Create an Ent schema
  module      Create a full feature module
```

Think of `crew make` as your “feature builder”: it creates real Go code in the correct places, wires things together, and follows the same architectural conventions used by the initial project template.

---

### Creating a feature module

```bash
crew make module notes
```

This creates a new business feature under:

```
internal/modules/notes
```

With a complete, ready-to-use structure:

* `controller.go` — entry point for the feature; coordinates requests and responses
* `service.go` — business logic
* `repository.go` — data access layer
* `api/routes.go` — HTTP route definitions for the module
* `module.go` — dependency wiring for DI/container registration
* `model.go` — base domain model for the feature

In other words, one command gives you a full vertical slice of a new domain feature, already aligned with the project architecture.

---

### Creating an Ent model (schema)

```bash
crew make model note
```

This creates a new Ent schema inside:

```
schema/
```

Example result:

```
schema/
└── note.go
```

This file defines your data model at the persistence level. After creating or modifying schemas, you typically run:

```bash
crew gen ent
```

to generate the Ent client and related code.

---

### Creating a config module

```bash
crew make config jwt
```

This creates a new configuration component inside:

```
app/config/
```

For example:

```
app/config/jwt.go
```

Use this for structured application settings such as:

* JWT configuration
* external service credentials
* feature flags
* module-specific config blocks

Each generated config integrates into the app’s configuration system and DI wiring.

---

### Why use `crew make`?

It ensures that:

* every feature follows the same structure
* modules are generated fully wired and ready to implement
* Ent schemas live in a predictable location
* configuration stays centralized and structured

This removes manual setup and prevents architectural drift as the project grows.

---

## Code Generation with `crew gen`

`crew gen` is responsible for generating derived code from your source definitions.

```bash
crew gen
```

```text
Code generators

Usage:
  crew gen [command]

Available Commands:
  all         Run all generators
  ent         Run go generate for ent schemas
  openapi     Generate openapi.yaml
  server      Generate HTTP server adapters for all modules
```

Each command has a specific responsibility.

---

### `crew gen ent`

```bash
crew gen ent
```

Runs `go generate` for all Ent schemas located in:

```
schema/
```

This produces:

* the typed Ent client
* query builders
* schema migrations metadata

Run this whenever you:

* add a new schema
* modify fields or relations
* change indexes or constraints

---

### `crew gen openapi`

```bash
crew gen openapi
```

Generates an `openapi.yaml` specification based on your modules and their HTTP contracts.

This spec reflects:

* registered routes
* request/response DTOs
* endpoint structure

Use this to:

* power Swagger / API docs
* generate clients
* keep API documentation always in sync with the code

---

### `crew gen server`

```bash
crew gen server
```

Generates HTTP transport adapters for all modules.

This includes the code that:

* connects routes to controllers
* handles request parsing
* maps DTOs to domain inputs
* writes HTTP responses

Because this layer is generated, your business logic stays fully transport-agnostic.

---

### `crew gen all`

```bash
crew gen all
```

Runs all generators in sequence:

1. HTTP server adapter generation
2. OpenAPI spec generation
3. Ent code generation

This is the most common command during development.

---

### Typical workflow

When adding a new feature:

```bash
crew make module notes
crew make model note
```

After defining fields in the Ent schema:

```bash
crew gen all
```

This ensures:

* database layer is generated
* HTTP layer is wired
* OpenAPI spec is up to date

At this point, you can focus purely on implementing business logic.

## Defining Routes

In Runway, HTTP routes are declared in a structured, declarative way using a `RoutesProvider`.
From this definition, the framework generates:

* HTTP transport adapters
* request binding code
* OpenAPI specification entries

You describe your endpoints once, and Runway derives the transport layer from it.

Import the core types from:

```go
import "github.com/cryingcatscloud/runway"
```

---

## Route Model

A route is described using the `runway.Route` struct:

```go
type Route struct {
	Method      string
	Path        string
	Request     any
	Response    any
	Middlewares []echo.MiddlewareFunc

	Summary     string
	Description string
	Tags        []string
}
```

### Fields

**Method**
HTTP method (`GET`, `POST`, etc.), compatible with `net/http` constants.

**Path**
Echo-style route path:

```
/notes
/notes/:id
/users/:userId/posts
```

**Request**
A DTO struct used for request binding.
Runway will pass it to `echo.Context.Bind()` automatically.

If a route has no body (e.g. `GET`), this can be omitted.

**Response**
A DTO struct that represents the response shape.
Used for OpenAPI generation and type clarity.

**Middlewares**
Optional Echo middlewares applied to this specific route.

Helper:

```go
runway.MW(m1, m2)
```

**Summary / Description**
Used in OpenAPI documentation.

**Tags**
Logical grouping labels for OpenAPI.
Can be anything (e.g. `"notes"`, `"admin"`, `"auth"`).

---

## RoutesProvider

Each module exposes its routes by implementing:

```go
type RoutesProvider interface {
	Routes() map[string]Route
}
```

The map key is the **operation name**.
It is used internally for:

* OpenAPI operation IDs
* generated server wiring
* handler mapping

---

## Example: Notes Routes

```go
package notes

import (
	"net/http"

	"github.com/cryingcatscloud/runway"
)

type Routes struct{}

func (Routes) Routes() map[string]runway.Route {
	return map[string]runway.Route{
		"CreateNote": {
			Method:      http.MethodPost,
			Path:        "/notes",
			Request:     CreateNoteRequest{},
			Response:    NoteResponse{},
			Summary:     "Create note",
			Description: "Creates a new note",
			Tags:        []string{"notes"},
		},
		"GetNote": {
			Method:   http.MethodGet,
			Path:     "/notes/:id",
			Response: NoteResponse{},
			Summary:  "Get note by ID",
			Tags:     []string{"notes"},
		},
	}
}
```

This definition is enough for Runway to:

* register HTTP routes in Echo
* bind incoming requests to DTOs
* include the endpoints in OpenAPI

---

## Validation Tags

Runway supports the full tag set from:

```
github.com/go-playground/validator/v10
```

Example:

```go
Title string `json:"title" validate:"required,min=3,max=120"`
```

Important notes:

* Runway does **not** perform validation automatically.
* Validation is fully controlled by the application.
* Tags are still useful because they are:

  * included in OpenAPI schema generation
  * visible in API documentation
  * consistent with validator/v10 if you use it manually

This gives you flexibility:

* use validator/v10
* use your own validation logic
* skip validation entirely

---

## Route Tags (OpenAPI Grouping)

The `Tags` field is used to group endpoints in OpenAPI:

```go
Tags: []string{"notes"}
```

You can use any logical grouping:

* `"notes"`
* `"admin"`
* `"public"`
* `"internal"`

There are no restrictions — tags are purely organizational.

---

## Middlewares Per Route

You can attach Echo middlewares to a specific route:

```go
"CreateNote": {
	Method:      http.MethodPost,
	Path:        "/notes",
	Request:     CreateNoteRequest{},
	Response:    NoteResponse{},
	Middlewares: runway.MW(AuthMiddleware),
}
```

This allows fine-grained control without affecting the whole server.

---

## How Data Reaches the Controller

Runway generates the HTTP layer, and controllers themselves receive a context and the request DTO:

```go
func (c *Controller) Create(ctx context.Context, dto CreateNoteRequest) error
```

All request data is bound into your DTO using:

```go
ctx.Bind(&dto)
```

This means:

* body
* query parameters
* path parameters
* form data

are all populated into the same struct.

You then work with that DTO inside your controller.

---

## Using DTOs for Body, Query, and Path Parameters

DTOs are not limited to JSON body fields.
Because Runway relies on Echo binding, you can describe all incoming data in one place.

Example:

```go
type CreateNoteRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=120"`
	Content string `json:"content"`
}
```

This struct will be populated from the request body.

But you can also include query and path parameters.

---

### Example: Path + Query + Body

Route:

```go
"UpdateNote": {
	Method:   http.MethodPut,
	Path:     "/notes/:id",
	Request:  UpdateNoteRequest{},
	Response: NoteResponse{},
	Tags:     []string{"notes"},
}
```

DTO:

```go
type UpdateNoteRequest struct {
	ID      string `param:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Publish bool   `query:"publish"`
}
```

How it works:

* `param:"id"` → `/notes/:id`
* `json:"title"` → request body
* `json:"content"` → request body
* `query:"publish"` → `?publish=true`

Example request:

```
PUT /notes/123?publish=true
Content-Type: application/json

{
  "title": "Updated",
  "content": "New text"
}
```

After binding, the DTO will contain:

```go
UpdateNoteRequest{
	ID:      "123",
	Title:   "Updated",
	Content: "New text",
	Publish: true,
}
```

---

## Why This Approach?

This design keeps things predictable:

* Controllers receive only `context.Context` and DTO
* No generated handler signatures
* DTOs define the full contract

Everything that comes from the request should be described in the DTO:

* body fields
* query params
* path params

This makes the DTO the single source of truth for:

* request shape
* OpenAPI schema
* validation hints
* controller input structure

## Architecture & Modularity

Runway is opinionated about structure, but intentionally unopinionated about business logic.

It takes full ownership of the transport layer (HTTP, routing, binding, OpenAPI, adapters) and stays completely out of your domain logic. The goal is to let you focus only on the application itself while keeping a clean, scalable architecture by default.

---

## Core Principles

### Transport layer is fully handled

Runway generates and owns everything related to HTTP:

* route registration
* request binding
* DTO → handler wiring
* OpenAPI generation
* middleware attachment

You define contracts and routes — Runway generates the glue code.

Your business logic:

* does not know about Echo
* does not depend on HTTP
* works with plain Go types and `context.Context`

---

### Business logic is untouched

Runway does not:

* enforce domain patterns
* impose validation strategy
* dictate repository structure
* introduce base classes or magic abstractions

You write:

* controllers
* services
* repositories
* domain models

as normal Go code.

Runway only connects them to the outside world.

---

## Feature-Based Modularity

The primary building block in Runway is a **module**.

A module is a self-contained feature that includes everything needed for one domain area:

```
internal/modules/<feature>
```

Example:

```
internal/modules/notes
├── api/
│   ├── dto.go
│   └── routes.go
├── controller.go
├── service.go
├── repository.go
└── module.go
```

This structure creates a vertical slice of the system:

* API contracts
* orchestration
* business logic
* persistence logic
* dependency wiring

all live together.

---

## Modules Are Self-Contained

A module is designed to be independent.

Recommended rules:

* A module should not import other feature modules
* Communication between features should happen through:

  * shared infrastructure
  * services exposed via DI
* Each module owns its domain

This keeps the system:

* predictable
* decoupled
* easy to scale as features grow

You can think of modules as internal microservices inside one binary.

---

## `module.go`: The Composition Root of a Feature

Each folder that represents a logical unit contains a `module.go`.

Examples:

```
internal/modules/notes/module.go
internal/infra/postgres/module.go
app/module.go
```

These files define how components are wired together.

Runway uses Uber Fx for dependency injection and lifecycle management.
Every `module.go` registers constructors and dependencies that belong to that part of the system.

Typical responsibilities of a feature module:

* register controller
* register service
* register repository
* expose routes
* connect dependencies

Conceptually:

```go
var Module = fx.Module(
	"notes",
	fx.Provide(
		NewService,
		NewRepository,
		NewController,
	),
	fx.Invoke(
		RegisterRoutes,
	),
)
```

You don’t manually stitch the app together in one place.
Each module declares its own dependencies and capabilities.

Fx collects and builds the final application graph.

---

## Layer Overview

Runway projects naturally split into layers:

### App layer

```
app/
```

Responsible for:

* starting the application
* loading config
* initializing HTTP server
* assembling root modules

This is where the application lifecycle is defined.

---

### Infrastructure layer

```
internal/infra/
```

Contains shared technical components:

* Postgres
* Redis
* logging
* external clients

These are exposed as Fx modules and can be used by any feature.

---

### Feature layer

```
internal/modules/
```

Each directory is one domain feature.

This is where you spend most of your time.

A feature module typically includes:

* controller
* service
* repository
* routes
* DTOs

---

### Persistence layer

```
schema/
```

Contains Ent schemas.

These are database models and are independent from:

* HTTP
* DTOs
* controllers

---

## Dependency Direction

Runway encourages a simple, stable direction of dependencies:

```
infra → modules → app
```

* Infra provides tools (DB, Redis, logger)
* Modules use infra
* App assembles everything

Modules should avoid depending on each other directly.

---

## How Everything Gets Assembled

At startup:

1. `app/module.go` registers root modules
2. Each feature registers its own `module.go`
3. Infrastructure modules register DB, Redis, etc.
4. Fx builds the dependency graph
5. Runway collects all of the routes and starts the server

No manual wiring is needed.

---

## The Runway Philosophy

Runway has a very clear boundary:

**It owns the HTTP layer.**
**It does not own your business logic.**

It will:

* generate adapters
* bind DTOs
* wire routes
* produce OpenAPI

It will NOT:

* dictate how services should look
* enforce repository patterns
* run validation automatically
* control your domain models
* mix HTTP concerns into your logic

The result is a system where:

* transport is automated
* architecture is consistent
* domain code stays clean and explicit

You focus on solving business problems.
Runway takes care of everything around the edges.

---

Cleared for takeoff, commander!
© 2026, Crying Cats Cloud