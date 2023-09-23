package writer

import (
	"github.com/johnfercher/go-arch/pkg"
	"github.com/johnfercher/go-arch/pkg/loader"
	"os"
	"strings"
)

type interfaceWriter struct {
	loader loader.Loader
	file   string
	goMod  string
	node   *pkg.Node
}

func NewInterfaceWriter(loader loader.Loader, node *pkg.Node) *interfaceWriter {
	return &interfaceWriter{
		loader: loader,
		file:   "templates/code/interface.txt",
		goMod:  "go.mod",
		node:   node,
	}
}

func (i *interfaceWriter) WriteFile(path string, dir string) error {
	bytes, err := i.loader.LoadFile(i.goMod)
	if err != nil {
		return err
	}

	firstLine := strings.Split(string(bytes), "\n")[0]
	goMod := strings.ReplaceAll(firstLine, "module ", "")

	bytes, err = i.loader.LoadFile(i.file)
	if err != nil {
		return err
	}

	stringValue := string(bytes)
	stringValue = strings.ReplaceAll(stringValue, "{package}", dir)
	stringValue = strings.ReplaceAll(stringValue, "{interface}", "Interface")
	stringValue = strings.ReplaceAll(stringValue, "{module}", goMod)

	return os.WriteFile(path+"/"+dir+"/interface.go", []byte(stringValue), os.ModePerm)
}
