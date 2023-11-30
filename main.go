package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type To struct {
	Email string `json:"email"`
}

type From struct {
	Email string `json:"email"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Personalizations struct {
	To []To `json:"to"`
}

type MsgBody struct {
	Personalizations []Personalizations `json:"personalizations"`
	From             From               `json:"from"`
	Subject          string             `json:"subject"`
	Content          []Content          `json:"content"`
}

func NewMsgBody(to []string, from, subject, body string) *MsgBody {
	var receivers []To
	for _, address := range to {
		receivers = append(receivers, To{
			Email: address,
		})
	}

	sender := From{
		Email: from,
	}

	message := Content{
		Type:  "text/plain",
		Value: body,
	}

	return &MsgBody{
		Personalizations: []Personalizations{
			{
				To: receivers,
			},
		},
		From:    sender,
		Subject: subject,
		Content: []Content{
			message,
		},
	}
}

func main() {
	msg := NewMsgBody([]string{"mail@example.com"}, "test@example.com", "Test subject", "Test body")

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

	r.Header.Set("Authorization", "Bearer")
	r.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
