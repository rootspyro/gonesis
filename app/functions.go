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

	// create README.md
  CreateFile("README.md", "# " + name)
	log.Info("README.md created")

	// create Makefile
  CreateFile("Makefile", GetMakefileContent(name))
  log.Info("Makefile created")

	// init git repo
	RunCommand("git", []string{"init"})

	// create .gitignore
  CreateFile(".gitignore", GetGitignoreContent())

	log.Info("Git repo initialized")

	// create .env.example
	CreateFile(".env.example", GetEnvContent(name))
	RunCommand("cp", []string{".env.example", ".env"})
  log.Info(".env.example created")


	// create sqlc.yaml
  CreateFile("sqlc.yaml", GetSQLCContent(name))
  log.Info("sqlc.yaml created")

	// create dockerfile
  CreateFile("Dockerfile", GetDockerfileContent(name))
	CreateFile(".dockerignore", GetDockerIgnoreContent())
  log.Info("Dockerfile created")

	// create pkg directory
  CreateDir("pkg")
	CreateDir("pkg/config")
	CreateDir("pkg/parser")

	CreateFile("pkg/parser/parser.go", GetPerserContent(name))

  // create config.go
  CreateFile("pkg/config/config.go", GetConfigContent(name))

	log.Info("pkg directory created")

	// create src directory
  CreateDir("src")

	// create db directory
  CreateDir("src/db")
	CreateFile("src/db/conn.go", GetDBConnContent(name))
  CreateDir("src/db/migrations")
	CreateDir("src/db/sqlc")
	CreateDir("src/db/sqlc/" + name)

	CreateFile("src/db/sqlc/" + name + "/query.sql", "")
	CreateFile("src/db/sqlc/" + name + "/schema.sql", "")
	
  CreateDir("src/repositories")
	CreateDir("src/services")

	CreateFile("src/services/commong.srv.go", GetServiceContent())

	CreateDir("src/api")

	CreateDir("src/api/handlers")
	CreateDir("src/api/handlers/common")

	CreateFile("src/api/handlers/common/pipes.go", GetCommonPipesContent())
	CreateFile("src/api/handlers/common/handler.go", GetCommonHandler(name))

	CreateFile("src/api/routes.go", GetAPIContent(name))

  log.Info("src directory created")

	// create cmd directory
	CreateDir("cmd")
	CreateDir("cmd/" + name)

	// create main.go
  CreateFile("cmd/" + name + "/main.go", GetMainContent(name))
	log.Info("cmd directory created")

	// go mod tidy
  RunCommand("go", []string{"mod", "tidy"})
  log.Info("Go modules tidied")

}
