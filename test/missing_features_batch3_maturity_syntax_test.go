package test

import "testing"

func TestArrowFunctionSingleParam(t *testing.T) {
	input := `
	dhoro double = x => x * 2;
	double(5);
	`
	testNumberObject(t, testEval(input), 10)
}

func TestArrowFunctionBlockBody(t *testing.T) {
	input := `
	dhoro inc = x => { ferao x + 1; };
	inc(9);
	`
	testNumberObject(t, testEval(input), 10)
}

func TestForOfLoopArray(t *testing.T) {
	input := `
	dhoro sum = 0;
	ghuriye (x of [1, 2, 3, 4]) {
		sum = sum + x;
	}
	sum;
	`
	testNumberObject(t, testEval(input), 10)
}

func TestForInLoopMap(t *testing.T) {
	input := `
	dhoro obj = {a: 1, b: 2};
	dhoro count = 0;
	ghuriye (k in obj) {
		count = count + 1;
	}
	count;
	`
	testNumberObject(t, testEval(input), 2)
}
