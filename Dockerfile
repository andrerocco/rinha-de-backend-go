FROM golang:1.21.0

WORKDIR /usr/src/app

# Copy the source code into the container
COPY . .

# Set up Go modules
RUN go mod download
