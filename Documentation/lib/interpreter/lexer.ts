// BanglaCode Lexer - Tokenizes source code

export enum TokenType {
  // Literals
  NUMBER = "NUMBER",
  STRING = "STRING",
  IDENTIFIER = "IDENTIFIER",

  // Keywords
  DHORO = "DHORO",       // let/var
  JODI = "JODI",         // if
  NAHOLE = "NAHOLE",     // else
  JOTOKKHON = "JOTOKKHON", // while
  GHURIYE = "GHURIYE",   // for
  KAJ = "KAJ",           // function
  FERAO = "FERAO",       // return
  SRENI = "SRENI",       // class
  SHURU = "SHURU",       // constructor
  NOTUN = "NOTUN",       // new
  SOTTI = "SOTTI",       // true
  MITTHA = "MITTHA",     // false
  KHALI = "KHALI",       // null
  EBONG = "EBONG",       // and
  BA = "BA",             // or
  NA = "NA",             // not
  THAMO = "THAMO",       // break
  CHHARO = "CHHARO",     // continue
  ANO = "ANO",           // import
  PATHAO = "PATHAO",     // export
  HISABE = "HISABE",     // as
  CHESTA = "CHESTA",     // try
  DHORO_BHUL = "DHORO_BHUL", // catch
  SHESH = "SHESH",       // finally
  FELO = "FELO",         // throw
  EI = "EI",             // this

  // Operators
  PLUS = "PLUS",
  MINUS = "MINUS",
  ASTERISK = "ASTERISK",
  SLASH = "SLASH",
  PERCENT = "PERCENT",
  POWER = "POWER",

  ASSIGN = "ASSIGN",
  PLUS_ASSIGN = "PLUS_ASSIGN",
  MINUS_ASSIGN = "MINUS_ASSIGN",
  ASTERISK_ASSIGN = "ASTERISK_ASSIGN",
  SLASH_ASSIGN = "SLASH_ASSIGN",

  EQ = "EQ",
  NOT_EQ = "NOT_EQ",
  LT = "LT",
  GT = "GT",
  LT_EQ = "LT_EQ",
  GT_EQ = "GT_EQ",

  BANG = "BANG",

  // Delimiters
  COMMA = "COMMA",
  SEMICOLON = "SEMICOLON",
  COLON = "COLON",
  DOT = "DOT",

  LPAREN = "LPAREN",
  RPAREN = "RPAREN",
  LBRACE = "LBRACE",
  RBRACE = "RBRACE",
  LBRACKET = "LBRACKET",
  RBRACKET = "RBRACKET",

  // Special
  EOF = "EOF",
  ILLEGAL = "ILLEGAL",
}

export interface Token {
  type: TokenType;
  literal: string;
  line: number;
  column: number;
}

const KEYWORDS: Record<string, TokenType> = {
  dhoro: TokenType.DHORO,
  jodi: TokenType.JODI,
  nahole: TokenType.NAHOLE,
  jotokkhon: TokenType.JOTOKKHON,
  ghuriye: TokenType.GHURIYE,
  kaj: TokenType.KAJ,
  ferao: TokenType.FERAO,
  sreni: TokenType.SRENI,
  shuru: TokenType.SHURU,
  notun: TokenType.NOTUN,
  sotti: TokenType.SOTTI,
  mittha: TokenType.MITTHA,
  khali: TokenType.KHALI,
  ebong: TokenType.EBONG,
  ba: TokenType.BA,
  na: TokenType.NA,
  thamo: TokenType.THAMO,
  chharo: TokenType.CHHARO,
  ano: TokenType.ANO,
  pathao: TokenType.PATHAO,
  hisabe: TokenType.HISABE,
  chesta: TokenType.CHESTA,
  dhoro_bhul: TokenType.DHORO_BHUL,
  shesh: TokenType.SHESH,
  felo: TokenType.FELO,
  ei: TokenType.EI,
};

export class Lexer {
  private input: string;
  private position: number = 0;
  private readPosition: number = 0;
  private ch: string = "";
  private line: number = 1;
  private column: number = 0;

  constructor(input: string) {
    this.input = input;
    this.readChar();
  }

  private readChar(): void {
    if (this.readPosition >= this.input.length) {
      this.ch = "\0";
    } else {
      this.ch = this.input[this.readPosition];
    }
    this.position = this.readPosition;
    this.readPosition++;
    this.column++;

    if (this.ch === "\n") {
      this.line++;
      this.column = 0;
    }
  }

  private peekChar(): string {
    if (this.readPosition >= this.input.length) {
      return "\0";
    }
    return this.input[this.readPosition];
  }

  nextToken(): Token {
    this.skipWhitespace();
    this.skipComment();
    this.skipWhitespace();

    const line = this.line;
    const column = this.column;

    let token: Token;

    switch (this.ch) {
      case "=":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.EQ, literal: "==", line, column };
        } else {
          token = { type: TokenType.ASSIGN, literal: "=", line, column };
        }
        break;
      case "+":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.PLUS_ASSIGN, literal: "+=", line, column };
        } else {
          token = { type: TokenType.PLUS, literal: "+", line, column };
        }
        break;
      case "-":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.MINUS_ASSIGN, literal: "-=", line, column };
        } else {
          token = { type: TokenType.MINUS, literal: "-", line, column };
        }
        break;
      case "*":
        if (this.peekChar() === "*") {
          this.readChar();
          token = { type: TokenType.POWER, literal: "**", line, column };
        } else if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.ASTERISK_ASSIGN, literal: "*=", line, column };
        } else {
          token = { type: TokenType.ASTERISK, literal: "*", line, column };
        }
        break;
      case "/":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.SLASH_ASSIGN, literal: "/=", line, column };
        } else {
          token = { type: TokenType.SLASH, literal: "/", line, column };
        }
        break;
      case "%":
        token = { type: TokenType.PERCENT, literal: "%", line, column };
        break;
      case "!":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.NOT_EQ, literal: "!=", line, column };
        } else {
          token = { type: TokenType.BANG, literal: "!", line, column };
        }
        break;
      case "<":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.LT_EQ, literal: "<=", line, column };
        } else {
          token = { type: TokenType.LT, literal: "<", line, column };
        }
        break;
      case ">":
        if (this.peekChar() === "=") {
          this.readChar();
          token = { type: TokenType.GT_EQ, literal: ">=", line, column };
        } else {
          token = { type: TokenType.GT, literal: ">", line, column };
        }
        break;
      case ",":
        token = { type: TokenType.COMMA, literal: ",", line, column };
        break;
      case ";":
        token = { type: TokenType.SEMICOLON, literal: ";", line, column };
        break;
      case ":":
        token = { type: TokenType.COLON, literal: ":", line, column };
        break;
      case ".":
        token = { type: TokenType.DOT, literal: ".", line, column };
        break;
      case "(":
        token = { type: TokenType.LPAREN, literal: "(", line, column };
        break;
      case ")":
        token = { type: TokenType.RPAREN, literal: ")", line, column };
        break;
      case "{":
        token = { type: TokenType.LBRACE, literal: "{", line, column };
        break;
      case "}":
        token = { type: TokenType.RBRACE, literal: "}", line, column };
        break;
      case "[":
        token = { type: TokenType.LBRACKET, literal: "[", line, column };
        break;
      case "]":
        token = { type: TokenType.RBRACKET, literal: "]", line, column };
        break;
      case '"':
      case "'":
        token = { type: TokenType.STRING, literal: this.readString(), line, column };
        break;
      case "\0":
        token = { type: TokenType.EOF, literal: "", line, column };
        break;
      default:
        if (this.isLetter(this.ch)) {
          const literal = this.readIdentifier();
          const type = KEYWORDS[literal] || TokenType.IDENTIFIER;
          return { type, literal, line, column };
        } else if (this.isDigit(this.ch)) {
          return { type: TokenType.NUMBER, literal: this.readNumber(), line, column };
        } else {
          token = { type: TokenType.ILLEGAL, literal: this.ch, line, column };
        }
    }

    this.readChar();
    return token;
  }

  private skipWhitespace(): void {
    while (this.ch === " " || this.ch === "\t" || this.ch === "\n" || this.ch === "\r") {
      this.readChar();
    }
  }

  private skipComment(): void {
    if (this.ch === "/" && this.peekChar() === "/") {
      while (this.ch !== "\n" && this.ch !== "\0") {
        this.readChar();
      }
    }
  }

  private readIdentifier(): string {
    const startPos = this.position;
    while (this.isLetter(this.ch) || this.isDigit(this.ch) || this.ch === "_") {
      this.readChar();
    }
    return this.input.slice(startPos, this.position);
  }

  private readNumber(): string {
    const startPos = this.position;
    while (this.isDigit(this.ch)) {
      this.readChar();
    }
    if (this.ch === "." && this.isDigit(this.peekChar())) {
      this.readChar();
      while (this.isDigit(this.ch)) {
        this.readChar();
      }
    }
    return this.input.slice(startPos, this.position);
  }

  private readString(): string {
    const quote = this.ch;
    this.readChar();
    const startPos = this.position;
    while (this.ch !== quote && this.ch !== "\0") {
      if (this.ch === "\\") {
        this.readChar();
      }
      this.readChar();
    }
    const str = this.input.slice(startPos, this.position);
    return str;
  }

  private isLetter(ch: string): boolean {
    return /[a-zA-Z_]/.test(ch);
  }

  private isDigit(ch: string): boolean {
    return /[0-9]/.test(ch);
  }
}
