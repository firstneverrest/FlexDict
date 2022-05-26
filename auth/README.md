# Auth service

## Installation

```bash
# create go.mod
go mod init github.com/firstneverrest/auth

# install go fiber
go get github.com/gofiber/fiber/v2
```

## Run

```
go run cmd/*.go
```

## Send Request with cURL

```
# GET
curl localhost:8000/user -i

# POST
curl localhost:8000/signin -i -X POST
```
