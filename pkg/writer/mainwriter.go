package writer

import (
	"github.com/johnfercher/go-arch/pkg"
	"github.com/johnfercher/go-arch/pkg/loader"
	"os"
)

type mainWriter struct {
	loader loader.Loader
	file   string
	node   *pkg.Node
}

func NewMainWriter(loader loader.Loader, node *pkg.Node) *mainWriter {
	return &mainWriter{
		loader: loader,
		file:   "templates/code/main.txt",
		node:   node,
	}
}

func (m *mainWriter) WriteFile(path string, dir string) error {
	bytes, err := m.loader.LoadFile(m.file)
	if err != nil {
		return err
	}

	filePath := path + "/" + dir + "/main.go"

	return os.WriteFile(filePath, bytes, os.ModePerm)
}
