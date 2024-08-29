package siteconfig

import (
	"1li/db"
	"1li/ent"
	"1li/ent/siteconfig"
	"context"
	"os"
)

func Get(ctx context.Context, site string) (*ent.SiteConfig, error) {
	return db.Client.SiteConfig.Query().Where(siteconfig.Site(site)).Only(ctx)
}

// TODO: A TUI to set config
func NewSite(ctx context.Context) (*ent.SiteConfig, error) {
	site := db.Client.SiteConfig.Create()

	site.SetSite(os.Getenv("SITE"))
	site.SetOrigin(os.Getenv("ORIGIN"))
	site.SetGhToken(os.Getenv("GH_TOKEN"))
	site.SetTgToken(os.Getenv("TG_TOKEN"))

	return site.Save(ctx)
}
