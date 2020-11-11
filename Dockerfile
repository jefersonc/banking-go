# Development stage
FROM golang:1.15.4-alpine3.12 AS dev

ADD . /code

WORKDIR /code

RUN go run main.go

# Compile stage
FROM golang:1.15.4-alpine3.12 AS build

ADD . /code

WORKDIR /code

RUN go build -o /main

# Final stage
FROM alpine:latest as prod

EXPOSE 8000

WORKDIR /

COPY --from=build /main /

CMD ["/main"]