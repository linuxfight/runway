{{- if .Options.Infra.Postgres }}
package postgres

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"{{ .ModulePath }}/app/config"
)

func NewSQLDB(cfg *config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxConnMaxLifetime) * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
{{- end }}