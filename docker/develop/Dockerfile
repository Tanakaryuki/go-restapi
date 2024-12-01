FROM golang:1.23.3-alpine3.19

WORKDIR /go/src/go-restapi

RUN apk update && apk add git
RUN go install github.com/air-verse/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]