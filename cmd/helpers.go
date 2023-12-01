package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (app *Config) GetSendgridKey() {
	key := getKeyFromEnv()
	if key != "" {
		app.SendgridKey = key
		return
	}
	app.SendgridKey = getKeyFromConfig()
}

func getKeyFromEnv() string {
	return os.Getenv("SENDGRID_KEY")
}

func getKeyFromConfig() string {
	var key string

	f, err := os.Open("/etc/sendgrid-mail.conf")
	if err != nil {
		fmt.Println("The config file /etc/sendgrid-mail.conf does not exist or does not have correct permissions set!")
		fmt.Print("Please ensure that config file exists and has the right permissions or create environment variable instead.\n\n")
		os.Exit(1)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		p := strings.Split(sc.Text(), ":")
		if p[0] == "apiKey" {
			key = removeChars(p[1])
		}
	}

	if key == "" {
		fmt.Println("Can't find sendgrid api key!")
		fmt.Print("Please ensure that \"apiKey\" option is specified in config file or create environment variable instead.\n\n")
		os.Exit(1)
	}

	return key
}

func removeChars(s string) string {
	var ns string

	for _, c := range s {
		if c != ' ' && c != '"' {
			ns += string(c)
		}
	}

	return ns
}
