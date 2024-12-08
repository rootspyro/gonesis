package app

import (
	"os"

	"github.com/rootspyro/gonesis/pkg/log"
)

func CreateProject(name string) {
	// create project directory
	CreateDir(name)
  log.Info("Project directory created")

  // change directory
	os.Chdir(name)

	// init go mod
	RunCommand("go", []string{"mod", "init", name})
	log.Info("Go mod initialized")

	// init git repo
	RunCommand("git", []string{"init"})
	log.Info("Git repo initialized")

	// create README.md
  CreateFile("README.md", "# " + name)
	log.Info("README.md created")

	// create .env.example
	CreateFile(".env.example", GetEnvContent(name))
  log.Info(".env.example created")

	// create Makefile
  CreateFile("Makefile", GetMakefileContent(name))
  log.Info("Makefile created")

	// create pkg directory
  CreateDir("pkg")
	CreateDir("pkg/config")

  // create config.go
  CreateFile("pkg/config/config.go", GetConfigContent(name))
  log.Info("config.go created")
	

	// create cmd directory
	CreateDir("cmd")
	CreateDir("cmd/" + name)

	// create main.go
  CreateFile("cmd/" + name + "/main.go", GetMainContent())
  log.Info("main.go created")

	// go mod tidy
  RunCommand("go", []string{"mod", "tidy"})
  log.Info("Go modules tidied")

}
