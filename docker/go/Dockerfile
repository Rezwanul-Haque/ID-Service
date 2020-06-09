# Start from golang base image
FROM golang:1.14-alpine3.11 AS builder

#ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Rezwanul Haque <rezwanul.cse@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Configure the repo url
ENV REPO_URL=github.com/Rezwanul-Haque/ID-Service

# Setup our $GOPATH
ENV GOPATH=/app

# Set the current working directory inside the container
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src/

COPY src $WORKPATH

WORKDIR $WORKPATH

RUN pwd
RUN ls

# Copy go mod and sum files
COPY ./src/go.mod .
COPY ./src/go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o id-service .

# Start a new stage from scratch
#FROM alpine:latest

#RUN ls bin/
## Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
#COPY --from=builder /app/id-service .
#COPY --from=builder /app/.env .

# Expose port 7001 to the outside world
EXPOSE 7001

# Command to run the executable
ENTRYPOINT ["./id-service"]