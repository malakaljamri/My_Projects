package main

import (
	"bufio"
	"fmt"
	"goreloaded/Converter"
	"os"

	//"io/ioutil"
	//"os"
	"regexp"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: go run . (input_file) (output_file)")
		return
	}
	inputFile := os.Args[1]
	resultFile := os.Args[2]

	fileText, err := os.Open(inputFile)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer fileText.Close()

	textScanner := bufio.NewScanner(fileText)
	if err := textScanner.Err(); err != nil {
		fmt.Print("failed to read the file")
		return
	}

	var input string
	for textScanner.Scan() {
		input += textScanner.Text()
		input += "\n"
	}
	// input := "1E (hex) files were added Ready, set, go (up) !I should stop SHOUTING (low)  Welcome to the Brooklyn bridge (cap) dodo do (cap, 55) last but not least"
	// to find the hexadecimal part
	/////////////////////////////////////////////////////

	// re := regexp.MustCompile(`\b[0-9A-Fa-f]+\b`)
	// hexValue := re.FindString(input)
	// decimalString := Converter.ReplaceHex(hexValue)
	// input = strings.Replace(input, hexValue+" (hex)", decimalString, 1)

	re := regexp.MustCompile(`\b[0-9A-Fa-f]+\b \(hex\)`)
	input = re.ReplaceAllStringFunc(input, Converter.ReplaceHex)
	reBin := regexp.MustCompile(`\b[0-1]+\b \(bin\)`)
	input = reBin.ReplaceAllStringFunc(input, Converter.ReplaceBin)

	/////////////////////////////////////////////////////////////////////////
	// fmt.Println(input)

	//fmt.Println(decimalString)

	//m2
	// Bin := "It has been 10 (bin) years"
	// re = regexp.MustCompile(``)

	//m3
	// input := "Ready, set, go (up) !"
	input = Converter.ConvertToUpper(input)
	// fmt.Println(input)

	//m4
	// input = "  I should stop SHOUTING (low)"
	input = Converter.ConvertToLower(input)
	// fmt.Println(output)

	// input = "  Welcome to the Brooklyn bridge (cap) "
	input = Converter.Cap(input)
	// fmt.Println(output)

	//newwwwwwwwwwwwww
	// input_test := "dodo do (cap, 55) last but not least" // there is a panic in (cap .55)
	strings_list := strings.Fields(input)
	// Applying Numlow, Numup, and Numcap functions on the input list
	for m, word := range strings_list {
		switch {
		case strings.HasPrefix(word, "(low,"):
			Converter.Numlow(strings_list, m)
		case strings.HasPrefix(word, "(up,"):
			Converter.Numup(strings_list, m)
		case strings.HasPrefix(word, "(cap,"):
			Converter.Numcap(strings_list, m)
		}
	}
	// Printing the modified list after applying the functions
	// fmt.Println(strings.Join(strings_list, " "))
	cleaned := []string{}
	for _, s := range strings_list {
		if s != "" {
			cleaned = append(cleaned, s)
		}
	}
	input = strings.Join(cleaned, " ")

	//punction without space
	// text := "I was sitting over there ,and then BAMM ??"
	// re = regexp.MustCompile(`\s*([.,!?;:])\s*`)
	// input = re.ReplaceAllString(input, "$1")
	///////////////////////////////////////////////////////
	// re = regexp.MustCompile(`\s*([.,!?;:])\s*`)
	// input = re.ReplaceAllString(input, "$1")

	// punctuationMarks := `[.,!?;:]`
	// re = regexp.MustCompile(`\s*(` + punctuationMarks + `)(\s*)`)
	// input= re.ReplaceAllString(input, "$1$2 ")

	// input := `[.,!?:;]+`
	// re := regexp.MustCompile(`(\s*)` + input + `(\s*)`)

	// text = "i was takkio ...."
	re = regexp.MustCompile(`\s*([.,;:?!]+)\s*`)
	input = re.ReplaceAllString(input, "$1 ")

	// fmt.Println(result)

	//m3

	// input = "I am exactly how they describe me: ' awesome ' and ' incredible '."

	re = regexp.MustCompile(`' (.*?) '`)
	input = re.ReplaceAllString(input, "'$1'")

	// fmt.Println(output)

	// 7
	// input = "As Elton John said: ' I am the most well-known homosexual in the world '"

	re = regexp.MustCompile(`'(\s*)(.*?)(\s*)'`)
	input = re.ReplaceAllString(input, "'$2'")

	// fmt.Println(output)

	//lastttt

	// input = "There it was. A amazing rock! A apple fell down. A hour passed."

	//////////////////////////////

	//re = regexp.MustCompile(`\bA\s+(?i:[aeiouh])`)
	re = regexp.MustCompile(`\b[AIai]\s+(?i:[aeiouh])`)
	input = re.ReplaceAllStringFunc(input, func(match string) string {
		secondWord := strings.Split(match, " ")[1]
		if strings.ToLower(match)[0] == 'a' {
			return string(match[0]) + "n " + secondWord
		}
		return match
	})


	input = strings.TrimSpace(input)
	/////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////
	err = os.WriteFile(resultFile, []byte(input), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	//////////////////////////////////////////
	fmt.Println("Content written to file successfully.")
	// fmt.Println(input)

}
