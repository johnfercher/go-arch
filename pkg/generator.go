package pkg

import (
	"errors"
	"os"
)

func Generate(dir string, node *Node) error {
	if node.Key != "structure" {
		return errors.New("invalid sub tree")
	}

	content := dir
	for _, inner := range node.Nodes {
		err := generate(content, inner)
		if err != nil {
			return err
		}
	}

	return nil
}

func generate(dir string, node *Node) error {
	content := dir + "/" + node.Key
	err := os.Mkdir(content, os.ModePerm)
	if err != nil {
		return err
	}

	for _, inner := range node.Nodes {
		err := generate(content, inner)
		if err != nil {
			return err
		}
	}

	return nil
}
