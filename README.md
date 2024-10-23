# Project go-blogger

This is a RESTful API for performing CRUD blog post operation. The http service built based on Golang net/http.

Several libraries/tools are used along with Golang stdlib:

- github.com/dotenv-org/godotenvvault (for loading .env in a more secure way)
- <s>github.com/golang-jwt/jwt/v5 (for JWT auth)</s>
- SQLC (for database things)

## Getting Started

To run this project locally, it is necessary to have golang installed in local machine and an instance of Postgres,

1. clone this project

```bash
git clone https://github.com/frsfahd/go-blogger.git
```

2. create .env file in project root based on env-sample (fill with your own data). I'm using [dotenvault](https://www.dotenv.org/) for managing env var in remote environment (so ignore the .env.vault)
3. install dependencies

```bash
go mod download
```

4. run with `make run` or `make watch` (see the Makefile)

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

Live reload the application:

```bash
make watch
```

Clean up binary from the last build:

```bash
make clean
```
