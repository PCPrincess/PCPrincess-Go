package main


import (
	"fmt"
	"os"
)

func main() {

	var input string
	fmt.Print("Enter a string of characters: ")
	fmt.Scan(&input)
	outputText, err := os.Create("Test.txt")
	if err != nil {
		 panic(err)
	}
	defer outputText.Close()
	outputText.WriteString(input)

}
