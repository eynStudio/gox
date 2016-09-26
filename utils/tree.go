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

func newTreeNode(m ITreeNode) *TreeNode {
	return &TreeNode{m.GetMc(), m.GetNs(), m, make([]*TreeNode, 0)}
}

func BuildTree(src interface{}) []*TreeNode {
	root := &TreeNode{"", "", nil, make([]*TreeNode, 0)}
	buildTreeNodes(src, root, "")
	return root.Nodes
}

func buildTreeNodes(src interface{}, r *TreeNode, prefix string) {
	results, _ := queryChildren(linq.From(src), prefix)
	for _, it := range results {
		child := newTreeNode(it.(ITreeNode))
		r.Nodes = append(r.Nodes, child)
		buildTreeNodes(src, child, it.(ITreeNode).GetNs()+".")
	}
}

func queryChildren(q linq.Query, prefix string) ([]linq.T, error) {
	return q.Where(func(s linq.T) (bool, error) {
		last := strings.TrimPrefix(s.(ITreeNode).GetNs(), prefix)
		return strings.HasPrefix(s.(ITreeNode).GetNs(), prefix) && !strings.Contains(last, "."), nil
	}).OrderBy(func(a, b linq.T) bool {
		return a.(ITreeNode).GetQz() > b.(ITreeNode).GetQz()
	}).Results()
}
