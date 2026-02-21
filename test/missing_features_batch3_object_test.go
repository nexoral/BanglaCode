package test

import "testing"

func TestObjectMaturityBuiltins(t *testing.T) {
	testBooleanObject(t, testEval(`nijer_ache({a: 1}, "a")`), true)
	testBooleanObject(t, testEval(`nijer_ache({a: 1}, "b")`), false)

	out1 := testEval(`jora_theke([["a", 1], ["b", 2]])["b"]`)
	testNumberObject(t, out1, 2)

	testBooleanObject(t, testEval(`ekoi_ki(1, 1)`), true)
	testBooleanObject(t, testEval(`ekoi_ki(1, 2)`), false)

	out2 := testEval(`notun_map({a: 1}, {b: 2})["b"]`)
	testNumberObject(t, out2, 2)

	out3 := testEval(`joma({x: 1})["x"]`)
	testNumberObject(t, out3, 1)
}
