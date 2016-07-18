package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	testString := "Huzzah, We Passed The Test!"
	// May also use arguments @ command line: var := os.Args[1]

	// Very simple HTML formatting using only stdout
	fmt.Println(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` +
		testString +
		`</h1>
	</body>
	</html>
	`)

	strHTML := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` +
		testString +
		`</h1>
	</body>
	</html>
	`)

	indexFile, err := os.Create("index.html")

	if err != nil {
		log.Fatal("Error Creating File", err)
	}
	defer indexFile.Close()

	// Copying the 'string HTML' into a HTML file (index.html)
	io.Copy(indexFile, strings.NewReader(strHTML))

}
