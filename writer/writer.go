package writer

type Writer interface {
	SetCode(code string) Writer
	Write(content []byte) (int, error)
}
