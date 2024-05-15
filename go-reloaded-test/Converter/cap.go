package Converter

import (
	"strings"
)

func Cap(malak string) string {
	// fmt.Printf("'%s'\n", malak)
	word := strings.Fields(malak)

	for i := 0; i < len(word); i++ {
		if word[i] == "(cap)" {
			word[i] = ""
			word[i-1] = strings.Title(word[i-1])
		}
	}
	return strings.Join(word, " ")
}
