# Auth service

## Technology

### Auth service

- mysql
- Go

### FlexDict service

### Vocabulary service

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
#  Sign up
curl http://localhost:8000/v1/signup -d '{"username":"first", "password":"123"}' -H "Content-Type: application/json" -i -X POST

# Sign in

# Sign out
```
