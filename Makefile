DOCKER_RUN = docker-compose run --rm runner

run:
	- ${DOCKER_RUN} go run main.go

install-dependencies:
	- ${DOCKER_RUN} go mod vendor

sh:
	- ${DOCKER_RUN} sh