{{- if .Options.ORM.Ent }}
package postgres

import (
	"context"
	stdsql "database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"go.uber.org/fx"

	"{{ .ModulePath }}/ent"
)

func NewEntClient(db *stdsql.DB) *ent.Client {
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func RegisterEntLifecycle(lc fx.Lifecycle, client *ent.Client) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
}
{{- end}}