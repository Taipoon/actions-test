FROM golang:1.19-alpine3.13 AS builder

WORKDIR /go/src

ENV GO111MODULE=on
ENV GOPATH=

COPY go.mod go.sum *.go ./

RUN go mod download

COPY ./main.go ./

RUN go build -o /go/bin/main -ldflags '-s -w' ./main.go


FROM alpine:3.13

WORKDIR /app

COPY --from=builder /go/bin/main .

COPY .env .

CMD [ "/app/main" ]
