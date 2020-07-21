FROM golang:latest
COPY . ./files
WORKDIR ./files
RUN go get github.com/gorilla/mux 
CMD go run main.go
