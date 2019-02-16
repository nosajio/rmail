# rmail: A Messaging API

rmail is a lightweight API for routing form data from client side applications to various message outlets

## Supported outlets
- Email (SendGrid)
- Slack webhook

## Install
1. Clone this repo.
2. Run `go get` to get dependencies.
3. Add environment variables to the current env:
  ```bash
  export PORT=3344
  export SENDGRID_API_KEY=
  export RECIPIENT_EMAIL=
  export RECIPIENT_SLACK_WEBHOOK_URL=
  ```

_Note: omitting the slack webhook url will disable slack functionality_

## Anatomy of the message request
