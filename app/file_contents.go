package app

import "fmt"

func GetMainContent() string {
	return `package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })
  app.Listen(":3000")
}
`
}

func GetConfigContent(name string) string {
	return fmt.Sprintf(`package config

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
