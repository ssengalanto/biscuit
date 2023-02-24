FROM golang:1.20

ARG BIN

ENV BIN=${BIN}
ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Create a location in the container for the source code.
RUN mkdir -p /app

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go.* /app/

# TODO: update viper config to look for a yaml file instead of .env
COPY .env.development /app/

WORKDIR /app
RUN go mod download
RUN go mod verify

# Copy the source code into the container.
COPY pkg pkg
COPY cmd/"$BIN" cmd/"$BIN"

# Start the application.
CMD go run ./cmd/account/main.go
