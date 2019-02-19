btp: build tag push

build:
	@echo "Building ${IMAGE_NAME} for release..."
	docker build \
		--build-arg PORT=${PORT} \
		--build-arg SENDGRID_API_KEY=${SENDGRID_API_KEY} \
		--build-arg RECIPIENT_EMAIL=${RECIPIENT_EMAIL} \
		--build-arg RECIPIENT_NAME=${RECIPIENT_NAME} \
		--build-arg RECIPIENT_SLACK_WEBHOOK_URL=${RECIPIENT_SLACK_WEBHOOK_URL} \
		--build-arg ALLOWED_ORIGINS=${ALLOWED_ORIGINS} \
		-t ${IMAGE_NAME}:latest .
	
	docker build -t ${IMAGE_NAME}:latest .

tag:
	@echo "Tagging..."
	docker tag ${IMAGE_NAME}:latest ${REPOSITORY_URL}/${REPOSITORY_NAME}:latest

push:
	@echo "Pushing..."
	docker push ${REPOSITORY_URL}/${REPOSITORY_NAME}:latest