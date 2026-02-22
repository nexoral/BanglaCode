package test

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

// Helper function to evaluate input code
func evalPathInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

// TestPathResolve tests path_resolve function
func TestPathResolve(t *testing.T) {
	input := `path_resolve(".")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Should return an absolute path
	if !filepath.IsAbs(str.Value) {
		t.Errorf("Expected absolute path, got %s", str.Value)
	}
}

// TestPathResolveMultiple tests path_resolve with multiple paths
func TestPathResolveMultiple(t *testing.T) {
	input := `path_resolve(".", "test", "file.txt")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Should contain "test" and "file.txt" in path
	if !strings.Contains(str.Value, "test") || !strings.Contains(str.Value, "file.txt") {
		t.Errorf("Expected path to contain 'test' and 'file.txt', got %s", str.Value)
	}
}

// TestPathNormalize tests path_normalize function
func TestPathNormalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`path_normalize("/foo/bar//baz")`, "/foo/bar/baz"},
		{`path_normalize("/foo/./bar")`, "/foo/bar"},
		{`path_normalize("/foo/../bar")`, "/bar"},
		{`path_normalize("./test")`, "test"},
	}

	for _, tt := range tests {
		result := evalPathInput(tt.input)
		str, ok := result.(*object.String)
		if !ok {
			t.Errorf("Expected String for %s, got %T", tt.input, result)
			continue
		}

		// Normalize the expected path for the current OS
		expected := filepath.Clean(strings.ReplaceAll(tt.expected, "/", string(filepath.Separator)))

		if str.Value != expected {
			t.Errorf("Expected %s = %s, got %s", tt.input, expected, str.Value)
		}
	}
}

// TestPathRelative tests path_relative function
func TestPathRelative(t *testing.T) {
	input := `path_relative("/home/user", "/home/user/projects/app")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Expected path depends on OS
	expected := filepath.Join("projects", "app")
	if str.Value != expected {
		t.Errorf("Expected relative path %s, got %s", expected, str.Value)
	}
}

// TestPathRelativeSameDir tests path_relative with same directory
func TestPathRelativeSameDir(t *testing.T) {
	input := `path_relative("/home/user", "/home/user")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "." {
		t.Errorf("Expected '.', got %s", str.Value)
	}
}

// TestPathRelativeParent tests path_relative going to parent
func TestPathRelativeParent(t *testing.T) {
	input := `path_relative("/home/user/projects", "/home/user")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != ".." {
		t.Errorf("Expected '..', got %s", str.Value)
	}
}

// TestPathSepConstant tests PATH_SEP constant
func TestPathSepConstant(t *testing.T) {
	input := `PATH_SEP`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Should match OS path separator
	expected := string(filepath.Separator)
	if str.Value != expected {
		t.Errorf("Expected PATH_SEP = %s, got %s", expected, str.Value)
	}
}

// TestPathDelimiterConstant tests PATH_DELIMITER constant
func TestPathDelimiterConstant(t *testing.T) {
	input := `PATH_DELIMITER`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Should be ":" on Unix, ";" on Windows
	if runtime.GOOS == "windows" {
		if str.Value != ";" {
			t.Errorf("Expected PATH_DELIMITER = ';', got %s", str.Value)
		}
	} else {
		if str.Value != ":" {
			t.Errorf("Expected PATH_DELIMITER = ':', got %s", str.Value)
		}
	}
}

// TestPathJoin tests path_joro function
func TestPathJoin(t *testing.T) {
	input := `path_joro("home", "user", "documents", "file.txt")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	expected := filepath.Join("home", "user", "documents", "file.txt")
	if str.Value != expected {
		t.Errorf("Expected %s, got %s", expected, str.Value)
	}
}

// TestPathDirname tests directory_naam function
func TestPathDirname(t *testing.T) {
	input := `directory_naam("/home/user/file.txt")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Convert to OS-specific path
	expected := filepath.Dir("/home/user/file.txt")
	if str.Value != expected {
		t.Errorf("Expected %s, got %s", expected, str.Value)
	}
}

// TestPathBasename tests path_naam function
func TestPathBasename(t *testing.T) {
	input := `path_naam("/home/user/file.txt")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "file.txt" {
		t.Errorf("Expected 'file.txt', got %s", str.Value)
	}
}

// TestPathExt tests file_ext function
func TestPathExt(t *testing.T) {
	input := `file_ext("/home/user/file.txt")`
	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != ".txt" {
		t.Errorf("Expected '.txt', got %s", str.Value)
	}
}

// TestPathCombination tests combining path operations
func TestPathCombination(t *testing.T) {
	input := `
		dhoro base = path_normalize("/home/user/./projects");
		dhoro filename = "app.txt";
		dhoro fullPath = path_joro(base, filename);
		fullPath
	`

	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	// Should contain normalized path
	expected := filepath.Join(filepath.Clean("/home/user/projects"), "app.txt")
	if str.Value != expected {
		t.Errorf("Expected %s, got %s", expected, str.Value)
	}
}

// TestPathConstantsInCode tests using path constants in code
func TestPathConstantsInCode(t *testing.T) {
	input := `
		dhoro parts = ["home", "user", "file.txt"];
		dhoro path = parts[0] + PATH_SEP + parts[1] + PATH_SEP + parts[2];
		path
	`

	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	expected := filepath.Join("home", "user", "file.txt")
	if str.Value != expected {
		t.Errorf("Expected %s, got %s", expected, str.Value)
	}
}

// TestPathResolveAbsolute tests path_resolve with absolute path
func TestPathResolveAbsolute(t *testing.T) {
	input := `
		dhoro base = path_resolve(".");
		dhoro relative = path_relative(base, base);
		relative
	`

	result := evalPathInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "." {
		t.Errorf("Expected '.', got %s", str.Value)
	}
}
