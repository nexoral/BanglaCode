# 🇮🇳 BanglaCode - Bengali Programming Language

> **Write code in your native language. Think in Bengali. Code in Bengali.**

BanglaCode is a **Bengali-syntax programming language** that brings programming education and development to Bengali speakers worldwide. Instead of writing code in English, you write it in **Banglish** (Bengali words in English script), making programming more intuitive and accessible.

Created by **Ankan Saha** from **West Bengal, India** 🏠

---

## 🌍 What is BanglaCode?

BanglaCode is a **complete, feature-rich programming language** with:

- 🗣️ **Bengali Keywords**: Use words like `dhoro` (variable), `kaj` (function), `jodi` (if)
- 🎯 **Modern Features**: Classes, async/await, error handling, modules, and more
- ⚡ **Fast Execution**: Tree-walking interpreter written in Go
- 📚 **Rich Ecosystem**: 45+ built-in functions for strings, arrays, math, I/O, networking
- 🔧 **Production-Ready**: Used for real applications, fully tested

### Why BanglaCode?

**"আমি একজন বাংলা মাধ্যমের ছাত্র। আমি logic তৈরি করতে পারি, কিন্তু সেই logic validate করতে Programming language এর syntax শিখতে হয়।"**

*"I am a Bengali medium student. I can create logic, but to validate that logic I need to learn programming language syntax."*

**The Problem:**
- Bengali speakers must learn English syntax to code
- Creates a barrier for students learning in Bengali
- Loses cultural and linguistic connection to technology

**The Solution:**
- Program entirely in Bengali/Banglish
- Same power as JavaScript, Python, Go
- Accessible to millions of Bengali speakers globally
- Preserves cultural heritage in technology

---

## 🚀 Key Features

### Language Features

#### 📘 Core Language
- **Variables**: `dhoro` (let), `sthir` (const), `bishwo` (global)
- **Data Types**: Numbers, Strings, Booleans, Arrays, Objects, null
- **Functions**: First-class, closures, arrow functions
- **Classes & OOP**: `sreni` (class), `notun` (new), `ei` (this)
- **Control Flow**: `jodi`/`nahole` (if/else), `jotokkhon` (while), `ghuriye` (for)
- **Error Handling**: `chesta`/`dhoro_bhul`/`shesh` (try/catch/finally)
- **Async/Await**: `proyash` (async), `opekha` (await), Promises

#### 📦 Advanced Features
- **Modules**: `ano` (import), `pathao` (export), `hisabe` (as)
- **Pattern Matching**: Destructuring for arrays and objects
- **Operators**: Full arithmetic, logical, bitwise support
- **Built-in Functions**: 45+ functions for strings, arrays, math, I/O, networking
- **File I/O**: Read/write files, directory operations
- **HTTP/Networking**: Client/server, JSON parsing, requests
- **Advanced Loops**: `break`, `continue`, labeled loops

### VS Code Extension Features

#### ✨ Syntax Highlighting
- Beautiful, intuitive color scheme for all BanglaCode syntax
- Highlights keywords, functions, strings, numbers, and comments
- Custom icons for `.bang`, `.bangla`, `.bong` files
- Dark & light theme support

#### 🔧 IntelliSense & Code Completion
- **All Keywords**: `dhoro`, `jodi`, `kaj`, `ferao`, `class`, `proyash`, etc.
- **45+ Built-in Functions**: With descriptions and usage examples
- **Auto-completion**: Variables, functions, and classes detected from your code
- **Parameter Hints**: See function parameters as you type
- **Smart Suggestions**: Context-aware completions

#### 📝 Code Snippets
Quick templates for common patterns (type prefix + Tab):

| Snippet | Expands To |
|---------|-----------|
| `dhoro` | Variable declaration |
| `sthir` | Constant declaration |
| `jodi` | If statement |
| `jodi-nahole` | If-else block |
| `ghuriye` | For loop |
| `jotokkhon` | While loop |
| `kaj` | Function definition |
| `arrow` | Arrow function |
| `sreni` | Class definition |
| `chesta` | Try-catch-finally |
| `ano` | Import statement |
| `pathao-kaj` | Export function |
| `server` | HTTP server template |
| `client` | HTTP client template |
| `main` | Main program template |

#### 📚 Hover Documentation
Hover over **any** keyword or function to see:
- **Bengali Meaning**: বাংলা অর্থ with Bangla script
- **English Description**: What it does
- **Usage Examples**: Real code examples
- **Parameters**: For functions and methods
- **Return Type**: What the function returns

#### 🎨 Visual Enhancements
- Custom file icons for BanglaCode files
- Syntax-aware bracket/parenthesis matching
- Code formatting support
- Comment toggling
- Smart indentation

## 📦 Installation

### Option 1: VS Code Marketplace (Recommended)

1. Open VS Code
2. Press `Ctrl+Shift+X` (Windows/Linux) or `Cmd+Shift+X` (Mac)
3. Search for: **BanglaCode**
4. Click **Install**
5. Done! Start creating `.bang` files

**[Install from Marketplace →](https://marketplace.visualstudio.com/items?itemName=AnkanSaha.banglacode)**

### Option 2: Open VSX Registry (VSCodium, etc.)

For users of **VSCodium**, **Theia**, or other VS Code forks:

1. Open Extensions in your editor
2. Search for: **BanglaCode**
3. Click **Install**

**[Install from Open VSX →](https://open-vsx.org/extension/AnkanSaha/banglacode)**

### Option 3: Local VSIX Installation

For development or offline installation:

1. Clone the repository:
   ```bash
   git clone https://github.com/nexoral/BanglaCode.git
   cd BanglaCode/Extension
   ```

2. Package the extension:
   ```bash
   npm install
   npx vsce package
   ```

3. Install in VS Code:
   - Press `Ctrl+Shift+P`
   - Search: "Install from VSIX"
   - Select the generated `.vsix` file

### Option 4: Manual Installation

1. Copy the `Extension` folder to:
   - **Windows**: `%USERPROFILE%\.vscode\extensions\AnkanSaha.banglacode`
   - **macOS**: `~/.vscode/extensions/AnkanSaha.banglacode`
   - **Linux**: `~/.vscode/extensions/AnkanSaha.banglacode`

2. Restart VS Code

## 💻 Getting Started

### Quick Start: Hello World

1. **Create a file**: `hello.bang`
2. **Type this code**:
```banglacode
dekho("Namaskar, West Bengal!");
```

3. **Run it**:
```bash
banglacode hello.bang
```

Output:
```
Namaskar, West Bengal!
```

### Common Code Examples

#### Variables & Data Types
```banglacode
dhoro naam = "Ankan";           // String
dhoro age = 25;                  // Number
dhoro active = sotti;            // Boolean (true)
dhoro items = [1, 2, 3];        // Array
dhoro person = {naam: "Ankan"}; // Object
```

#### Functions
```banglacode
kaj greet(name) {
    ferao "Hello, " + name;
}

dekho(greet("World"));

// Arrow function
dhoro add = (a, b) => a + b;
dekho(add(5, 3));  // Output: 8
```

#### Control Flow
```banglacode
jodi (age >= 18) {
    dekho("You are an adult");
} nahole {
    dekho("You are a minor");
}

ghuriye (dhoro i = 0; i < 5; i = i + 1) {
    dekho(i);
}
```

#### Classes & OOP
```banglacode
sreni Person {
    shuru(naam, boyosh) {
        ei.naam = naam;
        ei.boyosh = boyosh;
    }

    greet() {
        ferao "Hello, I am " + ei.naam;
    }
}

dhoro person = notun Person("Ankan", 25);
dekho(person.greet());
```

#### Error Handling
```banglacode
chesta {
    dhoro result = jSum(risky_data);
} dhoro_bhul (error) {
    dekho("Error:", error);
} shesh {
    dekho("Cleanup done");
}
```

#### Async/Await
```banglacode
proyash getData() {
    dhoro response = opekha fetch("https://api.example.com/data");
    ferao response;
}
```

## 🔑 Language Keywords

All keywords are in **Banglish** (Bengali words written in English script):

| Keyword | Bengali | Purpose |
|---------|---------|---------|
| **Variables** | | |
| `dhoro` | ধরো | Let/variable declaration |
| `sthir` | স্থির | Const/constant declaration |
| `bishwo` | বিশ্ব | Global variable |
| **Control Flow** | | |
| `jodi` | যদি | If statement |
| `nahole` | নাহলে | Else clause |
| `jotokkhon` | যতক্ষণ | While loop |
| `ghuriye` | ঘুরিয়ে | For loop |
| `thamo` | থামো | Break statement |
| `chharo` | ছাড়ো | Continue statement |
| **Functions & Classes** | | |
| `kaj` | কাজ | Function definition |
| `ferao` | ফেরাও | Return statement |
| `sreni` | স্রেণী | Class definition |
| `shuru` | শুরু | Constructor |
| `notun` | নতুন | New (instantiation) |
| `ei` | ঈ | This (self reference) |
| **Async/Promise** | | |
| `proyash` | প্রয়াশ | Async function |
| `opekha` | অপেক্ষা | Await promise |
| **Error Handling** | | |
| `chesta` | চেষ্টা | Try block |
| `dhoro_bhul` | ধরো ভুল | Catch block |
| `shesh` | শেষ | Finally block |
| `felo` | ফেলো | Throw error |
| **Modules** | | |
| `ano` | আনো | Import |
| `pathao` | পাঠাও | Export |
| `hisabe` | হিসাবে | As (alias) |
| **Literals** | | |
| `sotti` | সত্যি | True |
| `mittha` | মিথ্যা | False |
| `khali` | খালি | Null/undefined |

**[View complete syntax guide →](https://banglacode.nexoral.in/docs/syntax)**

## 📚 Built-in Functions

BanglaCode includes **45+ powerful built-in functions** for common tasks:

### I/O Functions
| Function | Purpose |
|----------|---------|
| `dekho()` | Print to console |
| `input()` | Read user input |
| `logg()` | Log with timestamp |

### String Functions
| Function | Purpose |
|----------|---------|
| `boroHater()` | Convert to UPPERCASE |
| `chotoHater()` | Convert to lowercase |
| `khondKoro()` | Split string |
| `joinKoro()` | Join array to string |
| `shikal_udhar()` | Trim whitespace |
| `antra()` | Get character at index |

### Array Functions
| Function | Purpose |
|----------|---------|
| `dorghyo()` | Get length |
| `dhokao()` | Add element (push) |
| `berKoro()` | Remove element (pop) |
| `filter()` | Filter elements |
| `map()` | Transform elements |
| `reduce()` | Aggregate elements |
| `find()` | Find element |

### Math Functions
| Function | Purpose |
|----------|---------|
| `borgomul()` | Square root |
| `ghon()` | Cube |
| `abs()` | Absolute value |
| `round()` | Round number |
| `floor()` | Floor division |
| `ceil()` | Ceiling |
| `random()` | Random number |
| `sin()`, `cos()`, `tan()` | Trigonometry |

### File I/O Functions
| Function | Purpose |
|----------|---------|
| `file_read()` | Read file |
| `file_write()` | Write file |
| `file_append()` | Append to file |
| `dir_list()` | List directory |

### HTTP/Network Functions
| Function | Purpose |
|----------|---------|
| `fetch()` | Make HTTP request |
| `json_parse()` | Parse JSON |
| `json_stringify()` | Convert to JSON |
| `server_chalu()` | Start HTTP server |

### Type Functions
| Function | Purpose |
|----------|---------|
| `prokar()` | Get type |
| `akorBhumi()` | Convert to integer |
| `dashomikBhumi()` | Convert to float |
| `shobolKoro()` | Convert to string |

**[View complete function reference →](https://banglacode.nexoral.in/docs/functions)**

## ⌨️ Code Snippets

Type snippet prefixes and press `Tab` to expand:

| Prefix | Expands To |
|--------|-----------|
| `dhoro` | Variable declaration with assignment |
| `sthir` | Constant declaration |
| `jodi` | If statement |
| `jodi-nahole` | If-else block |
| `ghuriye` | For loop with counter |
| `jotokkhon` | While loop |
| `kaj` | Function definition |
| `arrow` | Arrow function |
| `sreni` | Class with constructor |
| `chesta` | Try-catch-finally block |
| `ano` | Import statement |
| `pathao` | Export statement |
| `pathao-kaj` | Export function |
| `server` | HTTP server template |
| `client` | HTTP client example |
| `main` | Main program template |

## 💾 File Support

The extension supports **three file extensions**:

| Extension | Usage |
|-----------|-------|
| `.bang` | Primary extension (recommended) |
| `.bangla` | Alternative (বাংলা) |
| `.bong` | Alternative (বং) |

All extensions provide identical features with full syntax highlighting and IntelliSense.

## 📋 Requirements

- **VS Code**: 1.74.0 or higher
- **BanglaCode Interpreter** (optional): To run `.bang` files
  - [Download BanglaCode →](https://github.com/nexoral/BanglaCode/releases)
  - Or install via:
    ```bash
    # Linux/macOS
    curl -fsSL https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.sh | bash

    # Windows
    irm https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.ps1 | iex
    ```

## ⚙️ Extension Features

This extension provides complete language support:

- ✅ **Syntax Highlighting**: Beautiful colors for all BanglaCode syntax
- ✅ **IntelliSense**: Auto-complete for keywords, functions, variables
- ✅ **Code Snippets**: 15+ templates for common patterns
- ✅ **Hover Docs**: Information on keywords and built-ins
- ✅ **File Icons**: Custom icons for `.bang` files
- ✅ **Error Squiggles**: Instant feedback (requires interpreter)
- ✅ **Smart Indentation**: Automatic formatting
- ✅ **Bracket Matching**: Easy parenthesis tracking
- ✅ **Comments**: Toggle comments with Ctrl+/

## 🔧 Configuration

No configuration needed! The extension works out of the box.

**Optional**: To run code directly in VS Code, install the [BanglaCode Interpreter](https://github.com/nexoral/BanglaCode).

## 📖 Learning Resources

### Official Documentation
- **[Language Syntax](https://banglacode.nexoral.in/docs/syntax)** - Complete language reference
- **[Functions Reference](https://banglacode.nexoral.in/docs/functions)** - All built-in functions
- **[Examples](https://github.com/nexoral/BanglaCode/tree/main/examples)** - Code examples
- **[Tutorial](https://banglacode.nexoral.in/docs/introduction)** - Getting started guide

### Community & Support
- **[GitHub Repository](https://github.com/nexoral/BanglaCode)** - Source code
- **[Report Issues](https://github.com/nexoral/BanglaCode/issues)** - Bug reports
- **[Discussions](https://github.com/nexoral/BanglaCode/discussions)** - Q&A

## 🐛 Known Issues & Feedback

Found a bug? Have a suggestion? **[Report it on GitHub →](https://github.com/nexoral/BanglaCode/issues)**

We actively maintain BanglaCode and welcome your feedback!

## 🌟 About BanglaCode

### Mission

Make programming accessible to **500 million+ Bengali speakers** worldwide by removing the language barrier.

### Why It Matters

**The Problem:**
- Bengali students learn in Bengali but must code in English
- Creates cognitive overload: thinking in one language, coding in another
- Loses millions of potential programmers due to language barriers
- Technology remains English-centric despite global diversity

**Our Solution:**
- Write complete, production-quality code in Bengali
- Same power as JavaScript, Python, or Go
- Accessible to students, professionals, and hobbyists
- Preserves cultural heritage in technology
- Opens doors for millions previously excluded from programming

### Creator's Vision

> *"আমি একজন বাংলা মাধ্যমের ছাত্র। আমি logic তৈরি করতে পারি, কিন্তু সেই logic validate করতে Programming language এর syntax শিখতে হয়।"*
>
> *"I am a Bengali medium student. I can create logic, but to validate that logic I need to learn programming language syntax. This language is for people like me."*
>
> — **Ankan Saha**, Creator

## 🇮🇳 আমাদের ভাষায় কোড লিখুন

বাংলা ভাষায় প্রোগ্রামিং শিখুন। বাংলায় চিন্তা করুন। বাংলায় কোড লিখুন।

BanglaCode হল একটি সম্পূর্ণ প্রোগ্রামিং ভাষা যা বাংলা ভাষায় তৈরি।

---

## 🤝 Contributing

BanglaCode is open-source! Contributions are welcome:

- **Report bugs**: [GitHub Issues](https://github.com/nexoral/BanglaCode/issues)
- **Suggest features**: [GitHub Discussions](https://github.com/nexoral/BanglaCode/discussions)
- **Contribute code**: [GitHub PRs](https://github.com/nexoral/BanglaCode)
- **Improve docs**: Help translate or improve documentation

---

## 📄 License

MIT License - Free for personal and commercial use

---

**Created with ❤️ in West Bengal, India**

**[Visit Official Website →](https://banglacode.nexoral.in)** | **[GitHub Repository →](https://github.com/nexoral/BanglaCode)**
