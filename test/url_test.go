package test

import (
	"BanglaCode/src/object"
	"testing"
)

// TestURLParse tests basic URL parsing
func TestURLParse(t *testing.T) {
	input := `
	dhoro url = url_parse("https://example.com:8080/path?query=value#fragment");
	url
	`

	result := testEval(input)
	urlObj, ok := result.(*object.URL)
	if !ok {
		t.Errorf("Expected URL object, got %s", result.Type())
		return
	}

	if urlObj.Protocol != "https:" {
		t.Errorf("Expected protocol 'https:', got '%s'", urlObj.Protocol)
	}
	if urlObj.Hostname != "example.com" {
		t.Errorf("Expected hostname 'example.com', got '%s'", urlObj.Hostname)
	}
	if urlObj.Port != "8080" {
		t.Errorf("Expected port '8080', got '%s'", urlObj.Port)
	}
	if urlObj.Pathname != "/path" {
		t.Errorf("Expected pathname '/path', got '%s'", urlObj.Pathname)
	}
	if urlObj.Search != "?query=value" {
		t.Errorf("Expected search '?query=value', got '%s'", urlObj.Search)
	}
	if urlObj.Hash != "#fragment" {
		t.Errorf("Expected hash '#fragment', got '%s'", urlObj.Hash)
	}
}

// TestURLParseSimple tests simple URL without port
func TestURLParseSimple(t *testing.T) {
	input := `
	dhoro url = url_parse("http://example.com/api/users");
	url.Hostname
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "example.com" {
		t.Errorf("Expected 'example.com', got %v", result.Inspect())
	}
}

// TestURLParseWithAuth tests URL with username and password
func TestURLParseWithAuth(t *testing.T) {
	input := `
	dhoro url = url_parse("https://user:pass@example.com/path");
	url.Username + ":" + url.Password
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "user:pass" {
		t.Errorf("Expected 'user:pass', got %v", result.Inspect())
	}
}

// TestURLQueryParams tests creating URLSearchParams from string
func TestURLQueryParams(t *testing.T) {
	input := `
	dhoro params = url_query_params("key1=value1&key2=value2");
	params
	`

	result := testEval(input)
	if result.Type() != object.URL_PARAMS_OBJ {
		t.Errorf("Expected URLSearchParams, got %s", result.Type())
	}
}

// TestURLQueryParamsFromURL tests creating URLSearchParams from URL
func TestURLQueryParamsFromURL(t *testing.T) {
	input := `
	dhoro url = url_parse("https://example.com?name=John&age=30");
	dhoro params = url_query_params(url);
	url_query_get(params, "name")
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "John" {
		t.Errorf("Expected 'John', got %v", result.Inspect())
	}
}

// TestURLQueryGet tests getting query parameter
func TestURLQueryGet(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice&age=25");
	url_query_get(params, "name")
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Alice" {
		t.Errorf("Expected 'Alice', got %v", result.Inspect())
	}
}

// TestURLQueryGetNonExistent tests getting non-existent parameter
func TestURLQueryGetNonExistent(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice");
	url_query_get(params, "age")
	`

	result := testEval(input)
	if result != object.NULL {
		t.Errorf("Expected null, got %v", result.Inspect())
	}
}

// TestURLQuerySet tests setting query parameter
func TestURLQuerySet(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice");
	url_query_set(params, "age", "30");
	url_query_get(params, "age")
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "30" {
		t.Errorf("Expected '30', got %v", result.Inspect())
	}
}

// TestURLQueryAppend tests appending query parameter
func TestURLQueryAppend(t *testing.T) {
	input := `
	dhoro params = url_query_params("tag=javascript");
	url_query_append(params, "tag", "nodejs");
	url_query_toString(params)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	// Should contain both tags (order may vary)
	if !contains(str.Value, "tag=javascript") || !contains(str.Value, "tag=nodejs") {
		t.Errorf("Expected both tags, got '%s'", str.Value)
	}
}

// TestURLQueryDelete tests deleting query parameter
func TestURLQueryDelete(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice&age=30");
	url_query_delete(params, "age");
	url_query_has(params, "age")
	`

	result := testEval(input)
	boolean, ok := result.(*object.Boolean)
	if !ok || boolean.Value {
		t.Errorf("Expected false, got %v", result.Inspect())
	}
}

// TestURLQueryHas tests checking if parameter exists
func TestURLQueryHas(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice&age=30");
	url_query_has(params, "name")
	`

	result := testEval(input)
	boolean, ok := result.(*object.Boolean)
	if !ok || !boolean.Value {
		t.Errorf("Expected true, got %v", result.Inspect())
	}
}

// TestURLQueryKeys tests getting all keys
func TestURLQueryKeys(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice&age=30");
	dhoro keys = url_query_keys(params);
	dorghyo(keys)
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 2 {
		t.Errorf("Expected 2 keys, got %v", result.Inspect())
	}
}

// TestURLQueryValues tests getting all values
func TestURLQueryValues(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice&age=30");
	dhoro values = url_query_values(params);
	dorghyo(values)
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 2 {
		t.Errorf("Expected 2 values, got %v", result.Inspect())
	}
}

// TestURLQueryToString tests converting params to string
func TestURLQueryToString(t *testing.T) {
	input := `
	dhoro params = url_query_params("name=Alice&age=30");
	url_query_toString(params)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	// Should contain both parameters (order may vary)
	if !contains(str.Value, "name=Alice") || !contains(str.Value, "age=30") {
		t.Errorf("Expected both parameters, got '%s'", str.Value)
	}
}

// TestURLRealWorldExample tests a real-world URL parsing scenario
func TestURLRealWorldExample(t *testing.T) {
	input := `
	// Parse GitHub API URL
	dhoro url = url_parse("https://api.github.com/repos/owner/repo?page=2&per_page=50");
	
	// Extract components
	dhoro protocol = url.Protocol;
	dhoro host = url.Hostname;
	dhoro path = url.Pathname;
	
	// Parse query parameters
	dhoro params = url_query_params(url);
	dhoro page = url_query_get(params, "page");
	dhoro perPage = url_query_get(params, "per_page");
	
	// Verify (protocol already includes ":")
	protocol + "//" + host + path + "?page=" + page
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	expected := "https://api.github.com/repos/owner/repo?page=2"
	if str.Value != expected {
		t.Errorf("Expected '%s', got '%s'", expected, str.Value)
	}
}

// TestURLModifyParams tests modifying query parameters
func TestURLModifyParams(t *testing.T) {
	input := `
	dhoro url = url_parse("https://example.com?old=value");
	dhoro params = url_query_params(url);
	
	// Modify params
	url_query_delete(params, "old");
	url_query_set(params, "new", "data");
	url_query_append(params, "extra", "info");
	
	// Convert back to string
	url_query_toString(params)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	// Should not contain "old" but should contain "new" and "extra"
	if contains(str.Value, "old=value") {
		t.Errorf("Should not contain 'old=value', got '%s'", str.Value)
	}
	if !contains(str.Value, "new=data") || !contains(str.Value, "extra=info") {
		t.Errorf("Should contain 'new=data' and 'extra=info', got '%s'", str.Value)
	}
}

// TestURLInvalidURL tests error handling for invalid URLs
func TestURLInvalidURL(t *testing.T) {
	input := `
	url_parse("not a valid url ://")
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for invalid URL, got %s", result.Type())
	}
}

// TestURLComplexQuery tests complex query string with special characters
func TestURLComplexQuery(t *testing.T) {
	input := `
	dhoro params = url_query_params("search=hello world&filter=a+b");
	url_query_has(params, "search")
	`

	result := testEval(input)
	boolean, ok := result.(*object.Boolean)
	if !ok || !boolean.Value {
		t.Errorf("Expected true, got %v", result.Inspect())
	}
}

// TestURLMultipleValues tests multiple values for same key
func TestURLMultipleValues(t *testing.T) {
	input := `
	dhoro params = url_query_params("");
	url_query_append(params, "color", "red");
	url_query_append(params, "color", "blue");
	url_query_append(params, "color", "green");
	url_query_toString(params)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	// Should contain all three colors
	if !contains(str.Value, "color=red") || !contains(str.Value, "color=blue") || !contains(str.Value, "color=green") {
		t.Errorf("Expected all three colors, got '%s'", str.Value)
	}
}
