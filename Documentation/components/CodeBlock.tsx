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

// BanglaCode syntax highlighting using token-based approach
function highlightBanglaCode(code: string): string {
  // Keywords
  const keywords = new Set([
    "dhoro", "jodi", "nahole", "jotokkhon", "ghuriye", "kaj", "ferao",
    "sreni", "shuru", "notun", "sotti", "mittha", "khali", "ebong", "ba", "na",
    "thamo", "chharo", "ano", "pathao", "hisabe", "chesta", "dhoro_bhul", "shesh", "felo", "ei"
  ]);

  // Built-in functions
  const builtins = new Set([
    "dekho", "nao", "dhoron", "lipi", "sonkha", "dorghyo", "dhokao", "berKoro",
    "kato", "ulto", "ache", "saja", "boroHater", "chotoHater", "bhag", "joro",
    "chhanto", "khojo", "angsho", "bodlo", "borgomul", "ghat", "niche", "upore",
    "kache", "niratek", "choto", "boro", "chabi", "poro", "lekho", "somoy",
    "lotto", "ghum", "bondho", "server_chalu", "anun", "uttor", "json_uttor",
    "json_poro", "json_banao"
  ]);

  const result: string[] = [];
  let i = 0;

  while (i < code.length) {
    // Comments
    if (code[i] === '/' && code[i + 1] === '/') {
      let comment = '';
      while (i < code.length && code[i] !== '\n') {
        comment += code[i] === '<' ? '&lt;' : code[i] === '>' ? '&gt;' : code[i] === '&' ? '&amp;' : code[i];
        i++;
      }
      result.push(`<span class="text-gray-500">${comment}</span>`);
      continue;
    }

    // Strings
    if (code[i] === '"' || code[i] === "'") {
      const quote = code[i];
      let str = quote;
      i++;
      while (i < code.length && code[i] !== quote) {
        if (code[i] === '\\' && i + 1 < code.length) {
          str += code[i] + code[i + 1];
          i += 2;
        } else {
          str += code[i] === '<' ? '&lt;' : code[i] === '>' ? '&gt;' : code[i] === '&' ? '&amp;' : code[i];
          i++;
        }
      }
      if (i < code.length) {
        str += quote;
        i++;
      }
      result.push(`<span class="text-green-400">${str}</span>`);
      continue;
    }

    // Numbers
    if (/[0-9]/.test(code[i])) {
      let num = '';
      while (i < code.length && /[0-9.]/.test(code[i])) {
        num += code[i];
        i++;
      }
      result.push(`<span class="text-orange-400">${num}</span>`);
      continue;
    }

    // Identifiers (keywords, builtins, variables)
    if (/[a-zA-Z_]/.test(code[i])) {
      let ident = '';
      while (i < code.length && /[a-zA-Z0-9_]/.test(code[i])) {
        ident += code[i];
        i++;
      }
      
      // Check if followed by ( for function call
      let j = i;
      while (j < code.length && /\s/.test(code[j])) j++;
      const isCall = code[j] === '(';

      if (keywords.has(ident)) {
        result.push(`<span class="text-purple-400 font-semibold">${ident}</span>`);
      } else if (builtins.has(ident) && isCall) {
        result.push(`<span class="text-blue-400">${ident}</span>`);
      } else {
        result.push(ident);
      }
      continue;
    }

    // Operators
    const twoChar = code.slice(i, i + 2);
    const operators = ['+=', '-=', '*=', '/=', '==', '!=', '<=', '>=', '&&', '||'];
    if (operators.includes(twoChar)) {
      const escaped = twoChar.replace(/</g, '&lt;').replace(/>/g, '&gt;');
      result.push(`<span class="text-yellow-400">${escaped}</span>`);
      i += 2;
      continue;
    }

    const singleOps = ['+', '-', '*', '/', '%', '<', '>', '=', '!'];
    if (singleOps.includes(code[i])) {
      const escaped = code[i] === '<' ? '&lt;' : code[i] === '>' ? '&gt;' : code[i];
      result.push(`<span class="text-yellow-400">${escaped}</span>`);
      i++;
      continue;
    }

    // Everything else (whitespace, punctuation, etc.)
    if (code[i] === '<') {
      result.push('&lt;');
    } else if (code[i] === '>') {
      result.push('&gt;');
    } else if (code[i] === '&') {
      result.push('&amp;');
    } else {
      result.push(code[i]);
    }
    i++;
  }

  return result.join('');
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
