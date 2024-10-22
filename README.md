# Project go-blogger

This project intended to showcases the implementation of Golang standard library for building Restful API. These are available endpoints :

- /login
- /register
- /health

Several libraries/tools are used along with Golang stdlib:

- github.com/dotenv-org/godotenvvault (for loading .env in a more secure way)
- github.com/golang-jwt/jwt/v5 (for JWT auth)
- SQLC (for database things)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
