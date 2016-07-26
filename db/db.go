package db

import . "github.com/eynstudio/gox"

type Args struct {
	Sql  string
	Args []interface{}
}

func (p *Args) AddArgs(a ...interface{}) { p.Args = append(p.Args, a...) }

type Paging struct {
	Total int
	Items T
}
