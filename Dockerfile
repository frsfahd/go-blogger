#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./cmd/api/main.go

#final stage
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /go/src/app/.env.vault /
# LABEL Name=goblogger Version=0.0.1
EXPOSE 8080
ENTRYPOINT [ "/app" ]
