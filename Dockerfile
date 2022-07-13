FROM golang:alpine

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=8888

WORKDIR /go/src/mburoja-api

COPY . ./
# Run the two commands below to install git and dependencies for the project. 
# RUN apk update && apk add --no-cache git
# RUN go get ./...


RUN go build mburoja-api .

EXPOSE $PORT

ENTRYPOINT ["./mburoja-api"]
