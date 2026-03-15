# AGENTS.md

OpenAI Codex CLI Instructions for BanglaCode

## Project Overview

**BanglaCode** - Educational Programming Language in Bengali

- **Language**: Go ≥1.20
- **Type**: Tree-walking interpreter
- **Keywords**: 29 Bengali keywords (Banglish)
- **Built-ins**: 135+ functions
- **Purpose**: Accessible programming for Bengali speakers

## Build Commands

```bash
# Development
go build -o banglacode main.go
./banglacode examples/hello.bang
./banglacode                    # REPL

# Test
go test ./...
go fmt ./...
go vet ./...

# Cross-compile
GOOS=windows GOARCH=amd64 go build -o banglacode.exe .
```

## Core Principles

### 1. Bengali-First
- All keywords in Banglish (Bengali in English script)
- Error messages in Bengali
- Documentation bilingual (English + Bengali)

### 2. Educational Focus
- Clear, learnable syntax
- Comprehensive built-in functions
- Real-world capabilities (HTTP, DB, async)

### 3. Compatibility
- Don't break existing programs
- Maintain keyword consistency
- Preserve built-in function signatures

## Architecture

### Interpreter Pipeline
```
Source → Lexer → Parser → AST → Evaluator → Result
```

### Components
```
src/
├── lexer/          # Tokenization
├── parser/         # AST building
├── ast/            # Node types
├── object/         # Runtime types
├── evaluator/      # Execution
│   ├── builtins.go     # 135+ functions
│   ├── async.go        # Promises
│   ├── classes.go      # OOP
│   ├── modules.go      # Import/export
│   └── errors.go       # Try/catch
└── repl/           # Interactive shell
```

## Go Standards

### Error Handling
```go
// ✅ GOOD - Bengali messages
if err != nil {
    return newError("ভুল: %s", err.Error())
}

// ✅ Context
return newError("%d লাইনে: অজানা চিহ্ন '%s'", node.Line, token)
```

### Type Safety
```go
// ✅ Type assertions
intObj, ok := obj.(*object.Integer)
if !ok {
    return newError("সংখ্যা প্রত্যাশিত, পেয়েছি: %s", obj.Type())
}
```

## Key Features

### Keywords (29)
- `dhoro` (let), `jodi` (if), `kaj` (function)
- `proyash` (async), `opekha` (await)
- `dol` (class), `notun` (new), `ei` (this)
- `ano` (import), `pathao` (export)
- `chesta` (try), `dhoro_bhul` (catch)

### Built-ins (135+)
- I/O: `dekho`, `input`, file operations
- String/Array: `dorghyo`, `dhokao`, `chhino`
- HTTP: `http_server`, `http_get`, `http_post`
- Database: PostgreSQL, MySQL, MongoDB, Redis
- Async: `ghumaao`, promises

## Documentation

Update when features change:
- README.md - Overview
- SYNTAX.md - Language reference
- ARCHITECTURE.md - Design
- Documentation/ - Website
- Extension/ - VS Code extension

## Testing

- Unit tests: Lexer, parser, evaluator
- Integration: Full programs
- Examples: `examples/` directory
- Cross-platform: Linux, macOS, Windows

## Anti-Patterns

❌ English keywords
❌ Breaking existing programs
❌ Non-Bengali error messages
❌ Missing documentation
❌ Generic errors

## Success Criteria

- ✅ `go build` passes
- ✅ `go test ./...` passes
- ✅ Existing programs work
- ✅ Docs updated
- ✅ Bengali-friendly errors
- ✅ Cross-platform tested
