package writer

import (
	"github.com/johnfercher/go-arch/pkg"
	"github.com/johnfercher/go-arch/pkg/loader"
	"os"
	"strings"
)

type entityWriter struct {
	loader loader.Loader
	file   string
	node   *pkg.Node
}

func NewEntityWriter(loader loader.Loader, node *pkg.Node) *entityWriter {
	return &entityWriter{
		loader: loader,
		file:   "templates/code/entity.txt",
		node:   node,
	}
}

func (e *entityWriter) WriteFile(path string, dir string, value string) error {
	bytes, err := e.loader.LoadFile(e.file)
	if err != nil {
		return err
	}

	stringValue := string(bytes)
	stringValue = strings.ReplaceAll(stringValue, "{package}", dir)
	stringValue = strings.ReplaceAll(stringValue, "{struct}", "Entity")

	filePath := path + "/" + dir + "/" + value + ".go"

	return os.WriteFile(filePath, []byte(stringValue), os.ModePerm)
}
