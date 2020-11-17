DOCKER_RUN = docker-compose run --rm runner

run: install-dependencies
	- docker-compose up 

test:
	- ${DOCKER_CONTAINER} go test -covermode="count" ./...

install-dependencies:
	- ${DOCKER_RUN} go mod vendor
