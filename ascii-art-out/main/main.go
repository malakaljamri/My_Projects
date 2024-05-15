package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"flag"
)

func main() {
	

	msg := flag.String("output", "", "Output file name")
	flag.Parse()

	if *msg != "" {
		outputFile, err := os.Create(*msg)
		if err != nil {
			fmt.Println("ERROR: Could not create the output file:", err)
			return
		}
		defer outputFile.Close()
		os.Stdout = outputFile
	}


		var file1 string
	if len(flag.Args()) > 2 { // if its > 2 we must print what file we got
		file1 = flag.Arg(0) // ex shadow
	} else {
		file1 = "standard" // if nothing just the string
	} 

	if file1 != "standard" && file1 != "thinkertoy" && file1 != "shadow" {
		fmt.Println("ERROR: Invalid file name!")
		return
	}

	if len(flag.Args()) < 1 {
		fmt.Println("PLEASE PROVIDE A STRING :( ") // if we type just go run . without any thing
		os.Exit(1)
	}

	stringmalak := flag.Arg(1)
	if stringmalak == "" {
		os.Exit(0)
	}

	//
	file, err := os.Open("../files/" + file1 + ".txt") // we exist the files folder - the err for ex if we type shadow but the file not exist
	if err != nil {                                    // if we have err
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // close the file

	reader := bufio.NewScanner(file) // to read
	//var reader bufio.Scanner //read ithe scanner
////////////////////////
	// stringmalak := os.Args[1] // to find the string
	// if stringmalak == "" {
	// 	os.Exit(0)
	// }
	
	// to count the /n and -1 them becouse 1 /n will create $$ we want 1
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
