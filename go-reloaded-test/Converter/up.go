package Converter

import (
	"strings"
)

// func ConvertToUpper(malak string) string {
// 	words :=strings.Fields(input)
// 	for i, word :=range words{
// 		if word=="(up)"
// 		continue
// 	}
// 	if strings.HasSuffix(word, "(up)")

//			}
//		}
//		return strings.Join(word, "")
//	}
func ConvertToUpper(malak string) string {
	word := strings.Fields(malak)

	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" {
			word[i] = ""
			word[i-1] = strings.ToUpper(word[i-1])
		}

	}
	return strings.Join(word, " ")

}
