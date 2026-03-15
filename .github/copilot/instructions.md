# GitHub Copilot Instructions for BanglaCode

## Project Overview

**BanglaCode** - Educational Programming Language in Bengali
- **Language**: Go ≥1.20
- **Type**: Tree-walking interpreter
- **Keywords**: 29 Bengali keywords (Banglish)
- **Built-ins**: 135+ functions
- **Target**: Bengali-speaking students and developers

## Core Rules

### 1. Bengali-First
All keywords must be Bengali/Banglish:
- `dhoro` (let), `jodi` (if), `kaj` (function)
- `proyash` (async), `dol` (class), `ano` (import)

### 2. Error Messages in Bengali
```go
// ✅ GOOD
return newError("ভুল: %s", err.Error())
return newError("%d লাইনে: অজানা ফাংশন '%s'", line, name)

// ❌ BAD
return newError("Error: %s", err)
```

### 3. Compatibility
Don't break existing BanglaCode programs

## Commands

```bash
go build -o banglacode main.go
./banglacode examples/hello.bang
./banglacode                      # REPL
go test ./...
```

## Architecture

```
Source → Lexer → Parser → AST → Evaluator → Result

src/
├── lexer/          # Tokenization
├── parser/         # AST building
├── ast/            # Node types
├── object/         # Runtime types
├── evaluator/      # Execution
│   ├── builtins.go     # 135+ functions
│   ├── async.go        # Promises
│   ├── classes.go      # OOP
│   └── modules.go      # Import/export
└── repl/           # Interactive shell
```

## Go Standards

### Error Handling
```go
// ✅ GOOD - Bengali context
if err != nil {
    return newError("ভুল: %s", err.Error())
}
```

### Type Safety
```go
// ✅ Safe assertions
intObj, ok := obj.(*object.Integer)
if !ok {
    return newError("সংখ্যা প্রত্যাশিত, পেয়েছি: %s", obj.Type())
}
```

## Key Features

### 29 Bengali Keywords
`dhoro`, `jodi`, `nahole`, `kaj`, `firao`, `proyash`, `opekha`, `dol`, `notun`, `ei`, `ano`, `pathao`, `chesta`, `dhoro_bhul`, `shesh`

### 135+ Built-in Functions
- I/O: `dekho`, `input`
- String/Array: `dorghyo`, `dhokao`, `chhino`
- HTTP: `http_server`, `http_get`, `http_post`
- Database: `postgres_connect`, `mysql_connect`, `mongo_connect`, `redis_connect`
- Async: `ghumaao`, `proyash_solve`

## Documentation

Update when features change:
- SYNTAX.md - Language reference
- ARCHITECTURE.md - Design docs
- Documentation/ - Website
- Extension/ - VS Code extension

## Testing

- Unit tests: Lexer, parser, evaluator
- Integration: Full programs
- Examples: `examples/` directory
- Cross-platform: Linux, macOS, Windows

## Success Criteria

- ✅ `go build` passes
- ✅ `go test ./...` passes
- ✅ Existing programs work
- ✅ Docs updated
- ✅ Bengali-friendly errors
- ✅ Cross-platform tested
