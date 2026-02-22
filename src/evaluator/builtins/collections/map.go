package collections

import (
	"BanglaCode/src/object"
)

// MapBuiltins contains all ES6Map-related built-in functions
var MapBuiltins = map[string]*object.Builtin{
	"map_srishti": {Fn: mapCreate},
	"map_set":     {Fn: mapSet},
	"map_get":     {Fn: mapGet},
	"map_has":     {Fn: mapHas},
	"map_delete":  {Fn: mapDelete},
	"map_clear":   {Fn: mapClear},
	"map_akar":    {Fn: mapSize},
	"map_keys":    {Fn: mapKeys},
	"map_values":  {Fn: mapValues},
	"map_entries": {Fn: mapEntries},
	"map_foreach": {Fn: mapForEach},
}

// mapCreate creates a new ES6 Map
// Usage: dhoro myMap = map_srishti();
//
//	dhoro myMap = map_srishti([[key1, value1], [key2, value2]]);
func mapCreate(args ...object.Object) object.Object {
	if len(args) > 1 {
		return &object.Error{Message: "map_srishti() expects 0 or 1 argument (optional array of [key, value] pairs)"}
	}

	m := &object.ES6Map{
		Pairs: make(map[string]object.Object),
		Keys:  make(map[string]object.Object),
		Order: []string{},
	}

	// If array of [key, value] pairs provided, add them
	if len(args) == 1 {
		if arr, ok := args[0].(*object.Array); ok {
			for _, pairObj := range arr.Elements {
				if pair, ok := pairObj.(*object.Array); ok {
					if len(pair.Elements) == 2 {
						key := pair.Elements[0]
						value := pair.Elements[1]
						keyHash := hashObject(key)

						// Only add if key doesn't exist
						if _, exists := m.Pairs[keyHash]; !exists {
							m.Order = append(m.Order, keyHash)
						}

						m.Keys[keyHash] = key
						m.Pairs[keyHash] = value
					} else {
						return &object.Error{Message: "map_srishti() each entry must be [key, value] array"}
					}
				} else {
					return &object.Error{Message: "map_srishti() argument must be array of [key, value] arrays"}
				}
			}
		} else {
			return &object.Error{Message: "map_srishti() argument must be an array"}
		}
	}

	return m
}

// mapSet sets a key-value pair in the map
// Usage: map_set(myMap, "key", "value");
func mapSet(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "map_set() expects 3 arguments: map, key, value"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_set() first argument must be a Map"}
	}

	key := args[1]
	value := args[2]
	keyHash := hashObject(key)

	// Add to order if new key
	if _, exists := m.Pairs[keyHash]; !exists {
		m.Order = append(m.Order, keyHash)
	}

	m.Keys[keyHash] = key
	m.Pairs[keyHash] = value

	return m
}

// mapGet gets a value from the map by key
// Usage: dhoro value = map_get(myMap, "key");
func mapGet(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "map_get() expects 2 arguments: map, key"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_get() first argument must be a Map"}
	}

	key := args[1]
	keyHash := hashObject(key)

	if value, exists := m.Pairs[keyHash]; exists {
		return value
	}

	return object.NULL
}

// mapHas checks if a key exists in the map
// Usage: jodi (map_has(myMap, "key")) { ... }
func mapHas(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "map_has() expects 2 arguments: map, key"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_has() first argument must be a Map"}
	}

	key := args[1]
	keyHash := hashObject(key)

	if _, exists := m.Pairs[keyHash]; exists {
		return object.TRUE
	}
	return object.FALSE
}

// mapDelete removes a key from the map
// Usage: map_delete(myMap, "key");
func mapDelete(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "map_delete() expects 2 arguments: map, key"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_delete() first argument must be a Map"}
	}

	key := args[1]
	keyHash := hashObject(key)

	if _, exists := m.Pairs[keyHash]; exists {
		delete(m.Pairs, keyHash)
		delete(m.Keys, keyHash)

		// Remove from order
		for i, hash := range m.Order {
			if hash == keyHash {
				m.Order = append(m.Order[:i], m.Order[i+1:]...)
				break
			}
		}
		return object.TRUE
	}
	return object.FALSE
}

// mapClear removes all entries from the map
// Usage: map_clear(myMap);
func mapClear(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "map_clear() expects 1 argument: map"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_clear() first argument must be a Map"}
	}

	m.Pairs = make(map[string]object.Object)
	m.Keys = make(map[string]object.Object)
	m.Order = []string{}

	return object.NULL
}

// mapSize returns the number of entries in the map
// Usage: dhoro count = map_akar(myMap);
func mapSize(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "map_akar() expects 1 argument: map"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_akar() first argument must be a Map"}
	}

	return &object.Number{Value: float64(len(m.Order))}
}

// mapKeys returns all keys in the map as an array
// Usage: dhoro keys = map_keys(myMap);
func mapKeys(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "map_keys() expects 1 argument: map"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_keys() first argument must be a Map"}
	}

	keys := make([]object.Object, 0, len(m.Order))
	for _, keyHash := range m.Order {
		keys = append(keys, m.Keys[keyHash])
	}

	return &object.Array{Elements: keys}
}

// mapValues returns all values in the map as an array
// Usage: dhoro values = map_values(myMap);
func mapValues(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "map_values() expects 1 argument: map"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_values() first argument must be a Map"}
	}

	values := make([]object.Object, 0, len(m.Order))
	for _, keyHash := range m.Order {
		values = append(values, m.Pairs[keyHash])
	}

	return &object.Array{Elements: values}
}

// mapEntries returns all [key, value] pairs in the map as an array of arrays
// Usage: dhoro entries = map_entries(myMap);
func mapEntries(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "map_entries() expects 1 argument: map"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_entries() first argument must be a Map"}
	}

	entries := make([]object.Object, 0, len(m.Order))
	for _, keyHash := range m.Order {
		entry := &object.Array{
			Elements: []object.Object{
				m.Keys[keyHash],
				m.Pairs[keyHash],
			},
		}
		entries = append(entries, entry)
	}

	return &object.Array{Elements: entries}
}

// mapForEach iterates over all entries in the map
// Usage: map_foreach(myMap, kaj(value, key) { dekho(key, "=>", value); });
func mapForEach(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "map_foreach() expects 2 arguments: map, callback"}
	}

	m, ok := args[0].(*object.ES6Map)
	if !ok {
		return &object.Error{Message: "map_foreach() first argument must be a Map"}
	}

	callback, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: "map_foreach() second argument must be a function"}
	}

	// Need eval function (will be set by evaluator)
	if evalFunc == nil {
		return &object.Error{Message: "Internal error: eval function not set"}
	}

	for _, keyHash := range m.Order {
		key := m.Keys[keyHash]
		value := m.Pairs[keyHash]
		result := evalFunc(callback, []object.Object{value, key})
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return object.NULL
}
