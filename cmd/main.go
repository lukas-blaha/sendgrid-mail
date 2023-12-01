package main

import "fmt"

func main() {
	d := NewData()

	fmt.Println(d)

	// msg := NewMsgBody([]string{"mail@example.com"}, "test@example.com", "Test subject", "Test body")
	//
	// j, err := json.Marshal(msg)
	//
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	// body := bytes.NewBuffer(j)
	//
	// client := &http.Client{}
	//
	// r, err := http.NewRequest(http.MethodPost, "https://api.sendgrid.com/v3/mail/send", body)
	//
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	// r.Header.Set("Authorization", "Bearer")
	// r.Header.Set("Content-Type", "application/json")
	//
	// resp, err := client.Do(r)
	//
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	// defer resp.Body.Close()
	//
	// fmt.Println(resp.Status)
}
