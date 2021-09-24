package help

import "fmt"

const HelpText string = `
	help text goes here
`

func PrintHelp() {
	fmt.Print(HelpText)
}
