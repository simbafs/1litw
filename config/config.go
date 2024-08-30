package config

import "os"

type Config struct {
	TgToken string
	Origin  string
	Addr    string

	// github repo
	// GitHubToken string
	// User   string
	// Repo   string
	// Branch string

	// local writer
	Base string
}

func FromEnv() Config {
	return Config{
		TgToken: os.Getenv("TG_TOKEN"),
		Addr:    os.Getenv("ADDR"),

		Origin: os.Getenv("ORIGIN"),
		Base:   os.Getenv("BASE"),

		// GitHubToken: os.Getenv("GITHUB_TOKEN"),
		// User:        os.Getenv("GITHUB_USER"),
		// Repo:        os.Getenv("GITHUB_REPO"),
		// Branch:      os.Getenv("GITHUB_BRANCH"),
	}
}
