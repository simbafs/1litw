package main

import (
	"io/fs"
	"net/http"
	"os"
)

// initServer inits basic files and directories for the server.
func initServer(local string) error {
	// mkdir ./local
	if err := os.MkdirAll(local, 0755); err != nil {
		return err
	}

	// TODO: inject embed fs to ./local
	return nil
}

// fileServer serves files from the provided fs.FS.
func fileServer(addr string, files fs.FS) error {
	http.Handle("/", http.FileServer(http.FS(files)))
	return http.ListenAndServe(addr, nil)
}
