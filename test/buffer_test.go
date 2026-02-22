package test

import (
	"BanglaCode/src/object"
	"strings"
	"testing"
)

// TestBufferCreation tests creating buffers
func TestBufferCreation(t *testing.T) {
	input := `
	dhoro buf = buffer_banao(10);
	buf
	`

	result := testEval(input)
	if result.Type() != object.BUFFER_OBJ {
		t.Errorf("Expected Buffer object, got %s", result.Type())
	}

	buf := result.(*object.Buffer)
	if len(buf.Data) != 10 {
		t.Errorf("Expected buffer length 10, got %d", len(buf.Data))
	}
}

// TestBufferFromString tests creating buffer from string
func TestBufferFromString(t *testing.T) {
	input := `
	dhoro buf = buffer_theke("Hello");
	buffer_text(buf)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Hello" {
		t.Errorf("Expected 'Hello', got %v", result.Inspect())
	}
}

// TestBufferFromArray tests creating buffer from byte array
func TestBufferFromArray(t *testing.T) {
	input := `
	dhoro buf = buffer_theke([72, 101, 108, 108, 111]);
	buffer_text(buf)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Hello" {
		t.Errorf("Expected 'Hello', got %v", result.Inspect())
	}
}

// TestBufferConcat tests concatenating buffers
func TestBufferConcat(t *testing.T) {
	input := `
	dhoro buf1 = buffer_theke("Hello ");
	dhoro buf2 = buffer_theke("World");
	dhoro combined = buffer_joro(buf1, buf2);
	buffer_text(combined)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Hello World" {
		t.Errorf("Expected 'Hello World', got %v", result.Inspect())
	}
}

// TestBufferWrite tests writing to buffer
func TestBufferWrite(t *testing.T) {
	input := `
	dhoro buf = buffer_banao(20);
	buffer_lekho(buf, "Hello", 0);
	buffer_lekho(buf, "World", 6);
	buffer_text(buf, "utf8")
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	// Check if contains both Hello and World
	if !strings.Contains(str.Value, "Hello") || !strings.Contains(str.Value, "World") {
		t.Errorf("Expected buffer to contain 'Hello' and 'World', got %v", str.Value)
	}
}

// TestBufferSlice tests slicing buffers
func TestBufferSlice(t *testing.T) {
	input := `
	dhoro buf = buffer_theke("Hello World");
	dhoro slice = buffer_angsho(buf, 0, 5);
	buffer_text(slice)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Hello" {
		t.Errorf("Expected 'Hello', got %v", result.Inspect())
	}
}

// TestBufferSliceFromOffset tests slicing from offset to end
func TestBufferSliceFromOffset(t *testing.T) {
	input := `
	dhoro buf = buffer_theke("Hello World");
	dhoro slice = buffer_angsho(buf, 6);
	buffer_text(slice)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "World" {
		t.Errorf("Expected 'World', got %v", result.Inspect())
	}
}

// TestBufferCompare tests comparing buffers
func TestBufferCompare(t *testing.T) {
	input := `
	dhoro buf1 = buffer_theke("ABC");
	dhoro buf2 = buffer_theke("ABC");
	dhoro buf3 = buffer_theke("DEF");
	
	dhoro result1 = buffer_tulona(buf1, buf2);
	dhoro result2 = buffer_tulona(buf1, buf3);
	dhoro result3 = buffer_tulona(buf3, buf1);
	
	[result1, result2, result3]
	`

	result := testEval(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Errorf("Expected array, got %s", result.Type())
		return
	}

	// result1 should be 0 (equal)
	if num, ok := arr.Elements[0].(*object.Number); !ok || num.Value != 0 {
		t.Errorf("Expected 0 for equal buffers, got %v", arr.Elements[0].Inspect())
	}

	// result2 should be -1 (buf1 < buf3)
	if num, ok := arr.Elements[1].(*object.Number); !ok || num.Value != -1 {
		t.Errorf("Expected -1 for buf1 < buf3, got %v", arr.Elements[1].Inspect())
	}

	// result3 should be 1 (buf3 > buf1)
	if num, ok := arr.Elements[2].(*object.Number); !ok || num.Value != 1 {
		t.Errorf("Expected 1 for buf3 > buf1, got %v", arr.Elements[2].Inspect())
	}
}

// TestBufferToHex tests converting buffer to hex
func TestBufferToHex(t *testing.T) {
	input := `
	dhoro buf = buffer_theke([255, 0, 127]);
	buffer_hex(buf)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "ff007f" {
		t.Errorf("Expected 'ff007f', got %v", result.Inspect())
	}
}

// TestBufferCopy tests copying between buffers
func TestBufferCopy(t *testing.T) {
	input := `
	dhoro target = buffer_banao(10);
	dhoro source = buffer_theke("Hello");
	
	dhoro written = buffer_copy(target, source, 0);
	dhoro text = buffer_text(target);
	
	written
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 5 {
		t.Errorf("Expected 5 bytes written, got %v", result.Inspect())
	}
}

// TestBufferMultipleConcat tests concatenating multiple buffers
func TestBufferMultipleConcat(t *testing.T) {
	input := `
	dhoro buf1 = buffer_theke("One ");
	dhoro buf2 = buffer_theke("Two ");
	dhoro buf3 = buffer_theke("Three");
	dhoro combined = buffer_joro(buf1, buf2, buf3);
	buffer_text(combined)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "One Two Three" {
		t.Errorf("Expected 'One Two Three', got %v", result.Inspect())
	}
}

// TestBufferBinaryData tests working with binary data
func TestBufferBinaryData(t *testing.T) {
	input := `
	// Create buffer with specific byte values
	dhoro buf = buffer_theke([0, 1, 2, 3, 4, 5]);
	
	// Slice it
	dhoro slice = buffer_angsho(buf, 2, 5);
	
	// Get the sliced buffer and check if it's valid
	slice
	`

	result := testEval(input)
	// The result should be a Buffer object
	if result.Type() != object.BUFFER_OBJ {
		t.Errorf("Expected Buffer object, got %s: %v", result.Type(), result.Inspect())
	}
}

// TestBufferWriteOffset tests writing at specific offset
func TestBufferWriteOffset(t *testing.T) {
	input := `
	dhoro buf = buffer_banao(15);
	buffer_lekho(buf, "Hello", 0);
	buffer_lekho(buf, "World", 10);
	
	dhoro slice1 = buffer_angsho(buf, 0, 5);
	dhoro slice2 = buffer_angsho(buf, 10, 15);
	
	buffer_text(slice1) + " " + buffer_text(slice2)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Errorf("Expected string, got %s", result.Type())
		return
	}

	if !strings.Contains(str.Value, "Hello") || !strings.Contains(str.Value, "World") {
		t.Errorf("Expected result with Hello and World, got %v", str.Value)
	}
}

// TestBufferRealWorldExample tests a real-world buffer usage
func TestBufferRealWorldExample(t *testing.T) {
	input := `
	// Simulate building a protocol message
	dhoro header = buffer_theke([255, 1, 2]);  // Magic + version + type
	dhoro payload = buffer_theke("Hello Protocol");
	dhoro checksum = buffer_theke([171, 205]);
	
	// Combine all parts
	dhoro message = buffer_joro(header, payload, checksum);
	
	// Verify structure
	dhoro headerPart = buffer_angsho(message, 0, 3);
	dhoro payloadPart = buffer_angsho(message, 3, 17);
	
	buffer_text(payloadPart)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Hello Protocol" {
		t.Errorf("Expected 'Hello Protocol', got %v", result.Inspect())
	}
}
