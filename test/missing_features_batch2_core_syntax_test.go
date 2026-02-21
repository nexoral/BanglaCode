package test

import "testing"

func TestDoWhileLoop(t *testing.T) {
	input := `
	dhoro x = 0;
	do {
		x = x + 1;
	} jotokkhon (x < 3);
	x;
	`
	testNumberObject(t, testEval(input), 3)
}

func TestDoWhileRunsAtLeastOnce(t *testing.T) {
	input := `
	dhoro x = 0;
	do {
		x = x + 1;
	} jotokkhon (mittha);
	x;
	`
	testNumberObject(t, testEval(input), 1)
}

func TestInOperator(t *testing.T) {
	testBooleanObject(t, testEval(`"a" in {a: 1, b: 2}`), true)
	testBooleanObject(t, testEval(`"z" in {a: 1, b: 2}`), false)
	testBooleanObject(t, testEval(`1 in [10, 20, 30]`), true)
	testBooleanObject(t, testEval(`3 in [10, 20, 30]`), false)
	testBooleanObject(t, testEval(`2 in "bangla"`), true)
}

func TestInstanceofOperator(t *testing.T) {
	input := `
	sreni Manush {
		shuru(naam) { ei.naam = naam; }
	}
	dhoro p = notun Manush("Ankan");
	p instanceof Manush;
	`
	testBooleanObject(t, testEval(input), true)
}

func TestDeleteOperator(t *testing.T) {
	input := `
	dhoro obj = {a: 1, b: 2};
	delete obj.a;
	"a" in obj;
	`
	testBooleanObject(t, testEval(input), false)
}

func TestDeleteArrayIndex(t *testing.T) {
	input := `
	dhoro arr = [1, 2, 3];
	delete arr[1];
	arr[1] == khali;
	`
	testBooleanObject(t, testEval(input), true)
}
