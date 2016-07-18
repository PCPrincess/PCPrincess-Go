package main

// IMPORTANT: FUNCTIONS AND/OR VARIABLES MUST BE CAPITALIZED IF THEY WILL BE
// USED OUTSIDE THE PACKAGE (SEE: 'Plus' and 'ValExists' Below)
import (
	"fmt"
	"goWeb/variadic"
	"goWeb/commaOkIdiom"

)

const (
	A = iota	//0
	B           //1
	C           //2
)

func main() {

	var userInput string
	fmt.Print("Enter A Value Below 100: ")
	fmt.Scan(&userInput)
	fmt.Println(A, B, C, userInput, variadic.Plus(A, B, C))
	commaOkIdiom.ValExists()

}

