package writer

type Writer interface {
	WriteFile(path string, dir string, value string) error
}
