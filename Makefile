btp: build tag push

br: build run

build:
	@echo "Building ${IMAGE_NAME} for release..."
	source $(shell pwd)/.env.production && docker build \
		--build-arg PORT=${PORT} \
		--build-arg SENDGRID_API_KEY=${SENDGRID_API_KEY} \
		--build-arg RECIPIENT_EMAIL=${RECIPIENT_EMAIL} \
		--build-arg RECIPIENT_NAME=${RECIPIENT_NAME} \
		--build-arg RECIPIENT_SLACK_WEBHOOK_URL=${RECIPIENT_SLACK_WEBHOOK_URL} \
		--build-arg ALLOWED_ORIGINS=${ALLOWED_ORIGINS} \
		-t ${IMAGE_NAME}:latest $(shell pwd)

tag:
	@echo "Tagging..."
	docker tag ${IMAGE_NAME}:latest ${REPOSITORY_URL}/${REPOSITORY_NAME}:latest

push:
	@echo "Pushing..."
	docker push ${REPOSITORY_URL}/${REPOSITORY_NAME}:latest

run: 
	@echo "Running temporary image..."
	docker run --rm -p ${PORT}:${PORT} ${IMAGE_NAME}:latest