# BanglaCode Feature List - Implemented Features

**Last Updated**: February 2026 (v7.0.11 - Path Utilities Enhanced)

> Verification note: Parts of older Phase 2 entries were out of sync with the codebase.
> The items listed under "v7.0.5 - Verified Batch 1" below are now code-verified and tested.

This document lists all features that are **currently implemented** in BanglaCode, organized by category and implementation phase.

---

## Table of Contents

1. [v7.0.9 - Event-Driven, Streaming & URL Parsing](#v709---event-driven-streaming--url-parsing)
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

## v7.0.11 - Path Utilities Enhanced

### ✅ COMPLETED: Path Manipulation & Constants

| Feature | BanglaCode | Status |
|---------|------------|--------|
| **Resolve Absolute Path** | `path_resolve(...paths)` - Convert to absolute path | ✅ DONE |
| **Normalize Path** | `path_normalize(path)` - Clean path, resolve . and .. | ✅ DONE |
| **Relative Path** | `path_relative(base, target)` - Compute relative path | ✅ DONE |
| **Path Separator Constant** | `PATH_SEP` - Platform-specific separator (/ or \\) | ✅ DONE |
| **Path Delimiter Constant** | `PATH_DELIMITER` - Platform delimiter (: or ;) | ✅ DONE |

**Implementation:**
- 3 new functions: `path_resolve()`, `path_normalize()`, `path_relative()`
- 2 constants: `PATH_SEP`, `PATH_DELIMITER` added to global environment
- All path operations are cross-platform (Windows/Unix/Linux/macOS)
- 17 comprehensive tests (all passing, 456 total tests)
- Full VS Code extension support (syntax + 5 snippets)
- Comprehensive documentation with real-world examples

**Tests:**
- Path resolution with single and multiple paths
- Path normalization (redundant slashes, . and ..)
- Relative path calculation (subdirs, parent dirs, same dir)
- Path constants verification (platform-specific)
- Path analysis utility combining multiple operations
- Cross-platform path building

---

## v7.0.9 - Event-Driven, Streaming & URL Parsing

### ✅ COMPLETED: EventEmitter, Worker Threads, Streams, Buffer, URL Parsing

| Feature | BanglaCode | Status |
|---------|------------|--------|
| **EventEmitter** | `ghotona_srishti()` - Create event emitter | ✅ DONE |
| Event listening | `ghotona_shuno(emitter, "event", handler)` - Listen to events | ✅ DONE |
| One-time listener | `ghotona_ekbar(emitter, "event", handler)` - Listen once | ✅ DONE |
| Emit events | `ghotona_prokash(emitter, "event", data)` - Emit event | ✅ DONE |
| Remove listener | `ghotona_bondho(emitter, "event", handler)` - Remove listener | ✅ DONE |
| Remove all listeners | `ghotona_sob_bondho(emitter, "event")` - Remove all | ✅ DONE |
| Get listeners | `ghotona_shrotara(emitter, "event")` - Get all listeners | ✅ DONE |
| Get event names | `ghotona_naam_sob(emitter)` - Get event names | ✅ DONE |
| **Worker Threads** | `kaj_kormi_srishti(fn, data)` - Create worker thread | ✅ DONE |
| Post message | `kaj_kormi_pathao(worker, msg)` - Send to worker | ✅ DONE |
| Terminate worker | `kaj_kormi_bondho(worker)` - Stop worker | ✅ DONE |
| Listen to worker | `kaj_kormi_shuno(worker, handler)` - Receive from worker | ✅ DONE |
| Worker data | `kaj_kormi_tothya` - Initial data in worker | ✅ DONE |
| **Streams API** | `stream_readable_srishti()` - Create readable stream | ✅ DONE |
| Writable stream | `stream_writable_srishti(hwm?)` - Create writable stream | ✅ DONE |
| Write to stream | `stream_lekho(stream, data)` - Write data | ✅ DONE |
| Read from stream | `stream_poro(stream, size?)` - Read data | ✅ DONE |
| Close stream | `stream_bondho(stream)` - Close stream | ✅ DONE |
| End stream | `stream_shesh(stream)` - Signal end | ✅ DONE |
| Pipe streams | `stream_pipe(readable, writable)` - Connect streams | ✅ DONE |
| Stream events | `stream_on(stream, event, handler)` - Event handlers | ✅ DONE |
| Backpressure | High water mark (default 16KB) | ✅ DONE |
| **Buffer API** | `buffer_banao(size)` - Allocate buffer | ✅ DONE |
| Buffer from data | `buffer_theke(data)` - Create from string/array | ✅ DONE |
| Concat buffers | `buffer_joro(buf1, buf2, ...)` - Join buffers | ✅ DONE |
| Buffer to string | `buffer_text(buf, encoding?)` - Convert to text | ✅ DONE |
| Write to buffer | `buffer_lekho(buf, str, offset?)` - Write string | ✅ DONE |
| Buffer slice | `buffer_angsho(buf, start, end)` - Extract portion | ✅ DONE |
| Compare buffers | `buffer_tulona(buf1, buf2)` - Compare (-1, 0, 1) | ✅ DONE |
| Buffer to hex | `buffer_hex(buf)` - Convert to hex string | ✅ DONE |
| Copy buffer | `buffer_copy(target, source, offset)` - Copy data | ✅ DONE |
| **URL Parsing** | `url_parse(urlString)` - Parse URL into object | ✅ DONE |
| URL properties | Access via `url.Hostname`, `url.Port`, `url.Pathname`, etc. | ✅ DONE |
| URL components | Protocol, Username, Password, Host, Search, Hash, Origin | ✅ DONE |
| Query parameters | `url_query_params(queryOrURL)` - Create URLSearchParams | ✅ DONE |
| Get parameter | `url_query_get(params, key)` - Get parameter value | ✅ DONE |
| Set parameter | `url_query_set(params, key, value)` - Set parameter | ✅ DONE |
| Append parameter | `url_query_append(params, key, value)` - Append value | ✅ DONE |
| Delete parameter | `url_query_delete(params, key)` - Remove parameter | ✅ DONE |
| Has parameter | `url_query_has(params, key)` - Check if exists | ✅ DONE |
| Get all keys | `url_query_keys(params)` - Get all keys | ✅ DONE |
| Get all values | `url_query_values(params)` - Get all values | ✅ DONE |
| Query to string | `url_query_toString(params)` - Convert to query string | ✅ DONE |

### Phase 6: Collections (Set & Map) - v7.0.10

**Set Functions (8):**
| Feature | Function | Status |
|---------|----------|--------|
| Create Set | `set_srishti([array])` - Create new Set from optional array | ✅ DONE |
| Add element | `set_add(set, element)` - Add element to Set | ✅ DONE |
| Check exists | `set_has(set, element)` - Check if element exists | ✅ DONE |
| Delete element | `set_delete(set, element)` - Remove element from Set | ✅ DONE |
| Clear all | `set_clear(set)` - Remove all elements | ✅ DONE |
| Get size | `set_akar(set)` - Get Set size | ✅ DONE |
| Get values | `set_values(set)` - Get all values as array | ✅ DONE |
| Iterate | `set_foreach(set, callback)` - Iterate over elements | ✅ DONE |

**Map Functions (11):**
| Feature | Function | Status |
|---------|----------|--------|
| Create Map | `map_srishti([entries])` - Create new Map from optional [[k,v]] entries | ✅ DONE |
| Set entry | `map_set(map, key, value)` - Set key-value pair (any key type) | ✅ DONE |
| Get value | `map_get(map, key)` - Get value by key (returns khali if not found) | ✅ DONE |
| Check key | `map_has(map, key)` - Check if key exists | ✅ DONE |
| Delete entry | `map_delete(map, key)` - Remove key-value pair | ✅ DONE |
| Clear all | `map_clear(map)` - Remove all entries | ✅ DONE |
| Get size | `map_akar(map)` - Get Map size | ✅ DONE |
| Get keys | `map_keys(map)` - Get all keys as array | ✅ DONE |
| Get values | `map_values(map)` - Get all values as array | ✅ DONE |
| Get entries | `map_entries(map)` - Get all [key, value] pairs as array | ✅ DONE |
| Iterate | `map_foreach(map, callback)` - Iterate over entries | ✅ DONE |

**Statistics:**
- **67 new functions** implemented across 6 APIs
- **93 tests** written and passing (14 EventEmitter + 10 Worker + 14 Streams + 14 Buffer + 19 URL + 22 Collections)
- **6 documentation pages** created with comprehensive examples
- **Full VS Code extension support** (syntax highlighting + snippets)
- **Event-driven architecture** now fully supported
- **Parallel processing** enabled with Worker Threads
- **Memory-efficient streaming** for large data processing
- **Binary data handling** with Buffer API
- **URL parsing and manipulation** with URLSearchParams
- **Modern ES6 collections** with Set and Map (any type keys)

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

---

## Phase 11: OOP Enhancements (v7.0.14)

### Getters Implementation
Files:
- `src/lexer/token.go`: Added PAO token
- `src/ast/statements.go`: Extended ClassDeclaration with Getters map
- `src/parser/statements.go`: parseGetter() function
- `src/evaluator/classes.go`: Getter evaluation
- `src/evaluator/expressions_member.go`: Getter invocation on access
- `src/object/object.go`: Extended Class struct with Getters map

**Features:**
- `pao propertyName() { }`: Define getter methods (পাও - obtain/get)
- Zero parameters, must return value
- Executed automatically when property is accessed
- Full closure support with `ei` (this) binding
- Perfect for computed properties and derived data

**Example:**
```banglacode
sreni Person {
    shuru(naam, boichhor) {
        ei.naam = naam;
        ei.boichhor = boichhor;
    }

    pao boshi() {
        dhoro current = 2026;
        ferao current - ei.boichhor;
    }
}

dhoro p = notun Person("Ankan", 1995);
dekho(p.boshi); // 31 (computed automatically)
```

### Setters Implementation
Files:
- `src/lexer/token.go`: Added SET token
- `src/ast/statements.go`: Extended ClassDeclaration with Setters map
- `src/parser/statements.go`: parseSetter() function
- `src/evaluator/classes.go`: Setter evaluation
- `src/evaluator/expressions_member.go`: Setter invocation on assignment
- `src/object/object.go`: Extended Class struct with Setters map

**Features:**
- `set propertyName(value) { }`: Define setter methods
- Exactly one parameter required
- Executed automatically when property is assigned
- Full closure support with `ei` (this) binding
- Enable validation and transformation before assignment

**Example:**
```banglacode
sreni Temperature {
    shuru() {
        ei._celsius = 0;
    }

    pao celsius() {
        ferao ei._celsius;
    }

    set celsius(c) {
        ei._celsius = c;
    }

    pao fahrenheit() {
        ferao (ei._celsius * 9 / 5) + 32;
    }

    set fahrenheit(f) {
        ei._celsius = (f - 32) * 5 / 9;
    }
}

dhoro temp = notun Temperature();
temp.celsius = 100;
dekho(temp.fahrenheit); // 212

temp.fahrenheit = 32;
dekho(temp.celsius); // 0
```

### Static Properties Implementation
Files:
- `src/ast/statements.go`: Extended ClassDeclaration with StaticProperties map
- `src/parser/statements.go`: parseStaticProperty() function
- `src/evaluator/classes.go`: Static property evaluation
- `src/evaluator/expressions_member.go`: accessClassMember() and assignClassMember()
- `src/object/object.go`: Extended Class struct with StaticProperties map

**Features:**
- `sthir propertyName = value`: Define static properties (স্থির - static/constant)
- Belong to the class, not instances
- Access via `ClassName.property`
- Can be modified at runtime: `ClassName.property = newValue`
- Shared across all instances
- Perfect for constants and class-level counters

**Example:**
```banglacode
sreni Circle {
    sthir PI = 3.14159;

    shuru(radius) {
        ei.radius = radius;
    }

    kaj area() {
        ferao Circle.PI * ei.radius * ei.radius;
    }
}

dekho(Circle.PI); // 3.14159
dhoro c = notun Circle(10);
dekho(c.area()); // 314.159
```

### Private Fields Implementation
Files:
- `src/object/object.go`: Extended Instance struct with PrivateFields map
- `src/evaluator/expressions_member.go`: Private field access/assignment with `_` prefix check
- `src/evaluator/classes.go`: Instance initialization with PrivateFields

**Features:**
- Use underscore prefix (`_field`) for private fields by convention
- Separate storage from public properties
- Signals "internal use only" to other developers
- Access via getters/setters for encapsulation
- Full support for all data types

**Example:**
```banglacode
sreni BankAccount {
    shuru(balance) {
        ei._balance = balance;  // Private field
    }

    kaj deposit(amount) {
        ei._balance = ei._balance + amount;
    }

    kaj withdraw(amount) {
        jodi (amount <= ei._balance) {
            ei._balance = ei._balance - amount;
            ferao sotti;
        }
        ferao mittha;
    }

    pao balance() {
        ferao ei._balance;  // Controlled access
    }
}

dhoro account = notun BankAccount(1000);
account.deposit(500);
dekho(account.balance); // 1500 (via getter)
```

### Testing
File: `test/oop_enhancements_test.go` (494 lines)

**Tests:**
- TestClassGetter: Basic getter functionality
- TestClassSetter: Basic setter functionality
- TestGetterSetterTogether: Bidirectional conversion (Celsius/Fahrenheit)
- TestFahrenheitSetter: Setter with calculation
- TestStaticProperties: Class-level property access
- TestStaticPropertyInMethod: Static property usage in instance methods
- TestMultipleStaticProperties: Multiple static properties
- TestPrivateFields: Private field encapsulation with methods
- TestPrivateFieldDirectAccess: Direct access to private fields
- TestComplexGetter: Getter with conditional logic
- TestSetterWithValidation: Setter with validation logic
- TestCombinedOOPFeatures: Getters, setters, static properties, and private fields together
- TestProductWithComputedProperties: Real-world product with computed total/tax

**Results:** 13 tests passing

### VS Code Extension Updates
Files:
- `Extension/syntaxes/banglacode.tmLanguage.json`: Added `pao` and `set` keyword highlighting
- `Extension/snippets/banglacode.json`: Added 6 OOP snippets

**Snippets:**
- `pao`: Getter method template
- `set`: Setter method template
- `pao-set`: Getter and setter together
- `sthir-prop`: Static property template
- `sreni-getter-setter`: Class with getter/setter template
- `sreni-static`: Class with static properties template

### Documentation Updates
File: `Documentation/app/docs/oop/page.tsx`

**Content:**
- Comprehensive getter/setter documentation
- Static properties with examples
- Private fields convention
- Real-world examples: Temperature converter, Product with tax
- Best practices section
- Feature summary with use cases

### Architecture Notes
- **Getters**: Zero-parameter functions stored in Class.Getters map, invoked automatically on property access
- **Setters**: Single-parameter functions stored in Class.Setters map, invoked automatically on property assignment
- **Static Properties**: Stored in Class.StaticProperties map, accessed via Class object member access
- **Private Fields**: Stored in Instance.PrivateFields map, identified by `_` prefix convention
- **Evaluation Order**: When accessing property: getters → regular properties → methods
- **Assignment Order**: When assigning property: setters → private fields (if `_` prefix) → normal properties

### Performance Considerations
- Getters computed on every access (no caching) - keep logic lightweight
- Setters execute on every assignment - validate efficiently
- Static properties shared across instances - single memory allocation
- Private fields use separate map - no performance impact on normal properties

---

**Phase 11 Summary:**
- 4 major OOP features implemented
- 13 comprehensive tests (all passing)
- 470 total tests passing
- Full VS Code extension support with syntax highlighting and snippets
- Complete documentation with real-world examples
- Compatible with existing OOP features (inheritance, methods, constructors)

**Version:** 7.0.14  
**Status:** Phase 11 Complete ✅

---

## Phase 12: File System Enhancements (v7.0.15)

**Goal:** Add comprehensive file system operations including append, delete, copy, and file watching capabilities for complete file management.

### File Operations (6 functions)

#### 1. **file_jog()** - Append to File (জোগ = add)
```bangla
file_jog(path, content)
```
- **Parameters:**
  - `path` (string) - File path
  - `content` (string) - Content to append
- **Returns:** boolean (success/failure)
- **Creates file if doesn't exist**

**Example:**
```bangla
// Append to log file
file_jog("app.log", "2026-02-22 12:00:00 - User logged in\n");
file_jog("app.log", "2026-02-22 12:00:01 - Data processed\n");

// Incremental data collection
file_jog("results.csv", "John,25,Engineer\n");
file_jog("results.csv", "Sarah,30,Doctor\n");
```

#### 2. **file_mochho()** - Delete File (মোছো = erase)
```bangla
file_mochho(path)
```
- **Parameters:**
  - `path` (string) - File path to delete
- **Returns:** boolean (success/failure)
- **Permanently deletes the file**

**Example:**
```bangla
// Delete temporary file
file_mochho("temp.txt");

// Clean up old logs
file_mochho("old_log.txt");
```

#### 3. **file_nokol()** - Copy File (নকল = duplicate)
```bangla
file_nokol(source, destination)
```
- **Parameters:**
  - `source` (string) - Source file path
  - `destination` (string) - Destination file path
- **Returns:** boolean (success/failure)
- **Efficiently copies large files using io.Copy**

**Example:**
```bangla
// Backup configuration
file_nokol("config.json", "config.backup.json");

// Duplicate data
file_nokol("data.txt", "data_copy.txt");

// Create template copy
file_nokol("template.html", "index.html");
```

#### 4. **folder_mochho()** - Delete Folder (ফোল্ডার মোছো)
```bangla
folder_mochho(path, [recursive])
```
- **Parameters:**
  - `path` (string) - Folder path
  - `recursive` (boolean, optional) - Delete contents recursively
- **Returns:** boolean (success/failure)
- **If recursive = sotti, deletes folder and all contents**
- **If recursive = mittha/omitted, only deletes empty folder**

**Example:**
```bangla
// Delete empty folder
folder_mochho("empty_dir");

// Delete folder with contents
folder_mochho("old_project", sotti);

// Clean up build artifacts
folder_mochho("build", sotti);
folder_mochho("dist", sotti);
```

#### 5. **file_dekhun()** - Watch File (দেখুন = watch)
```bangla
file_dekhun(path, callback)
```
- **Parameters:**
  - `path` (string) - File path to watch
  - `callback` (function) - Function called on change: `kaj(event, filename) { ... }`
- **Returns:** watcher map with `path` and `active` properties
- **Polls file every 1 second for changes**
- **Callback receives: event ("change"), filename**

**Example:**
```bangla
// Watch configuration file
dhoro watcher = file_dekhun("config.json", kaj(event, filename) {
  dekho("Config file changed:", event, filename);
  
  // Reload configuration
  dhoro newConfig = json_poro(poro("config.json"));
  dekho("Reloaded:", newConfig);
});

// Watch data file for updates
dhoro dataWatcher = file_dekhun("data.txt", kaj(event, filename) {
  dekho("Data updated at:", somoy());
});
```

#### 6. **file_dekhun_bondho()** - Stop Watching (দেখুন বন্ধ = stop watching)
```bangla
file_dekhun_bondho(watcher)
```
- **Parameters:**
  - `watcher` (map) - Watcher object returned by `file_dekhun()`
- **Returns:** boolean (success/failure)
- **Stops the file watching goroutine**

**Example:**
```bangla
// Watch file temporarily
dhoro watcher = file_dekhun("temp.log", kaj(event, filename) {
  dekho("Log changed");
});

// Stop watching after 10 seconds
ghumaao(10000).tarpor(kaj() {
  file_dekhun_bondho(watcher);
  dekho("Stopped watching");
});
```

### Real-World Use Cases

#### Use Case 1: Log File Management
```bangla
// Initialize log file
lekho("app.log", "Application started at " + somoy() + "\n");

// Append logs throughout runtime
file_jog("app.log", "User login: admin\n");
file_jog("app.log", "Query executed: SELECT * FROM users\n");
file_jog("app.log", "Response sent: 200 OK\n");

// Backup old logs
file_nokol("app.log", "app_" + somoy() + ".log.backup");

// Clear current log
lekho("app.log", "New session started\n");
```

#### Use Case 2: Data Pipeline with File Watching
```bangla
// Watch input directory for new files
dhoro inputWatcher = file_dekhun("input/data.csv", kaj(event, filename) {
  dekho("New data file detected:", filename);
  
  // Read and process data
  dhoro data = poro("input/data.csv");
  dhoro processed = processData(data);
  
  // Write to output
  lekho("output/processed.csv", processed);
  
  // Backup input file
  file_nokol("input/data.csv", "archive/data_" + somoy() + ".csv");
  
  // Clean up input
  file_mochho("input/data.csv");
});

// Clean up on exit
file_dekhun_bondho(inputWatcher);
```

#### Use Case 3: Configuration Hot Reload
```bangla
// Load initial configuration
dhoro config = json_poro(poro("config.json"));

// Watch for configuration changes
dhoro configWatcher = file_dekhun("config.json", kaj(event, filename) {
  dekho("Configuration changed, reloading...");
  
  // Backup old config before reload
  file_nokol("config.json", "config.backup.json");
  
  // Reload configuration
  chesta {
    config = json_poro(poro("config.json"));
    dekho("Configuration reloaded successfully");
  } dhoro_bhul(err) {
    dekho("Failed to reload config:", err);
    // Restore from backup
    file_nokol("config.backup.json", "config.json");
  }
});
```

#### Use Case 4: Temporary File Management
```bangla
// Create temporary directory
folder_banao("temp_processing");

// Process files
ghuriye (dhoro i = 0; i < 10; i = i + 1) {
  dhoro filename = "temp_processing/file_" + lipi(i) + ".txt";
  lekho(filename, "Processing data " + lipi(i));
  // ... process ...
}

// Clean up all temporary files
folder_mochho("temp_processing", sotti);
dekho("Cleaned up temporary files");
```

### Technical Details

**File Append (`file_jog`):**
- Uses `os.OpenFile` with `O_APPEND|O_CREATE|O_WRONLY` flags
- Mode 0644 (rw-r--r--)
- Creates file if doesn't exist
- Efficient for incremental writes (logs, data collection)

**File Delete (`file_mochho`):**
- Uses `os.Remove` to delete single file
- Returns error if file doesn't exist
- Permanent deletion (no recovery)

**File Copy (`file_nokol`):**
- Uses `io.Copy` for efficient copying
- Handles large files efficiently (streaming copy)
- Proper resource cleanup with defer Close()
- Creates destination if doesn't exist

**Folder Delete (`folder_mochho`):**
- Recursive = sotti: uses `os.RemoveAll` (deletes all contents)
- Recursive = mittha: uses `os.Remove` (only empty folder)
- Returns error if folder doesn't exist or not empty (non-recursive)

**File Watching (`file_dekhun`):**
- Polling-based: checks file ModTime every 1 second
- Runs in goroutine (non-blocking)
- Callback invoked with ("change", filename) on modification
- Returns watcher map: `{ path: string, active: boolean }`
- Stops automatically if file is deleted

**Stop Watching (`file_dekhun_bondho`):**
- Sets watcher active flag to false
- Goroutine exits on next iteration
- Safe to call multiple times

### Performance Notes

- **File append** is optimized for sequential writes (no seek operations)
- **File copy** uses buffered I/O (efficient for large files)
- **File watching** uses 1-second polling (low CPU overhead)
- **Folder delete** (recursive) handles deep directory trees efficiently

### Error Handling

All functions return boolean (sotti/mittha) for success/failure:
```bangla
// Check if operation succeeded
jodi (file_mochho("file.txt")) {
  dekho("File deleted successfully");
} nahole {
  dekho("Failed to delete file");
}

// Safe file copy with error handling
jodi (file_nokol("source.txt", "dest.txt")) {
  dekho("File copied");
} nahole {
  dekho("Copy failed - check if source exists");
}
```

**Phase 12 Summary:**
- ✅ 6 file operation functions implemented
- ✅ 10 comprehensive tests (all passing)
- ✅ Append, delete, copy, folder delete, file watching
- ✅ Efficient I/O operations (io.Copy, goroutines)
- ✅ Cross-platform file operations
- ✅ VS Code extension updated (6 functions, 6 snippets)
- ✅ Documentation updated with 4 real-world examples

**Version:** 7.0.15  
**Status:** Phase 12 Complete ✅

---

## Phase 13: Error Handling Enhancements (v7.0.16)

**Goal:** Implement JavaScript-compatible error types (TypeError, ReferenceError, RangeError, SyntaxError) with stack traces and error utility functions for robust error handling.

### Custom Error Types (5 error constructors)

BanglaCode now provides standard JavaScript error types as constructors that return error objects:

#### 1. **Error()** - Generic Error Constructor
```bangla
Error(message)
```
- **Parameters:** `message` (string) - Error description
- **Returns:** Error map with `name`, `message`, and `stack` properties
- **Use case:** General purpose errors

**Example:**
```bangla
dhoro err = Error("Something went wrong");
felo err;

// In catch block:
chesta {
    felo Error("Operation failed");
} dhoro_bhul(e) {
    dekho(bhul_naam(e));     // "Error"
    dekho(bhul_message(e));  // "Operation failed"
}
```

#### 2. **TypeError()** - Type Mismatch Errors
```bangla
TypeError(message)
```
- **Use case:** When wrong type is provided
- **Returns:** Error map with `name: "TypeError"`

**Example:**
```bangla
kaj divide(a, b) {
    jodi (dhoron(a) != "NUMBER" ba dhoron(b) != "NUMBER") {
        felo TypeError("Arguments must be numbers");
    }
    ferao a / b;
}

chesta {
    divide("10", 2);
} dhoro_bhul(e) {
    dekho(bhul_naam(e));  // "TypeError"
}
```

#### 3. **ReferenceError()** - Undefined Variable Errors
```bangla
ReferenceError(message)
```
- **Use case:** When required variable/property doesn't exist
- **Returns:** Error map with `name: "ReferenceError"`

**Example:**
```bangla
kaj getUser(userId) {
    dhoro users = {"1": "Rahim", "2": "Karim"};
    
    dhoro userKeys = chabi(users);
    dhoro found = mittha;
    ghuriye (dhoro i = 0; i < kato(userKeys); i = i + 1) {
        jodi (userKeys[i] == userId) {
            found = sotti;
        }
    }
    
    jodi (!found) {
        felo ReferenceError("User not found: " + userId);
    }
    
    ferao users[userId];
}
```

#### 4. **RangeError()** - Out of Range Errors
```bangla
RangeError(message)
```
- **Use case:** When value is outside acceptable range
- **Returns:** Error map with `name: "RangeError"`

**Example:**
```bangla
kaj validateAge(age) {
    jodi (age < 0) {
        felo RangeError("Age cannot be negative");
    }
    jodi (age > 150) {
        felo RangeError("Age must be <= 150");
    }
    ferao sotti;
}

chesta {
    validateAge(-5);
} dhoro_bhul(e) {
    dekho(bhul_naam(e));     // "RangeError"
    dekho(bhul_message(e));  // "Age cannot be negative"
}
```

#### 5. **SyntaxError()** - Syntax/Parse Errors
```bangla
SyntaxError(message)
```
- **Use case:** When parsing or syntax validation fails
- **Returns:** Error map with `name: "SyntaxError"`

**Example:**
```bangla
kaj parseJSON(jsonString) {
    chesta {
        ferao json_poro(jsonString);
    } dhoro_bhul(e) {
        felo SyntaxError("Invalid JSON format");
    }
}

chesta {
    parseJSON("{invalid json");
} dhoro_bhul(e) {
    dekho(bhul_naam(e));  // "SyntaxError"
}
```

### Error Utility Functions (4 functions)

#### 1. **bhul_message(error)** - Get Error Message (বুল = error, message)
```bangla
bhul_message(error)
```
- **Parameters:** `error` - Error object
- **Returns:** String containing error message
- **Works with:** Error objects, Exception objects, any object with `message` property

**Example:**
```bangla
dhoro err = TypeError("Expected number");
dekho(bhul_message(err));  // "Expected number"
```

#### 2. **bhul_naam(error)** - Get Error Name/Type (নাম = name)
```bangla
bhul_naam(error)
```
- **Parameters:** `error` - Error object
- **Returns:** String containing error type name
- **Returns:** "Error", "TypeError", "ReferenceError", "RangeError", or "SyntaxError"

**Example:**
```bangla
dhoro err1 = TypeError("test");
dhoro err2 = RangeError("test");
dekho(bhul_naam(err1));  // "TypeError"
dekho(bhul_naam(err2));  // "RangeError"
```

#### 3. **bhul_stack(error)** - Get Stack Trace (স্ট্যাক = stack)
```bangla
bhul_stack(error)
```
- **Parameters:** `error` - Error object
- **Returns:** String containing stack trace information
- **Shows:** Function call chain and line numbers

**Example:**
```bangla
chesta {
    felo Error("Test error");
} dhoro_bhul(e) {
    dekho(bhul_stack(e));
    // Stack trace:
    //   at <throw statement> (line X)
}
```

#### 4. **is_error(value)** - Check if Value is Error
```bangla
is_error(value)
```
- **Parameters:** `value` - Any value to check
- **Returns:** Boolean (sotti/mittha)
- **Recognizes:** All error types (Error, TypeError, etc.)

**Example:**
```bangla
dhoro err = TypeError("test");
dhoro num = 42;

dekho(is_error(err));  // sotti
dekho(is_error(num));  // mittha
```

### Real-World Use Cases

#### Use Case 1: API Input Validation
```bangla
kaj validateAPIRequest(request) {
    // Type checking
    jodi (dhoron(request) != "MAP") {
        felo TypeError("Request must be an object");
    }
    
    // Required fields
    dhoro keys = chabi(request);
    dhoro hasMethod = mittha;
    dhoro hasUrl = mittha;
    
    ghuriye (dhoro i = 0; i < kato(keys); i = i + 1) {
        jodi (keys[i] == "method") { hasMethod = sotti; }
        jodi (keys[i] == "url") { hasUrl = sotti; }
    }
    
    jodi (!hasMethod) {
        felo ReferenceError("Missing required field: method");
    }
    jodi (!hasUrl) {
        felo ReferenceError("Missing required field: url");
    }
    
    // Value validation
    dhoro method = request["method"];
    dhoro validMethods = ["GET", "POST", "PUT", "DELETE"];
    dhoro isValidMethod = mittha;
    
    ghuriye (dhoro i = 0; i < kato(validMethods); i = i + 1) {
        jodi (method == validMethods[i]) {
            isValidMethod = sotti;
        }
    }
    
    jodi (!isValidMethod) {
        felo RangeError("Invalid HTTP method: " + method);
    }
    
    ferao sotti;
}

// Usage
chesta {
    validateAPIRequest({"method": "GET", "url": "/api/users"});
    dekho("Request is valid");
} dhoro_bhul(e) {
    dekho("Validation error:");
    dekho("  Type:", bhul_naam(e));
    dekho("  Message:", bhul_message(e));
}
```

#### Use Case 2: Safe Data Processing
```bangla
kaj processUserData(data) {
    // Step 1: Type validation
    chesta {
        jodi (dhoron(data) != "ARRAY") {
            felo TypeError("Data must be an array");
        }
    } dhoro_bhul(e) {
        dekho("Type Error:", bhul_message(e));
        ferao khali;
    }
    
    // Step 2: Range validation
    chesta {
        jodi (kato(data) == 0) {
            felo RangeError("Data array cannot be empty");
        }
        jodi (kato(data) > 1000) {
            felo RangeError("Data array too large (max 1000)");
        }
    } dhoro_bhul(e) {
        dekho("Range Error:", bhul_message(e));
        ferao khali;
    }
    
    // Step 3: Process data
    dhoro processed = [];
    ghuriye (dhoro i = 0; i < kato(data); i = i + 1) {
        chesta {
            dhoro item = data[i];
            jodi (dhoron(item) != "NUMBER") {
                felo TypeError("Array item must be number");
            }
            dhokao(processed, item * 2);
        } dhoro_bhul(e) {
            dekho("Skipping invalid item at index", i, ":", bhul_message(e));
        }
    }
    
    ferao processed;
}

// Usage
dhoro result = processUserData([1, 2, 3, 4, 5]);
dekho(result);  // [2, 4, 6, 8, 10]
```

#### Use Case 3: Error Type Discrimination
```bangla
kaj handleError(error) {
    jodi (!is_error(error)) {
        dekho("Not an error object");
        ferao;
    }
    
    dhoro errorType = bhul_naam(error);
    dhoro errorMsg = bhul_message(error);
    
    // Different handling based on error type
    jodi (errorType == "TypeError") {
        dekho("Type Error - Check data types:");
        dekho("  →", errorMsg);
    } nahole jodi (errorType == "ReferenceError") {
        dekho("Reference Error - Check if variable/property exists:");
        dekho("  →", errorMsg);
    } nahole jodi (errorType == "RangeError") {
        dekho("Range Error - Check value is within bounds:");
        dekho("  →", errorMsg);
    } nahole jodi (errorType == "SyntaxError") {
        dekho("Syntax Error - Check input format:");
        dekho("  →", errorMsg);
    } nahole {
        dekho("Generic Error:");
        dekho("  →", errorMsg);
    }
}

// Test different error types
handleError(TypeError("Expected string"));
handleError(RangeError("Index out of bounds"));
handleError(ReferenceError("Variable undefined"));
```

#### Use Case 4: Nested Function Error Propagation
```bangla
kaj readConfig() {
    dhoro config = poro("config.json");
    jodi (config == khali) {
        felo ReferenceError("Config file not found");
    }
    ferao config;
}

kaj parseConfig() {
    dhoro content = readConfig();  // May throw ReferenceError
    dhoro parsed = json_poro(content);
    jodi (parsed == khali) {
        felo SyntaxError("Invalid config format");
    }
    ferao parsed;
}

kaj validateConfig(config) {
    dhoro keys = chabi(config);
    dhoro hasPort = mittha;
    
    ghuriye (dhoro i = 0; i < kato(keys); i = i + 1) {
        jodi (keys[i] == "port") { hasPort = sotti; }
    }
    
    jodi (!hasPort) {
        felo ReferenceError("Config missing 'port' field");
    }
    
    dhoro port = config["port"];
    jodi (port < 1024 ba port > 65535) {
        felo RangeError("Port must be between 1024 and 65535");
    }
}

kaj loadConfig() {
    chesta {
        dhoro config = parseConfig();
        validateConfig(config);
        ferao config;
    } dhoro_bhul(e) {
        dekho("Failed to load config:");
        dekho("  Error:", bhul_naam(e));
        dekho("  Reason:", bhul_message(e));
        dekho("  Stack:", bhul_stack(e));
        ferao khali;
    }
}
```

### Technical Details

**Error Object Structure:**
- Error objects are Maps with three properties:
  - `name` (string): Error type ("Error", "TypeError", etc.)
  - `message` (string): Error description
  - `stack` (string): Stack trace (populated when thrown)

**Stack Trace:**
- Captured when error is thrown via `felo`
- Shows function name, file, line, and column
- Format: `at <function> (file:line:col)`

**Error Propagation:**
- Errors propagate up call stack until caught
- Uncaught errors terminate program execution
- Exception object wraps error for internal handling

**Try-Catch Integration:**
- Caught error parameter receives error object (not just message)
- Error utilities work with caught errors
- Stack trace preserved through catch chain

### Performance Notes

- Error creation is lightweight (Map allocation)
- Stack trace capture is minimal overhead
- Error utilities are O(1) map/property access
- No performance impact on non-error code paths

### Best Practices

1. **Use specific error types:**
   ```bangla
   // ✅ Good - specific error type
   felo TypeError("Expected number");
   
   // ❌ Bad - generic error
   felo Error("Wrong type");
   ```

2. **Provide descriptive messages:**
   ```bangla
   // ✅ Good - actionable message
   felo RangeError("Age must be between 0 and 150, got -5");
   
   // ❌ Bad - vague message
   felo RangeError("Bad value");
   ```

3. **Check error types in handlers:**
   ```bangla
   chesta {
       processData(input);
   } dhoro_bhul(e) {
       dhoro errorType = bhul_naam(e);
       jodi (errorType == "TypeError") {
           // Type-specific recovery
       }
   }
   ```

4. **Validate before processing:**
   ```bangla
   kaj processValue(val) {
       // Validate first
       jodi (dhoron(val) != "NUMBER") {
           felo TypeError("Value must be number");
       }
       jodi (val < 0) {
           felo RangeError("Value must be >= 0");
       }
       
       // Then process
       ferao val * 2;
   }
   ```

**Phase 13 Summary:**
- ✅ 5 error type constructors (Error, TypeError, ReferenceError, RangeError, SyntaxError)
- ✅ 4 error utility functions (bhul_message, bhul_naam, bhul_stack, is_error)
- ✅ Stack trace support for thrown errors
- ✅ 15 comprehensive tests (all passing, 547 total tests)
- ✅ JavaScript-compatible error handling
- ✅ Enhanced try-catch with typed errors
- ✅ VS Code extension updated (9 error types/functions, 7 snippets)
- ✅ Documentation updated with 4 real-world examples

**Version:** 7.0.16  
**Status:** Phase 13 Complete ✅
