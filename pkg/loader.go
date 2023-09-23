package pkg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func Load(file string) (*Node, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	node := &Node{}

	for key, value := range data {
		inner := BuildTree(key, value)
		node.Nodes = append(node.Nodes, inner)
	}

	return node, nil
}
