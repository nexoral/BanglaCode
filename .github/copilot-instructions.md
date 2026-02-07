# Copilot Instructions for BanglaCode

BanglaCode is a Bengali-syntax programming language interpreter written in Go. It uses Banglish (Bengali words in English script) keywords and follows C-like strict syntax.

## Build and Run

```bash
# Build the interpreter
go build -o banglacode main.go

# Run REPL
./banglacode

# Run a file
./banglacode examples/hello.bang

# Alternative using go run
go run main.go examples/hello.bang
```

## Architecture

The interpreter follows a classic pipeline: **Lexer → Parser → AST → Evaluator**

```
src/
├── lexer/       # Tokenizer - converts source to tokens
│   ├── token.go     # Token type definitions and Banglish keyword mappings
│   └── lexer.go     # Tokenizer implementation
├── parser/      # Parser - builds AST from tokens
│   └── parser.go    # Recursive descent parser
├── ast/         # AST node definitions
│   └── ast.go
├── object/      # Runtime value types and environment
│   ├── object.go        # Object interface and types (Number, String, Boolean, etc.)
│   └── environment.go   # Variable scopes (lexical scoping)
├── evaluator/   # Tree-walking interpreter
│   ├── evaluator.go     # Core evaluation logic
│   └── builtins.go      # Built-in functions (dekho, dorghyo, json_poro, etc.)
└── repl/        # Interactive shell
    └── repl.go
```

## Key Conventions

### Banglish Keywords (src/lexer/token.go)
When adding new keywords, update both the `const` block and the `keywords` map:
- `dhoro` = variable declaration
- `kaj` = function
- `jodi`/`nahole` = if/else
- `jotokkhon`/`ghuriye` = while/for loops
- `ferao` = return
- `ei` = this (class context)

### Adding Built-in Functions (src/evaluator/builtins.go)
Built-in functions use Banglish names. Register new functions in the `builtins` map with Bengali-inspired names that are intuitive for Bengali speakers. After adding, update:
1. The `builtins` map in `builtins.go`
2. REPL help text in `src/repl/repl.go`
3. VSCode syntax highlighting in `VSCode_Extension/syntaxes/banglacode.tmLanguage.json`
4. VSCode snippets in `VSCode_Extension/snippets/banglacode.json`

### Object Types (src/object/object.go)
Runtime values implement the `Object` interface with `Type()` and `Inspect()` methods. The null value displays as "khali".

### File Extension
BanglaCode source files use the `.bang` extension.

### Module System
- `ano "file.bang"` imports a module
- `pathao kaj fn() {}` exports a function
- Modules are resolved relative to the importing file's directory

### JSON Functions
- `json_poro(str)` - Parse JSON string to BanglaCode object
- `json_banao(obj)` - Convert BanglaCode object to JSON string

### HTTP Response Helpers
- `uttor(res, body, [status], [contentType])` - Simple response helper
- `json_uttor(res, data, [status])` - JSON response with auto Content-Type
