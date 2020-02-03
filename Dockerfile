FROM golang:1.10-alpine

WORKDIR /go/src/api-di

COPY . .

RUN go build -o ~/go/bin/api_di

CMD ~/go/bin/api_di

