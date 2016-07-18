package main

import (
	"fmt"
	"go/token"
)

func main() {
	tokenChecker := token.Lookup("b")
	fmt.Println(tokenChecker) // output: IDENT
	tokenChecker.String()
	fmt.Println(tokenChecker) // output: IDENT
	fmt.Printf("%T\n", tokenChecker) // type: token.Token
	// Need More Research~
}
