# Application environment
# Possible values: production, staging, development
APP_ENV=development

# ------------------------------------------------------------------------------
# Logging
# ------------------------------------------------------------------------------

# Log level: debug, info, warn, error
LOG_LEVEL=info

# Enable JSON logs (true for production, false for local dev)
LOG_JSON=true

# ------------------------------------------------------------------------------
# HTTP Server
# ------------------------------------------------------------------------------

# Public bind address for HTTP server
HTTP_PUBLIC_BIND_ADDR=:8080

# Read timeout in seconds
HTTP_READ_TIMEOUT_SEC=15

# Write timeout in seconds
HTTP_WRITE_TIMEOUT_SEC=15

# Idle timeout in seconds
HTTP_IDLE_TIMEOUT_SEC=60

{{- if .Options.Infra.Postgres }}

# ------------------------------------------------------------------------------
# Postgres
# ------------------------------------------------------------------------------

# Postgres connection string
# Example:
# postgres://user:password@localhost:5432/dbname?sslmode=disable
POSTGRES_DSN=postgres://user:password@localhost:5432/db?sslmode=disable

# Maximum number of open connections
POSTGRES_MAX_OPEN_CONNS=50

# Maximum number of idle connections
POSTGRES_MAX_IDLE_CONNS=10

# Maximum connection lifetime (in minutes)
POSTGRES_CONN_MAX_LIFETIME_MINUTES=30
{{- end }}

{{- if .Options.Infra.Redis }}

# ------------------------------------------------------------------------------
# Redis
# ------------------------------------------------------------------------------

# Redis server address
REDIS_ADDR=localhost:6379

# Redis password (leave empty if none)
REDIS_PASSWORD=

# Redis database index
REDIS_DB=0
{{- end }}
