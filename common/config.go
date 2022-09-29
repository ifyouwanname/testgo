package common

import (
	"github.com/fatih/color"
)

func ColorBlue(colorStr string) {
	colorPrint := color.New(color.FgBlue).Add(color.Bold)
	colorPrint.Println(colorStr)
}
