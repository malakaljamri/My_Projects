package Converter

import (
	"strconv"
	"strings"
)

func Numlow(stringlist []string, m int) {
	num_str := strings.TrimRight(stringlist[m+1], ")")
	num, _ := strconv.Atoi(num_str)
	for i := m; (i > m-num) && (i > 0); i-- {
		stringlist[i-1] = strings.ToLower(stringlist[i-1])
	}
	if stringlist[m+2] == " " {
		stringlist[m+2] = ""
	}
	stringlist[m] = ""
	stringlist[m+1] = ""

}

func Numup(stringlist []string, m int) {
	num_str := strings.Split(stringlist[m+1], ")")
	num, _ := strconv.Atoi(num_str[0])
	// m= indx

	for i := m; (i > m-num) && (i > 0); i-- {
		str := stringlist[i-1]
		stringlist[i-1] = strings.ToUpper(string(str[0]) + str[1:])

	}
	stringlist[m] = ""
	stringlist[m+1] = ""
}

func Numcap(stringlist []string, m int) {
	num_str := strings.Split(stringlist[m+1], ")")
	num, _ := strconv.Atoi(num_str[0])

	for i := m; (i > m-num) && (i > 0); i-- {
		str := stringlist[i-1]
		stringlist[i-1] = strings.ToUpper(string(str[0])) + str[1:]
	}
	stringlist[m] = ""
	stringlist[m+1] = ""
}