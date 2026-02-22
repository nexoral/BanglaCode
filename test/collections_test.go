package test

import (
	"testing"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

// TestSetCreate tests creating a new Set
func TestSetCreate(t *testing.T) {
	input := `
		dhoro mySet = set_srishti();
		mySet
	`

	result := evalInput(input)
	set, ok := result.(*object.Set)
	if !ok {
		t.Fatalf("Expected Set, got %T (%+v)", result, result)
	}

	if len(set.Order) != 0 {
		t.Errorf("Expected empty set, got %d elements", len(set.Order))
	}
}

// TestSetCreateWithArray tests creating a Set from an array
func TestSetCreateWithArray(t *testing.T) {
	input := `
		dhoro mySet = set_srishti([1, 2, 3, 2, 1]);
		set_akar(mySet)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 3 {
		t.Errorf("Expected set size 3 (duplicates removed), got %f", num.Value)
	}
}

// TestSetAdd tests adding elements to a Set
func TestSetAdd(t *testing.T) {
	input := `
		dhoro mySet = set_srishti();
		set_add(mySet, 1);
		set_add(mySet, 2);
		set_add(mySet, 3);
		set_add(mySet, 2); // Duplicate - should not be added
		set_akar(mySet)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 3 {
		t.Errorf("Expected set size 3 (duplicate not added), got %f", num.Value)
	}
}

// TestSetHas tests checking if an element exists in a Set
func TestSetHas(t *testing.T) {
	input := `
		dhoro mySet = set_srishti([1, 2, 3]);
		set_has(mySet, 2)
	`

	result := evalInput(input)
	if result != object.TRUE {
		t.Errorf("Expected TRUE for element that exists, got %v", result)
	}

	input2 := `
		dhoro mySet = set_srishti([1, 2, 3]);
		set_has(mySet, 5)
	`

	result2 := evalInput(input2)
	if result2 != object.FALSE {
		t.Errorf("Expected FALSE for element that doesn't exist, got %v", result2)
	}
}

// TestSetDelete tests removing elements from a Set
func TestSetDelete(t *testing.T) {
	input := `
		dhoro mySet = set_srishti([1, 2, 3]);
		set_delete(mySet, 2);
		set_akar(mySet)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 2 {
		t.Errorf("Expected set size 2 after deletion, got %f", num.Value)
	}

	// Verify element was actually removed
	input2 := `
		dhoro mySet = set_srishti([1, 2, 3]);
		set_delete(mySet, 2);
		set_has(mySet, 2)
	`

	result2 := evalInput(input2)
	if result2 != object.FALSE {
		t.Errorf("Expected FALSE for deleted element, got %v", result2)
	}
}

// TestSetClear tests clearing all elements from a Set
func TestSetClear(t *testing.T) {
	input := `
		dhoro mySet = set_srishti([1, 2, 3]);
		set_clear(mySet);
		set_akar(mySet)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 0 {
		t.Errorf("Expected set size 0 after clear, got %f", num.Value)
	}
}

// TestSetValues tests getting all values from a Set
func TestSetValues(t *testing.T) {
	input := `
		dhoro mySet = set_srishti([1, 2, 3]);
		dhoro values = set_values(mySet);
		values
	`

	result := evalInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T (%+v)", result, result)
	}

	if len(arr.Elements) != 3 {
		t.Errorf("Expected 3 values, got %d", len(arr.Elements))
	}
}

// TestSetForEach tests iterating over Set elements
func TestSetForEach(t *testing.T) {
	input := `
		dhoro sum = 0;
		dhoro mySet = set_srishti([1, 2, 3]);
		set_foreach(mySet, kaj(value) {
			sum = sum + value;
		});
		sum
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 6 {
		t.Errorf("Expected sum 6, got %f", num.Value)
	}
}

// TestSetWithStrings tests Set with string elements
func TestSetWithStrings(t *testing.T) {
	input := `
		dhoro mySet = set_srishti(["apple", "banana", "apple"]);
		set_akar(mySet)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 2 {
		t.Errorf("Expected set size 2 (duplicate string removed), got %f", num.Value)
	}
}

// TestMapCreate tests creating a new Map
func TestMapCreate(t *testing.T) {
	input := `
		dhoro myMap = map_srishti();
		myMap
	`

	result := evalInput(input)
	m, ok := result.(*object.ES6Map)
	if !ok {
		t.Fatalf("Expected ES6Map, got %T (%+v)", result, result)
	}

	if len(m.Order) != 0 {
		t.Errorf("Expected empty map, got %d entries", len(m.Order))
	}
}

// TestMapCreateWithEntries tests creating a Map with initial entries
func TestMapCreateWithEntries(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		map_akar(myMap)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 2 {
		t.Errorf("Expected map size 2, got %f", num.Value)
	}
}

// TestMapSetAndGet tests setting and getting map entries
func TestMapSetAndGet(t *testing.T) {
	input := `
		dhoro myMap = map_srishti();
		map_set(myMap, "name", "Ankan");
		map_set(myMap, "age", 25);
		map_get(myMap, "name")
	`

	result := evalInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T (%+v)", result, result)
	}

	if str.Value != "Ankan" {
		t.Errorf("Expected 'Ankan', got %s", str.Value)
	}
}

// TestMapHas tests checking if a key exists in a Map
func TestMapHas(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"]]);
		map_has(myMap, "name")
	`

	result := evalInput(input)
	if result != object.TRUE {
		t.Errorf("Expected TRUE for key that exists, got %v", result)
	}

	input2 := `
		dhoro myMap = map_srishti([["name", "Ankan"]]);
		map_has(myMap, "age")
	`

	result2 := evalInput(input2)
	if result2 != object.FALSE {
		t.Errorf("Expected FALSE for key that doesn't exist, got %v", result2)
	}
}

// TestMapDelete tests removing entries from a Map
func TestMapDelete(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		map_delete(myMap, "age");
		map_akar(myMap)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 1 {
		t.Errorf("Expected map size 1 after deletion, got %f", num.Value)
	}

	// Verify entry was actually removed
	input2 := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		map_delete(myMap, "age");
		map_has(myMap, "age")
	`

	result2 := evalInput(input2)
	if result2 != object.FALSE {
		t.Errorf("Expected FALSE for deleted key, got %v", result2)
	}
}

// TestMapClear tests clearing all entries from a Map
func TestMapClear(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		map_clear(myMap);
		map_akar(myMap)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 0 {
		t.Errorf("Expected map size 0 after clear, got %f", num.Value)
	}
}

// TestMapKeys tests getting all keys from a Map
func TestMapKeys(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		dhoro keys = map_keys(myMap);
		keys
	`

	result := evalInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T (%+v)", result, result)
	}

	if len(arr.Elements) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(arr.Elements))
	}
}

// TestMapValues tests getting all values from a Map
func TestMapValues(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		dhoro values = map_values(myMap);
		values
	`

	result := evalInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T (%+v)", result, result)
	}

	if len(arr.Elements) != 2 {
		t.Errorf("Expected 2 values, got %d", len(arr.Elements))
	}
}

// TestMapEntries tests getting all [key, value] pairs from a Map
func TestMapEntries(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"], ["age", 25]]);
		dhoro entries = map_entries(myMap);
		entries
	`

	result := evalInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T (%+v)", result, result)
	}

	if len(arr.Elements) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(arr.Elements))
	}

	// Each entry should be a 2-element array
	for i, entry := range arr.Elements {
		entryArr, ok := entry.(*object.Array)
		if !ok {
			t.Errorf("Entry %d should be an array, got %T", i, entry)
			continue
		}
		if len(entryArr.Elements) != 2 {
			t.Errorf("Entry %d should have 2 elements, got %d", i, len(entryArr.Elements))
		}
	}
}

// TestMapForEach tests iterating over Map entries
func TestMapForEach(t *testing.T) {
	input := `
		dhoro sum = 0;
		dhoro myMap = map_srishti([["a", 1], ["b", 2], ["c", 3]]);
		map_foreach(myMap, kaj(value, key) {
			sum = sum + value;
		});
		sum
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 6 {
		t.Errorf("Expected sum 6, got %f", num.Value)
	}
}

// TestMapWithObjectKeys tests Map with object keys (not just strings)
func TestMapWithObjectKeys(t *testing.T) {
	input := `
		dhoro myMap = map_srishti();
		dhoro key1 = [1, 2];
		dhoro key2 = [3, 4];
		map_set(myMap, key1, "first");
		map_set(myMap, key2, "second");
		map_akar(myMap)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 2 {
		t.Errorf("Expected map size 2 with object keys, got %f", num.Value)
	}
}

// TestMapUpdateValue tests updating an existing map entry
func TestMapUpdateValue(t *testing.T) {
	input := `
		dhoro myMap = map_srishti([["name", "Ankan"]]);
		map_set(myMap, "name", "Saha");
		map_get(myMap, "name")
	`

	result := evalInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T (%+v)", result, result)
	}

	if str.Value != "Saha" {
		t.Errorf("Expected updated value 'Saha', got %s", str.Value)
	}

	// Size should still be 1 (updated, not added)
	input2 := `
		dhoro myMap = map_srishti([["name", "Ankan"]]);
		map_set(myMap, "name", "Saha");
		map_akar(myMap)
	`

	result2 := evalInput(input2)
	num, ok := result2.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result2, result2)
	}

	if num.Value != 1 {
		t.Errorf("Expected map size 1 after update, got %f", num.Value)
	}
}

// TestSetComplexTypes tests Set with complex types (arrays, maps)
func TestSetComplexTypes(t *testing.T) {
	input := `
		dhoro mySet = set_srishti();
		set_add(mySet, [1, 2, 3]);
		set_add(mySet, [1, 2, 3]); // Same content - should not be added
		set_add(mySet, [4, 5, 6]); // Different - should be added
		set_akar(mySet)
	`

	result := evalInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T (%+v)", result, result)
	}

	if num.Value != 2 {
		t.Errorf("Expected set size 2 (duplicate array not added), got %f", num.Value)
	}
}

// Helper function to evaluate input code
func evalInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return evaluator.Eval(program, env)
}
