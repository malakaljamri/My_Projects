package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"errors"
	"strings"
	"os"
	"bufio"

)

type Fonts struct {
	Art                string
	Hidden             string
}

const (
	fileLen = 855
)

const port = ":2324"

func printHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedTemplate, _ := template.ParseFiles("../../static/405.html")
		parsedTemplate.Execute(w, nil)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		parsedTemplate, _ := template.ParseFiles("../../static/404.html")
		parsedTemplate.Execute(w, nil)
		return
	}

	parsedTemplate, err := template.ParseFiles("../../static/index.html")
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedTemplate, _ := template.ParseFiles("../../static/405.html")
		parsedTemplate.Execute(w, nil)
		return
	} else {
		name := r.FormValue("name")
		errName := IsValid(name)
		if errName != nil {
			w.WriteHeader(http.StatusBadRequest)
			parsedTemplate, _ := template.ParseFiles("../../static/400.html")
			err := parsedTemplate.Execute(w, nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				parsedTemplate, _ := template.ParseFiles("../../static/500.html")
				err = parsedTemplate.Execute(w, nil)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}
			return
		}
		
		banner := r.FormValue("banner")
		if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
			w.WriteHeader(http.StatusNotFound)
			parsedTemplate, _ := template.ParseFiles("../../static/404.html")
			parsedTemplate.Execute(w, nil)
			return
		}
		fmt.Println(banner)
		art, err := AsciiArt(banner, name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			parsedTemplate, _ := template.ParseFiles("../../static/404.html")
			parsedTemplate.Execute(w, nil)
			return
		}
		fonts := Fonts{Art: art, Hidden: "false"}
		parsedTemplate, err := template.ParseFiles("../../static/index.html")
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}

		err = parsedTemplate.Execute(w, fonts)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	}
}

func AsciiArt(banner string, fontStr string) (string, error) { // name string, font string
	argsArr := strings.Split(strings.ReplaceAll(fontStr, "\r", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("../../internal/asciiart/fonts/" + banner + ".txt")
	if err != nil {
		defer readFile.Close()
		return "", errors.New("failed to open file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		
		return "", errors.New("file is corrupted")
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	return PrintBanners(argsArr, arr), nil
}

func IsValid(s string) error {
	if len(s) > 400 {
		return errors.New("large input text")
	}
	for _, ch := range s {
		if (ch < ' ' && ch != 10 || ch > '~') && ch != '\r' {
			return errors.New("invalid character in input text")
		}
	}
	return nil
}

func PrintBanners(banners, arr []string) string {
	art := ""
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
			for _, j := range ch {
				n := (j-32)*9 + 1
				art += arr[int(n)+i]

			}
			art += "\n"

		}
		art += "\n"

	}
	return art
}



func main() {
	http.HandleFunc("/", printHandler)
	http.HandleFunc("/ascii-art", formHandler)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../static/style.css")
	})

	fmt.Printf("Starting server at http://localhost" + port + "/\n")
	if err := http.ListenAndServe(":2324", nil); err != nil {
		log.Fatal(err)
	}
}
