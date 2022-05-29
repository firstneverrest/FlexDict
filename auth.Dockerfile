FROM golang:1.18-alpine

WORKDIR /app

COPY ./auth/go.mod .
COPY ./auth/go.sum .
RUN go mod download

COPY ./auth .

RUN go build -o ./bin/auth ./cmd/*.go

EXPOSE 8080

CMD ["./bin/auth"]