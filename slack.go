package main

import ()

// Slack describes data for a slack message
type Slack struct {
	MessageChannel
	WebhookURL string
}

// Send will send the slack message
func (s *Slack) Send() error {
	if s.WebhookURL == "" {
		return nil
	}
	return nil
}
