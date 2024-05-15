package Converter
import(
	// "fmt"
	"regexp"
	// "Strings"
)

func formatPunctuation(text string) string {
	// Define a regular expression to match spaces before and after the specified punctuation marks
	re := regexp.MustCompile(`\s*([.,!?;:])\s*`)

	// Replace matched patterns with the formatted punctuation
	formattedText := re.ReplaceAllString(text, "$1")

	return formattedText
}