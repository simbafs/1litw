package main

import (
	"1li/ent"
	"log"
)

func StaticGenOne(rec *ent.Record) error {
	log.Printf("Generating static for %s -> %s\n", rec.Code, rec.Target)

	return nil
}
