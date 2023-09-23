package main

import (
	"github.com/johnfercher/go-arch/pkg"
	"log"
)

func main() {
	node, err := pkg.Load("templates/architecture/hexagonal.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	structure, found := node.GetSubTree("structure")
	if !found {
		log.Fatal("not found")
	}

	apiName := "hexagonal"

	err = pkg.Generate("docs/examples", apiName, structure)
	if err != nil {
		log.Fatal(err.Error())
	}
}
