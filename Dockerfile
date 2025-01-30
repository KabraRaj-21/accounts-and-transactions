# Use an official Golang image as a build stage
FROM golang:1.23 AS build-env

ENV GO111MODULE=on

# Automatically use the environment's OS and Architecture
ARG TARGETOS
ARG TARGETARCH

ADD . /src/transaction
WORKDIR /src/transaction

COPY go.mod .
COPY go.sum .

RUN go mod tidy

# Copy the entire project
COPY . .

# Build the binary
RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -tags=musl,dynamic -o transaction ./cmd/main.go

RUN chmod +x transaction

EXPOSE 8080
CMD ["./transaction"]
