FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY app .

RUN go mod init github.com/brtmvdl/cycle

RUN go mod tidy

RUN go build main.go

CMD ./main
