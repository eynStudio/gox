package filter

import x "github.com/eynstudio/gox/x"

// Rule 过滤规则
type Rule struct {
	F  string
	O  string
	V1 string
	V2 string
}

// Group 过滤规则组
type Group struct {
	Con    string
	Rules  []Rule
	Groups []Group
}

// AddRule 添加Rule
func (p *Group) AddRule(r Rule) { p.Rules = append(p.Rules, r) }

// AddGroup 添加规则组
func (p *Group) AddGroup(g Group) { p.Groups = append(p.Groups, g) }

// NewAndGroup 创建AND规则组
func NewAndGroup() (fg Group) { return Group{Con: "and"} }

// NewOrGroup 创建OR规则组
func NewOrGroup() (fg Group) { return Group{Con: "or"} }

// Filter 过滤类
type Filter struct {
	Group
	Keyword string //可用于模糊搜索
	Ext     x.M    // 其它信息，便于扩展
}

// PageFilter 带分页信息过滤类
type PageFilter struct {
	Filter
	PageIndex int
	PageSize  int
}

// Offset 计算Offset
func (p *PageFilter) Offset() int { return (p.PageIndex - 1) * p.PageSize }

// NewPageFilter NewPageFilter
func NewPageFilter(pageIndex, pageSize int) *PageFilter {
	p := &PageFilter{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}
	return p
}
