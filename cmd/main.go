package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	SendgridKey string
}

func main() {
	app := Config{}
	app.GetSendgridKey()

	d := NewData()

	msg := NewMsgBody(d.To, d.From, d.Subject, d.Body)

	j, err := json.Marshal(msg)

	if err != nil {
		log.Panic(err)
	}

	body := bytes.NewBuffer(j)

	client := &http.Client{}

	r, err := http.NewRequest(http.MethodPost, "https://api.sendgrid.com/v3/mail/send", body)

	if err != nil {
		log.Panic(err)
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app.SendgridKey))
	r.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(r)

	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
