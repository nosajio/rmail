# rmail: A Messaging API

rmail (rest + email) is a lightweight API for creating and routing messages from client side applications to various outlets.

## Supported outlets
- Email (SendGrid smtp api)
- Slack (webhook)

## Run
1. Clone this repo.
2. Run `go get` to get dependencies.
3. Add environment variables to the current env:
```bash
export PORT=8080
export SENDGRID_API_KEY=
export RECIPIENT_EMAIL=
export RECIPIENT_SLACK_WEBHOOK_URL=
export ALLOWED_ORIGINS=*
```
4. After exporting env variables, build and run with `go run *.go`

_Note: omitting the slack webhook url will disable slack functionality_

## Anatomy of the message request
```http
POST /message
Content-Type: application/json

{
    "email": "somebody@example.com",
    "message": "The body of the message....",
    "name": "Somebody",
    "subject": "Work enquiry"
}
```

## Docker deploys
Build the image with:
```bash
docker build \
  --build-arg PORT=8080 \
  --build-arg SENDGRID_API_KEY=xxxxxx \
  --build-arg RECIPIENT_EMAIL=you@example.com \
  --build-arg RECIPIENT_NAME=xxx \
  --build-arg RECIPIENT_SLACK_WEBHOOK_URL=https://xxx \
  --build-arg ALLOWED_ORIGINS=* \
  -t rmail:latest .
```