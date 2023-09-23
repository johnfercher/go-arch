package loader

import (
	"github.com/johnfercher/go-arch/pkg"
	"gopkg.in/yaml.v3"
	"os"
)

type Loader interface {
	LoadArchitecture(file string) (*pkg.Node, error)
	LoadFile(file string) ([]byte, error)
}

type loader struct {
}

func New() Loader {
	return &loader{}
}

func (l *loader) LoadArchitecture(file string) (*pkg.Node, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	node := &pkg.Node{}

	for key, value := range data {
		inner := BuildTree(key, value)
		node.Nodes = append(node.Nodes, inner)
	}

	return node, nil
}

func (l *loader) LoadFile(file string) ([]byte, error) {
	return os.ReadFile(file)
}
