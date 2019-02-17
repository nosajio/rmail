FROM golang:alpine

# Configure ENV variables. Run go build with --build-arg PORT=8080 --build-arg SENDGRID_API_KEY=... etc
ENV PORT                        = ${PORT}
ENV SENDGRID_API_KEY            = ${SENDGRID_API_KEY}
ENV RECIPIENT_EMAIL             = ${RECIPIENT_EMAIL}
ENV RECIPIENT_SLACK_WEBHOOK_URL = ${RECIPIENT_SLACK_WEBHOOK_URL}

RUN mkdir -p /usr/rmail/bin
ADD . /usr/rmail
WORKDIR /usr/rmail

# Compile the binary
RUN go build -o /usr/rmail/bin/production .

# Same as PORT environment variable. The listening port
EXPOSE ${PORT}

# Run the binary
CMD ["./bin/production"]