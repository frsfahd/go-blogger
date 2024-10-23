#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/main -v ./cmd/api/main.go

#final stage
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/bin/main /go/src/app/.env.vault /go/src/app/docs/ ./
# LABEL Name=goblogger Version=0.0.1
EXPOSE 8080
ENTRYPOINT [ "/main" ]
