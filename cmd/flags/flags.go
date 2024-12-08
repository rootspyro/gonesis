package flags

import (
	"flag"
)

type flags struct {
	Version 	 bool 
	NewProject string
}


func ParseFlags() flags {

	var flags = flags{}

	flag.BoolVar(&flags.Version, "version", false, "Print the version number")
  flag.StringVar(&flags.NewProject, "create", "", "Create a new project. Usage: -create <name>")
  flag.Parse()

	// create a project
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
  }

	return flags
}
