const vscode = require('vscode');

/**
 * BanglaCode Language Extension
 * Provides IntelliSense and code completion for BanglaCode
 * Created by Ankan from West Bengal, India
 */

// All BanglaCode keywords
const keywords = [
    { label: 'dhoro', kind: vscode.CompletionItemKind.Keyword, detail: 'Variable declaration (ধরো)', insertText: 'dhoro ${1:name} = ${2:value};' },
    { label: 'jodi', kind: vscode.CompletionItemKind.Keyword, detail: 'If condition (যদি)', insertText: 'jodi (${1:condition}) {\n\t${2}\n}' },
    { label: 'nahole', kind: vscode.CompletionItemKind.Keyword, detail: 'Else (নাহলে)', insertText: 'nahole {\n\t${1}\n}' },
    { label: 'jotokkhon', kind: vscode.CompletionItemKind.Keyword, detail: 'While loop (যতক্ষণ)', insertText: 'jotokkhon (${1:condition}) {\n\t${2}\n}' },
    { label: 'ghuriye', kind: vscode.CompletionItemKind.Keyword, detail: 'For loop (ঘুরিয়ে)', insertText: 'ghuriye (dhoro ${1:i} = 0; ${1:i} < ${2:10}; ${1:i} = ${1:i} + 1) {\n\t${3}\n}' },
    { label: 'kaj', kind: vscode.CompletionItemKind.Keyword, detail: 'Function (কাজ)', insertText: 'kaj ${1:name}(${2:params}) {\n\t${3}\n}' },
    { label: 'ferao', kind: vscode.CompletionItemKind.Keyword, detail: 'Return (ফেরাও)', insertText: 'ferao ${1:value};' },
    { label: 'class', kind: vscode.CompletionItemKind.Keyword, detail: 'Class definition (ক্লাস)', insertText: 'class ${1:Name} {\n\tkaj init(${2:params}) {\n\t\t${3}\n\t}\n}' },
    { label: 'notun', kind: vscode.CompletionItemKind.Keyword, detail: 'New instance (নতুন)', insertText: 'notun ${1:ClassName}(${2:args})' },
    { label: 'sotti', kind: vscode.CompletionItemKind.Constant, detail: 'True (সত্যি)', insertText: 'sotti' },
    { label: 'mittha', kind: vscode.CompletionItemKind.Constant, detail: 'False (মিথ্যা)', insertText: 'mittha' },
    { label: 'khali', kind: vscode.CompletionItemKind.Constant, detail: 'Null (খালি)', insertText: 'khali' },
    { label: 'ebong', kind: vscode.CompletionItemKind.Operator, detail: 'Logical AND (এবং)', insertText: 'ebong' },
    { label: 'ba', kind: vscode.CompletionItemKind.Operator, detail: 'Logical OR (বা)', insertText: 'ba' },
    { label: 'na', kind: vscode.CompletionItemKind.Operator, detail: 'Logical NOT (না)', insertText: 'na' },
    { label: 'thamo', kind: vscode.CompletionItemKind.Keyword, detail: 'Break (থামো)', insertText: 'thamo;' },
    { label: 'chharo', kind: vscode.CompletionItemKind.Keyword, detail: 'Continue (ছাড়ো)', insertText: 'chharo;' },
    { label: 'ei', kind: vscode.CompletionItemKind.Keyword, detail: 'This reference (এই)', insertText: 'ei.' },
    { label: 'ano', kind: vscode.CompletionItemKind.Keyword, detail: 'Import module (আনো)', insertText: 'ano "${1:module.bang}";' },
    { label: 'pathao', kind: vscode.CompletionItemKind.Keyword, detail: 'Export (পাঠাও)', insertText: 'pathao ' },
    { label: 'chesta', kind: vscode.CompletionItemKind.Keyword, detail: 'Try block (চেষ্টা)', insertText: 'chesta {\n\t${1}\n} dhoro_bhul (${2:err}) {\n\t${3}\n}' },
    { label: 'dhoro_bhul', kind: vscode.CompletionItemKind.Keyword, detail: 'Catch block (ধরো ভুল)', insertText: 'dhoro_bhul (${1:err}) {\n\t${2}\n}' },
    { label: 'shesh', kind: vscode.CompletionItemKind.Keyword, detail: 'Finally block (শেষ)', insertText: 'shesh {\n\t${1}\n}' },
    { label: 'felo', kind: vscode.CompletionItemKind.Keyword, detail: 'Throw error (ফেলো)', insertText: 'felo "${1:error message}";' },
];

// Built-in functions
const builtinFunctions = [
    // Output
    { label: 'dekho', kind: vscode.CompletionItemKind.Function, detail: 'দেখো - Print values', insertText: 'dekho(${1:value});', documentation: 'Prints values to the console' },
    
    // Type functions
    { label: 'dhoron', kind: vscode.CompletionItemKind.Function, detail: 'ধরন - Get type', insertText: 'dhoron(${1:value})', documentation: 'Returns the type of a value' },
    { label: 'lipi', kind: vscode.CompletionItemKind.Function, detail: 'লিপি - To string', insertText: 'lipi(${1:value})', documentation: 'Converts value to string' },
    { label: 'sonkha', kind: vscode.CompletionItemKind.Function, detail: 'সংখ্যা - To number', insertText: 'sonkha(${1:value})', documentation: 'Converts value to number' },
    { label: 'dorghyo', kind: vscode.CompletionItemKind.Function, detail: 'দৈর্ঘ্য - Get length', insertText: 'dorghyo(${1:value})', documentation: 'Returns the length of string or array' },
    
    // String functions
    { label: 'boroHater', kind: vscode.CompletionItemKind.Function, detail: 'বড় হাতের - Uppercase', insertText: 'boroHater(${1:str})', documentation: 'Converts string to uppercase' },
    { label: 'chotoHater', kind: vscode.CompletionItemKind.Function, detail: 'ছোট হাতের - Lowercase', insertText: 'chotoHater(${1:str})', documentation: 'Converts string to lowercase' },
    { label: 'chhanto', kind: vscode.CompletionItemKind.Function, detail: 'ছাঁটো - Trim', insertText: 'chhanto(${1:str})', documentation: 'Removes leading and trailing whitespace' },
    { label: 'bhag', kind: vscode.CompletionItemKind.Function, detail: 'ভাগ - Split string', insertText: 'bhag(${1:str}, ${2:separator})', documentation: 'Splits string into array' },
    { label: 'joro', kind: vscode.CompletionItemKind.Function, detail: 'জোড়ো - Join array', insertText: 'joro(${1:arr}, ${2:separator})', documentation: 'Joins array into string' },
    { label: 'khojo', kind: vscode.CompletionItemKind.Function, detail: 'খোঁজো - Index of', insertText: 'khojo(${1:str}, ${2:substr})', documentation: 'Finds index of substring (-1 if not found)' },
    { label: 'angsho', kind: vscode.CompletionItemKind.Function, detail: 'অংশ - Substring', insertText: 'angsho(${1:str}, ${2:start}, ${3:end})', documentation: 'Extracts substring from start to end' },
    { label: 'bodlo', kind: vscode.CompletionItemKind.Function, detail: 'বদলো - Replace', insertText: 'bodlo(${1:str}, ${2:old}, ${3:new})', documentation: 'Replaces all occurrences' },
    
    // Array functions
    { label: 'dhokao', kind: vscode.CompletionItemKind.Function, detail: 'ঢোকাও - Push to array', insertText: 'dhokao(${1:arr}, ${2:value})', documentation: 'Adds element to array' },
    { label: 'berKoro', kind: vscode.CompletionItemKind.Function, detail: 'বের করো - Pop from array', insertText: 'berKoro(${1:arr})', documentation: 'Removes and returns last element' },
    { label: 'kato', kind: vscode.CompletionItemKind.Function, detail: 'কাটো - Slice', insertText: 'kato(${1:arr}, ${2:start}, ${3:end})', documentation: 'Extracts portion of array' },
    { label: 'ulto', kind: vscode.CompletionItemKind.Function, detail: 'উল্টো - Reverse', insertText: 'ulto(${1:arr})', documentation: 'Reverses array (returns new array)' },
    { label: 'saja', kind: vscode.CompletionItemKind.Function, detail: 'সাজা - Sort', insertText: 'saja(${1:arr})', documentation: 'Sorts array (returns new array)' },
    { label: 'ache', kind: vscode.CompletionItemKind.Function, detail: 'আছে - Includes', insertText: 'ache(${1:arr}, ${2:value})', documentation: 'Checks if value exists in array' },
    { label: 'chabi', kind: vscode.CompletionItemKind.Function, detail: 'চাবি - Get keys', insertText: 'chabi(${1:map})', documentation: 'Returns array of map keys' },
    
    // Math functions
    { label: 'borgomul', kind: vscode.CompletionItemKind.Function, detail: 'বর্গমূল - Square root', insertText: 'borgomul(${1:x})', documentation: 'Returns square root' },
    { label: 'ghat', kind: vscode.CompletionItemKind.Function, detail: 'ঘাত - Power', insertText: 'ghat(${1:base}, ${2:exp})', documentation: 'Returns base raised to exponent' },
    { label: 'niche', kind: vscode.CompletionItemKind.Function, detail: 'নিচে - Floor', insertText: 'niche(${1:x})', documentation: 'Rounds down to integer' },
    { label: 'upore', kind: vscode.CompletionItemKind.Function, detail: 'উপরে - Ceiling', insertText: 'upore(${1:x})', documentation: 'Rounds up to integer' },
    { label: 'kache', kind: vscode.CompletionItemKind.Function, detail: 'কাছে - Round', insertText: 'kache(${1:x})', documentation: 'Rounds to nearest integer' },
    { label: 'niratek', kind: vscode.CompletionItemKind.Function, detail: 'নিরপেক্ষ - Absolute', insertText: 'niratek(${1:x})', documentation: 'Returns absolute value' },
    { label: 'choto', kind: vscode.CompletionItemKind.Function, detail: 'ছোট - Minimum', insertText: 'choto(${1:a}, ${2:b})', documentation: 'Returns minimum value' },
    { label: 'boro', kind: vscode.CompletionItemKind.Function, detail: 'বড় - Maximum', insertText: 'boro(${1:a}, ${2:b})', documentation: 'Returns maximum value' },
    { label: 'lotto', kind: vscode.CompletionItemKind.Function, detail: 'লটো - Random', insertText: 'lotto()', documentation: 'Returns random number between 0 and 1' },
    
    // Utility functions
    { label: 'somoy', kind: vscode.CompletionItemKind.Function, detail: 'সময় - Current time', insertText: 'somoy()', documentation: 'Returns current timestamp in milliseconds' },
    { label: 'ghum', kind: vscode.CompletionItemKind.Function, detail: 'ঘুম - Sleep', insertText: 'ghum(${1:milliseconds})', documentation: 'Pauses execution for milliseconds' },
    { label: 'nao', kind: vscode.CompletionItemKind.Function, detail: 'নাও - Input', insertText: 'nao(${1:"Enter value: "})', documentation: 'Reads user input from console' },
    { label: 'bondho', kind: vscode.CompletionItemKind.Function, detail: 'বন্ধ - Exit', insertText: 'bondho(${1:0})', documentation: 'Exits program with code' },
    
    // File functions
    { label: 'poro', kind: vscode.CompletionItemKind.Function, detail: 'পড়ো - Read file', insertText: 'poro(${1:"filename"})', documentation: 'Reads file contents as string' },
    { label: 'lekho', kind: vscode.CompletionItemKind.Function, detail: 'লেখো - Write file', insertText: 'lekho(${1:"filename"}, ${2:content})', documentation: 'Writes content to file' },
    
    // HTTP functions
    { label: 'server_chalu', kind: vscode.CompletionItemKind.Function, detail: 'সার্ভার চালু - Start HTTP server', insertText: 'server_chalu(${1:3000}, ${2:handler})', documentation: 'Starts HTTP server on port' },
    { label: 'anun', kind: vscode.CompletionItemKind.Function, detail: 'আনুন - HTTP GET', insertText: 'anun(${1:"url"})', documentation: 'Makes HTTP GET request' },
];

function activate(context) {
    console.log('BanglaCode extension activated!');

    // Register completion provider
    const completionProvider = vscode.languages.registerCompletionItemProvider(
        'banglacode',
        {
            provideCompletionItems(document, position, token, context) {
                const completions = [];
                
                // Add keywords
                for (const kw of keywords) {
                    const item = new vscode.CompletionItem(kw.label, kw.kind);
                    item.detail = kw.detail;
                    item.insertText = new vscode.SnippetString(kw.insertText);
                    completions.push(item);
                }
                
                // Add built-in functions
                for (const fn of builtinFunctions) {
                    const item = new vscode.CompletionItem(fn.label, fn.kind);
                    item.detail = fn.detail;
                    item.insertText = new vscode.SnippetString(fn.insertText);
                    if (fn.documentation) {
                        item.documentation = new vscode.MarkdownString(fn.documentation);
                    }
                    completions.push(item);
                }
                
                // Extract variables from current document
                const text = document.getText();
                const varRegex = /dhoro\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=/g;
                let match;
                const seenVars = new Set();
                while ((match = varRegex.exec(text)) !== null) {
                    const varName = match[1];
                    if (!seenVars.has(varName)) {
                        seenVars.add(varName);
                        const item = new vscode.CompletionItem(varName, vscode.CompletionItemKind.Variable);
                        item.detail = 'Variable';
                        completions.push(item);
                    }
                }
                
                // Extract functions from current document
                const funcRegex = /kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(/g;
                const seenFuncs = new Set();
                while ((match = funcRegex.exec(text)) !== null) {
                    const funcName = match[1];
                    if (!seenFuncs.has(funcName) && funcName !== 'init') {
                        seenFuncs.add(funcName);
                        const item = new vscode.CompletionItem(funcName, vscode.CompletionItemKind.Function);
                        item.detail = 'User function';
                        item.insertText = new vscode.SnippetString(funcName + '(${1})');
                        completions.push(item);
                    }
                }
                
                // Extract classes from current document
                const classRegex = /class\s+([a-zA-Z_][a-zA-Z0-9_]*)/g;
                const seenClasses = new Set();
                while ((match = classRegex.exec(text)) !== null) {
                    const className = match[1];
                    if (!seenClasses.has(className)) {
                        seenClasses.add(className);
                        const item = new vscode.CompletionItem(className, vscode.CompletionItemKind.Class);
                        item.detail = 'User class';
                        completions.push(item);
                    }
                }
                
                return completions;
            }
        },
        '' // Trigger on any character
    );

    // Register hover provider
    const hoverProvider = vscode.languages.registerHoverProvider('banglacode', {
        provideHover(document, position, token) {
            const range = document.getWordRangeAtPosition(position);
            if (!range) return null;
            
            const word = document.getText(range);
            
            // Check keywords
            const kw = keywords.find(k => k.label === word);
            if (kw) {
                return new vscode.Hover([
                    `**${kw.label}** - ${kw.detail}`,
                    'BanglaCode keyword'
                ]);
            }
            
            // Check built-in functions
            const fn = builtinFunctions.find(f => f.label === word);
            if (fn) {
                return new vscode.Hover([
                    `**${fn.label}** - ${fn.detail}`,
                    fn.documentation || '',
                    'BanglaCode built-in function'
                ]);
            }
            
            return null;
        }
    });

    // Register signature help provider
    const signatureProvider = vscode.languages.registerSignatureHelpProvider(
        'banglacode',
        {
            provideSignatureHelp(document, position, token, context) {
                const line = document.lineAt(position.line).text;
                const beforeCursor = line.substring(0, position.character);
                
                // Find function name
                const match = beforeCursor.match(/([a-zA-Z_][a-zA-Z0-9_]*)\s*\([^)]*$/);
                if (!match) return null;
                
                const funcName = match[1];
                const fn = builtinFunctions.find(f => f.label === funcName);
                if (!fn) return null;
                
                const sigHelp = new vscode.SignatureHelp();
                const sig = new vscode.SignatureInformation(fn.insertText.replace(/\$\{\d+:?/g, '').replace(/\}/g, ''), fn.documentation);
                sigHelp.signatures = [sig];
                sigHelp.activeSignature = 0;
                
                return sigHelp;
            }
        },
        '(', ','
    );

    context.subscriptions.push(completionProvider, hoverProvider, signatureProvider);
}

function deactivate() {}

module.exports = { activate, deactivate };
