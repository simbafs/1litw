package writer

type Writer interface {
	Write(content []byte) (n int, err error)
	CD(path string) Writer
}
