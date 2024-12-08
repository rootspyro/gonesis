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

  if flags.NewProject != "" {
		log.Info("Creating project " + flags.NewProject)
		app.CreateProject(flags.NewProject)
		fmt.Printf("\n")

		log.Success("Project created!\n")

		fmt.Println(colors.Green +  " • " + colors.Reset + "Run " + colors.Cyan + "cd " + flags.NewProject + colors.Reset + " and " + colors.Cyan + "make run" + colors.Reset + " to start the server.\n")

		os.Exit(0)
  }

}
