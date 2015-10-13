package tesla

import (
	"github.com/fatih/color"
)

var (
	Ok     = color.New(color.Bold, color.FgGreen).PrintlnFunc()
	Fail   = color.New(color.Bold, color.FgRed).Add(color.BgWhite).PrintlnFunc()
	Notice = color.New(color.FgCyan).Add(color.Underline).PrintlnFunc()
)
