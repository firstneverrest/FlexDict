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

## APIs

### Sign up

| Request | Path       | Header                         | Body                                           | Response |
| ------- | ---------- | ------------------------------ | ---------------------------------------------- | -------- |
| POST    | /v1/signup | Content-Type: application/json | {"username":"carlos", "password":"mypassword"} | 201      |

### Sign in

| Request | Path       | Header                         | Body                                           | Response |
| ------- | ---------- | ------------------------------ | ---------------------------------------------- | -------- |
| POST    | /v1/signup | Content-Type: application/json | {"username":"carlos", "password":"mypassword"} | 201      |

## Send Request with cURL

```
#  Sign up
curl http://localhost:8000/v1/signup -d '{"username":"first", "password":"123123"}' -H "Content-Type: application/json" -i -X POST

# Sign in
curl http://localhost:8000/v1/signin -d '{"username":"first", "password":"123123"}' -H "Content-Type: application/json" -i -X POST

# Get vocabulary

# Add vocabulary

# Edit vocabulary

# Delete vocabulary
```
