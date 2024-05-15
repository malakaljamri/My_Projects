package Converter

import (
	"fmt"
	"strconv"
	"strings"
)

func ReplaceBin(Bin string) string {
	Bin = strings.TrimSuffix(Bin, " (bin)")
	dec, err := strconv.ParseInt(Bin, 2, 64)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//we use the sprintf = to convert it to string :)
	return fmt.Sprintf("%d", dec)
}
