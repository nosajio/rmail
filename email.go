package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"html"
	"os"
	"regexp"
	"text/template"
)

// Email represents the email type
type Email struct {
	MessageChannel
}

// Send sends an email to the RECIPIENT_EMAIL using the Sendgrid go SDK
func (e *Email) Send() error {
	if e.FromName == "" || e.FromEmail == "" {
		return errors.New("Name and email must be provided")
	}
	nlReg := regexp.MustCompile(`(?m)\n|\r`)
	// Ensure the string is properly prepared for sending over HTTP
	e.TextBody = html.EscapeString(
		string(nlReg.ReplaceAll([]byte(e.TextBody), []byte("<br/>"))))
	// Make the api request
	host := "https://api.sendgrid.com"
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	req := sendgrid.GetRequest(sendgridAPIKey, "/v3/mail/send", host)
	req.Method = "POST"
	req.Body = []byte(e.sendgridReqJSON())
	res, err := sendgrid.API(req)
	if err != nil || res == nil || res.StatusCode < 200 || res.StatusCode >= 300 {
		fmt.Printf("There was a problem sending the email. Logging the message instead: %+v", e)
		return err
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
			"type": "text/html",
			"value": "{{.TextBody}}"
		}]
	}`
	t := template.Must(template.New("email").Parse(s))
	t.Execute(w, e)
	return w.String()
}
