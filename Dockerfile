FROM golang:1.21.0

WORKDIR /usr/src/app

# Sets up Go's hot reload
RUN go install github.com/cosmtrek/air@latest
 
COPY . .
RUN go mod download
