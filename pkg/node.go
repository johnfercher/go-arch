package pkg

import "fmt"

type Node struct {
	Key    string
	Nodes  []*Node
	Values []string
}

func (n *Node) String(depth int) string {
	var content string

	var space string
	for i := 0; i < depth; i++ {
		space += "--"
	}

	content += fmt.Sprintf("%s%s\n", space, n.Key)
	content += fmt.Sprintf("%s%v\n", space, n.Values)

	for _, node := range n.Nodes {
		content += node.String(depth + 1)
	}

	return content
}

func (n *Node) GetSubTree(key string) (*Node, bool) {
	if n.Key == key {
		return n, true
	}

	for _, inner := range n.Nodes {
		_, matches := inner.GetSubTree(key)
		if matches {
			return inner, matches
		}
	}

	return nil, false
}
