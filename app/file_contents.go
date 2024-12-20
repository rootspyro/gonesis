package app

import "fmt"

func GetMainContent(name string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"

	"%s/src/api"
	"%s/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	config.Setup()

	app.Use(logger.New())

	api.Setup(app)

	appSocket := fmt.Sprintf("%%s:%%s", config.App.Host, config.App.Port)
	if err := app.Listen(appSocket); err != nil {
		log.Error(err)
  }
}
`, name, name)
}

func GetGitignoreContent() string {
	return `
# Created by https://www.toptal.com/developers/gitignore/api/go
# Edit at https://www.toptal.com/developers/gitignore?templates=go

### Go ###
# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
bin/
.env
`
}

func GetMakefileContent(name string) string {
	return fmt.Sprintf(`APP=%s
GCO_ENABLED=0
GOOS=linux
GOARCH=amd64
MIGRATIONS_PATH=./src/db/migrations
PSQL_CONN="postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

include .env

run:
	go run cmd/$(APP)/main.go

build:
	GCO_ENABLED=$(GCO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP) cmd/$(APP)/main.go

start:
	./bin/$(APP)

migration_create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filename)

migration_up:
	migrate -path $(MIGRATIONS_PATH) -database $(PSQL_CONN) -verbose up

migration_down:
	migrate -path $(MIGRATIONS_PATH) -database $(PSQL_CONN) -verbose down 1

migration_fix:
	migrate -path $(MIGRATIONS_PATH) -database $(PSQL_CONN) -verbose force $(version)
`, name)
}

func GetConfigContent(name string) string {
	return fmt.Sprintf(`package config

import(
	"os"

  "github.com/joho/godotenv"
)

type app struct {
	Name    string
	Version string
	Host    string
	Port    string
}

type db struct {
  Host     string
  Port     string
  User     string
  Password string
  Name     string
}

var App = app{}
var Database = db{}

func Setup() {
	godotenv.Load()
	App.Name = readEnv("APP_NAME", "%s")
	App.Version = readEnv("APP_VERSION", "1.0.0")
	App.Host = readEnv("APP_HOST", "localhost")
  App.Port = readEnv("APP_PORT", "3000")

  Database.Host = readEnv("DB_HOST", "localhost")
  Database.Port = readEnv("DB_PORT", "5432")
  Database.User = readEnv("DB_USER", "postgres")
  Database.Password = readEnv("DB_PASSWORD", "postgres")
  Database.Name = readEnv("DB_NAME", "postgres")
}


func readEnv(key, defaultValue string) string {
	value := os.Getenv(key)

  if value == "" {
		value = defaultValue
  }

	return value
}
`, name)
}

func GetPerserContent(name string) string {
	return fmt.Sprintf(`package parser

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ---------------------
// 		RESPONSE STRUCTS
// ---------------------
type statusResponse struct {
  success string 
	error   string
}` +
 "\n type response struct {\n" + 
 "\tStatus string `json:\"status\"`\n" + 
 "\tStatusCode int `json:\"statusCode\"`\n" + 
 "\tData any `json:\"data\"`\n" + 
 "}\n" + 
 "\n type errorResponse struct {\n" + 
 "\tStatus string `json:\"status\"`\n" + 
 "\tStatusCode int `json:\"statusCode\"`\n" + 
 "\tError errorItem `json:\"error\"`\n" + 
 "}\n" + 
 "\n type errorItem struct {\n" + 
 "\tCode string `json:\"code\"`\n" + 
 "\tMessage string `json:\"message\"`\n" + 
 "\tDetails any `json:\"details\"`\n" + 
 "\tTimestamp time.Time `json:\"timestamp\"`\n" + 
 "\tPath string `json:\"path\"`\n" + 
 "}\n" + `

// ---------------------
// 		STATUS CODES	
// ---------------------

var status = statusResponse{
  success: "success",
  error:   "error",
}

// ---------------------
// 		  FUNCTIONS 
// ---------------------

func jsonFormatter(statusCode int, data any, c *fiber.Ctx) error {
  return c.Status(statusCode).JSON(response{
    Status: status.success,
    StatusCode: statusCode,
    Data: data,
  })
}

func errorFormatter(statusCode int, error errorItem, c *fiber.Ctx) error {
  return c.Status(statusCode).JSON(errorResponse{
    Status: status.error,
    StatusCode: statusCode,
    Error: error,
  })
}

// OK - 200
func OK(data any, c *fiber.Ctx) error {
  return jsonFormatter(http.StatusOK, data, c)
}

// Created - 201
func Created(data any, c *fiber.Ctx) error {
  return jsonFormatter(http.StatusCreated, data, c)
}

// BAD REQUEST - 400
func BadRequest(data any, c *fiber.Ctx) error {
  errorResponse := errorItem{
    Code: "BAD_REQUEST",
		Message: "Invalid request. Please check your request parameters and try again.",
    Details: data,
    Timestamp: timestamp(),
    Path: c.Path(),
  }
  return errorFormatter(http.StatusBadRequest, errorResponse, c)
}

// UNAUTHORIZED - 401
func Unauthorized(data any, c *fiber.Ctx) error {
  errorResponse := errorItem{
    Code: "UNAUTHORIZED",
    Message: "Authentication failed. Please check your credentials and try again.",
    Details: data,
    Timestamp: timestamp(),
    Path: c.Path(),
  }
  return errorFormatter(http.StatusUnauthorized, errorResponse, c)
}

// FORBIDDEN - 403
func Forbidden(data any, c *fiber.Ctx) error {
  errorResponse := errorItem{
    Code: "FORBIDDEN",
    Message: "Access denied. You do not have permission to access this resource.",
    Details: data,
		Timestamp: timestamp(),
    Path: c.Path(),
  }
  return errorFormatter(http.StatusForbidden, errorResponse, c)
}

// NOT FOUND - 404
func NotFound(data any, c *fiber.Ctx) error {
	errorResponse := errorItem{
    Code: "NOT_FOUND",
    Message: "Resource not found. The requested resource could not be located.",
    Details: data,
    Timestamp: timestamp(),
    Path: c.Path(),
  }
  return errorFormatter(http.StatusNotFound, errorResponse, c)
}

// CONFLICT - 409
func Conflict(data any, c *fiber.Ctx) error {
  errorResponse := errorItem{
    Code: "CONFLICT",
    Message: "Conflict. The request cannot be processed due to a conflict with the current state of the resource.",
    Details: data,
    Timestamp: timestamp(),
    Path: c.Path(),
  }
  return errorFormatter(http.StatusConflict, errorResponse, c)
}

// INTERNAL SERVER ERROR - 500
func InternalServerError(data any, c *fiber.Ctx) error {
  errorResponse := errorItem{
    Code: "INTERNAL_SERVER_ERROR",
    Message: "An unexpected error occurred while processing your request.",
    Details: data,
    Timestamp: timestamp(),
    Path: c.Path(),
  }
  return errorFormatter(http.StatusInternalServerError, errorResponse, c)
}

// ---------------------
// 				UTILS 
// ---------------------
func timestamp() time.Time {
  return time.Now().Local()
}
`)
}

func GetEnvContent(name string) string {
  return fmt.Sprintf(`APP_NAME="%s"
APP_VERSION="1.0.0"
APP_HOST="localhost"
APP_PORT="3000"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="postgres"
`, name)
}

func GetSQLCContent(name string) string {
	return fmt.Sprintf(`version: "2"
sql:
  - engine: "postgresql"
    queries: "./src/db/sqlc/%s/query.sql"
    schema: "./src/db/sqlc/%s/schema.sql"
    gen:
      go:
        package: "%srepo"
        out: "./src/repositories/%s_repo"
        sql_package: "pgx/v5"
`, name, name, name, name)
}

func GetServiceContent() string {
	return fmt.Sprintf(`package services
type CommonSrv struct {}

func NewCommonSrv() *CommonSrv {
  return &CommonSrv{}
}

func (s *CommonSrv) GetEndpoints() map[string]string {
	endpoints := map[string]string {
  	"GET /": "Index page",	
    "GET /health": "Health check",
	}

	return endpoints
}
`)
} 

func GetCommonPipesContent() string {
	return fmt.Sprintf(`package common

type IndexResponse struct {
` + "\t" + `Message string ` + "`json:\"message\"`\n" +
	"\t" + "Version string" + "`json:\"version\"`\n" +
	"\t" + "Endpoints map[string]string" + "`json:\"endpoints\"`\n" +
	"\t" + "Documentation string" + "`json:\"documentation\"`" + `
}
`)
}

func GetCommonHandler(name string) string {
	return fmt.Sprintf(`package common

import (
	"%s/pkg/parser"
	"%s/src/services"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	srv *services.CommonSrv
}

func New(services * services.CommonSrv) *Handler {
	return &Handler{
		srv : services,
	}
}

func(h *Handler) Index(c *fiber.Ctx) error {
	response := IndexResponse{
  	Message: "Welcome to %s",	
		Version: "v10.0.0",
		Endpoints: h.srv.GetEndpoints(),
		Documentation: "localhost:3000/docs",
	}
	return parser.OK(response, c)
}
`, name, name, name)
}

func GetAPIContent(name string) string {
	return fmt.Sprintf(`package api

import (
	"context"
	"os"

	"%s/src/api/handlers/common"
	"%s/src/db"
	"%s/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Setup(app *fiber.App) {

	// db
	pool := db.New()
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
    log.Error(err)
    os.Exit(1)
  }

	// services
	commonSrv := services.NewCommonSrv()

	// handlers
	commonH := common.New(commonSrv)

  // routes
  app.Get("/", commonH.Index)

}
`,name, name, name)
}

func GetDBConnContent(name string) string {
	return fmt.Sprintf(`package db

import (
	"context"
	"fmt"
	"%s/pkg/config"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New() *pgxpool.Pool {
	url := fmt.Sprintf("postgres://%%s:%%s@%%s:%%s/%%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	pool, err := pgxpool.New(context.Background(), url)
  if err != nil {
		log.Error(err)
		os.Exit(1)
  }

	return pool
}
`, name)
}

func GetDockerfileContent(name string) string {
	return fmt.Sprintf(`# Stage 1: Build stage 
FROM golang:1.23 AS build

WORKDIR /src

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o %s ./cmd/%s/main.go

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /src/%s .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set environment variables
ENV APP_NAME="%s"
ENV APP_VERSION="1.0.0"
ENV APP_HOST="0.0.0.0"
ENV APP_PORT="3000"
ENV DB_HOST=
ENV DB_PORT=
ENV DB_USER=
ENV DB_PASSWORD=
ENV DB_NAME=

EXPOSE 3000

CMD ["./%s"]
`, name, name, name, name, name)
}

func GetDockerIgnoreContent() string {
	return fmt.Sprintln(`.env
.env.example
.git
.gitignore
Makefile
README.md
build
dist
	`)
}

func GetREADMEContent(name string) string {
	return fmt.Sprintf(`# %s
Flexible and customizable base project generated by [Gonesis](https://github.com/rootspyro/gonesis) 

## Index
- [Tech Stack](#tech-stack)
- [Get Started](#get-started)
    - [Environment Variables](#environment-variables)
    - [Sqlc](#sqlc)
    - [Migrations](#migrations)
    - [Start the server](#start-the-server)
- [Docker Setup](#docker-setup)

## Tech Stack

- [Golang](https://go.dev) - Programming Language
- [Fiber](https://gofiber.io) - Web Framework
- [Sqlc](https://sqlc.dev) - SQL Database Abstraction Layer
- [Swagger](https://swagger.io) - API Documentation
- [Postgres](https://www.postgresql.org) - Database
- [Docker](https://www.docker.com) - Containerization

## Get Started

### Environment Variables
The %s file is readed by the [configuration](pkg/config/config.go) file. Please update the variables with your database credentials.
%s

### Sqlc 

This project uses [Sqlc](https://github.com/kyleconroy/sqlc) for database abstraction. It is configured to use pgxv5 under the hood.

Schemas and queries are stored in the %s directory, where %s is the main package for this project, but can be changed. Please check the [sqlc.yaml](sqlc.yaml) configration file.

%s

For example, you can create the following schemas:

%s

%s

To generate the repositories, run the following command:
%s 

### Migrations
Once the connection to the database is established and your schemas are generated, this project uses [go-migrate]() for migrations management.

__1. Create a new migration__
%s

Following the last example, you can create the following migrations:
%s

%s

__2. Run Migrations__
%s

__3. Rollback Migrations__
%s

__IMPORTANT:__ This command will revert only the last migration created.

### Start the server

To start the server, run the following command:
%s

For production environment, please check the [Docker Setup](#docker-setup) section.

## Docker Setup

1. __Build the docker image.__
%s

2. Run the docker image.
%s
`,
  	name,
		markdownHighlight(".env"),
		markdownCodeSection(
			fmt.Sprintf("APP_NAME=%s\nAPP_VERSION=1.0.0\nAPP_HOST=localhost\nAPP_PORT=3000\nDB_HOST=localhost\nDB_PORT=5432\nDB_USER=postgres\nDB_PASSWORD=postgres\nDB_NAME=postgres", name),
			"shell",
		),
		markdownHighlight("db/sqlc/" + name),
		markdownHighlight(name),
		markdownCodeSection(
			fmt.Sprintf(`version: "2"
sql:
  - engine: "postgresql"
    queries: "./src/db/sqlc/boilerplate/query.sql"
    schema: "./src/db/sqlc/boilerplate/schema.sql"
    gen:
      go:
        package: "boilerplaterepo"
        out: "./src/repositories/boilerplate_repo"
        sql_package: "pgx/v5"`),
			"yaml",
		),
		markdownCodeSection(fmt.Sprintf(`-- ./src/db/sqlc/%s/schema.sql
CREATE TABLE hello_world (
  id SERIAL PRIMARY KEY,
  message VARCHAR(255)
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);`, name), "sql"),
		markdownCodeSection(fmt.Sprintf(`-- ./src/db/sqlc/%s/query.sql

-- name: GetMessages :many
SELECT * FROM hello_world;

-- name: GetSingleMessage :one
SELECT * FROM hello_world WHERE id = $1;`, name), "sql"),
		markdownCodeSection("$ sqlc generate", "bash"),
		markdownCodeSection(fmt.Sprintf(`$ make migration_create filename=first_table

# migrate create -ext sql -dir ./src/db/migrations -seq first_table
# .../%s/src/db/migrations/000001_first_table.up.sql
# .../%s/src/db/migrations/000001_first_table.down.sql`, name, name), "sql"),
		markdownCodeSection(fmt.Sprintf(`-- ./src/db/migrations/000001_first_table.up.sql

CREATE TABLE hello_world (
  id SERIAL PRIMARY KEY,
  message VARCHAR(255)
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);`), "sql"),
		markdownCodeSection(fmt.Sprintf(`-- ./src/db/migrations/000001_first_table.down.sql

DELETE FROM hello_world;
DROP TABLE IF EXISTS hello_world;`), "sql"),
		markdownCodeSection("$ make migration_up", "bash"),
		markdownCodeSection("$ make migration_down", "bash"),
		markdownCodeSection(fmt.Sprintf(`$ make run

# ┌───────────────────────────────────────────────────┐ 
# │                   Fiber v2.52.5                   │ 
# │               http://127.0.0.1:3000               │ 
# │       (bound on host 0.0.0.0 and port 3000)       │ 
# │                                                   │ 
# │ Handlers ............. 3  Processes ........... 1 │ 
# │ Prefork ....... Disabled  PID ............. 81488 │ 
# └───────────────────────────────────────────────────┘ `), "bash"),
		markdownCodeSection(fmt.Sprintf(`$ make build -t %s .`, name), "bash"),
		markdownCodeSection(fmt.Sprintf(`$ docker run -d -p 3000:3000 --name %s -e DB_HOST=localhost -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=postgres -e DB_NAME=%s %s`, name, name, name), "bash"),
	)
}

func markdownCodeSection(content string, language string) string {
  return fmt.Sprintf("```%s\n%s\n```\n", language, content)
}

func markdownHighlight(content string) string {
	return "`" + content + "`"
}
