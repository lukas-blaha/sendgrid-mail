package main

import (
	"os"
)

func (app *Config) GetSendgridKey() {
	key := getKeyFromEnv()
	if key != "" {
		app.SendgridKey = key
		return
	}
}

func getKeyFromEnv() string {
	return os.Getenv("SENDGRID_KEY")
}
