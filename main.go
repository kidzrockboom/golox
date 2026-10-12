package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func readFileError(err error) {
	if err != nil {
		panic(err)
	}
}

func runFile(path string) {
	path = filepath.Join(path)

	data, err := os.ReadFile(path)
	readFileError(err)

	run(string(data)) //convert []byte to string

}

func runPrompt() {
	// create a reader to listen to standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Golox. type 'exit' to quit")
	for {
		fmt.Print("golox> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}

		if input == "" {
			continue
		}

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		run(input)

	}
}

func run(source string) {
	//todo
}

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Usage: Golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
