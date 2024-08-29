package ssg

import (
	"1li/db"
	"1li/ent"
	"1li/errorCollector"
	"1li/writer"
	"context"
	"fmt"
	"io"
	"log"
	"sync"
)

// StaticGenOne generate a static page for a record.
func StaticGenOne(rec *ent.Record, w io.Writer) error {
	_, err := fmt.Fprintf(w, "<meta http-equiv=\"refresh\" content=\"0; url=%s\"><p>Redirecting to <a href=\"%s\">%s</a></p>", rec.Target, rec.Target, rec.Target)
	if err != nil {
		return err
	}

	// // mkdir ./local/rec.Code
	// if err := os.MkdirAll(path.Join(".", "static", rec.Code), 0755); err != nil {
	// 	return err
	// }
	//
	// // write to ./local/rec.Code/index.html
	// if err := os.WriteFile(path.Join(".", "static", rec.Code, "index.html"), content, 0644); err != nil {
	// 	return err
	// }

	// log.Printf("Generate link for %s -> %s\n", rec.Code, rec.Target)

	return nil
}

// SyncFromDB generate static pages for all records in the database.
func SyncFromDB(w writer.Writer) error {
	recs, err := db.Client.Record.Query().All(context.Background())
	if err != nil {
		return err
	}

	errs := errorCollector.New()

	wg := sync.WaitGroup{}
	wg.Add(len(recs))

	for _, rec := range recs {
		go func(rec *ent.Record) {
			defer wg.Done()
			err := StaticGenOne(rec, w.CD(rec.Code+"/index.html"))
			errs.Add(err)
		}(rec)
	}

	wg.Wait()

	log.Println("sync done", err)
	return errs.Join()
}
