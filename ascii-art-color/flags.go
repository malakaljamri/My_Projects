package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

func CheckOutput(Args []string) ([]string, string, bool) {
	//// --color, supStr , --str , output ,fs
	output := ""
	var hasColor bool
	for i, arg := range Args {
		// if strings.HasPrefix(arg, "--color") {
		// 	color := strings.TrimPrefix(arg, "--color=")
		// 	if color == "" {
		// 		fmt.Println("usage: go run . --color=<color> <letters to be colored> 'something'")
		// 		os.Exit(0)
		// 	}
		// 	Args[i] = ""
		// }
		if strings.HasPrefix(arg, "--output=") {
			output = strings.TrimPrefix(arg, "--output=")
			if output == "" {
				fmt.Println("usage: go run . --output=<fileName.txt> something standard")
				os.Exit(0)
			}
			Args[i] = ""
		} else if strings.HasPrefix(arg, "--") && !strings.HasPrefix(arg, "--color=") {
			fmt.Println("usage: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		} else if strings.HasPrefix(arg, "--color=") {
			hasColor = true
			Args[i] = strings.TrimPrefix(arg, "--color=")
		}
	}
	Args = CleanInput(Args)
	return Args, output, hasColor
}
