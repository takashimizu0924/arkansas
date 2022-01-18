FROM golang:1.15.7-alpine as dev

RUN apk update && apk add git

RUN mkdir /var/src

WORKDIR /var/src

# COPY go.mod /var/src
# COPY go.sum /var/src/go.sum

# RUN go mod download

ADD . /var/src

# CMD [ "go","run","main.go" ]
