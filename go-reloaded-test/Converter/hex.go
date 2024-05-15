package Converter

import (
	"fmt"
	"strconv"
	"strings"
)

func ReplaceHex(hexa string) string {
	hexa = strings.TrimSuffix(hexa, " (hex)")
	dec, err := strconv.ParseInt(hexa, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%d", dec)
}
