FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app

FROM alpine:3.20

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app", "server"]

