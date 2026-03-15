---
name: banglacode-development
description: Development rules for BanglaCode programming language interpreter
version: 1.0.0
tags: [golang, interpreter, bengali, language-design, education]
author: BanglaCode Team
---

# BanglaCode Development Skill

## Project Identity

**BanglaCode** - Educational Programming Language in Bengali
- Go ≥1.20
- Tree-walking interpreter
- 29 Bengali keywords (Banglish)
- 135+ built-in functions
- Target: 300+ million Bengali speakers

## Mandatory Workflows

### After EVERY Code Change
```bash
go build -o banglacode main.go
go test ./...
go fmt ./...
go vet ./...
```

### For ANY Language Feature
1. Update lexer if new keyword
2. Update parser if new syntax
3. Update evaluator for execution
4. Add tests in `test/`
5. Update SYNTAX.md
6. Update Documentation/

## Definition of "Done"

- ✅ `go build` passes
- ✅ `go test ./...` passes
- ✅ Existing programs work
- ✅ SYNTAX.md updated
- ✅ Documentation updated
- ✅ Error messages in Bengali
- ✅ Examples work
- ✅ Cross-platform tested

## Architecture

### Interpreter Pipeline
```
Source Code (.bang)
    ↓
Lexer (Tokenization)
    ↓
Parser (AST Building)
    ↓
Evaluator (Execution)
    ↓
Result
```

### Components
```
src/
├── lexer/
│   ├── lexer.go        # Tokenizer
│   └── token.go        # 29 Bengali keywords
├── parser/
│   ├── parser.go       # Pratt parsing
│   └── precedence.go   # Operator precedence
├── ast/
│   └── ast.go          # AST node types
├── object/
│   ├── object.go       # Runtime types
│   └── environment.go  # Variable scopes
├── evaluator/
│   ├── evaluator.go    # Main Eval()
│   ├── builtins.go     # 135+ functions
│   ├── async.go        # Promises
│   ├── classes.go      # OOP
│   ├── modules.go      # Import/export
│   └── errors.go       # Try/catch/finally
└── repl/
    └── repl.go         # Interactive shell
```

## Go Standards (STRICT)

### Error Handling - Bengali Messages
```go
// ✅ REQUIRED - Bengali-friendly
if err != nil {
    return newError("ভুল: %s", err.Error())
}

// ✅ Context with line numbers
return newError("%d লাইনে: অজানা ফাংশন '%s'", node.Line, fnName)

// ❌ FORBIDDEN - English only
if err != nil {
    return newError("Error: %s", err)
}
```

### Type Assertions
```go
// ✅ GOOD - Safe type assertion
intObj, ok := obj.(*object.Integer)
if !ok {
    return newError("সংখ্যা প্রত্যাশিত, পেয়েছি: %s", obj.Type())
}

// ❌ BAD - Unsafe
intObj := obj.(*object.Integer)  // Can panic!
```

### Object Types
```go
// ✅ Define runtime types
type Integer struct {
    Value int64
}

type String struct {
    Value string
}

// ✅ Implement Object interface
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
```

## Key Patterns

### Bengali Keywords
```go
// lexer/token.go
var keywords = map[string]TokenType{
    // Variables
    "dhoro":        LET,         // let/var
    "sthir":        CONST,       // const
    "protiti":      FOR,         // for

    // Control Flow
    "jodi":         IF,          // if
    "nahole":       ELSE,        // else
    "jokhon":       WHILE,       // while
    "bhango":       BREAK,       // break

    // Functions
    "kaj":          FUNCTION,    // function
    "firao":        RETURN,      // return
    "proyash":      ASYNC,       // async
    "opekha":       AWAIT,       // await

    // OOP
    "dol":          CLASS,       // class
    "notun":        NEW,         // new
    "ei":           THIS,        // this
    "super":        SUPER,       // super

    // Modules
    "ano":          IMPORT,      // import
    "theke":        FROM,        // from
    "pathao":       EXPORT,      // export
    "hisabe":       AS,          // as

    // Error Handling
    "chesta":       TRY,         // try
    "dhoro_bhul":   CATCH,       // catch
    "shesh":        FINALLY,     // finally
    "chhar":        THROW,       // throw
}
```

### Built-in Functions
```go
// evaluator/builtins.go
var builtins = map[string]*object.Builtin{
    // I/O
    "dekho":           &object.Builtin{Fn: dekhoPrint},           // print
    "input":           &object.Builtin{Fn: inputRead},            // input

    // String/Array
    "dorghyo":         &object.Builtin{Fn: dorghyoLength},       // length
    "dhokao":          &object.Builtin{Fn: dhokaoPush},          // push
    "chhino":          &object.Builtin{Fn: chhinoPop},           // pop

    // HTTP
    "http_server":     &object.Builtin{Fn: httpServer},          // HTTP server
    "http_get":        &object.Builtin{Fn: httpGet},             // GET request
    "http_post":       &object.Builtin{Fn: httpPost},            // POST request

    // Database
    "postgres_connect":&object.Builtin{Fn: postgresConnect},     // PostgreSQL
    "mysql_connect":   &object.Builtin{Fn: mysqlConnect},        // MySQL
    "mongo_connect":   &object.Builtin{Fn: mongoConnect},        // MongoDB
    "redis_connect":   &object.Builtin{Fn: redisConnect},        // Redis

    // Async
    "ghumaao":         &object.Builtin{Fn: ghumaaoSleep},        // sleep
    "proyash_solve":   &object.Builtin{Fn: proyashSolve},        // Promise.resolve

    // ... 135+ total
}
```

### AST Nodes
```go
// ast/ast.go
type LetStatement struct {
    Token token.Token  // DHORO token
    Name  *Identifier
    Value Expression
}

type IfExpression struct {
    Token       token.Token  // JODI token
    Condition   Expression
    Consequence *BlockStatement
    Alternative *BlockStatement
}

type FunctionLiteral struct {
    Token      token.Token  // KAJ token
    Parameters []*Identifier
    Body       *BlockStatement
    Async      bool         // proyash flag
}
```

## Testing Requirements

### Unit Tests
```go
func TestLetStatement(t *testing.T) {
    input := `dhoro x = 5;`

    program := parseProgram(input)
    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements বিবৃতি ১টি থাকা উচিত")
    }

    stmt := program.Statements[0].(*ast.LetStatement)
    if stmt.Name.Value != "x" {
        t.Errorf("নাম 'x' হওয়া উচিত, পেয়েছি %s", stmt.Name.Value)
    }
}
```

### Integration Tests
```go
// test/integration_test.go
func TestHttpServer(t *testing.T) {
    input := `
    dhoro server = http_server(8080, kaj() {
        dekho("সার্ভার চালু");
    });
    `

    evaluated := testEval(input)
    if !isNull(evaluated) {
        t.Errorf("সার্ভার তৈরি ব্যর্থ")
    }
}
```

## Documentation Requirements

### SYNTAX.md
Update when language features change:
- New keywords
- New operators
- New built-in functions
- Syntax changes

### ARCHITECTURE.md
Update when internals change:
- New AST node types
- Parser modifications
- Evaluator changes
- Performance improvements

### Documentation/ (Website)
Update for user-facing changes:
- Tutorials
- API reference
- Examples
- Best practices

### Extension/ (VS Code)
Update when language features added:
- Syntax highlighting
- IntelliSense snippets
- Keyword completions

## Security Standards

### Input Validation
```go
// ✅ Validate file paths
func validatePath(path string) error {
    if strings.Contains(path, "..") {
        return fmt.Errorf("পথে '..' অনুমোদিত নয়")
    }
    return nil
}
```

### Safe Execution
```go
// ✅ Timeout for async operations
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

## Anti-Patterns (FORBIDDEN)

❌ English keywords (must be Bengali/Banglish)
❌ Breaking existing BanglaCode programs
❌ Non-Bengali error messages
❌ Missing tests for new features
❌ Undocumented syntax changes
❌ Unsafe type assertions (without ok check)
❌ Generic error messages
❌ Missing examples

## Workflow Guidelines

### When Adding New Keyword
1. Add to `lexer/token.go` keywords map
2. Add token type constant
3. Update parser to handle new syntax
4. Update evaluator for execution
5. Add tests
6. Update SYNTAX.md
7. Update VS Code extension syntax

### When Adding Built-in Function
1. Implement in `evaluator/builtins.go`
2. Add to `builtins` map
3. Follow naming: Bengali verb (e.g., `dekho`, `dhokao`)
4. Add validation and error handling
5. Add tests
6. Document in SYNTAX.md
7. Add VS Code snippet

### When Fixing Bugs
1. Write failing test
2. Fix bug
3. Verify test passes
4. Check for similar issues
5. Update CHANGELOG.md

## Success Criteria

Every task must meet ALL:
- ✅ `go build` passes
- ✅ `go test ./...` passes
- ✅ Existing programs still work
- ✅ SYNTAX.md updated
- ✅ Documentation updated
- ✅ Error messages in Bengali
- ✅ Examples work
- ✅ Cross-platform tested (Linux, macOS, Windows)
- ✅ VS Code extension updated
