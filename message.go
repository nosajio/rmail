package main

// Message manages formatting message data and sending it off to a service
type Message interface {
	WriteMessage(toName string, toEmail string, fromName string, fromEmail string, subject string, body string) error
	Send() error
}

// MessageChannel is the datatype used by Message implementations
type MessageChannel struct {
	FromEmail, FromName, ToName, ToEmail, Subject, TextBody string
}

// WriteMessage writes client data to an MessageChannel
func (e *MessageChannel) WriteMessage(toName string, toEmail string, fromName string, fromEmail string, subject string, body string) error {
	e.FromEmail = fromEmail
	e.FromName = fromName
	e.Subject = subject
	e.ToEmail = toEmail
	e.ToName = toName
	e.TextBody = body
	return nil
}
