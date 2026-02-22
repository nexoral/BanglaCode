export const metadata = {
  title: 'Collections (Set & Map) - BanglaCode Documentation',
  description: 'Learn how to use Set and ES6 Map collections in BanglaCode for managing unique values and key-value pairs with any type as key',
}

# Collections: Set & Map

BanglaCode provides modern ES6-style **Set** and **Map** collections for efficient data management. Unlike JavaScript objects that only support string keys, BanglaCode's Map supports **any type as key** including arrays and objects.

## Quick Start

```banglacode
// Set - Unique values collection
dhoro mySet = set_srishti([1, 2, 3, 2, 1]);
dekho(set_akar(mySet)); // 3 (duplicates removed)

// Map - Key-value pairs with any key type
dhoro myMap = map_srishti();
map_set(myMap, "name", "Ankan");
map_set(myMap, [1, 2], "array key");
dekho(map_get(myMap, "name")); // "Ankan"
```

## Core Concepts

### Set
A **Set** is a collection of **unique values**. Adding a duplicate value has no effect. Sets maintain insertion order and use SHA-256 hashing to determine uniqueness.

**Key Features:**
- ✅ Automatic duplicate removal
- ✅ Maintains insertion order
- ✅ Supports any value type (numbers, strings, arrays, objects)
- ✅ Fast membership testing with `set_has()`

### Map
A **Map** is an ES6-style key-value collection that supports **any type as key**, not just strings. This is a major advantage over regular BanglaCode objects which only allow string keys.

**Key Features:**
- ✅ Any type as key (arrays, objects, numbers, strings)
- ✅ Maintains insertion order
- ✅ Separate keys and values tracking
- ✅ Fast lookups with `map_get()` and `map_has()`

### Set vs Array
| Feature | Set | Array |
|---------|-----|-------|
| Duplicates | ❌ Removed automatically | ✅ Allowed |
| Order | ✅ Insertion order | ✅ Index order |
| Lookup | O(1) with `set_has()` | O(n) with loop |
| Use Case | Unique values | Ordered list |

### Map vs Object
| Feature | Map | Object |
|---------|-----|--------|
| Key Types | ✅ Any type | ❌ Strings only |
| Size | `map_akar()` | Manual count |
| Iteration | `map_foreach()` | Loop over keys |
| Use Case | Complex keys | Simple string keys |

---

## Set API Reference

### `set_srishti()`
Creates a new empty Set or initializes it from an array.

**Syntax:**
```banglacode
dhoro mySet = set_srishti();
dhoro mySet = set_srishti(array);
```

**Parameters:**
- `array` (optional): Array of initial values (duplicates will be removed)

**Returns:** New Set object

**Examples:**
```banglacode
// Empty Set
dhoro emptySet = set_srishti();
dekho(set_akar(emptySet)); // 0

// From array (removes duplicates)
dhoro numbers = set_srishti([1, 2, 3, 2, 1]);
dekho(set_akar(numbers)); // 3

// With strings
dhoro fruits = set_srishti(["apple", "banana", "apple"]);
dekho(set_akar(fruits)); // 2
```

---

### `set_add(set, element)`
Adds an element to the Set. If the element already exists, no action is taken.

**Syntax:**
```banglacode
set_add(mySet, element);
```

**Parameters:**
- `set`: The Set object
- `element`: Value to add (any type)

**Returns:** The Set (for chaining)

**Examples:**
```banglacode
dhoro mySet = set_srishti();
set_add(mySet, 1);
set_add(mySet, 2);
set_add(mySet, 1); // No effect - already exists
dekho(set_akar(mySet)); // 2

// Add complex types
set_add(mySet, [1, 2, 3]);
set_add(mySet, {naam: "Ankan"});
```

---

### `set_has(set, element)`
Checks if an element exists in the Set.

**Syntax:**
```banglacode
dhoro exists = set_has(mySet, element);
```

**Parameters:**
- `set`: The Set object
- `element`: Value to check

**Returns:** `sotti` (true) if exists, `mittha` (false) otherwise

**Examples:**
```banglacode
dhoro mySet = set_srishti([1, 2, 3]);

jodi (set_has(mySet, 2)) {
  dekho("Found!"); // Prints "Found!"
}

jodi (set_has(mySet, 5)) {
  dekho("Not found");
} nahole {
  dekho("5 doesn't exist"); // Prints this
}
```

---

### `set_delete(set, element)`
Removes an element from the Set.

**Syntax:**
```banglacode
dhoro deleted = set_delete(mySet, element);
```

**Parameters:**
- `set`: The Set object
- `element`: Value to remove

**Returns:** `sotti` if deleted, `mittha` if element didn't exist

**Examples:**
```banglacode
dhoro mySet = set_srishti([1, 2, 3]);
set_delete(mySet, 2);
dekho(set_akar(mySet)); // 2

// Check if deleted
dhoro deleted = set_delete(mySet, 2);
dekho(deleted); // mittha (already removed)
```

---

### `set_clear(set)`
Removes all elements from the Set.

**Syntax:**
```banglacode
set_clear(mySet);
```

**Parameters:**
- `set`: The Set object

**Returns:** `khali` (null)

**Examples:**
```banglacode
dhoro mySet = set_srishti([1, 2, 3]);
dekho(set_akar(mySet)); // 3

set_clear(mySet);
dekho(set_akar(mySet)); // 0
```

---

### `set_akar(set)`
Returns the number of elements in the Set.

**Syntax:**
```banglacode
dhoro size = set_akar(mySet);
```

**Parameters:**
- `set`: The Set object

**Returns:** Number of elements

**Examples:**
```banglacode
dhoro mySet = set_srishti([1, 2, 3]);
dekho(set_akar(mySet)); // 3

set_add(mySet, 4);
dekho(set_akar(mySet)); // 4

set_delete(mySet, 1);
dekho(set_akar(mySet)); // 3
```

---

### `set_values(set)`
Returns all values in the Set as an array.

**Syntax:**
```banglacode
dhoro values = set_values(mySet);
```

**Parameters:**
- `set`: The Set object

**Returns:** Array of all values (in insertion order)

**Examples:**
```banglacode
dhoro mySet = set_srishti([3, 1, 2]);
dhoro values = set_values(mySet);
dekho(values); // [3, 1, 2] (insertion order preserved)

// Convert back to array after removing duplicates
dhoro arr = [1, 2, 3, 2, 1];
dhoro unique = set_values(set_srishti(arr));
dekho(unique); // [1, 2, 3]
```

---

### `set_foreach(set, callback)`
Iterates over all elements in the Set and calls the callback for each.

**Syntax:**
```banglacode
set_foreach(mySet, kaj(value) {
  // Process value
});
```

**Parameters:**
- `set`: The Set object
- `callback`: Function receiving `value` parameter

**Returns:** `khali` (null)

**Examples:**
```banglacode
dhoro mySet = set_srishti([1, 2, 3]);

// Print each value
set_foreach(mySet, kaj(value) {
  dekho(value);
});

// Sum all values
dhoro sum = 0;
set_foreach(mySet, kaj(value) {
  sum = sum + value;
});
dekho(sum); // 6
```

---

## Map API Reference

### `map_srishti()`
Creates a new empty Map or initializes it from an array of `[key, value]` pairs.

**Syntax:**
```banglacode
dhoro myMap = map_srishti();
dhoro myMap = map_srishti(entries);
```

**Parameters:**
- `entries` (optional): Array of `[key, value]` arrays

**Returns:** New Map object

**Examples:**
```banglacode
// Empty Map
dhoro emptyMap = map_srishti();

// From entries
dhoro myMap = map_srishti([
  ["name", "Ankan"],
  ["age", 25],
  ["city", "Kolkata"]
]);
dekho(map_akar(myMap)); // 3
```

---

### `map_set(map, key, value)`
Sets a key-value pair in the Map. If key exists, value is updated.

**Syntax:**
```banglacode
map_set(myMap, key, value);
```

**Parameters:**
- `map`: The Map object
- `key`: Key (any type)
- `value`: Value (any type)

**Returns:** The Map (for chaining)

**Examples:**
```banglacode
dhoro myMap = map_srishti();
map_set(myMap, "name", "Ankan");
map_set(myMap, "age", 25);

// Update existing key
map_set(myMap, "age", 26);
dekho(map_get(myMap, "age")); // 26

// Use array as key
dhoro keyArr = [1, 2];
map_set(myMap, keyArr, "array value");
```

---

### `map_get(map, key)`
Gets the value associated with the key.

**Syntax:**
```banglacode
dhoro value = map_get(myMap, key);
```

**Parameters:**
- `map`: The Map object
- `key`: Key to lookup

**Returns:** Value if key exists, `khali` (null) otherwise

**Examples:**
```banglacode
dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
dekho(map_get(myMap, "name")); // "Ankan"
dekho(map_get(myMap, "city")); // khali (null)

// Check before using
dhoro city = map_get(myMap, "city");
jodi (city == khali) {
  dekho("City not set");
}
```

---

### `map_has(map, key)`
Checks if a key exists in the Map.

**Syntax:**
```banglacode
dhoro exists = map_has(myMap, key);
```

**Parameters:**
- `map`: The Map object
- `key`: Key to check

**Returns:** `sotti` (true) if exists, `mittha` (false) otherwise

**Examples:**
```banglacode
dhoro myMap = map_srishti([["name", "Ankan"]]);

jodi (map_has(myMap, "name")) {
  dekho("Name exists");
}

jodi (na map_has(myMap, "age")) {
  dekho("Age not set");
}
```

---

### `map_delete(map, key)`
Removes a key-value pair from the Map.

**Syntax:**
```banglacode
dhoro deleted = map_delete(myMap, key);
```

**Parameters:**
- `map`: The Map object
- `key`: Key to remove

**Returns:** `sotti` if deleted, `mittha` if key didn't exist

**Examples:**
```banglacode
dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
map_delete(myMap, "age");
dekho(map_akar(myMap)); // 1

// Check if deleted
dhoro deleted = map_delete(myMap, "age");
dekho(deleted); // mittha (already removed)
```

---

### `map_clear(map)`
Removes all entries from the Map.

**Syntax:**
```banglacode
map_clear(myMap);
```

**Parameters:**
- `map`: The Map object

**Returns:** `khali` (null)

**Examples:**
```banglacode
dhoro myMap = map_srishti([["a", 1], ["b", 2]]);
dekho(map_akar(myMap)); // 2

map_clear(myMap);
dekho(map_akar(myMap)); // 0
```

---

### `map_akar(map)`
Returns the number of entries in the Map.

**Syntax:**
```banglacode
dhoro size = map_akar(myMap);
```

**Parameters:**
- `map`: The Map object

**Returns:** Number of entries

**Examples:**
```banglacode
dhoro myMap = map_srishti([["a", 1], ["b", 2]]);
dekho(map_akar(myMap)); // 2

map_set(myMap, "c", 3);
dekho(map_akar(myMap)); // 3
```

---

### `map_keys(map)`
Returns all keys in the Map as an array.

**Syntax:**
```banglacode
dhoro keys = map_keys(myMap);
```

**Parameters:**
- `map`: The Map object

**Returns:** Array of all keys (in insertion order)

**Examples:**
```banglacode
dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
dhoro keys = map_keys(myMap);
dekho(keys); // ["name", "age"]

// Iterate over keys
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
  dhoro key = keys[i];
  dekho(key, "=>", map_get(myMap, key));
}
```

---

### `map_values(map)`
Returns all values in the Map as an array.

**Syntax:**
```banglacode
dhoro values = map_values(myMap);
```

**Parameters:**
- `map`: The Map object

**Returns:** Array of all values (in insertion order)

**Examples:**
```banglacode
dhoro myMap = map_srishti([["a", 1], ["b", 2], ["c", 3]]);
dhoro values = map_values(myMap);
dekho(values); // [1, 2, 3]

// Calculate sum of values
dhoro sum = 0;
ghuriye (dhoro i = 0; i < dorghyo(values); i = i + 1) {
  sum = sum + values[i];
}
dekho(sum); // 6
```

---

### `map_entries(map)`
Returns all `[key, value]` pairs in the Map as an array of arrays.

**Syntax:**
```banglacode
dhoro entries = map_entries(myMap);
```

**Parameters:**
- `map`: The Map object

**Returns:** Array of `[key, value]` arrays (in insertion order)

**Examples:**
```banglacode
dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
dhoro entries = map_entries(myMap);
dekho(entries); // [["name", "Ankan"], ["age", 25]]

// Iterate over entries
ghuriye (dhoro i = 0; i < dorghyo(entries); i = i + 1) {
  dhoro entry = entries[i];
  dhoro key = entry[0];
  dhoro value = entry[1];
  dekho(key, "=>", value);
}
```

---

### `map_foreach(map, callback)`
Iterates over all entries in the Map and calls the callback for each.

**Syntax:**
```banglacode
map_foreach(myMap, kaj(value, key) {
  // Process key-value pair
});
```

**Parameters:**
- `map`: The Map object
- `callback`: Function receiving `value` and `key` parameters

**Returns:** `khali` (null)

**Examples:**
```banglacode
dhoro myMap = map_srishti([["a", 1], ["b", 2], ["c", 3]]);

// Print each entry
map_foreach(myMap, kaj(value, key) {
  dekho(key, "=>", value);
});

// Sum all values
dhoro sum = 0;
map_foreach(myMap, kaj(value, key) {
  sum = sum + value;
});
dekho(sum); // 6
```

---

## Real-World Examples

### Example 1: Remove Duplicates from Array

```banglacode
// Remove duplicates using Set
dhoro numbers = [1, 2, 3, 2, 1, 4, 3, 5];
dhoro uniqueSet = set_srishti(numbers);
dhoro unique = set_values(uniqueSet);

dekho("Original:", numbers);
dekho("Unique:", unique); // [1, 2, 3, 4, 5]
```

### Example 2: Count Occurrences

```banglacode
// Count how many times each word appears
dhoro words = ["apple", "banana", "apple", "cherry", "banana", "apple"];
dhoro countMap = map_srishti();

// Count occurrences
ghuriye (dhoro i = 0; i < dorghyo(words); i = i + 1) {
  dhoro word = words[i];
  dhoro count = map_get(countMap, word);
  
  jodi (count == khali) {
    map_set(countMap, word, 1);
  } nahole {
    map_set(countMap, word, count + 1);
  }
}

// Print results
dekho("Word counts:");
map_foreach(countMap, kaj(count, word) {
  dekho(word, "appears", count, "times");
});
```

### Example 3: Caching with Map

```banglacode
// Simple cache using Map
dhoro cache = map_srishti();

kaj getDataWithCache(id) {
  // Check cache first
  jodi (map_has(cache, id)) {
    dekho("Cache hit for", id);
    ferao map_get(cache, id);
  }
  
  // Simulate expensive operation
  dekho("Cache miss - fetching data for", id);
  dhoro data = "Data for ID " + id;
  
  // Store in cache
  map_set(cache, id, data);
  ferao data;
}

// Usage
dekho(getDataWithCache(1)); // Cache miss
dekho(getDataWithCache(2)); // Cache miss
dekho(getDataWithCache(1)); // Cache hit
dekho(getDataWithCache(2)); // Cache hit
```

### Example 4: Set Operations (Union, Intersection)

```banglacode
// Union of two sets
kaj setUnion(set1, set2) {
  dhoro result = set_srishti();
  
  // Add all from set1
  set_foreach(set1, kaj(value) {
    set_add(result, value);
  });
  
  // Add all from set2
  set_foreach(set2, kaj(value) {
    set_add(result, value);
  });
  
  ferao result;
}

// Intersection of two sets
kaj setIntersection(set1, set2) {
  dhoro result = set_srishti();
  
  set_foreach(set1, kaj(value) {
    jodi (set_has(set2, value)) {
      set_add(result, value);
    }
  });
  
  ferao result;
}

// Usage
dhoro setA = set_srishti([1, 2, 3, 4]);
dhoro setB = set_srishti([3, 4, 5, 6]);

dhoro unionSet = setUnion(setA, setB);
dekho("Union:", set_values(unionSet)); // [1, 2, 3, 4, 5, 6]

dhoro intersectSet = setIntersection(setA, setB);
dekho("Intersection:", set_values(intersectSet)); // [3, 4]
```

---

## Best Practices

### ✅ Do's

1. **Use Set to remove duplicates**
   ```banglacode
   dhoro unique = set_values(set_srishti(arrayWithDuplicates));
   ```

2. **Use Map for complex keys**
   ```banglacode
   dhoro map = map_srishti();
   map_set(map, [1, 2], "array key"); // Object/array keys
   ```

3. **Check existence before accessing**
   ```banglacode
   jodi (map_has(myMap, key)) {
     dhoro value = map_get(myMap, key);
   }
   ```

4. **Use `_akar()` instead of manual counting**
   ```banglacode
   dekho("Size:", set_akar(mySet)); // Fast O(1)
   ```

5. **Clear collections when done**
   ```banglacode
   set_clear(mySet); // Free memory
   map_clear(myMap);
   ```

6. **Use `foreach` for iteration**
   ```banglacode
   set_foreach(mySet, kaj(value) { dekho(value); });
   map_foreach(myMap, kaj(val, key) { dekho(key, val); });
   ```

7. **Use Map for caching**
   ```banglacode
   dhoro cache = map_srishti();
   map_set(cache, userId, userData);
   ```

8. **Preserve insertion order**
   ```banglacode
   // Both Set and Map maintain insertion order
   dhoro ordered = set_srishti([3, 1, 2]);
   dekho(set_values(ordered)); // [3, 1, 2]
   ```

### ❌ Don'ts

1. **Don't modify Set/Map during iteration**
   ```banglacode
   // BAD - modifying during iteration
   set_foreach(mySet, kaj(value) {
     set_delete(mySet, value); // Unpredictable behavior
   });
   ```

2. **Don't use regular objects for complex keys**
   ```banglacode
   // BAD - only string keys
   dhoro obj = {naam: "Ankan"};
   
   // GOOD - any type as key
   dhoro map = map_srishti();
   map_set(map, [1, 2], "value");
   ```

3. **Don't forget to check for khali**
   ```banglacode
   // BAD
   dhoro value = map_get(myMap, "key");
   dekho(value + 10); // Error if khali!
   
   // GOOD
   dhoro value = map_get(myMap, "key");
   jodi (value != khali) {
     dekho(value + 10);
   }
   ```

4. **Don't assume Set ordering is sorted**
   ```banglacode
   // Set maintains insertion order, NOT sorted order
   dhoro set = set_srishti([3, 1, 2]);
   dekho(set_values(set)); // [3, 1, 2], not [1, 2, 3]
   ```

5. **Don't create unnecessary Sets/Maps**
   ```banglacode
   // BAD - creating new Set each time
   ghuriye (...) {
     dhoro temp = set_srishti();
   }
   
   // GOOD - reuse or create once
   dhoro temp = set_srishti();
   ghuriye (...) {
     set_clear(temp); // Reuse
   }
   ```

6. **Don't use Set for ordered lists**
   ```banglacode
   // BAD - Set doesn't support indexing
   dhoro set = set_srishti([1, 2, 3]);
   dekho(set[0]); // Error!
   
   // GOOD - use array
   dhoro arr = [1, 2, 3];
   dekho(arr[0]); // 1
   ```

7. **Don't ignore return values**
   ```banglacode
   // Check if deletion was successful
   dhoro deleted = set_delete(mySet, value);
   jodi (deleted) {
     dekho("Removed successfully");
   }
   ```

8. **Don't store Sets/Maps in regular objects**
   ```banglacode
   // BAD - regular object can't properly store Sets
   dhoro obj = {mySet: set_srishti()};
   
   // GOOD - use Map to store collections
   dhoro collections = map_srishti();
   map_set(collections, "mySet", set_srishti());
   ```

---

## Performance Tips

### ⚡ Optimize for Speed

1. **Pre-allocate when possible**
   ```banglacode
   // Create from array at once (faster)
   dhoro set = set_srishti([1, 2, 3, 4, 5]);
   
   // Instead of adding one by one
   dhoro set = set_srishti();
   set_add(set, 1);
   set_add(set, 2);
   // ... slower
   ```

2. **Use `_has()` for lookups** - O(1) constant time
   ```banglacode
   // Fast Set lookup
   jodi (set_has(mySet, value)) { ... } // O(1)
   
   // Instead of array search
   jodi (arrayContains(myArr, value)) { ... } // O(n)
   ```

3. **Batch operations when possible**
   ```banglacode
   // Create Map with all entries at once
   dhoro map = map_srishti([["a", 1], ["b", 2], ["c", 3]]);
   
   // Instead of one by one
   dhoro map = map_srishti();
   map_set(map, "a", 1);
   map_set(map, "b", 2);
   map_set(map, "c", 3);
   ```

4. **Use `_foreach()` over manual iteration**
   ```banglacode
   // Faster - optimized iteration
   set_foreach(mySet, kaj(value) { dekho(value); });
   
   // Slower - array conversion + loop
   dhoro arr = set_values(mySet);
   ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
     dekho(arr[i]);
   }
   ```

5. **Clear instead of recreate**
   ```banglacode
   // Faster - reuse memory
   set_clear(mySet);
   
   // Slower - allocate new Set
   mySet = set_srishti();
   ```

---

## Common Patterns

### Pattern 1: Unique Filter
```banglacode
// Filter array to unique values
kaj uniqueFilter(arr) {
  ferao set_values(set_srishti(arr));
}

dhoro numbers = [1, 2, 2, 3, 3, 3, 4];
dekho(uniqueFilter(numbers)); // [1, 2, 3, 4]
```

### Pattern 2: Frequency Counter
```banglacode
// Count frequency of elements
kaj frequencyCounter(arr) {
  dhoro freq = map_srishti();
  
  ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
    dhoro item = arr[i];
    dhoro count = map_get(freq, item);
    map_set(freq, item, (count == khali ? 1 : count + 1));
  }
  
  ferao freq;
}

dhoro letters = ["a", "b", "a", "c", "b", "a"];
dhoro freq = frequencyCounter(letters);
map_foreach(freq, kaj(count, letter) {
  dekho(letter, "=>", count);
});
```

### Pattern 3: Memoization
```banglacode
// Memoize expensive function results
dhoro memo = map_srishti();

kaj fibonacci(n) {
  jodi (map_has(memo, n)) {
    ferao map_get(memo, n);
  }
  
  dhoro result;
  jodi (n <= 1) {
    result = n;
  } nahole {
    result = fibonacci(n - 1) + fibonacci(n - 2);
  }
  
  map_set(memo, n, result);
  ferao result;
}

dekho(fibonacci(10)); // Fast with memoization
```

### Pattern 4: Grouping
```banglacode
// Group items by property
kaj groupBy(arr, property) {
  dhoro groups = map_srishti();
  
  ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
    dhoro item = arr[i];
    dhoro key = item[property];
    dhoro group = map_get(groups, key);
    
    jodi (group == khali) {
      group = [];
      map_set(groups, key, group);
    }
    
    group[dorghyo(group)] = item;
  }
  
  ferao groups;
}

// Usage
dhoro people = [
  {naam: "Ankan", city: "Kolkata"},
  {naam: "Raj", city: "Delhi"},
  {naam: "Priya", city: "Kolkata"}
];

dhoro byCity = groupBy(people, "city");
map_foreach(byCity, kaj(people, city) {
  dekho(city, "has", dorghyo(people), "people");
});
```

---

## Related APIs

- **[Arrays](/docs/syntax#arrays)** - Ordered lists with duplicates
- **[Objects](/docs/syntax#objects)** - Simple key-value pairs (string keys only)
- **[Control Flow](/docs/control-flow)** - Loops and conditionals for iteration
- **[Functions](/docs/syntax#functions)** - Create reusable logic with `kaj`

---

## Summary

Collections (Set & Map) provide **modern ES6-style data structures** in BanglaCode:

### Set Benefits:
✅ Automatic duplicate removal  
✅ Fast membership testing (O(1))  
✅ Maintains insertion order  
✅ Supports any value type  

### Map Benefits:
✅ **Any type as key** (not just strings)  
✅ Fast key lookups (O(1))  
✅ Maintains insertion order  
✅ Proper size tracking with `map_akar()`  

**Use Set for:** Unique values, membership testing, duplicate removal  
**Use Map for:** Complex keys, caching, counting, grouping  

Collections are essential for **efficient data management** in BanglaCode! 🚀
