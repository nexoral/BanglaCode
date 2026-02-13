import Link from "next/link";
import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";
import { Download, Terminal, Box } from "lucide-react";

export default function Installation() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Getting Started
        </span>
      </div>

      <h1>Installation</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Install BanglaCode on your machine and start coding in Bengali.
        Choose from pre-built binaries or build from source.
      </p>

      <h2>Quick Install Options</h2>

      <div className="grid sm:grid-cols-2 gap-4 my-6">
        <a
          href="https://github.com/nexoral/BanglaCode/releases"
          target="_blank"
          rel="noopener noreferrer"
          className="group p-4 rounded-lg border border-border hover:border-primary/50 hover:bg-secondary/30 transition-colors"
        >
          <Download className="w-8 h-8 text-primary mb-2" />
          <h4 className="font-medium group-hover:text-primary transition-colors">
            Download Binary
          </h4>
          <p className="text-sm text-muted-foreground mt-1">
            Pre-built binaries for Windows, macOS, and Linux
          </p>
        </a>
        <div className="p-4 rounded-lg border border-border bg-secondary/20">
          <Box className="w-8 h-8 text-primary mb-2" />
          <h4 className="font-medium">Build from Source</h4>
          <p className="text-sm text-muted-foreground mt-1">
            Requires Go 1.20+ installed
          </p>
        </div>
      </div>

      <h2>Prerequisites</h2>

      <p>
        To build BanglaCode from source, you need <strong>Go 1.20 or later</strong> installed
        on your system. Download it from <a href="https://go.dev/dl/" target="_blank" rel="noopener noreferrer" className="text-primary hover:underline">go.dev/dl</a>.
      </p>

      <p>Verify Go installation:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`go version
# Should show: go version go1.20+ ...`}
      />

      <h2>Building from Source</h2>

      <h3>Step 1: Clone the Repository</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`git clone https://github.com/nexoral/BanglaCode.git
cd BanglaCode`}
      />

      <h3>Step 2: Build the Binary</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`go build -o banglacode main.go`}
      />

      <h3>Step 3: Verify Installation</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`./banglacode --version`}
      />

      <h3>Optional: Add to PATH</h3>

      <p>To run <code>banglacode</code> from anywhere:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Linux/macOS
sudo mv banglacode /usr/local/bin/

# Or add to your PATH in ~/.bashrc or ~/.zshrc
export PATH="$PATH:/path/to/banglacode"

# Windows (PowerShell as Admin)
# Move to a folder in your PATH or add the folder to PATH`}
      />

      <h2>Cross-Compilation</h2>

      <p>Build for different platforms:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Windows
GOOS=windows GOARCH=amd64 go build -o banglacode.exe main.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o banglacode-mac main.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o banglacode-mac-arm main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o banglacode-linux main.go`}
      />

      <h2>Running Your First Program</h2>

      <h3>Create a File</h3>

      <p>Create a file named <code>hello.bang</code>:</p>

      <CodeBlock
        filename="hello.bang"
        code={`// Your first BanglaCode program
dekho("Namaskar, BanglaCode!");

dhoro naam = "World";
dekho("Hello,", naam);`}
      />

      <h3>Run the Program</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`./banglacode hello.bang`}
      />

      <p>Output:</p>

      <CodeBlock
        language="output"
        showLineNumbers={false}
        code={`Namaskar, BanglaCode!
Hello, World`}
      />

      <h2>Interactive REPL</h2>

      <p>Start the interactive mode by running without arguments:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`./banglacode`}
      />

      <p>You&apos;ll see a prompt where you can type BanglaCode directly:</p>

      <CodeBlock
        language="output"
        showLineNumbers={false}
        code={`BanglaCode REPL v1.0
Type 'help' for commands, 'exit' to quit

>> dhoro x = 5
>> dekho(x * 2)
10
>> kaj double(n) { ferao n * 2; }
>> double(21)
42`}
      />

      <h2>VS Code Extension</h2>

      <p>
        For syntax highlighting and code snippets, install the BanglaCode VS Code extension:
      </p>

      <ol>
        <li>Open VS Code</li>
        <li>Go to Extensions (Ctrl+Shift+X)</li>
        <li>Search for &quot;BanglaCode&quot;</li>
        <li>Click Install</li>
      </ol>

      <p>
        Or install from the Extension folder in the repository:
      </p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Copy extension to VS Code extensions folder
cp -r Extension ~/.vscode/extensions/banglacode`}
      />

      <h2>Troubleshooting</h2>

      <h3>Command not found</h3>

      <p>Make sure the binary is in your PATH or use the full path:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Use full path
/path/to/banglacode hello.bang

# Or run with go directly
go run main.go hello.bang`}
      />

      <h3>Permission denied (Linux/macOS)</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`chmod +x banglacode`}
      />

      <h3>Build errors</h3>

      <p>Make sure you have Go 1.20+ and all dependencies:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`go mod tidy
go build -o banglacode main.go`}
      />

      <DocNavigation currentPath="/docs/installation" />
    </div>
  );
}
