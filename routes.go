package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
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

var (
	slackChan = &Slack{WebhookURL: os.Getenv("SLACK_WEBHOOK_URL")}
	emailChan = &Email{}
	channels  = []Message{emailChan, slackChan}
)

// HandlePostMessage handles sending an email to the address in the RECIPIENT_EMAIL
// environment variable
func HandlePostMessage(w http.ResponseWriter, r *http.Request) {
	var res []byte
	b, _ := readRequestParams(r.Body)
	m, _ := initChannels(b, channels...)
	err := sendMessages(m)
	if err != nil {
		err := err.Error()
		res, _ = json.Marshal(MsgResBody{Sent: false, Error: &err})
	} else {
		res, _ = json.Marshal(MsgResBody{Sent: true, Error: nil})
	}

	w.Write(res)
}

func sendMessages(m []Message) error {
	for _, mm := range m {
		if mm == nil {
			continue
		}
		err := mm.Send()
		if err != nil {
			fmt.Printf("Error sending message: \"%s\" (%s)\n", reflect.TypeOf(mm), err.Error())
			return err
		}
	}
	return nil
}

func initChannels(b *MsgReqBody, messages ...Message) (m []Message, err error) {
	var o []Message
	for _, m := range messages {
		m.WriteMessage(os.Getenv("RECIPIENT_NAME"), os.Getenv("RECIPIENT_EMAIL"), b.Name, b.Email, b.Subject, b.Message)
		o = append(o, m)
	}
	return o, nil
}

// func sendEmail(b *MsgReqBody) error {
// 	e := Email{
// 		FromEmail: b.Email,
// 		FromName:  b.Name,
// 		Subject:   b.Subject,
// 		ToEmail:   os.Getenv("RECIPIENT_EMAIL"),
// 		ToName:    os.Getenv("RECIPIENT_NAME"),
// 	}
// 	e.WriteBody(b.Message)
// 	if err := e.Send(); err != nil {
// 		fmt.Printf("Email not sent. f: %s, m: %s\n---\nerr: %v\n---\n", e.FromEmail, e.TextBody, err)
// 		return errors.New("There was an issue sending the message")
// 	}
// 	return nil
// }

func readRequestParams(b io.ReadCloser) (e *MsgReqBody, err error) {
	bodyBytes, _ := ioutil.ReadAll(b)
	// Build the email
	var bodyFields MsgReqBody
	json.Unmarshal(bodyBytes, &bodyFields)
	return &bodyFields, nil
}
