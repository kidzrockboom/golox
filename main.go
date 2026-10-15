package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var hadError bool = false

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

	// Indicate an error in the exit code
	if hadError {
		os.Exit(65)
	}

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

		// reset flag so it doesnt kill the user session from one error
		hadError = false

	}
}

func run(source string) {
	scanner := Scanner(source)
	tokens := scanner.scanTokens()

	for _, token := range tokens {
		println(token)
	}
}

func erreur(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Errorf("[Line %d] Error %s: %s", line, where, message)
	hadError = true
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
