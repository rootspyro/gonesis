package log

import (
	"fmt"

	"github.com/rootspyro/gonesis/pkg/colors"
)

func Info(msg string) {
	fmt.Println(colors.Cyan,"[INFO] ", colors.Reset, msg)
}

func Error(msg string) {
  fmt.Println(colors.Red,"[ERROR] ", colors.Reset, msg)
}

func Success(msg string) {
  fmt.Println(colors.Green,"[SUCCESS] ", colors.Reset, msg)
}
