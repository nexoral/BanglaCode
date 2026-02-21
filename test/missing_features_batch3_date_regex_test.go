package test

import "BanglaCode/src/object"
import "testing"

func TestDateBuiltins(t *testing.T) {
	now := testEval(`tarikh_ekhon()`)
	if _, ok := now.(*object.Number); !ok {
		t.Fatalf("tarikh_ekhon should return NUMBER, got=%T", now)
	}

	testBooleanObject(t, testEval(`sonkhya_na(tarikh_parse("invalid date"))`), true)
	testBooleanObject(t, testEval(`sonkhya_na(tarikh_parse("2026-02-21"))`), false)

	out := testEval(`dhoro t = tarikh_parse("2026-02-21"); tarikh_format(t, "2006-01-02")`)
	testStringObject(t, out, "2026-02-21")
}

func TestRegexBuiltins(t *testing.T) {
	testBooleanObject(t, testEval(`regex_test("[a-z]+", "bangla")`), true)
	testNumberObject(t, testEval(`regex_search("la", "bangla")`), 4)
	testStringObject(t, testEval(`regex_replace("a", "banana", "x")`), "bxnxnx")

	match := testEval(`regex_match("b(ang)", "bangla")`)
	arr, ok := match.(*object.Array)
	if !ok {
		t.Fatalf("regex_match should return ARRAY, got=%T", match)
	}
	if len(arr.Elements) != 2 {
		t.Fatalf("regex_match expected 2 captures, got=%d", len(arr.Elements))
	}
}
