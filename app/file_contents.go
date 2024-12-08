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
	GCO_ENABLED=$(GCO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP) cmd/main.go

start:
	./bin/$(APP)

migration_create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filename)

migration_up:
	migrate -path $(MIGRATIONS_PATH) -database $(PSQL_CONN) -verbose up

migration_down:
	migrate -path $(MIGRATIONS_PATH) db/migrations -database $(PSQL_CONN) -verbose down 1

migration_fix:
	migrate -path $(MIGRATIONS_PATH) -database $(PSQL_CONN) -verbose fix $(version)
`, name)
}

func GetConfigContent(name string) string {
	return fmt.Sprintf(`package config

import "os"

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
  return fmt.Sprintf(`
APP_NAME="%s"
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
	"%s/src/api/handlers/common"
	"%s/src/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// services
	commonSrv := services.NewCommonSrv()

	// handlers
	commonH := common.New(commonSrv)

  // routes
  app.Get("/", commonH.Index)

}
`, name, name)
}
