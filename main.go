package main

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"BanglaCode/src/repl"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	// Check for command line arguments
	if len(os.Args) == 1 {
		// No arguments - start REPL
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Namaskar %s! Welcome to BanglaCode!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Check for flags
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	if os.Args[1] == "--version" || os.Args[1] == "-v" {
		printVersion()
		return
	}

	// Execute file
	filename := os.Args[1]
	runFile(filename)
}

func printHelp() {
	help := `
BanglaCode - Bengali Programming Language
Created by Ankan from West Bengal, India

Usage:
  banglacode                  Start interactive REPL
  banglacode <file.bang>      Execute a BanglaCode file
  banglacode --help, -h       Show this help message
  banglacode --version, -v    Show version information

Examples:
  banglacode                  # Start REPL
  banglacode hello.bang       # Run hello.bang file

For more information, see SYNTAX.md
`
	fmt.Println(help)
}

func printVersion() {
	fmt.Println("BanglaCode v2.0.0")
	fmt.Println("Bengali Programming Language")
	fmt.Println("Created by Ankan from West Bengal, India")
}

func runFile(filename string) {
	// Read file
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Set current directory for module imports
	absPath, _ := filepath.Abs(filename)
	evaluator.SetCurrentDir(filepath.Dir(absPath))

	// Create environment
	env := object.NewEnvironment()

	// Lex
	l := lexer.New(string(content))

	// Parse
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Fprintln(os.Stderr, "\033[31mParser errors:\033[0m")
		for _, msg := range p.Errors() {
			fmt.Fprintf(os.Stderr, "\t%s\n", msg)
		}
		os.Exit(1)
	}

	// Evaluate
	result := evaluator.Eval(program, env)

	if result != nil && result.Type() == object.ERROR_OBJ {
		fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m\n", result.Inspect())
		os.Exit(1)
	}
}
