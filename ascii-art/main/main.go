package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strings"
)

func main() {

	// if len(os.Args) < 3 && len(os.Args) >= 2 {
	// 	fmt.Println("Usage: go run main.go (input) (output)")
	// 	fmt.Println("done by na hehee malak:chaap")
	// 	return
	// }

	// +++__++++ qustion on optional ++++__++++ //

	var file1 string
	if len(os.Args) > 2 { // if its > 2 we must print what file we got
		file1 = os.Args[2] // ex shadow
	} else {
		file1 = "standard" // if nothing just the string
	}

	if file1 != "standard" && file1 != "thinkertoy" && file1 != "shadow" {
		fmt.Println("ERROR: Invalid file name!")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("PLEASE PROVIDE A STRING :( ") // if we type just go run . without any thing
		os.Exit(1)
	}

	file, err := os.Open("../files/" + file1 + ".txt") // we exist the files folder - the err for ex if we type shadow but the file not exist
	if err != nil {                                    // if we have err
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // close the file

	reader := bufio.NewScanner(file) // to read
	//var reader bufio.Scanner //read ithe scanner

	stringmalak := os.Args[1] // to find the string
	if stringmalak == "" {
		os.Exit(0)
	}
	// if stringmalak == "\n" {
	//  strings.Split(stringmalak, "\n")
	// }
	// stringmalak = strings.ReplaceAll(stringmalak, "\\n", "\n")
	split := strings.Split(stringmalak, "\\n")
	c := 0
	for _, a := range split {

		if a == "" {
			c++
		}
	}
	if c == len(split) {
		split = split[:c-1]
	}
	for _, s := range split {

		lines := make([]string, 8) // to safe & store 8
		// for i :=0 ; i<len(string_array); i++{
		for _, char := range s {
			index_of_char := 9*(int(char)-32) + 1

			file.Seek(0, 0)

			reader = bufio.NewScanner(file)
			for j := 0; j < index_of_char; j++ {
				reader.Scan()
			}
			for k := 0; k < 8 && reader.Scan(); k++ {
				lines[k] += reader.Text()
			}
		}
		// fmt.Println(lines)
		for _, line := range lines {
			// fmt.Println(i, line)
			if line != "" {
				fmt.Println(line)

			} else {
				fmt.Println()
				break
				//continue
			}
		}

	}

}
