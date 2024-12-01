FROM golang:1.23.3-alpine3.19 as go-builder

WORKDIR /go/src/techer-me-backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:3.19

WORKDIR /usr/src/techer-me-backend

COPY --from=go-builder /go/src/techer-me-backend/api api
COPY --from=go-builder /go/src/techer-me-backend/api .env
RUN chmod +x "/usr/src/techer-me-backend/api"

ENTRYPOINT ["/usr/src/techer-me-backend/api"]