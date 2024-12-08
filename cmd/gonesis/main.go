package main

import (
	"fmt"
	"os"

	"github.com/rootspyro/gonesis/app"
	f "github.com/rootspyro/gonesis/cmd/flags"
	"github.com/rootspyro/gonesis/pkg/colors"
	"github.com/rootspyro/gonesis/pkg/log"
)

func main() {
	fmt.Printf(colors.Cyan + `
   ____ _ ____   ____   ___   _____ (_)_____
  / __  // __ \ / __ \ / _ \ / ___// // ___/
 / /_/ // /_/ // / / //  __/(__  )/ /(__  ) 
 \__, / \____//_/ /_/ \___//____//_//____/  
/____/                                      

` + colors.Reset)

	flags := f.ParseFlags()

	if flags.Version {
		log.Info("Gonesis " + colors.Magenta + "v1.0.0")
		fmt.Printf("\n")
		os.Exit(0)
  }

  if flags.NewProject != "" {
		log.Info("Creating project " + flags.NewProject)
		app.CreateProject(flags.NewProject)
		fmt.Printf("\n")
		os.Exit(0)
		// log.Success("Project created! Run " + colors.Magenta + "cd " + flags.NewProject + colors.Reset + " and " colors.M" to start working on it")
  }

}
