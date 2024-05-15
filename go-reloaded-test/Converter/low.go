package Converter

import (
	"strings"
)

func ConvertToLower(malak string) string {

	word := strings.Fields(malak)

	for i := 0; i < len(word); i++ {
		if word[i] == "(low)" {
			word[i] = "" //to remove the (low)
			word[i-1] = strings.ToLower(word[i-1])
		}

	}
	return strings.Join(word, " ")

}
