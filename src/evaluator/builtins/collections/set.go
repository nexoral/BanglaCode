package collections

import (
	"crypto/sha256"
	"fmt"

	"BanglaCode/src/object"
)

// SetBuiltins contains all Set-related built-in functions
var SetBuiltins = map[string]*object.Builtin{
	"set_srishti": {Fn: setCreate},
	"set_add":     {Fn: setAdd},
	"set_has":     {Fn: setHas},
	"set_delete":  {Fn: setDelete},
	"set_clear":   {Fn: setClear},
	"set_akar":    {Fn: setSize},
	"set_values":  {Fn: setValues},
	"set_foreach": {Fn: setForEach},
}

// hashObject creates a hash string for any object to use as set key
func hashObject(obj object.Object) string {
	hash := sha256.New()
	hash.Write([]byte(obj.Type()))
	hash.Write([]byte(":"))
	hash.Write([]byte(obj.Inspect()))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// setCreate creates a new Set
// Usage: dhoro mySet = set_srishti();
//
//	dhoro mySet = set_srishti([1, 2, 3]);
func setCreate(args ...object.Object) object.Object {
	if len(args) > 1 {
		return &object.Error{Message: "set_srishti() expects 0 or 1 argument (optional array)"}
	}

	set := &object.Set{
		Elements: make(map[string]bool),
		Order:    []object.Object{},
	}

	// If array provided, add all elements
	if len(args) == 1 {
		if arr, ok := args[0].(*object.Array); ok {
			for _, elem := range arr.Elements {
				hash := hashObject(elem)
				if !set.Elements[hash] {
					set.Elements[hash] = true
					set.Order = append(set.Order, elem)
				}
			}
		} else {
			return &object.Error{Message: "set_srishti() argument must be an array"}
		}
	}

	return set
}

// setAdd adds an element to the set
// Usage: set_add(mySet, 42);
func setAdd(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "set_add() expects 2 arguments: set, element"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_add() first argument must be a Set"}
	}

	elem := args[1]
	hash := hashObject(elem)

	// Only add if not already present
	if !set.Elements[hash] {
		set.Elements[hash] = true
		set.Order = append(set.Order, elem)
	}

	return set
}

// setHas checks if an element exists in the set
// Usage: jodi (set_has(mySet, 42)) { ... }
func setHas(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "set_has() expects 2 arguments: set, element"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_has() first argument must be a Set"}
	}

	elem := args[1]
	hash := hashObject(elem)

	if set.Elements[hash] {
		return object.TRUE
	}
	return object.FALSE
}

// setDelete removes an element from the set
// Usage: set_delete(mySet, 42);
func setDelete(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "set_delete() expects 2 arguments: set, element"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_delete() first argument must be a Set"}
	}

	elem := args[1]
	hash := hashObject(elem)

	if set.Elements[hash] {
		delete(set.Elements, hash)
		// Remove from order array
		for i, orderedElem := range set.Order {
			if hashObject(orderedElem) == hash {
				set.Order = append(set.Order[:i], set.Order[i+1:]...)
				break
			}
		}
		return object.TRUE
	}
	return object.FALSE
}

// setClear removes all elements from the set
// Usage: set_clear(mySet);
func setClear(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "set_clear() expects 1 argument: set"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_clear() first argument must be a Set"}
	}

	set.Elements = make(map[string]bool)
	set.Order = []object.Object{}

	return object.NULL
}

// setSize returns the number of elements in the set
// Usage: dhoro count = set_akar(mySet);
func setSize(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "set_akar() expects 1 argument: set"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_akar() first argument must be a Set"}
	}

	return &object.Number{Value: float64(len(set.Order))}
}

// setValues returns all values in the set as an array
// Usage: dhoro values = set_values(mySet);
func setValues(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "set_values() expects 1 argument: set"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_values() first argument must be a Set"}
	}

	return &object.Array{Elements: set.Order}
}

// setForEach iterates over all elements in the set
// Usage: set_foreach(mySet, kaj(value) { dekho(value); });
func setForEach(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "set_foreach() expects 2 arguments: set, callback"}
	}

	set, ok := args[0].(*object.Set)
	if !ok {
		return &object.Error{Message: "set_foreach() first argument must be a Set"}
	}

	callback, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: "set_foreach() second argument must be a function"}
	}

	// Need eval function (will be set by evaluator)
	if evalFunc == nil {
		return &object.Error{Message: "Internal error: eval function not set"}
	}

	for _, elem := range set.Order {
		result := evalFunc(callback, []object.Object{elem})
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return object.NULL
}

// evalFunc is set by the evaluator to allow calling functions
var evalFunc func(*object.Function, []object.Object) object.Object
