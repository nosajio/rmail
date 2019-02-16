package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// EmailReqBody describes the body of the request sent to the handle fn
type EmailReqBody struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Subject string `json:"subject"`
}

// EmailResBody describes what the PostMail handler responds with
type EmailResBody struct {
	Sent  bool    `json:"sent"`
	Error *string `json:"error"`
}

// HandlePostMail handles sending an email to the address in the RECIPIENT_EMAIL
// environment variable
func HandlePostMail(w http.ResponseWriter, r *http.Request) {
	b, _ := readMailBody(r.Body)
	e := Email{
		FromEmail: b.Email,
		FromName:  b.Name,
		Subject:   b.Subject,
		ToEmail:   os.Getenv("RECIPIENT_EMAIL"),
		ToName:    os.Getenv("RECIPIENT_NAME"),
	}
	e.WriteBody(b.Message)
	var res []byte
	if err := e.Send(); err != nil {
		fmt.Printf("Email not sent. f: %s, m: %s\n---\nerr: %v\n---\n", e.FromEmail, e.TextBody, err)
		err := "There was an issue sending the message"
		res, _ = json.Marshal(EmailResBody{Sent: false, Error: &err})
	} else {
		res, _ = json.Marshal(EmailResBody{Sent: true, Error: nil})
	}
	w.Write(res)
}

func readMailBody(b io.ReadCloser) (e *EmailReqBody, err error) {
	bodyBytes, _ := ioutil.ReadAll(b)
	// Build the email
	var bodyFields EmailReqBody
	json.Unmarshal(bodyBytes, &bodyFields)
	return &bodyFields, nil
}
