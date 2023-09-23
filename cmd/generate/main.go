package main

import (
	"github.com/johnfercher/go-arch/pkg/generator"
	"github.com/johnfercher/go-arch/pkg/loader"
	"github.com/johnfercher/go-arch/pkg/writer"
	"log"
)

func main() {
	loader := loader.New()

	node, err := loader.LoadArchitecture("templates/architecture/hexagonal.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	factory := writer.New(loader, node)
	generator := generator.New(factory)

	structure, found := node.GetSubTree("structure")
	if !found {
		log.Fatal("not found")
	}

	apiName := "hexagonal"

	err = generator.Bootstrap("docs/examples", apiName, structure)
	if err != nil {
		log.Fatal(err.Error())
	}
}
