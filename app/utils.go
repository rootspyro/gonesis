package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rootspyro/gonesis/pkg/log"
)

func CreateDir(name string) {
  err := os.Mkdir(name, 0755)
  if err != nil {
		log.Error(err.Error())
    os.Exit(1)
  }
}


func RunCommand(name string, args []string) {
  cmd := exec.Command(name, args...)
  _, err := cmd.CombinedOutput()
  if err != nil {
		log.Error(err.Error())
    os.Exit(1)
  }
}

func CreateFile(name string, content string) {
  file, err := os.Create(name)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  _, err = file.WriteString(content)
  if err != nil {
		log.Error(err.Error())
    os.Exit(1)
  }

	file.Close() 
}
