build:
	docker build -t hjkelly/budget-category-service .

shell:
	docker run -v ~/.config:/root/.config --env GOOGLE_APPLICATION_CREDENTIALS=/root/.config/gcloud/application_default_credentials.json -itv `pwd`:/go/src/github.com/hjkelly/budget-category-service/ -p 8080:8080 hjkelly/budget-category-service:latest bash

server:
	docker run -v ~/.config:/root/.config --env GOOGLE_APPLICATION_CREDENTIALS=/root/.config/gcloud/application_default_credentials.json -v `pwd`:/go/src/github.com/hjkelly/budget-category-service/ -p 8080:8080 hjkelly/budget-category-service:latest

