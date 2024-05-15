package main

import (
	"asciiArt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Args := os.Args[1:]
	if len(Args) == 0 {
		fmt.Println("Error: no arguments.")
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
		os.Exit(0)
	} else if len(Args) > 5 {
		fmt.Println("Error: so many Arguments.")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
		os.Exit(0)
	} else {
		//// --color, supStr , --str , output ,fs
		output := "" // name of file from --output=
		input := ""  // input string to be converted to ascii art
		banner := "standard.txt"
		var color string  // the color from flag --color
		toBeColored := "" // sup text to be colored
		var hasColor bool // if flag --color is found

		Art := ""

		Args, output, hasColor = asciiArt.CheckOutput(Args)
		if len(Args) < 1 {
			fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
			os.Exit(0)
		}
		lastIndex := len(Args) - 1
		banner = asciiArt.CheckBanner(Args[lastIndex])
		// banner = asciiArt.CheckBanner("standard")
		////////////////////////////////////////////////////////////////////////////
		if len(Args) == 1 {
			input = Args[0]
			banner = asciiArt.CheckBanner("standard")
		} else if len(Args) == 2 {
			if !hasColor {
				input = Args[0]
				banner = asciiArt.CheckBanner(Args[1])
			} else {
				color = Args[0]
				input = Args[1]
				banner = asciiArt.CheckBanner("standard")
			}
		} else if len(Args) == 3 {
			if !hasColor {
				fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
				os.Exit(0)

			} else if banner == "NotFound" {
				color = Args[0]
				toBeColored = Args[1]
				input = Args[2]
				banner = asciiArt.CheckBanner("standard")
			} else {

				color = Args[0]
				input = Args[1]
			}
		} else if len(Args) == 4 {
			color = Args[0]
			toBeColored = Args[1]
			input = Args[2]
		}
		fmt.Println(output, color, toBeColored)
		////////////////////////////////////////////////////////////////
		var fileLines []string
		// input := os.Args[1]
		if len(fileLines) == 1 {
			file, err := os.Open("standard.txt")
			if err != nil {
				log.Fatal(err)
			}
			fileScanner := bufio.NewScanner(file)
			fileScanner.Split(bufio.ScanLines)
			for fileScanner.Scan() {
				fileLines = append(fileLines, fileScanner.Text())
				defer file.Close()
			}
		}
		if hasColor {
			color = asciiArt.PrintColor(color)
			if color == "" {
				log.Fatal("color not found")
			}
		}

		file, err := os.Open(banner)
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}
		file.Close()
		var fileload []int
		result := strings.ReplaceAll(input, "\\t", "   ")
		result = strings.ReplaceAll(result, "\\n", " \\n ")
		words := strings.Split(result, " \\n ")
		if hasColor {
			asciiArt.PrintBannersWithColors(toBeColored, color, words, fileLines)
		}else {
			for j := 0; j < len(words); j++ {
				if words[j] != "" {
					fileload = asciiArt.ConvertToASCIIArt(words[j])
					fileload1 := asciiArt.PrintLines(fileload, fileLines)
					for i := 0; i < len(fileload1); i++ {
						if strings.Join(fileload1[i], "") != "" {
							if hasColor {
								fmt.Println(color, strings.Join(fileload1[i], ""), "\033[0m")
							} else {
								fmt.Println(strings.Join(fileload1[i], ""))
							}
	
							Art = Art + strings.Join(fileload1[i], "") + "\n"
						}
					}
				} else if j != 1 && words[j] == "" {
					fmt.Println()
					Art = Art + "\n"
				}
		}
		
		}
		asciiArt.CreateFile(output, Art)
	}
}
