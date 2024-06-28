docker-build:
	docker build . -t snippetbox:1.0.0

compose-clean:
	docker-compose down && docker image rm snippetbox:1.0.0