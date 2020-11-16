DOCKER_RUN = docker-compose run --rm runner

run:
	- ${DOCKER_RUN} go run main.go

test:
	- ${DOCKER_CONTAINER} go test -covermode="count" ./...

install-dependencies:
	- ${DOCKER_RUN} go mod vendor

sh:
	- ${DOCKER_RUN} sh