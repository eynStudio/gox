package utils

import (
	"strings"

	linq "github.com/ahmetalpbalkan/go-linq"
)

type ITreeNode interface {
	GetMc() string
	GetNs() string
	GetQz() int
}

type TreeNode struct {
	Mc    string
	Ns    string
	M     interface{}
	Nodes []*TreeNode
}

type TreeNodes []*TreeNode

func newTreeNode(m ITreeNode) *TreeNode {
	return &TreeNode{m.GetMc(), m.GetNs(), m, make([]*TreeNode, 0)}
}

func BuildTree(src interface{}) TreeNodes {
	root := &TreeNode{"", "", nil, make([]*TreeNode, 0)}
	buildTreeNodes(src, root, "")
	return root.Nodes
}

func buildTreeNodes(src interface{}, r *TreeNode, prefix string) {
	results := queryChildren(src, prefix)
	for _, it := range results {
		child := newTreeNode(it.(ITreeNode))
		r.Nodes = append(r.Nodes, child)
		buildTreeNodes(src, child, it.(ITreeNode).GetNs()+".")
	}
}

func queryChildren(ss interface{}, prefix string) []interface{} {
	q := linq.From(ss).Where(func(s interface{}) bool {
		last := strings.TrimPrefix(s.(ITreeNode).GetNs(), prefix)
		return strings.HasPrefix(s.(ITreeNode).GetNs(), prefix) && !strings.Contains(last, ".")
	})
	if q.Any() {
		return q.OrderByDescending(func(a interface{}) interface{} {
			return a.(ITreeNode).GetQz()
		}).Results()
	}
	return q.Results()
}
