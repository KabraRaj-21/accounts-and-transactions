# Use an official Golang image as a build stage
FROM golang:1.23 AS build-env

# Automatically use the environment's OS and Architecture
ARG TARGETOS
ARG TARGETARCH

ENV GO111MODULE=on
ENV TZ=Asia/Kolkata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY . /accounts-and-transactions
WORKDIR /accounts-and-transactions

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the entire project
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o accounts-and-transactions ./cmd/main.go

RUN chmod +x accounts-and-transactions

EXPOSE 8080
CMD ["./accounts-and-transactions"]
