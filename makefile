ssh-dev:
	docker run -it --entrypoint /bin/bash ds-dit_api-dev

ssh-prod:
	docker run -it --entrypoint /bin/sh ds-dit_api-prod

dev-up:
	docker-compose -f docker-compose.dev.yml up

dev-build-up:
	docker-compose -f docker-compose.dev.yml up --build

prod-up:
	docker-compose -f docker-compose.prod.yml up

prod-build-up:
	docker-compose -f docker-compose.prod.yml up --build

images := $(shell docker images -a -q)
rmi: # remove all images
	docker rmi -f $(images)

