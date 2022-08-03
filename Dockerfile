# syntax=docker/dockerfile:1

# start go server
FROM golang:1.18-alpine
WORKDIR /app/backend
COPY ./backend ./
RUN go mod download 
RUN go build

# start react app
FROM node:17.0.0
WORKDIR /app/frontend
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend ./

WORKDIR /app

EXPOSE 3000

CMD [ "./backend/real-time-forum", "./frontend/npm", "./frontend/start" ]