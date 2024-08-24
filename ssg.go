package main

import (
	"1li/ent"
	"fmt"
	"log"
	"os"
	"path"
)

func StaticGenOne(rec *ent.Record) error {
	content := []byte(fmt.Sprintf("<meta http-equiv=\"refresh\" content=\"0; url=%s\"><p>Redirecting to <a href=\"%s\">%s</a></p>", rec.Target, rec.Target, rec.Target))

	// mkdir ./local/rec.Code
	if err := os.MkdirAll(path.Join(".", "static", rec.Code), 0755); err != nil {
		return err
	}

	// write to ./local/rec.Code/index.html
	if err := os.WriteFile(path.Join(".", "static", rec.Code, "index.html"), content, 0644); err != nil {
		return err
	}

	log.Printf("Generate link for %s -> %s\n", rec.Code, rec.Target)

	return nil
}
