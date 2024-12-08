package flags

import (
	"flag"
)

type flags struct {
	NewProject string
}


func ParseFlags() flags {

	var flags = flags{}

  flag.StringVar(&flags.NewProject, "create", "", "Create a new project. Usage: -create <name>")
  flag.Parse()

	// create a project
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
  }

	return flags
}
