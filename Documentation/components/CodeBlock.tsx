"use client";

import { useState } from "react";
import { Copy, Check } from "lucide-react";
import clsx from "clsx";

interface CodeBlockProps {
  code: string;
  language?: string;
  filename?: string;
  showLineNumbers?: boolean;
}

// BanglaCode syntax highlighting
function highlightBanglaCode(code: string): string {
  // Keywords
  const keywords = [
    "dhoro", "jodi", "nahole", "jotokkhon", "ghuriye", "kaj", "ferao",
    "sreni", "shuru", "notun", "sotti", "mittha", "khali", "ebong", "ba", "na",
    "thamo", "chharo", "ano", "pathao", "hisabe", "chesta", "dhoro_bhul", "shesh", "felo"
  ];

  // Built-in functions
  const builtins = [
    "dekho", "nao", "dhoron", "lipi", "sonkha", "dorghyo", "dhokao", "berKoro",
    "kato", "ulto", "ache", "saja", "boroHater", "chotoHater", "bhag", "joro",
    "chhanto", "khojo", "angsho", "bodlo", "borgomul", "ghat", "niche", "upore",
    "kache", "niratek", "choto", "boro", "chabi", "poro", "lekho", "somoy",
    "lotto", "ghum", "bondho", "server_chalu", "anun", "uttor", "json_uttor",
    "json_poro", "json_banao"
  ];

  let highlighted = code
    // Escape HTML
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    // Comments
    .replace(/(\/\/[^\n]*)/g, '<span class="text-gray-500">$1</span>')
    // Strings
    .replace(/("(?:[^"\\]|\\.)*")/g, '<span class="text-green-400">$1</span>')
    .replace(/('(?:[^'\\]|\\.)*')/g, '<span class="text-green-400">$1</span>')
    // Numbers
    .replace(/\b(\d+\.?\d*)\b/g, '<span class="text-orange-400">$1</span>');

  // Keywords (must come after string highlighting)
  keywords.forEach(kw => {
    const regex = new RegExp(`\\b(${kw})\\b(?![^<]*>)`, "g");
    highlighted = highlighted.replace(regex, '<span class="text-purple-400 font-semibold">$1</span>');
  });

  // Built-in functions
  builtins.forEach(fn => {
    const regex = new RegExp(`\\b(${fn})\\s*\\(`, "g");
    highlighted = highlighted.replace(regex, '<span class="text-blue-400">$1</span>(');
  });

  // Operators
  highlighted = highlighted
    .replace(/(\+\=|\-\=|\*\=|\/\=|\=\=|\!\=|\&lt;\=|\&gt;\=|\+|\-|\*|\/|\%|\&lt;|\&gt;)/g,
      '<span class="text-yellow-400">$1</span>');

  return highlighted;
}

export default function CodeBlock({ code, language = "banglacode", filename, showLineNumbers = true }: CodeBlockProps) {
  const [copied, setCopied] = useState(false);

  const copyToClipboard = async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const lines = code.split("\n");
  const highlighted = language === "banglacode" ? highlightBanglaCode(code) : code;

  return (
    <div className="relative group rounded-lg overflow-hidden border border-border bg-[#1e1e1e] my-4">
      {/* Header */}
      {filename && (
        <div className="flex items-center justify-between px-4 py-2 bg-[#2d2d2d] border-b border-border">
          <span className="text-xs font-mono text-muted-foreground">{filename}</span>
          <span className="text-xs text-muted-foreground uppercase">{language}</span>
        </div>
      )}

      {/* Copy button */}
      <button
        onClick={copyToClipboard}
        className="absolute top-2 right-2 p-2 rounded-md bg-secondary/50 opacity-0 group-hover:opacity-100 transition-opacity hover:bg-secondary"
        title="Copy code"
      >
        {copied ? (
          <Check className="w-4 h-4 text-green-400" />
        ) : (
          <Copy className="w-4 h-4 text-muted-foreground" />
        )}
      </button>

      {/* Code content */}
      <div className="overflow-x-auto">
        <pre className="p-4 text-sm font-mono leading-6">
          {showLineNumbers ? (
            <table className="w-full">
              <tbody>
                {lines.map((line, i) => (
                  <tr key={i} className="hover:bg-white/5">
                    <td className="pr-4 text-right text-gray-600 select-none w-8 align-top">
                      {i + 1}
                    </td>
                    <td
                      className="text-gray-300"
                      dangerouslySetInnerHTML={{
                        __html: language === "banglacode"
                          ? highlightBanglaCode(line)
                          : line.replace(/</g, "&lt;").replace(/>/g, "&gt;")
                      }}
                    />
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <code
              className="text-gray-300"
              dangerouslySetInnerHTML={{ __html: highlighted }}
            />
          )}
        </pre>
      </div>
    </div>
  );
}

// Simple inline code component
export function InlineCode({ children }: { children: React.ReactNode }) {
  return (
    <code className="px-1.5 py-0.5 bg-secondary/50 rounded text-primary font-mono text-sm">
      {children}
    </code>
  );
}
