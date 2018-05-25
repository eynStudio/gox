package filter

import (
	"fmt"
	"strings"

	"github.com/eynstudio/gox/db"
)

type VisitorResults []db.SqlArgs

func (p VisitorResults) Join(con string) (vr db.SqlArgs) {
	if con == "" {
		con = "and"
	}

	var sqls []string
	for _, it := range p {
		if it.Sql != "" {
			sqls = append(sqls, it.Sql)
			vr.Args = append(vr.Args, it.Args...)
		}
	}
	if len(sqls) > 0 {
		vr.Sql = "(" + strings.Join(sqls, " "+con+" ") + ")"
	}
	return
}

type Visitor struct {
	QuoteStr string
	Quote    func(v string) string
}

func (p *Visitor) quote(v string) string {
	if p.Quote != nil {
		return p.Quote(v)
	}
	return p.QuoteStr + v + p.QuoteStr
}

func (p *Visitor) Visitor(filter Group) db.SqlArgs {
	sqlArgs := p.VisitGroup(filter)
	if len(sqlArgs.Sql) > 0 {
		sqlArgs.Sql = " where " + sqlArgs.Sql
	}
	return sqlArgs
}

func (p *Visitor) VisitGroup(f Group) (vr db.SqlArgs) {
	var lst VisitorResults

	ruleResult := p.VisitRules(f.Con, f.Rules)
	lst = append(lst, ruleResult)

	for _, it := range f.Groups {
		lst = append(lst, p.VisitGroup(it))
	}

	return lst.Join(f.Con)
}

func (p *Visitor) VisitRules(con string, rules []Rule) (vr db.SqlArgs) {
	var lst VisitorResults
	for _, it := range rules {
		lst = append(lst, p.VisitRule(it))
	}
	return lst.Join(con)
}

func (p *Visitor) VisitRule(f Rule) (vr db.SqlArgs) {
	switch f.O {
	case "like", "start", "end":
		vr.Sql = fmt.Sprintf("%s like ?", p.quote(f.F))
		switch f.O {
		case "like":
			vr.AddArgs("%" + f.V1 + "%")
		case "start":
			vr.AddArgs(f.V1 + "%")
		case "end":
			vr.AddArgs("%" + f.V1)
		}
	case "!like":
		vr.Sql = fmt.Sprintf("%s not like ?", p.quote(f.F))
		vr.AddArgs("%" + f.V1 + "%")
	case "=", "<>", ">=", ">", "<", "<=":
		vr.Sql = fmt.Sprintf("%s %s ?", p.quote(f.F), f.O)
		vr.AddArgs(f.V1)
	case "empty":
		vr.Sql = fmt.Sprintf("%s = ?", p.quote(f.F))
		vr.AddArgs("")
	case "!empty":
		vr.Sql = fmt.Sprintf("%s <> ?", p.quote(f.F))
		vr.AddArgs("")
	case "in":
		var ss []string
		vlst := strings.Split(f.V1, ",")
		for _, it := range vlst {
			ss = append(ss, "?")
			vr.Args = append(vr.Args, it)
		}
		vs := strings.Join(ss, ",")
		vr.Sql = fmt.Sprintf("%s %s (%s)", p.quote(f.F), f.O, vs)
	case "between":
		vr.Sql = fmt.Sprintf("%s %s ? and ?", p.quote(f.F), f.O)
		vr.AddArgs(f.V1, f.V2)
	}
	return
}
