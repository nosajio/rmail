package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// MsgReqBody describes the body of the request sent to the handle fn
type MsgReqBody struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Subject string `json:"subject"`
}

// MsgResBody describes what the PostMail handler responds with
type MsgResBody struct {
	Sent  bool    `json:"sent"`
	Error *string `json:"error"`
}

// HandlePostMessage handles sending an email to the address in the RECIPIENT_EMAIL
// environment variable
func HandlePostMessage(w http.ResponseWriter, r *http.Request) {
	var res []byte
	b, _ := readRequestParams(r.Body)
	err := sendEmail(b)
	if err != nil {
		err := err.Error()
		res, _ = json.Marshal(MsgResBody{Sent: false, Error: &err})
	} else {
		res, _ = json.Marshal(MsgResBody{Sent: true, Error: nil})
	}
	w.Write(res)
}

// sendSlack will send the message in b to the defined slack webhook. Note that
// if the webhook is not defined this method will not run
func sendSlack(b *MsgReqBody, slackAPIKey string) error {
	if slackAPIKey == "" {
		return nil
	}
	return nil
}

func sendEmail(b *MsgReqBody) error {
	e := Email{
		FromEmail: b.Email,
		FromName:  b.Name,
		Subject:   b.Subject,
		ToEmail:   os.Getenv("RECIPIENT_EMAIL"),
		ToName:    os.Getenv("RECIPIENT_NAME"),
	}
	e.WriteBody(b.Message)
	if err := e.Send(); err != nil {
		fmt.Printf("Email not sent. f: %s, m: %s\n---\nerr: %v\n---\n", e.FromEmail, e.TextBody, err)
		return errors.New("There was an issue sending the message")
	}
	return nil
}

func readRequestParams(b io.ReadCloser) (e *MsgReqBody, err error) {
	bodyBytes, _ := ioutil.ReadAll(b)
	// Build the email
	var bodyFields MsgReqBody
	json.Unmarshal(bodyBytes, &bodyFields)
	return &bodyFields, nil
}
