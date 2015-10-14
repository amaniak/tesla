package tesla

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	Ok     = color.New(color.Bold, color.FgGreen).PrintlnFunc()
	Fail   = color.New(color.Bold, color.FgRed).Add(color.BgWhite).PrintlnFunc()
	Notice = color.New(color.FgCyan).Add(color.Underline).PrintlnFunc()
)

func Ask(input string) string {

	//Ask for
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input)
	return reader.ReadString('\n')
}
