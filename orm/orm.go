package orm

import (
	"database/sql"

	"github.com/eynstudio/gox"
	"github.com/eynstudio/gox/db/meta"
	"github.com/eynstudio/gox/orm/dialects"
)

type Orm struct {
	db      *sql.DB
	dialect dialects.Dialect
	mapper  MapperFn
	//	models  *models
}

func Open(driver, source string) (*Orm, error) {
	var err error

	orm := &Orm{dialect: dialects.NewDialect(driver)}
	//	orm.models = newModels(orm)
	orm.db, err = sql.Open(driver, source)

	if err == nil {
		err = orm.db.Ping()
	}
	return orm, err
}

func MustOpen(driver, source string) *Orm {
	o, e := Open(driver, source)
	gox.Must(e)
	return o
}
func (p *Orm) SetMapper(f MapperFn) *Orm {
	p.mapper = f
	return p
}
func (p *Orm) DB() *sql.DB            { return p.db }
func (p *Orm) LoadMeta() *meta.MetaDb { return p.dialect.LoadMeta(p.db) }
