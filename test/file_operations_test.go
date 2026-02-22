package test

import (
	"os"
	"testing"
	"time"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

func evalFileInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

// Test file append
func TestFileAppend(t *testing.T) {
	testFile := "/tmp/test_append.txt"
	defer os.Remove(testFile)

	input := `
		lekho("/tmp/test_append.txt", "First line\n");
		file_jog("/tmp/test_append.txt", "Second line\n");
		file_jog("/tmp/test_append.txt", "Third line\n");
		poro("/tmp/test_append.txt")
	`

	result := evalFileInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	expected := "First line\\nSecond line\\nThird line\\n"
	if str.Value != expected {
		t.Errorf("Expected %q, got %q", expected, str.Value)
	}
}

// Test file delete
func TestFileDelete(t *testing.T) {
	testFile := "/tmp/test_delete.txt"
	
	// Create file
	os.WriteFile(testFile, []byte("test content"), 0644)

	input := `
		file_mochho("/tmp/test_delete.txt")
	`

	result := evalFileInput(input)
	if result != object.TRUE {
		t.Errorf("Expected TRUE, got %s", result.Inspect())
	}

	// Verify file was deleted
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Error("File should have been deleted")
	}
}

// Test file copy
func TestFileCopy(t *testing.T) {
	srcFile := "/tmp/test_copy_src.txt"
	dstFile := "/tmp/test_copy_dst.txt"
	defer os.Remove(srcFile)
	defer os.Remove(dstFile)

	input := `
		lekho("/tmp/test_copy_src.txt", "Content to copy");
		file_nokol("/tmp/test_copy_src.txt", "/tmp/test_copy_dst.txt");
		poro("/tmp/test_copy_dst.txt")
	`

	result := evalFileInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "Content to copy" {
		t.Errorf("Expected 'Content to copy', got '%s'", str.Value)
	}
}

// Test folder delete (empty)
func TestFolderDeleteEmpty(t *testing.T) {
	testDir := "/tmp/test_folder_empty"
	os.Mkdir(testDir, 0755)

	input := `folder_mochho("/tmp/test_folder_empty")`

	result := evalFileInput(input)
	if result != object.TRUE {
		t.Errorf("Expected TRUE, got %s", result.Inspect())
	}

	// Verify folder was deleted
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Error("Folder should have been deleted")
	}
}

// Test folder delete (recursive)
func TestFolderDeleteRecursive(t *testing.T) {
	testDir := "/tmp/test_folder_recursive"
	os.MkdirAll(testDir+"/subdir", 0755)
	os.WriteFile(testDir+"/file.txt", []byte("test"), 0644)
	os.WriteFile(testDir+"/subdir/file2.txt", []byte("test2"), 0644)

	input := `folder_mochho("/tmp/test_folder_recursive", sotti)`

	result := evalFileInput(input)
	if result != object.TRUE {
		t.Errorf("Expected TRUE, got %s", result.Inspect())
	}

	// Verify folder and contents were deleted
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Error("Folder and contents should have been deleted")
	}
}

// Test file watching
func TestFileWatch(t *testing.T) {
	testFile := "/tmp/test_watch.txt"
	os.WriteFile(testFile, []byte("initial"), 0644)
	defer os.Remove(testFile)

	input := `
		dhoro changed = mittha;
		dhoro watcher = file_dekhun("/tmp/test_watch.txt", kaj(event, filename) {
			changed = sotti;
		});
		watcher
	`

	result := evalFileInput(input)
	watcher, ok := result.(*object.Map)
	if !ok {
		t.Fatalf("Expected Map (watcher), got %T", result)
	}

	// Check watcher has path
	if path, ok := watcher.Pairs["path"]; ok {
		if str, ok := path.(*object.String); ok {
			if str.Value != testFile {
				t.Errorf("Expected path '%s', got '%s'", testFile, str.Value)
			}
		}
	}

	// Check watcher is active
	if active, ok := watcher.Pairs["active"]; ok {
		if b, ok := active.(*object.Boolean); ok {
			if !b.Value {
				t.Error("Watcher should be active")
			}
		}
	}

	// Modify file to trigger watcher
	time.Sleep(500 * time.Millisecond)
	os.WriteFile(testFile, []byte("modified"), 0644)

	// Wait for watcher to detect change
	time.Sleep(2 * time.Second)

	// Stop watcher
	stopInput := `file_dekhun_bondho(watcher)`
	evalFileInput(stopInput)

	// Note: We can't easily test the callback execution in this test
	// because it runs in a goroutine and modifies a variable in a different environment
}

// Test append to non-existent file (creates new)
func TestFileAppendCreateNew(t *testing.T) {
	testFile := "/tmp/test_append_new.txt"
	defer os.Remove(testFile)

	input := `
		file_jog("/tmp/test_append_new.txt", "New content");
		poro("/tmp/test_append_new.txt")
	`

	result := evalFileInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "New content" {
		t.Errorf("Expected 'New content', got '%s'", str.Value)
	}
}

// Test multiple appends
func TestMultipleAppends(t *testing.T) {
	testFile := "/tmp/test_multi_append.txt"
	defer os.Remove(testFile)

	input := `
		file_jog("/tmp/test_multi_append.txt", "Line 1\n");
		file_jog("/tmp/test_multi_append.txt", "Line 2\n");
		file_jog("/tmp/test_multi_append.txt", "Line 3\n");
		file_jog("/tmp/test_multi_append.txt", "Line 4\n");
		poro("/tmp/test_multi_append.txt")
	`

	result := evalFileInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	expected := "Line 1\\nLine 2\\nLine 3\\nLine 4\\n"
	if str.Value != expected {
		t.Errorf("Expected %q, got %q", expected, str.Value)
	}
}

// Test copy large file
func TestCopyLargeFile(t *testing.T) {
	srcFile := "/tmp/test_large_src.txt"
	dstFile := "/tmp/test_large_dst.txt"
	defer os.Remove(srcFile)
	defer os.Remove(dstFile)

	// Create large content
	largeContent := ""
	for i := 0; i < 1000; i++ {
		largeContent += "This is line " + string(rune(i%10+'0')) + "\n"
	}

	// Write large file
	os.WriteFile(srcFile, []byte(largeContent), 0644)

	input := `
		file_nokol("/tmp/test_large_src.txt", "/tmp/test_large_dst.txt");
		poro("/tmp/test_large_dst.txt")
	`

	result := evalFileInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if len(str.Value) != len(largeContent) {
		t.Errorf("Expected length %d, got %d", len(largeContent), len(str.Value))
	}
}

// Test real-world scenario: Log file management
func TestLogFileManagement(t *testing.T) {
	logFile := "/tmp/test_app.log"
	backupFile := "/tmp/test_app.log.backup"
	defer os.Remove(logFile)
	defer os.Remove(backupFile)

	input := `
		// Write initial logs
		lekho("/tmp/test_app.log", "2026-02-22 12:00:00 - App started\n");
		file_jog("/tmp/test_app.log", "2026-02-22 12:00:01 - User logged in\n");
		file_jog("/tmp/test_app.log", "2026-02-22 12:00:02 - Data processed\n");

		// Backup log file
		file_nokol("/tmp/test_app.log", "/tmp/test_app.log.backup");

		// Clear current log and start fresh
		lekho("/tmp/test_app.log", "2026-02-22 12:01:00 - New session\n");

		// Read both files
		dhoro current = poro("/tmp/test_app.log");
		dhoro backup = poro("/tmp/test_app.log.backup");
		
		[current, backup]
	`

	result := evalFileInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T", result)
	}

	if len(arr.Elements) != 2 {
		t.Fatalf("Expected 2 elements, got %d", len(arr.Elements))
	}

	current := arr.Elements[0].(*object.String).Value
	backup := arr.Elements[1].(*object.String).Value

	if current != "2026-02-22 12:01:00 - New session\\n" {
		t.Errorf("Unexpected current log: %q (expected %q)", current, "2026-02-22 12:01:00 - New session\\n")
	}

	if !containsSubstr(backup, "App started") || !containsSubstr(backup, "User logged in") || !containsSubstr(backup, "Data processed") {
		t.Errorf("Backup log missing expected entries: %s", backup)
	}
}

func containsSubstr(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || containsSubstr(s[1:], substr)))
}
