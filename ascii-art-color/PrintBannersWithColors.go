package asciiArt

import (
	"fmt"
)

// Print the full outcome in the triminal
func PrintBannersWithColors(Str, colors string, banners, arr []string) {
	num := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			if Str == "" {
				for _, j := range ch {
					n := (j-32)*9 + 1
					fmt.Print(colors, arr[int(n)+i])
				}
			} else {
				h := 0
				count := 0
				match := false
				for _, j := range ch {
					
					if !match || count >= len(Str) {
						h = h + 1
					}
					
					check := true
					n := (j-32)*9 + 1
					for q := 0; q < len(Str); q++ {
						
							if rune(Str[q]) == j {
								
								word := ch
								if count < len(Str) {
									if  h <= ( 1+ len(word) - len(Str))  {
										if (Str == word[h-1:h+len(Str)-1] || (match && count < len(Str)) ){
											match = true
											count = count + 1
											
											fmt.Print(colors, arr[int(n)+i])
										
											check = false
										}

									}
								
									if count == len(Str) {
										count = 0
										match = false
									}
									break
								}
						
							}
					}
					if check == true {
						fmt.Print("\033[0m", arr[int(n)+i])
					}
					
				}
				count = 0
			}
			fmt.Println("\033[0m")
		}
	}
}
