# syntax=docker/dockerfile:1

# start go server
FROM golang:1.18-alpine

WORKDIR /app

RUN apk add build-base gcompat
RUN CGO_ENABLED=0
RUN GOOS=linux


# RUN apk add --no-cache libc6-compat 

COPY go.mod .
COPY go.sum .
RUN go mod download 

COPY . .
RUN go build -a -installsuffix cgo -o main 
RUN chmod +x main

EXPOSE 4000

# CMD ./build/main
ENTRYPOINT main