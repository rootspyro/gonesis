package app

import (
	"fmt"
	"os"
)

func CreateProject(name string) {
	// create project directory
	CreateDir(name)

  // change directory
	os.Chdir(name)

	// init go mod
	RunCommand("go", []string{"mod", "init", name})

	// init git repo
	RunCommand("git", []string{"init"})

	// create README.md
  CreateFile("README.md", "# " + name)

	// create cmd directory
	CreateDir("cmd")
	CreateDir("cmd/" + name)

	// create main.go
  file, err := os.Create("cmd/" + name + "/main.go")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  file.Close()

}
