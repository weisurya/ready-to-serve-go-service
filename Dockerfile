FROM golang:1.12-alpine

# Allow Go to retrieve the dependencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 admin
RUN mkdir /gopherconuk/ && chown admin /gopherconuk/
USER admin

ADD . /go/src/ready-to-serve-go-service
WORKDIR /go/src/ready-to-serve-go-service

RUN go get ready-to-serve-go-service

ENTRYPOINT [ "go/bin/ready-to-serve-go-service" ]