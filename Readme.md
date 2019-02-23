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
# Docker build and repo vars
export IMAGE_NAME=rmail
export REPOSITORY_NAME=rmail
export REPOSITORY_URL=

# App runtime vars
export PORT=8080
export SENDGRID_API_KEY=
export RECIPIENT_EMAIL=
export RECIPIENT_SLACK_WEBHOOK_URL=
export ALLOWED_ORIGINS=*
```
4. After exporting env variables, build and run with `go run *.go`

_Note: omit the slack webhook url to disable slack functionality_

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

#### Build the container
```bash
make build
```

#### Tag the container
```bash
make tag

```

#### Push the container 
```bash
make push
```

#### Build, tag, push
```bash
make btp
```