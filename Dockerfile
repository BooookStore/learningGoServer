FROM golang:1.9 AS builder

VOLUME /logs

# Copy project
RUN mkdir -p src/github.com/BooookStore/learningGoServer
WORKDIR src/github.com/BooookStore/learningGoServer
COPY ./ ./

# Build
RUN go build

# Run Server
ENTRYPOINT ./learningGoServer
