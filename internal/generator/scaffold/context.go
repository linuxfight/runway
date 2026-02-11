package scaffold

type ProjectOptions struct {
	ProjectName string

	Infra struct {
		Postgres bool
		Redis    bool
	}

	ORM struct {
		Ent bool
	}
}

type ProjectContext struct {
	ProjectName string
	ModulePath  string
	Options     ProjectOptions
}
