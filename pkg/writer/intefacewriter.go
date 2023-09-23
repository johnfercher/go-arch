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

func (i *interfaceWriter) WriteFile(path string, dir string, value string) error {
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

	values := strings.Split(value, ">")
	cleanedValue := strings.ReplaceAll(values[0], "Interface", "")

	upperCleaned := cleanedValue
	upperCleaned = strings.Replace(upperCleaned, string(upperCleaned[0]), strings.ToUpper(string(upperCleaned[0])), 1)

	stringValue := string(bytes)
	stringValue = strings.ReplaceAll(stringValue, "{package}", dir)
	stringValue = strings.ReplaceAll(stringValue, "{entity_package}", dir)
	stringValue = strings.ReplaceAll(stringValue, "{dir}", path)
	stringValue = strings.ReplaceAll(stringValue, "{path}", path)
	stringValue = strings.ReplaceAll(stringValue, "{interface}", upperCleaned)
	stringValue = strings.ReplaceAll(stringValue, "{module}", goMod)

	filePath := path + "/" + dir + "/" + cleanedValue + ".go"

	return os.WriteFile(filePath, []byte(stringValue), os.ModePerm)
}
