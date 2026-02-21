# BanglaCode Feature List - Implemented Features

**Last Updated**: February 2026 (v7.0.8 - Verified Batch 4)

> Verification note: Parts of older Phase 2 entries were out of sync with the codebase.
> The items listed under "v7.0.5 - Verified Batch 1" below are now code-verified and tested.

This document lists all features that are **currently implemented** in BanglaCode, organized by category and implementation phase.

---

## Table of Contents

1. [v7.0.8 - Verified Batch 4](#v708---verified-batch-4)
1. [v7.0.7 - Verified Batch 3](#v707---verified-batch-3)
1. [v7.0.6 - Verified Batch 2](#v706---verified-batch-2)
1. [v7.0.5 - Verified Batch 1](#v705---verified-batch-1)
1. [Phase 1 Features (v7.0.3)](#phase-1-features-v703)
2. [Core Language Features](#core-language-features)
3. [Data Types & Literals](#data-types--literals)
4. [Operators](#operators)
5. [Control Flow](#control-flow)
6. [Functions](#functions)
7. [Array Methods](#array-methods)
8. [String Methods](#string-methods)
9. [Object Methods](#object-methods)
10. [Built-in Functions](#built-in-functions)
11. [OOP Features](#oop-features)
12. [Async/Await](#asyncawait)
13. [Module System](#module-system)
14. [Error Handling](#error-handling)
15. [I/O Operations](#io-operations)
16. [Networking](#networking)
17. [Database Operations](#database-operations)

---

## v7.0.8 - Verified Batch 4

### ✅ COMPLETED: Maturity Pack Extensions

| Feature | BanglaCode | Status |
|---------|------------|--------|
| Destructuring (array) | `dhoro [a, b] = arr` | ✅ DONE |
| Destructuring (object) | `dhoro {x, y} = obj` | ✅ DONE |
| Multi-parameter arrows | `(a, b) => ...`, `() => ...` | ✅ DONE |
| Timers | `setTimeout`, `setInterval`, `clearTimeout`, `clearInterval` | ✅ DONE |
| Regex wrappers + flags | `match`, `matchAll`, `search`, optional flags in `regex_*` | ✅ DONE |

---

## v7.0.7 - Verified Batch 3

### ✅ COMPLETED: Maturity Syntax + APIs

| Feature | BanglaCode | Status |
|---------|------------|--------|
| Arrow functions | `x => expr`, `x => { ... }` | ✅ DONE |
| for...of loops | `ghuriye (x of iterable) { ... }` | ✅ DONE |
| for...in loops | `ghuriye (k in target) { ... }` | ✅ DONE |
| Date core | `tarikh_ekhon`, `tarikh_parse`, `tarikh_format` | ✅ DONE |
| RegExp core | `regex_test`, `regex_match`, `regex_match_all`, `regex_search`, `regex_replace` | ✅ DONE |
| Object utilities | `nijer_ache`, `jora_theke`, `ekoi_ki`, `notun_map`, `joma` | ✅ DONE (freeze semantic partial) |

---

## v7.0.6 - Verified Batch 2

### ✅ COMPLETED: Core Syntax Operators + Loop

| Feature | Syntax | JS Equivalent | Status |
|---------|--------|---------------|--------|
| do...while | `do { ... } jotokkhon (cond);` | `do { ... } while (cond)` | ✅ DONE |
| in operator | `"key" in obj` | `'key' in obj` | ✅ DONE |
| instanceof operator | `obj instanceof Class` | `obj instanceof Class` | ✅ DONE |
| delete operator | `delete obj.key` | `delete obj.key` | ✅ DONE |

### ✅ COMPLETED: Additional Utility Parity (Batch 2B)

| Feature | BanglaCode | JS Equivalent | Status |
|---------|------------|---------------|--------|
| reduceRight | `sonkuchito_dan()` | `arr.reduceRight()` | ✅ DONE |
| concat | `joro_array()` | `arr.concat()` | ✅ DONE |
| flat | `somtol()` | `arr.flat()` | ✅ DONE |
| codePointAt | `codepoint_at()` | `str.codePointAt()` | ✅ DONE |
| localeCompare | `tulona_text()` | `str.localeCompare()` | ✅ DONE |
| normalize | `shadharon_text()` | `str.normalize()` | ✅ DONE |

---

## v7.0.5 - Verified Batch 1

### ✅ COMPLETED: Additional Array Methods

| Function | Bengali | JS Equivalent | Status |
|----------|---------|---------------|--------|
| find | `khojo_prothom()` | `arr.find()` | ✅ DONE |
| findIndex | `khojo_index()` | `arr.findIndex()` | ✅ DONE |
| findLast | `khojo_shesh()` | `arr.findLast()` | ✅ DONE |
| findLastIndex | `khojo_shesh_index()` | `arr.findLastIndex()` | ✅ DONE |
| every | `prottek()` | `arr.every()` | ✅ DONE |
| some | `kono()` | `arr.some()` | ✅ DONE |
| flatMap | `somtol_manchitro()` | `arr.flatMap()` | ✅ DONE |
| at | `array_at()` | `arr.at()` | ✅ DONE |
| lastIndexOf | `shesh_index_of()` | `arr.lastIndexOf()` | ✅ DONE |

### ✅ COMPLETED: Additional String Methods

| Function | Bengali | JS Equivalent | Status |
|----------|---------|---------------|--------|
| includes | `ache_text()` | `str.includes()` | ✅ DONE |
| startsWith | `shuru_diye()` | `str.startsWith()` | ✅ DONE |
| endsWith | `shesh_diye()` | `str.endsWith()` | ✅ DONE |
| repeat | `baro()` | `str.repeat()` | ✅ DONE |
| padStart | `agey_bhoro()` | `str.padStart()` | ✅ DONE |
| padEnd | `pichoney_bhoro()` | `str.padEnd()` | ✅ DONE |
| charAt | `okkhor()` | `str.charAt()` | ✅ DONE |
| at | `text_at()` | `str.at()` | ✅ DONE |
| charCodeAt | `okkhor_code()` | `str.charCodeAt()` | ✅ DONE |
| trimStart | `chhanto_shuru()` | `str.trimStart()` | ✅ DONE |
| trimEnd | `chhanto_shesh()` | `str.trimEnd()` | ✅ DONE |
| lastIndexOf | `shesh_khojo()` | `str.lastIndexOf()` | ✅ DONE |

### ✅ COMPLETED: Numeric + URI Globals

| Function | Bengali | JS Equivalent | Status |
|----------|---------|---------------|--------|
| parseInt | `purno_sonkhya()` | `parseInt()` | ✅ DONE |
| parseFloat | `doshomik_sonkhya()` | `parseFloat()` | ✅ DONE |
| isNaN | `sonkhya_na()` | `isNaN()` | ✅ DONE |
| isFinite | `sonkhya_shimito()` | `isFinite()` | ✅ DONE |
| encodeURI | `uri_encode()` | `encodeURI()` | ✅ DONE |
| decodeURI | `uri_decode()` | `decodeURI()` | ✅ DONE |
| encodeURIComponent | `uri_ongsho_encode()` | `encodeURIComponent()` | ✅ DONE |
| decodeURIComponent | `uri_ongsho_decode()` | `decodeURIComponent()` | ✅ DONE |

---

## Phase 1 Features (v7.0.3)

### ✅ COMPLETED: Array Methods (4 new methods)

**Bengali Names & Descriptions:**

| Method | Bengali | Purpose | Example | Status |
|--------|---------|---------|---------|--------|
| **map()** | `manchitro()` | Transform each element | `manchitro(arr, kaj(x) { ferao x * 2; })` | ✅ DONE |
| **filter()** | `chhanno()` | Filter elements by condition | `chhanno(arr, kaj(x) { ferao x > 5; })` | ✅ DONE |
| **reduce()** | `sonkuchito()` | Reduce to single value | `sonkuchito(arr, kaj(a,b) { ferao a+b; })` | ✅ DONE |
| **forEach()** | `proti()` | Execute for each element | `proti(arr, kaj(x) { dekho(x); })` | ✅ DONE |

**Features:**
- ✅ Callback receives `(element, index, array)` parameters
- ✅ Performance optimized with pre-allocated arrays
- ✅ Error propagation from callbacks
- ✅ Works with nested arrays and objects
- ✅ Comprehensive test coverage (25+ tests)

---

### ✅ COMPLETED: Object Methods (3 new methods)

**Bengali Names & Descriptions:**

| Method | Bengali | Purpose | Example | Status |
|--------|---------|---------|---------|--------|
| **values()** | `maan()` | Extract object values | `maan(obj)` → array of values | ✅ DONE |
| **entries()** | `jora()` | Extract key-value pairs | `jora(obj)` → array of [key, value] pairs | ✅ DONE |
| **assign()** | `mishra()` | Merge objects (in-place) | `mishra(target, source1, source2)` | ✅ DONE |

**Features:**
- ✅ `maan()` - Returns array of all object values
- ✅ `jora()` - Returns array of [key, value] pairs
- ✅ `mishra()` - Merges objects (mutates target, supports multiple sources)
- ✅ Proper error handling for non-object arguments
- ✅ Comprehensive test coverage (20+ tests)

---

### ✅ COMPLETED: Switch/Case Control Flow

**Bengali Keywords:**

| Keyword | Bengali | Meaning | English Equivalent |
|---------|---------|---------|-------------------|
| **switch** | `bikolpo` | বিকল্প (alternative) | switch |
| **case** | `khetre` | ক্ষেত্রে (in case of) | case |
| **default** | `manchito` | মানচিত্র (default/standard) | default |
| **break** | `thamo` | থামো (stop) | break |

**Syntax & Features:**
```bangla
bikolpo (expression) {
    khetre value1: { /* code */ }
    khetre value2: { /* code */ }
    manchito: { /* default code */ }
}
```

- ✅ Type-safe comparison using `objectsEqual()`
- ✅ Break statement (`thamo`) support
- ✅ Default case handling
- ✅ Works with numbers, strings, booleans, null
- ✅ Comprehensive test coverage (15+ tests)
- ✅ No fall-through (each case is independent)

---

### ✅ COMPLETED: Template Literals

**Syntax:**

```bangla
`Hello ${name}!`
`Result: ${5 + 3}`
`Array length: ${dorghyo(arr)}`
```

**Features:**
- ✅ Backtick syntax with `${}` interpolation
- ✅ Support for expressions inside `${}`
- ✅ Function calls within expressions
- ✅ Nested objects/arrays with balanced brace counting
- ✅ Empty template support
- ✅ Special character handling (Unicode, newlines)
- ✅ Comprehensive test coverage (18+ tests)
- ✅ Error propagation from invalid expressions

---

## Summary Statistics

### Phase 1 Implementation Summary
- ✅ **9 features** completed
- ✅ **4 array methods** (manchitro, chhanno, sonkuchito, proti)
- ✅ **3 object methods** (maan, jora, mishra)
- ✅ **1 control flow** (bikolpo/khetre/manchito with thamo)
- ✅ **1 string feature** (template literals with ${} interpolation)
- ✅ **78+ test cases** written
- ✅ **All tests passing** (291/291)
- ✅ **VS Code extension updated** with syntax highlighting and snippets
- ✅ **Documentation website updated** with examples and usage

---

## Phase 2 Features (v7.0.4)

### ✅ COMPLETED: Core Language Essentials (Phase 2A - 7 features)

| Feature | Bengali | Purpose | Syntax | Status |
|---------|---------|---------|--------|--------|
| **Ternary operator** | — | Inline conditionals | `condition ? trueVal : falseVal` | ✅ DONE |
| **Optional chaining** | — | Safe property access | `obj?.prop`, `obj?.[expr]` | ✅ DONE |
| **Nullish coalescing** | — | Default for null | `left ?? right` | ✅ DONE |
| **Array find** | `khojo_prothom()` | Find first element | `khojo_prothom(arr, kaj(x) { ferao x > 5; })` | ✅ DONE |
| **Array findIndex** | `khojo_index()` | Find first index | `khojo_index(arr, kaj(x) { ferao x > 5; })` | ✅ DONE |
| **Array every** | `prottek()` | All pass test | `prottek(arr, kaj(x) { ferao x > 0; })` | ✅ DONE |
| **Array some** | `kono()` | Any pass test | `kono(arr, kaj(x) { ferao x > 10; })` | ✅ DONE |

---

### ✅ COMPLETED: String & Array Utility Methods (Phase 2B - 14 functions)

| Function | Bengali | JS Equivalent | Status |
|----------|---------|---------------|--------|
| **String includes** | `ache_text()` | `str.includes()` | ✅ DONE |
| **String startsWith** | `shuru_diye()` | `str.startsWith()` | ✅ DONE |
| **String endsWith** | `shesh_diye()` | `str.endsWith()` | ✅ DONE |
| **String repeat** | `baro()` | `str.repeat()` | ✅ DONE |
| **String padStart** | `agey_bhoro()` | `str.padStart()` | ✅ DONE |
| **String padEnd** | `pichoney_bhoro()` | `str.padEnd()` | ✅ DONE |
| **String charAt** | `okkhor()` | `str.charAt()` | ✅ DONE |
| **String trimStart** | `chhanto_shuru()` | `str.trimStart()` | ✅ DONE |
| **String trimEnd** | `chhanto_shesh()` | `str.trimEnd()` | ✅ DONE |
| **Array concat** | `joro_array()` | `arr.concat()` | ✅ DONE |
| **Array flat** | `somtol()` | `arr.flat()` | ✅ DONE |
| **parseInt** | `purno_sonkhya()` | `parseInt()` | ✅ DONE |
| **parseFloat** | `doshomik_sonkhya()` | `parseFloat()` | ✅ DONE |
| **isNaN** | `sonkhya_na()` | `isNaN()` | ✅ DONE |

---

### ✅ COMPLETED: OOP Enhancements (Phase 2C - 3 features)

| Feature | Bengali | JS Equivalent | Syntax | Status |
|---------|---------|---------------|--------|--------|
| **Class inheritance** | `theke` | `extends` | `sreni Child theke Parent { }` | ✅ DONE |
| **Super calls** | `upor` | `super` | `upor.method()` | ✅ DONE |
| **Static methods** | `sthir kaj` | `static` | `sthir kaj method() { }` | ✅ DONE |

---

### ✅ COMPLETED: HTTP Full Methods (Phase 2D - 4 methods)

| Method | Bengali | Usage | Status |
|--------|---------|-------|--------|
| **HTTP POST** | `pathao_post()` | `pathao_post(url, body, headers)` | ✅ DONE |
| **HTTP PUT** | `pathao_put()` | `pathao_put(url, body, headers)` | ✅ DONE |
| **HTTP DELETE** | `pathao_delete()` | `pathao_delete(url)` | ✅ DONE |
| **HTTP PATCH** | `pathao_patch()` | `pathao_patch(url, body, headers)` | ✅ DONE |

---

### ✅ COMPLETED: Crypto Module (Phase 2E - 7 functions)

| Function | Bengali | Purpose | Status |
|----------|---------|---------|--------|
| **SHA-256** | `hash_sha256()` | Hash string to hex | ✅ DONE |
| **SHA-512** | `hash_sha512()` | Hash string to hex | ✅ DONE |
| **MD5** | `hash_md5()` | Hash string to hex | ✅ DONE |
| **HMAC-SHA256** | `hmac_sha256()` | Keyed hash | ✅ DONE |
| **Random bytes** | `lotto_bytes()` | Crypto-secure random | ✅ DONE |
| **Base64 encode** | `base64_encode()` | Encode to base64 | ✅ DONE |
| **Base64 decode** | `base64_decode()` | Decode from base64 | ✅ DONE |

---

### Test Coverage
- Phase 1: 78+ tests
- Phase 2: 54 new tests
- Total: **345 passing tests**

### Code Quality Metrics
- File size compliance: All files under 500 lines (ideal < 300)
- No code violations: All CLAUDE.md rules followed
- Performance: No regression in existing features
- 0 test regressions across all phases

---

## Version History

### v7.0.4 (Phase 2 Complete) ✅
- ✅ Ternary operator, optional chaining (`?.`), nullish coalescing (`??`)
- ✅ Array search methods: `khojo_prothom`, `khojo_index`, `prottek`, `kono`
- ✅ 9 string methods: `ache_text`, `shuru_diye`, `shesh_diye`, `baro`, `agey_bhoro`, `pichoney_bhoro`, `okkhor`, `chhanto_shuru`, `chhanto_shesh`
- ✅ Array utilities: `joro_array` (concat), `somtol` (flat)
- ✅ Number parsing: `purno_sonkhya`, `doshomik_sonkhya`, `sonkhya_na`
- ✅ OOP: class inheritance (`theke`), super (`upor`), static methods (`sthir kaj`)
- ✅ HTTP methods: `pathao_post`, `pathao_put`, `pathao_delete`, `pathao_patch`
- ✅ Crypto: `hash_sha256`, `hash_sha512`, `hash_md5`, `hmac_sha256`, `lotto_bytes`, `base64_encode`, `base64_decode`
- ✅ 54 new tests, 345 total, 0 regressions

### v7.0.3 (Phase 1 Complete) ✅
- ✅ Added 4 array methods (manchitro, chhanno, sonkuchito, proti)
- ✅ Added 3 object methods (maan, jora, mishra)
- ✅ Added switch/case control flow (bikolpo/khetre/manchito/thamo)
- ✅ Added template literals (backtick syntax with ${} interpolation)
- ✅ 78+ new test cases with 100% pass rate
- ✅ Updated VS Code extension with syntax highlighting and snippets
- ✅ Updated documentation website with examples for all features
- ✅ All code follows CLAUDE.md standards (file size, architecture, performance)

### Earlier Versions
- v7.0.2: Network and database features
- v7.0.1: Core language features and basic OOP support

---

## Implementation Details

### Array Methods Implementation
File: `src/evaluator/builtins/builtins_array.go` (320 lines)

**Features:**
- `manchitro()`: Transform elements with callback(element, index, array)
- `chhanno()`: Filter elements with boolean callback
- `sonkuchito()`: Reduce to single value with optional initial value
- `proti()`: Iterate and execute for each element
- Pre-allocated arrays for optimal performance
- Error propagation from callbacks

### Object Methods Implementation
File: `src/evaluator/builtins/builtins_object.go` (85 lines)

**Features:**
- `maan()`: Extract and return array of object values
- `jora()`: Extract and return array of [key, value] pairs
- `mishra()`: Merge multiple objects into target (mutates target)
- Proper error handling for non-object arguments
- Maintains insertion order

### Switch/Case Statement Implementation
Files: 
- `src/parser/statements.go`: Parse switch syntax
- `src/evaluator/evaluator.go`: Evaluate switch cases
- `src/lexer/token.go`: Token definitions

**Features:**
- `bikolpo` (switch), `khetre` (case), `manchito` (default), `thamo` (break)
- Type-safe comparison using objectsEqual()
- No fall-through between cases
- Support for all data types (numbers, strings, booleans, null)

### Template Literals Implementation
Files:
- `src/lexer/lexer.go`: Parse backtick syntax
- `src/evaluator/expressions.go`: Evaluate template expressions
- `src/lexer/token.go`: Template token definition

**Features:**
- Backtick (`) syntax for template strings
- ${expression} interpolation with balanced brace counting
- Nested objects/arrays support
- Function calls within expressions
- Error propagation from invalid expressions

---

**Document Purpose**: Complete inventory of Phase 1+2 completed features for BanglaCode v7.0.4

**Last Updated**: February 2026
**Status**: Phase 2 Implementation Complete ✅
