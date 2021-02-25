package internal

import (
	"github.com/fatih/color"
)

func OutputToConsoleWithSuccessFormatting(text string) {
	c := color.New(color.FgGreen)
	c.Println(text)
}

func OutputToConsoleWithWarningFormatting(text string) {
	c := color.New(color.FgYellow)
	c.Println(text)
}

func OutputToConsoleWithAlertFormatting(text string) {
	c := color.New(color.FgRed)
	c.Println(text)
}