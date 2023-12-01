package main

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
