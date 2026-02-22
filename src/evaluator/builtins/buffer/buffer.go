package buffer

import (
	"BanglaCode/src/object"
	"encoding/hex"
	"fmt"
	"strings"
)

// Builtins exports all buffer-related built-in functions
var Builtins = map[string]*object.Builtin{
	"buffer_banao":  {Fn: createBuffer},
	"buffer_theke":  {Fn: createBufferFrom},
	"buffer_joro":   {Fn: concatBuffers},
	"buffer_text":   {Fn: bufferToString},
	"buffer_lekho":  {Fn: writeToBuffer},
	"buffer_angsho": {Fn: sliceBuffer},
	"buffer_tulona": {Fn: compareBuffers},
	"buffer_hex":    {Fn: bufferToHex},
	"buffer_copy":   {Fn: copyBuffer},
}

// createBuffer creates a new buffer with specified size
// Usage: dhoro buf = buffer_banao(10);  // Buffer of 10 bytes
func createBuffer(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "buffer_banao() expects 1 argument (size)"}
	}

	// Get size
	size, ok := args[0].(*object.Number)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("size must be number, got %s", args[0].Type())}
	}

	if size.Value < 0 {
		return &object.Error{Message: "buffer size cannot be negative"}
	}

	return object.CreateBuffer(int(size.Value))
}

// createBufferFrom creates a buffer from string, array, or another buffer
// Usage: dhoro buf = buffer_theke("Hello");      // From string
//
//	dhoro buf = buffer_theke([72, 101]);     // From byte array
func createBufferFrom(args ...object.Object) object.Object {
	if len(args) == 0 {
		return &object.Error{Message: "buffer_theke() expects at least 1 argument"}
	}

	switch arg := args[0].(type) {
	case *object.String:
		// Create buffer from string (UTF-8 encoding)
		return object.CreateBufferFrom([]byte(arg.Value))

	case *object.Array:
		// Create buffer from array of numbers (bytes)
		bytes := make([]byte, len(arg.Elements))
		for i, elem := range arg.Elements {
			num, ok := elem.(*object.Number)
			if !ok {
				return &object.Error{Message: fmt.Sprintf("array element %d must be number, got %s", i, elem.Type())}
			}
			if num.Value < 0 || num.Value > 255 {
				return &object.Error{Message: fmt.Sprintf("byte value must be 0-255, got %g", num.Value)}
			}
			bytes[i] = byte(num.Value)
		}
		return object.CreateBufferFrom(bytes)

	case *object.Buffer:
		// Create a copy of existing buffer
		arg.Mu.RLock()
		defer arg.Mu.RUnlock()
		return object.CreateBufferFrom(arg.Data)

	default:
		return &object.Error{Message: fmt.Sprintf("cannot create buffer from %s", arg.Type())}
	}
}

// concatBuffers concatenates multiple buffers
// Usage: dhoro combined = buffer_joro(buf1, buf2, buf3);
func concatBuffers(args ...object.Object) object.Object {
	if len(args) == 0 {
		return &object.Error{Message: "buffer_joro() expects at least 1 argument"}
	}

	// Calculate total size
	totalSize := 0
	for _, arg := range args {
		buf, ok := arg.(*object.Buffer)
		if !ok {
			return &object.Error{Message: fmt.Sprintf("all arguments must be buffers, got %s", arg.Type())}
		}
		buf.Mu.RLock()
		totalSize += len(buf.Data)
		buf.Mu.RUnlock()
	}

	// Allocate combined buffer
	combined := make([]byte, totalSize)
	offset := 0

	// Copy all buffers
	for _, arg := range args {
		buf := arg.(*object.Buffer)
		buf.Mu.RLock()
		copy(combined[offset:], buf.Data)
		offset += len(buf.Data)
		buf.Mu.RUnlock()
	}

	return object.CreateBufferFrom(combined)
}

// bufferToString converts buffer to string
// Usage: dhoro text = buffer_text(buf);              // UTF-8 (default)
//
//	dhoro text = buffer_text(buf, "utf8");      // UTF-8
//	dhoro text = buffer_text(buf, "hex");       // Hexadecimal
//	dhoro text = buffer_text(buf, "base64");    // Base64
func bufferToString(args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return &object.Error{Message: "buffer_text() expects 1 or 2 arguments (buffer, [encoding])"}
	}

	// Get buffer
	buf, ok := args[0].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be buffer, got %s", args[0].Type())}
	}

	// Get encoding (default: utf8)
	encoding := "utf8"
	if len(args) == 2 {
		enc, ok := args[1].(*object.String)
		if !ok {
			return &object.Error{Message: fmt.Sprintf("encoding must be string, got %s", args[1].Type())}
		}
		encoding = strings.ToLower(enc.Value)
	}

	buf.Mu.RLock()
	defer buf.Mu.RUnlock()

	switch encoding {
	case "utf8", "utf-8":
		return &object.String{Value: string(buf.Data)}

	case "hex":
		return &object.String{Value: hex.EncodeToString(buf.Data)}

	case "base64":
		// Note: We already have base64_encode function, but include here for compatibility
		return &object.String{Value: fmt.Sprintf("%x", buf.Data)}

	default:
		return &object.Error{Message: fmt.Sprintf("unsupported encoding: %s", encoding)}
	}
}

// writeToBuffer writes string or data to buffer at specified offset
// Usage: buffer_lekho(buf, "Hello", 0);     // Write at offset 0
//
//	buffer_lekho(buf, "World", 6);     // Write at offset 6
func writeToBuffer(args ...object.Object) object.Object {
	if len(args) < 2 || len(args) > 3 {
		return &object.Error{Message: "buffer_lekho() expects 2 or 3 arguments (buffer, data, [offset])"}
	}

	// Get buffer
	buf, ok := args[0].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be buffer, got %s", args[0].Type())}
	}

	// Get data to write
	var dataToWrite []byte
	switch data := args[1].(type) {
	case *object.String:
		dataToWrite = []byte(data.Value)
	case *object.Array:
		dataToWrite = make([]byte, len(data.Elements))
		for i, elem := range data.Elements {
			num, ok := elem.(*object.Number)
			if !ok {
				return &object.Error{Message: fmt.Sprintf("array element must be number, got %s", elem.Type())}
			}
			if num.Value < 0 || num.Value > 255 {
				return &object.Error{Message: fmt.Sprintf("byte value must be 0-255, got %g", num.Value)}
			}
			dataToWrite[i] = byte(num.Value)
		}
	default:
		return &object.Error{Message: fmt.Sprintf("data must be string or array, got %s", data.Type())}
	}

	// Get offset (default: 0)
	offset := 0
	if len(args) == 3 {
		off, ok := args[2].(*object.Number)
		if !ok {
			return &object.Error{Message: fmt.Sprintf("offset must be number, got %s", args[2].Type())}
		}
		offset = int(off.Value)
	}

	// Write to buffer (thread-safe)
	buf.Mu.Lock()
	defer buf.Mu.Unlock()

	if offset < 0 || offset >= len(buf.Data) {
		return &object.Error{Message: fmt.Sprintf("offset %d out of range [0, %d)", offset, len(buf.Data))}
	}

	// Copy as much as possible
	written := copy(buf.Data[offset:], dataToWrite)

	return &object.Number{Value: float64(written)}
}

// sliceBuffer returns a slice of the buffer
// Usage: dhoro slice = buffer_angsho(buf, 0, 5);  // Bytes 0-4
//
//	dhoro slice = buffer_angsho(buf, 5);     // From byte 5 to end
func sliceBuffer(args ...object.Object) object.Object {
	if len(args) < 2 || len(args) > 3 {
		return &object.Error{Message: "buffer_angsho() expects 2 or 3 arguments (buffer, start, [end])"}
	}

	// Get buffer
	buf, ok := args[0].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be buffer, got %s", args[0].Type())}
	}

	// Get start
	start, ok := args[1].(*object.Number)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("start must be number, got %s", args[1].Type())}
	}

	buf.Mu.RLock()
	defer buf.Mu.RUnlock()

	startIdx := int(start.Value)
	if startIdx < 0 {
		startIdx = 0
	}
	if startIdx >= len(buf.Data) {
		return object.CreateBufferFrom([]byte{})
	}

	// Get end (default: length)
	endIdx := len(buf.Data)
	if len(args) == 3 {
		end, ok := args[2].(*object.Number)
		if !ok {
			return &object.Error{Message: fmt.Sprintf("end must be number, got %s", args[2].Type())}
		}
		endIdx = int(end.Value)
		if endIdx > len(buf.Data) {
			endIdx = len(buf.Data)
		}
		if endIdx < startIdx {
			return object.CreateBufferFrom([]byte{})
		}
	}

	// Create new buffer with sliced data
	return object.CreateBufferFrom(buf.Data[startIdx:endIdx])
}

// compareBuffers compares two buffers
// Usage: dhoro result = buffer_tulona(buf1, buf2);
// Returns: -1 if buf1 < buf2, 0 if equal, 1 if buf1 > buf2
func compareBuffers(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "buffer_tulona() expects 2 arguments (buffer1, buffer2)"}
	}

	// Get buffers
	buf1, ok := args[0].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be buffer, got %s", args[0].Type())}
	}

	buf2, ok := args[1].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be buffer, got %s", args[1].Type())}
	}

	buf1.Mu.RLock()
	buf2.Mu.RLock()
	defer buf1.Mu.RUnlock()
	defer buf2.Mu.RUnlock()

	// Compare byte by byte
	minLen := len(buf1.Data)
	if len(buf2.Data) < minLen {
		minLen = len(buf2.Data)
	}

	for i := 0; i < minLen; i++ {
		if buf1.Data[i] < buf2.Data[i] {
			return &object.Number{Value: -1}
		}
		if buf1.Data[i] > buf2.Data[i] {
			return &object.Number{Value: 1}
		}
	}

	// If all bytes equal up to minLen, compare lengths
	if len(buf1.Data) < len(buf2.Data) {
		return &object.Number{Value: -1}
	}
	if len(buf1.Data) > len(buf2.Data) {
		return &object.Number{Value: 1}
	}

	return &object.Number{Value: 0}
}

// bufferToHex converts buffer to hex string
// Usage: dhoro hex = buffer_hex(buf);
func bufferToHex(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "buffer_hex() expects 1 argument (buffer)"}
	}

	buf, ok := args[0].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("argument must be buffer, got %s", args[0].Type())}
	}

	buf.Mu.RLock()
	defer buf.Mu.RUnlock()

	return &object.String{Value: hex.EncodeToString(buf.Data)}
}

// copyBuffer copies data from one buffer to another
// Usage: dhoro written = buffer_copy(target, source, [targetStart], [sourceStart], [sourceEnd]);
func copyBuffer(args ...object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "buffer_copy() expects at least 2 arguments (target, source)"}
	}

	// Get buffers
	target, ok := args[0].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be buffer, got %s", args[0].Type())}
	}

	source, ok := args[1].(*object.Buffer)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be buffer, got %s", args[1].Type())}
	}

	// Default values
	targetStart := 0
	sourceStart := 0

	source.Mu.RLock()
	sourceEnd := len(source.Data)
	source.Mu.RUnlock()

	// Parse optional arguments
	if len(args) >= 3 {
		ts, ok := args[2].(*object.Number)
		if !ok {
			return &object.Error{Message: "targetStart must be number"}
		}
		targetStart = int(ts.Value)
	}

	if len(args) >= 4 {
		ss, ok := args[3].(*object.Number)
		if !ok {
			return &object.Error{Message: "sourceStart must be number"}
		}
		sourceStart = int(ss.Value)
	}

	if len(args) >= 5 {
		se, ok := args[4].(*object.Number)
		if !ok {
			return &object.Error{Message: "sourceEnd must be number"}
		}
		sourceEnd = int(se.Value)
	}

	// Copy data (thread-safe)
	target.Mu.Lock()
	source.Mu.RLock()
	defer target.Mu.Unlock()
	defer source.Mu.RUnlock()

	if targetStart < 0 || targetStart >= len(target.Data) {
		return &object.Error{Message: "targetStart out of range"}
	}

	if sourceStart < 0 || sourceStart > len(source.Data) {
		return &object.Error{Message: "sourceStart out of range"}
	}

	if sourceEnd < sourceStart || sourceEnd > len(source.Data) {
		return &object.Error{Message: "sourceEnd out of range"}
	}

	written := copy(target.Data[targetStart:], source.Data[sourceStart:sourceEnd])

	return &object.Number{Value: float64(written)}
}
