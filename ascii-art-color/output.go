package asciiArt

import "os"

func CreateFile(file, art string) {
	os.WriteFile(file, []byte(art), 0644)
}

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func createTestFile(fileName, content string)  {
// 	ioutil.WriteFile(fileName, []byte(content), 0644)

// }

// func main() {
// 	fileName := "test.txt"
// 	content := "This is the content of the test file."

// 	err := createTestFile(fileName, content)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Test file created successfully!")
// }
