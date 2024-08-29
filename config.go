package main

import "os"

type Config struct {
	TgToken     string
	GitHubToken string
	Origin      string

	// github repo
	User   string
	Repo   string
	Branch string
}

func FromEnv() Config {
	return Config{
		TgToken:     os.Getenv("TG_TOKEN"),
		GitHubToken: os.Getenv("GITHUB_TOKEN"),
		Origin:      os.Getenv("ORIGIN"),
		User:        os.Getenv("GITHUB_USER"),
		Repo:        os.Getenv("GITHUB_REPO"),
		Branch:      os.Getenv("GITHUB_BRANCH"),
	}
}
