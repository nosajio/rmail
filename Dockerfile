# Stage: 1
FROM golang:alpine AS build

# Add missing packages that aren't in alpine
RUN apk update && apk add --no-cache git

# Setup work dir
RUN mkdir -p /go/src/rmail
WORKDIR /go/src/rmail/
ADD . .

# Build the binary
RUN go get . && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/rmail .

# Stage: 2
FROM alpine:latest

# Configure workdir and move binary from build step
WORKDIR /go/bin/
COPY --from=build /go/bin/rmail .

# Make SSL work
RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

# Configure ENV variables. 
ARG PORT
ARG SENDGRID_API_KEY
ARG RECIPIENT_EMAIL
ARG RECIPIENT_NAME
ARG RECIPIENT_SLACK_WEBHOOK_URL
ENV PORT                        ${PORT}
ENV SENDGRID_API_KEY            ${SENDGRID_API_KEY}
ENV RECIPIENT_EMAIL             ${RECIPIENT_EMAIL}
ENV RECIPIENT_NAME              ${RECIPIENT_NAME}
ENV RECIPIENT_SLACK_WEBHOOK_URL ${RECIPIENT_SLACK_WEBHOOK_URL}

# Run the binary
ENTRYPOINT ["./rmail"]

# Expose listening port
EXPOSE ${PORT}