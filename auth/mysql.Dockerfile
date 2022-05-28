# mysql
FROM mysql:5.7 

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./bin/auth ./cmd/*.go

EXPOSE 3306

CMD ["./bin/auth"]