package config

import (
  "os"

  "github.com/joho/godotenv"
)

type app struct {
  Version string
	Author  string
	Github  string
}

var App = app{}

func Setup() {

	godotenv.Load(".env")

  App.Version = readEnv("APP_VERSION", "1.0.0")
  App.Author = readEnv("APP_AUTHOR", "")
  App.Github = readEnv("APP_GITHUB", "")
}

func readEnv(key, defaultValue string) string {
  value := os.Getenv(key)
  if value == "" {
    return defaultValue
  }
  return value
}
