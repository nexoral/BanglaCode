# GEMINI.md

This file provides guidance to Gemini Code Assist when working with code in this repository.

## Project Overview

**BanglaCode** - Educational Programming Language in Bengali

- **Language**: Go ≥1.20
- **Type**: Tree-walking interpreter for Bengali-syntax programming
- **Platform**: Cross-platform (Linux, macOS, Windows)
- **Purpose**: Make programming accessible to 300+ million Bengali speakers

## Commands

```bash
# Build & Run
go build -o banglacode main.go
./banglacode examples/hello.bang
./banglacode                          # REPL

# Test
go test ./...
go fmt ./...
go vet ./...

# Cross-compile
GOOS=windows GOARCH=amd64 go build -o banglacode.exe .
GOOS=darwin GOARCH=arm64 go build -o banglacode .
```

## Core Rules (NON-NEGOTIABLE)

1. **Bengali keywords**: Use Banglish (Bengali in English script)
2. **ALWAYS test**: Run test files after changes
3. **ALWAYS build**: `go build` after code changes
4. **Maintain compatibility**: Don't break existing BanglaCode programs
5. **Clear errors**: Bengali-friendly error messages
6. **Update docs**: README, SYNTAX.md, Documentation/

## Architecture

### Interpreter Pipeline
```
Source Code → Lexer → Parser → AST → Evaluator → Result
```

### Structure
```
src/
├── lexer/          # Tokenization (29 Bengali keywords)
├── parser/         # Pratt parsing, AST building
├── ast/            # Node definitions
├── object/         # Runtime types, Environment
├── evaluator/      # Tree-walking interpreter
│   ├── evaluator.go    # Main Eval() switch
│   ├── builtins.go     # 135+ built-in functions
│   ├── async.go        # Async/await, promises
│   ├── classes.go      # OOP support
│   ├── modules.go      # Import/export
│   └── errors.go       # Try/catch/finally
└── repl/           # Interactive shell

examples/           # BanglaCode example programs
test/              # Test files
Extension/         # VS Code extension
```

## Go Standards

### Error Handling
```go
// ✅ GOOD
if err != nil {
    return newError("ভুল: %s", err.Error())
}

// ❌ BAD
if err != nil {
    return nil
}
```

### Type Safety
```go
// ✅ Use object types
type Integer struct {
    Value int64
}

// ✅ Type assertions
intObj, ok := obj.(*object.Integer)
if !ok {
    return newError("সংখ্যা প্রত্যাশিত")
}
```

## Key Patterns

### Bengali Keywords
```go
// Token definitions in lexer/token.go
var keywords = map[string]TokenType{
    "dhoro":     LET,        // let/var
    "jodi":      IF,         // if
    "nahole":    ELSE,       // else
    "kaj":       FUNCTION,   // function
    "firao":     RETURN,     // return
    "proyash":   ASYNC,      // async
    "opekha":    AWAIT,      // await
    // ... 29 keywords total
}
```

### Built-in Functions
```go
// evaluator/builtins.go
var builtins = map[string]*object.Builtin{
    "dekho":        &object.Builtin{Fn: dekhoPrint},     // print
    "dorghyo":      &object.Builtin{Fn: dorghyoLength}, // length
    "dhokao":       &object.Builtin{Fn: dhokaoPush},    // push
    // ... 135+ functions
}
```

### Error Messages
```go
// ✅ Bengali-friendly errors
return newError("'%s' চিহ্ন অপারেটর সমর্থন করে না: %s", operator, left.Type())

// ✅ Context in errors
return newError("%d লাইনে: অজানা ফাংশন '%s'", node.Line, fn.Name)
```

## Documentation

Update when features change:
- README.md - Overview, installation, usage
- SYNTAX.md - Language syntax reference
- ARCHITECTURE.md - Interpreter design
- Documentation/ - Website docs
- Extension/ - VS Code extension features

## Testing

- Unit tests for lexer, parser, evaluator
- Integration tests for full programs
- Example programs in `examples/`
- REPL testing
- Cross-platform testing

## Bengali Language Features

### Keywords (29 total)
- Variables: `dhoro`, `sthir`, `protiti`
- Control: `jodi`, `nahole`, `jokhon`, `bhango`
- Functions: `kaj`, `firao`, `proyash`, `opekha`
- OOP: `dol`, `notun`, `ei`, `super`
- Modules: `ano`, `theke`, `pathao`, `hisabe`
- Error: `chesta`, `dhoro_bhul`, `shesh`, `chhar`

### Built-in Functions (135+)
- I/O: `dekho`, `input`, `file_lekho`, `file_poro`
- String: `dorghyo`, `boro_hater`, `choto_hater`
- Array: `dhokao`, `chhino`, `jog_koro`, `filter_koro`
- Math: `gononaa`, `bolod`, `muladhon`, `ghuriye`
- HTTP: `http_server`, `http_get`, `http_post`
- Database: `postgres_connect`, `mysql_connect`, `mongo_connect`
- Async: `ghumaao`, `somoy_dekhao`, `proyash_solve`

## Definition of "Done"

- ✅ `go build` passes
- ✅ `go test ./...` passes
- ✅ Existing BanglaCode programs work
- ✅ Documentation updated
- ✅ Error messages in Bengali
- ✅ Cross-platform tested
