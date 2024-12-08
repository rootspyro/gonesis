package main

import (
	"fmt"
	"os"

	"github.com/rootspyro/gonesis/app"
	f "github.com/rootspyro/gonesis/cmd/flags"
)

func main() {
	fmt.Printf(`
   ____ _ ____   ____   ___   _____ (_)_____
  / __  // __ \ / __ \ / _ \ / ___// // ___/
 / /_/ // /_/ // / / //  __/(__  )/ /(__  ) 
 \__, / \____//_/ /_/ \___//____//_//____/  
/____/                                      

`)

	flags := f.ParseFlags()

	if flags.Version {
    fmt.Println("v1.0.0")
		os.Exit(0)
  }

  if flags.NewProject != "" {
		fmt.Println("\nCreating project: " + flags.NewProject)
		app.CreateProject(flags.NewProject)
  }

}
