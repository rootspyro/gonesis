package log

import (
	"fmt"

	"github.com/rootspyro/gonesis/pkg/colors"
)

func Info(msg string) {
	fmt.Println(colors.Cyan + colors.Bold ,"[INFO] ", colors.Reset, msg)
}

func Error(msg string) {
  fmt.Println(colors.Red + colors.Bold ,"[ERROR] ", colors.Reset, msg)
}

func Success(msg string) {
  fmt.Println(colors.Green + colors.Bold ,"[SUCCESS] ", colors.Reset, msg)
}
