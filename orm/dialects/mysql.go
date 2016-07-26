package dialects

import (
	"database/sql"
	"fmt"

	"github.com/eynstudio/gox/db/meta"
)

type mysql struct {
	commonDialect
}

func (p *mysql) Driver() string { return "mysql" }

func (mysql) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}

func (p *mysql) LoadMeta(db *sql.DB) *meta.MetaDb {
	m := &meta.MetaDb{}
	return m
}
