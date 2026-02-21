package test

import (
	"BanglaCode/src/object"
	"math"
	"testing"
)

func TestMissingBatchArrayFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`khojo_prothom([2,4,7,8], kaj(x) { ferao x > 5; })`, 7},
		{`khojo_index([2,4,7,8], kaj(x) { ferao x > 5; })`, 2},
		{`khojo_shesh([2,4,7,8], kaj(x) { ferao x % 2 == 0; })`, 8},
		{`khojo_shesh_index([2,4,7,8], kaj(x) { ferao x % 2 == 0; })`, 3},
		{`array_at([10,20,30], -1)`, 30},
		{`shesh_index_of([1,2,3,2,1], 2)`, 3},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestMissingBatchArrayBooleanFunctions(t *testing.T) {
	testBooleanObject(t, testEval(`prottek([2,4,6], kaj(x) { ferao x % 2 == 0; })`), true)
	testBooleanObject(t, testEval(`prottek([2,3,6], kaj(x) { ferao x % 2 == 0; })`), false)
	testBooleanObject(t, testEval(`kono([1,3,5], kaj(x) { ferao x % 2 == 0; })`), false)
	testBooleanObject(t, testEval(`kono([1,4,5], kaj(x) { ferao x % 2 == 0; })`), true)
}

func TestMissingBatchArrayFlatMap(t *testing.T) {
	input := `
	dhoro out = somtol_manchitro([1, 2, 3], kaj(x) {
		ferao [x, x * 10];
	});
	out;
	`
	evaluated := testEval(input)
	testArrayObject(t, evaluated, []float64{1, 10, 2, 20, 3, 30}, 0)
}

func TestMissingBatchArrayConcatFlatReduceRight(t *testing.T) {
	testArrayObject(t, testEval(`joro_array([1,2], [3], 4)`), []float64{1, 2, 3, 4}, 0)

	flatInput := `
	dhoro a = [1, [2, [3]]];
	somtol(a, 2);
	`
	testArrayObject(t, testEval(flatInput), []float64{1, 2, 3}, 0)

	testNumberObject(t, testEval(`sonkuchito_dan([1,2,3], kaj(acc, x) { ferao acc - x; }, 0)`), -6)
}

func TestMissingBatchStringFunctions(t *testing.T) {
	testBooleanObject(t, testEval(`ache_text("banglacode", "code")`), true)
	testBooleanObject(t, testEval(`shuru_diye("banglacode", "bang")`), true)
	testBooleanObject(t, testEval(`shesh_diye("banglacode", "code")`), true)

	s1 := testEval(`baro("ha", 3)`)
	testStringObject(t, s1, "hahaha")

	s2 := testEval(`agey_bhoro("7", 3, "0")`)
	testStringObject(t, s2, "007")

	s3 := testEval(`pichoney_bhoro("7", 3, "0")`)
	testStringObject(t, s3, "700")

	s4 := testEval(`okkhor("bangla", 2)`)
	testStringObject(t, s4, "n")

	s5 := testEval(`text_at("bangla", -1)`)
	testStringObject(t, s5, "a")

	s6 := testEval(`chhanto_shuru("   hi")`)
	testStringObject(t, s6, "hi")

	s7 := testEval(`chhanto_shesh("hi   ")`)
	testStringObject(t, s7, "hi")

	testNumberObject(t, testEval(`okkhor_code("A", 0)`), 65)
	testNumberObject(t, testEval(`shesh_khojo("banana", "na")`), 4)
	testNumberObject(t, testEval(`codepoint_at("A", 0)`), 65)
	testNumberObject(t, testEval(`tulona_text("a", "b")`), -1)
	testStringObject(t, testEval(`shadharon_text("text")`), "text")
}

func TestMissingBatchGlobalNumericFunctions(t *testing.T) {
	testNumberObject(t, testEval(`purno_sonkhya("42")`), 42)
	testNumberObject(t, testEval(`purno_sonkhya("0x10")`), 16)
	testNumberObject(t, testEval(`purno_sonkhya("111", 2)`), 7)
	testNumberObject(t, testEval(`doshomik_sonkhya("3.14")`), 3.14)

	testBooleanObject(t, testEval(`sonkhya_na("abc")`), true)
	testBooleanObject(t, testEval(`sonkhya_na("12.5")`), false)
	testBooleanObject(t, testEval(`sonkhya_shimito("12.5")`), true)
	testBooleanObject(t, testEval(`sonkhya_shimito("abc")`), false)
}

func TestMissingBatchNaNBehavior(t *testing.T) {
	out := testEval(`purno_sonkhya("abc")`)
	num, ok := out.(*object.Number)
	if !ok {
		t.Fatalf("expected Number for NaN case, got=%T", out)
	}
	if !math.IsNaN(num.Value) {
		t.Fatalf("expected NaN, got=%v", num.Value)
	}
}

func TestMissingBatchURIFunctions(t *testing.T) {
	testStringObject(t, testEval(`uri_ongsho_encode("hello world")`), "hello%20world")
	testStringObject(t, testEval(`uri_ongsho_decode("hello%20world")`), "hello world")
	testStringObject(t, testEval(`uri_encode("https://a.com/q=a b")`), "https://a.com/q=a%20b")
	testStringObject(t, testEval(`uri_decode("https://a.com/q=a%20b")`), "https://a.com/q=a b")
}
