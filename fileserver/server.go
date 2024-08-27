package fileserver

import (
	"io/fs"
	"net/http"
	"os"
)

// initServer inits basic files and directories for the server.
func Init(local string) error {
	// mkdir ./local
	if err := os.MkdirAll(local, 0755); err != nil {
		return err
	}

	// TODO: inject embed fs to ./local
	return nil
}

// ListenAndServe serves files from the provided fs.FS.
func ListenAndServe(addr string, files fs.FS) error {
	http.Handle("/", http.FileServer(http.FS(files)))
	return http.ListenAndServe(addr, nil)
}
