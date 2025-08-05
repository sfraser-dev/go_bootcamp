// go run main.go myfile.txt
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	expectedNumOfArgs := 2
	if len(os.Args) != expectedNumOfArgs {
		fmt.Println("number of input args does not equal", expectedNumOfArgs)
		os.Exit(1)
	}

	progName := os.Args[0]
	inputFilename := os.Args[1]
	fmt.Println("program name is", progName)
	fmt.Println("input file name is", inputFilename)

	// os.File (f) type returned by os.Open implements
	// Read() from the io.Reader interface, ie, it
	// is promoted to type Reader interface too
	f, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Copy Reader interface fh to the Writer interface Stdout,
	// ie, print the contents of the file to the terminal
	io.Copy(os.Stdout, f)
}
