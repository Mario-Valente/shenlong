FROM golang:1.23.3-alpine3.20

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app

ENTRYPOINT ["/go/bin/app", "server"]

