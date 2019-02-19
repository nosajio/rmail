package main

import (
	"bytes"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
	"text/template"
)

// Email represents the email type
type Email struct {
	MessageChannel
}

// Send sends an email to the RECIPIENT_EMAIL using the Sendgrid go SDK
func (e *Email) Send() error {
	host := "https://api.sendgrid.com"
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	req := sendgrid.GetRequest(sendgridAPIKey, "/v3/mail/send", host)
	req.Method = "POST"
	req.Body = []byte(e.sendgridReqJSON())
	res, err := sendgrid.API(req)
	if err != nil || res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("There was a problem sending the email (%s)", err.Error())
	}
	fmt.Printf("Email sent (%d) to: %s, from: %s\n", res.StatusCode, e.ToEmail, e.FromEmail)
	return nil
}

// Create the sendgrid json request object
func (e Email) sendgridReqJSON() string {
	w := new(bytes.Buffer)
	s := `{
		"personalizations": [
			{
				"to": [{
					"email": "{{.ToEmail}}",
					"name": "{{.ToName}}"
				}],
				"subject": "{{.Subject}}"
			}
		],
		"from" : {
			"email": "{{.FromEmail}}",
			"name": "{{.FromName}}"
		},
		"reply_to" : {
			"email": "{{.FromEmail}}",
			"name": "{{.FromName}}"
		},
		"content": [{
			"type": "text/plain",
			"value": "{{.TextBody}}"
		}]
	}`
	t := template.Must(template.New("email").Parse(s))
	t.Execute(w, e)
	return w.String()
}
