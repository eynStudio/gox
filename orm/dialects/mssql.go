package dialects

type mssql struct {
	commonDialect
}

func (p *mssql) Driver() string { return "mssql" }
