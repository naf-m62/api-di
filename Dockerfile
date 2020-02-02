FROM golang:1.10-alpine

WORKDIR /go/src/api-di

COPY . .

RUN go build -o ~/go/bin/api_di

ENV SERVER_PORT=8093
ENV MONGO_URL=mongo:27017

CMD ~/go/bin/api_di

