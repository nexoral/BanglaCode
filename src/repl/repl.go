package repl

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"bufio"
	"fmt"
	"io"
	"strings"
)

const PROMPT = ">> "

const BANNER = `
 ____                   _        _____          _
|  _ \                 | |      / ____|        | |
| |_) | __ _ _ __   __ | | __ _| |     ___   __| | ___
|  _ < / _' | '_ \ / _' |/ _' | |    / _ \ / _' |/ _ \
| |_) | (_| | | | | (_| | (_| | |___| (_) | (_| |  __/
|____/ \__,_|_| |_|\__, |\__,_|\_____\___/ \__,_|\___|
                    __/ |
                   |___/

Welcome to BanglaCode - Bengali Programming Language
Created by Ankan from West Bengal, India
Type 'exit' or Ctrl+C to quit
Type 'help' for a list of keywords
Type 'clear' to clear screen
`

const HELP = `
BanglaCode Keywords:
  dhoro        - variable declaration (let/var)
  jodi         - if condition
  nahole       - else
  jotokkhon    - while loop
  ghuriye      - for loop
  kaj          - function definition
  ferao        - return statement
  class        - class definition
  notun        - new instance
  sotti        - true
  mittha       - false
  khali        - null
  ebong        - logical and (&&)
  ba           - logical or (||)
  na           - logical not (!)
  thamo        - break
  chharo       - continue
  ano          - import module
  pathao       - export function/class
  chesta       - try block
  dhoro_bhul   - catch block
  shesh        - finally block
  felo         - throw error

Built-in Functions:
  dekho(...)        - print values
  length(x)         - get length of string/array
  push(arr, val)    - add value to array
  pop(arr)          - remove last value from array
  keys(map)         - get keys of map
  type(x)           - get type of value
  string(x)         - convert to string
  number(x)         - convert to number

String Functions:
  upper(s), lower(s), trim(s), split(s, sep), join(arr, sep)
  indexOf(s, sub), substring(s, start, end), replace(s, old, new)

Array Functions:
  slice(arr, start, end), reverse(arr), sort(arr), includes(arr, val)

Math Functions:
  sqrt(x), pow(base, exp), floor(x), ceil(x), round(x)
  abs(x), min(...), max(...), random()

Utility Functions:
  time(), sleep(ms), input(prompt), exit(code)
  readFile(path), writeFile(path, content)
  http_server(port, handler), http_get(url)

REPL Commands:
  help         - show this help
  clear/cls    - clear screen
  exit/quit    - exit REPL

Example:
  >> dhoro naam = "Ankan"
  >> dekho("Namaskar", naam)
  Namaskar Ankan
`

// Start begins the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Fprint(out, BANNER)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Handle special commands
		if line == "exit" || line == "quit" {
			fmt.Fprintln(out, "Dhonnobad! (Thank you!)")
			return
		}

		if line == "help" {
			fmt.Fprint(out, HELP)
			continue
		}

		if line == "clear" || line == "cls" {
			// Clear screen using ANSI escape codes
			fmt.Fprint(out, "\033[2J\033[H")
			continue
		}

		if line == "" {
			continue
		}

		// Handle multi-line input
		input := line
		if needsMoreInput(line) {
			input = readMultiLine(scanner, line)
		}

		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() != object.NULL_OBJ && evaluated.Type() != object.ERROR_OBJ {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			} else if evaluated.Type() == object.ERROR_OBJ {
				io.WriteString(out, "\033[31m") // Red color
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\033[0m\n") // Reset color
			}
		}
	}
}

func needsMoreInput(line string) bool {
	openBraces := strings.Count(line, "{")
	closeBraces := strings.Count(line, "}")
	openParens := strings.Count(line, "(")
	closeParens := strings.Count(line, ")")
	openBrackets := strings.Count(line, "[")
	closeBrackets := strings.Count(line, "]")

	return openBraces > closeBraces || openParens > closeParens || openBrackets > closeBrackets
}

func readMultiLine(scanner *bufio.Scanner, initial string) string {
	var builder strings.Builder
	builder.WriteString(initial)
	builder.WriteString("\n")

	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(line)
		builder.WriteString("\n")

		if !needsMoreInput(builder.String()) {
			break
		}
		fmt.Print(".. ")
	}

	return builder.String()
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\033[31m") // Red color
	io.WriteString(out, "Oops! Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, "\033[0m") // Reset color
}
