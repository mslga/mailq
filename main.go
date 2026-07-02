package main

import (
	"flag"
	"log"
	"os"
	"strings"

	mail "github.com/mslga/mailq/internal/mail"
)

func main() {
	optHost := flag.String("h", "smtp.company.com", "SMTP server hostname")
	optPort := flag.Int("p", 25, "SMTP server port number")
	optInsecure := flag.Bool("i", false, "Skip TLS certificate verification (use only for trusted SMTP servers)")
	optSender := flag.String("s", "mailq.sender@company.com", "Mail sender")
	optPass := flag.String("w", "", "SMTP password (optional)")
	optRecipients := flag.String("r", "recipient@company.com", "Mail recipients. Separate multiple recipients with spaces or commas")
	optSubject := flag.String("u", "Empty subject", "Mail subject")
	optMessage := flag.String("m", "Empty message", "Mail message in HTML format")

	flag.Parse()

	var recipients []string

	if strings.Contains(*optRecipients, ",") {
		for _, r := range strings.Split(*optRecipients, ",") {
			r = strings.TrimSpace(r)
			if r != "" {
				recipients = append(recipients, r)
			}
		}
	} else {
		recipients = strings.Fields(*optRecipients)
	}

	if len(recipients) == 0 {
		log.Println("no recipients specified")
		os.Exit(1)
	}

	if err := mail.SendMail(*optHost, *optPort, *optSender, *optPass, *optInsecure, recipients, *optSubject, *optMessage); err != nil {
		log.Fatalf("failed to send mail: %v", err)
	}
}
