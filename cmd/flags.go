package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	To      []string
	From    string
	Subject string
	Body    string
}

func NewData() *Data {
	d := Data{}
	var to string

	flag.StringVar(&to, "to", "", "Specify message recepients (comma separated list)")
	flag.StringVar(&d.From, "from", "", "Specify sender address")
	flag.StringVar(&d.Subject, "subject", "", "Specify message subject")
	flag.StringVar(&d.Body, "body", "", "Specify message body")

	flag.Parse()

	// Check for a double dash and handle positional arguments
	for i, arg := range os.Args {
		if arg == "--" {
			flag.CommandLine.Parse(os.Args[i+1:])
			break
		}
	}

	for _, address := range strings.Split(to, ",") {
		d.To = append(d.To, address)
	}

	if to == "" || d.From == "" {
		fmt.Println("You have to specify sender address and recepient(s)")
		fmt.Print("Please see help using -h or --help\n\n")
		os.Exit(0)
	}

	return &d
}
