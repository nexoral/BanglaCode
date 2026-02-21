package test

import "BanglaCode/src/object"
import "testing"

func TestDestructuringDeclarations(t *testing.T) {
	inputArray := `
	dhoro [a, b, c] = [10, 20];
	a + b;
	`
	testNumberObject(t, testEval(inputArray), 30)

	testBooleanObject(t, testEval(`dhoro [x, y] = [1]; y == khali;`), true)
	testNumberObject(t, testEval(`dhoro {a, b} = {a: 7, b: 9}; a + b;`), 16)
	testBooleanObject(t, testEval(`dhoro {a, b} = {a: 1}; b == khali;`), true)
}

func TestMultiParamArrowFunctions(t *testing.T) {
	testNumberObject(t, testEval(`dhoro add = (a, b) => a + b; add(3, 4);`), 7)
	testNumberObject(t, testEval(`dhoro z = () => 11; z();`), 11)
}

func TestTimerBuiltins(t *testing.T) {
	inputTimeout := `
	dhoro x = 0;
	setTimeout(kaj() { x = 5; }, 10);
	ghum(40);
	x;
	`
	testNumberObject(t, testEval(inputTimeout), 5)

	inputInterval := `
	dhoro c = 0;
	dhoro id = setInterval(kaj() { c = c + 1; }, 5);
	ghum(35);
	clearInterval(id);
	dhoro prev = c;
	ghum(20);
	c == prev;
	`
	testBooleanObject(t, testEval(inputInterval), true)
}

func TestRegexFlagsAndWrappers(t *testing.T) {
	testBooleanObject(t, testEval(`regex_test("bangla", "BANGLA", "i")`), true)
	testNumberObject(t, testEval(`search("BANGLA CODE", "code", "i")`), 7)

	out := testEval(`matchAll("a1 b2", "[a-z][0-9]")`)
	arr, ok := out.(*object.Array)
	if !ok {
		t.Fatalf("matchAll expected ARRAY, got=%T", out)
	}
	if len(arr.Elements) != 2 {
		t.Fatalf("matchAll expected 2 matches, got=%d", len(arr.Elements))
	}
}
