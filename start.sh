#!/bin/bash

pwd
cd ./frontend
npm install && npm start &
cd ./../backend
go mod download && go run main.go