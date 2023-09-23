package writer

import (
	"github.com/johnfercher/go-arch/pkg"
	"github.com/johnfercher/go-arch/pkg/loader"
	"strings"
)

type Factory interface {
	Create(value string) Writer
}

type factory struct {
	writers map[string]Writer
}

func New(loader loader.Loader, node *pkg.Node) *factory {
	writers := make(map[string]Writer)
	writers["main.go"] = NewMainWriter(loader, node)
	writers["entity"] = NewEntityWriter(loader, node)
	writers["interface"] = NewInterfaceWriter(loader, node)

	f := &factory{
		writers: writers,
	}

	return f
}

func (f factory) Create(value string) Writer {
	writer, ok := f.writers[value]
	if ok {
		return writer
	}

	arr := strings.Split(value, ">")

	if strings.Contains(arr[0], "Interface") {
		return f.writers["interface"]
	}

	return nil
}
