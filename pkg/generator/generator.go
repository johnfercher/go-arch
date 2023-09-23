package generator

import (
	"errors"
	"github.com/johnfercher/go-arch/pkg"
	"github.com/johnfercher/go-arch/pkg/writer"
	"os"
)

type Generator interface {
	Bootstrap(dir string, apiName string, node *pkg.Node) error
}

type generator struct {
	factory writer.Factory
}

func New(factory writer.Factory) Generator {
	return &generator{
		factory: factory,
	}
}

func (g *generator) Bootstrap(dir string, apiName string, node *pkg.Node) error {
	if node.Key != "structure" {
		return errors.New("invalid sub tree")
	}

	content := dir + "/" + apiName
	err := os.Mkdir(content, os.ModePerm)
	if err != nil {
		return err
	}

	for _, inner := range node.Nodes {
		err := g.generate(content, inner)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *generator) generate(dir string, node *pkg.Node) error {
	content := dir + "/" + node.Key
	err := os.Mkdir(content, os.ModePerm)
	if err != nil {
		return err
	}

	for _, inner := range node.Nodes {
		err := g.generate(content, inner)
		if err != nil {
			return err
		}
	}

	for _, value := range node.Values {
		err := g.createFile(dir, node.Key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *generator) createFile(path string, dir string, value string) error {
	writer := g.factory.Create(value)
	if writer != nil {
		return writer.WriteFile(path, dir)
	}
	return nil
}
