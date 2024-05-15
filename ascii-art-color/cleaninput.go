package asciiArt

import (
	"strings"
)

func CleanInput(args []string) []string {
	var cleanArgs []string
	for _, arg := range args {
		if arg != "" {
			cleanArgs = append(cleanArgs, arg)
		}
	}
	return cleanArgs
}

func CheckBanner(banner string) string {

	switch strings.ToLower(banner) {
	case "standard":
		banner = "standard.txt"
	case "shadow":
		banner = "shadow.txt"
	case "thinkertoy":
		banner = "thinkertoy.txt"
	default:
		banner = "NotFound"
	}
	return banner
}
