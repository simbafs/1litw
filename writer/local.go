package writer

import (
	"log"
	"net/http"
	"os"
	"path"
)

type Local struct {
	code string
	Base string
}

func (l Local) SetCode(code string) Writer {
	l.code = code
	return l
}

func (l Local) Write(content []byte) (int, error) {
	// if base/code doesn't exist, create it. if it does but bot a directory, return an error
	if _, err := os.Stat(path.Join(l.Base, l.code)); os.IsNotExist(err) {
		if err := os.MkdirAll(path.Join(l.Base, l.code), 0755); err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	// write content to base/l.code/index.html
	log.Println(l.Base, path.Join(l.Base, l.code, "index.html"))
	f, err := os.OpenFile(path.Join(l.Base, l.code, "index.html"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(content)
}

func (l Local) ListenAndServe(addr string) error {
	// if l.Base doesn't exist, create it
	if _, err := os.Stat(l.Base); os.IsNotExist(err) {
		if err := os.MkdirAll(l.Base, 0755); err != nil {
			return err
		}
	}

	http.Handle("/", http.FileServer(http.Dir(l.Base)))
	return http.ListenAndServe(addr, nil)
}
