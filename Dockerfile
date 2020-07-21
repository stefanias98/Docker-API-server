FROM golang:latest
COPY . ./files
WORKDIR ./files
CMD go run main.go
