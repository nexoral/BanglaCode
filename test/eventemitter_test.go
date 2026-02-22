package test

import (
	"BanglaCode/src/object"
	"strings"
	"testing"
)

// TestEventEmitterCreation tests creating an EventEmitter
func TestEventEmitterCreation(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	emitter
	`

	result := testEval(input)
	if result.Type() != object.EVENT_EMITTER_OBJ {
		t.Errorf("Expected EventEmitter object, got %s", result.Type())
	}
}

// TestEventEmitterOn tests adding event listeners
func TestEventEmitterOn(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro count = 0;
	
	ghotona_shuno(emitter, "test", kaj() {
		count = count + 1;
	});
	
	ghotona_prokash(emitter, "test");
	ghotona_prokash(emitter, "test");
	
	count
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 2 {
		t.Errorf("Expected count to be 2, got: %v", result.Inspect())
	}
}

// TestEventEmitterOnce tests once event listeners
func TestEventEmitterOnce(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro count = 0;
	
	ghotona_ekbar(emitter, "test", kaj() {
		count = count + 1;
	});
	
	ghotona_prokash(emitter, "test");
	ghotona_prokash(emitter, "test");
	ghotona_prokash(emitter, "test");
	
	count
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 1 {
		t.Errorf("Expected count to be 1 (once listener), got: %v", result.Inspect())
	}
}

// TestEventEmitterEmitWithData tests emitting events with data
func TestEventEmitterEmitWithData(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro receivedData = khali;
	
	ghotona_shuno(emitter, "data", kaj(data) {
		receivedData = data;
	});
	
	ghotona_prokash(emitter, "data", "Hello World");
	
	receivedData
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "Hello World" {
		t.Errorf("Expected 'Hello World', got: %v", result.Inspect())
	}
}

// TestEventEmitterMultipleData tests emitting with multiple data arguments
func TestEventEmitterMultipleData(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro sum = 0;
	
	ghotona_shuno(emitter, "add", kaj(a, b, c) {
		sum = a + b + c;
	});
	
	ghotona_prokash(emitter, "add", 10, 20, 30);
	
	sum
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 60 {
		t.Errorf("Expected sum to be 60, got: %v", result.Inspect())
	}
}

// TestEventEmitterMultipleListeners tests multiple listeners for same event
func TestEventEmitterMultipleListeners(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro count1 = 0;
	dhoro count2 = 0;
	
	ghotona_shuno(emitter, "test", kaj() {
		count1 = count1 + 1;
	});
	
	ghotona_shuno(emitter, "test", kaj() {
		count2 = count2 + 1;
	});
	
	ghotona_prokash(emitter, "test");
	
	count1 + count2
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 2 {
		t.Errorf("Expected total count to be 2 (both listeners called), got: %v", result.Inspect())
	}
}

// TestEventEmitterRemoveListener tests removing specific listeners
func TestEventEmitterRemoveListener(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro count = 0;
	
	dhoro listener = kaj() {
		count = count + 1;
	};
	
	ghotona_shuno(emitter, "test", listener);
	ghotona_prokash(emitter, "test");
	
	ghotona_bondho(emitter, "test", listener);
	ghotona_prokash(emitter, "test");
	
	count
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 1 {
		t.Errorf("Expected count to be 1 (listener removed after first call), got: %v", result.Inspect())
	}
}

// TestEventEmitterRemoveAllListeners tests removing all listeners
func TestEventEmitterRemoveAllListeners(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro count = 0;
	
	ghotona_shuno(emitter, "test", kaj() {
		count = count + 1;
	});
	
	ghotona_shuno(emitter, "test", kaj() {
		count = count + 1;
	});
	
	ghotona_prokash(emitter, "test");
	
	ghotona_sob_bondho(emitter, "test");
	ghotona_prokash(emitter, "test");
	
	count
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 2 {
		t.Errorf("Expected count to be 2 (listeners called once then removed), got: %v", result.Inspect())
	}
}

// TestEventEmitterGetListeners tests getting listeners for an event
func TestEventEmitterGetListeners(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	
	ghotona_shuno(emitter, "test", kaj() { });
	ghotona_shuno(emitter, "test", kaj() { });
	
	dhoro listeners = ghotona_shrotara(emitter, "test");
	dorghyo(listeners)
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 2 {
		t.Errorf("Expected 2 listeners, got: %v", result.Inspect())
	}
}

// TestEventEmitterGetEventNames tests getting all event names
func TestEventEmitterGetEventNames(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	
	ghotona_shuno(emitter, "event1", kaj() { });
	ghotona_shuno(emitter, "event2", kaj() { });
	ghotona_shuno(emitter, "event3", kaj() { });
	
	dhoro events = ghotona_naam_sob(emitter);
	dorghyo(events)
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 3 {
		t.Errorf("Expected 3 event names, got: %v", result.Inspect())
	}
}

// TestEventEmitterDifferentEvents tests multiple different events
func TestEventEmitterDifferentEvents(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro result = "";
	
	ghotona_shuno(emitter, "start", kaj() {
		result = result + "started ";
	});
	
	ghotona_shuno(emitter, "end", kaj() {
		result = result + "ended";
	});
	
	ghotona_prokash(emitter, "start");
	ghotona_prokash(emitter, "end");
	
	result
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "started ended" {
		t.Errorf("Expected 'started ended', got: %v", result.Inspect())
	}
}

// TestEventEmitterNoListeners tests emitting event with no listeners
func TestEventEmitterNoListeners(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	ghotona_prokash(emitter, "nonexistent")
	`

	result := testEval(input)
	if result != object.TRUE {
		t.Errorf("Expected sotti (true) for successful emit with no listeners, got: %v", result.Inspect())
	}
}

// TestEventEmitterChaining tests method chaining
func TestEventEmitterChaining(t *testing.T) {
	input := `
	dhoro emitter = ghotona_srishti();
	dhoro count = 0;
	
	ghotona_shuno(
		ghotona_shuno(emitter, "test1", kaj() {
			count = count + 1;
		}),
		"test2",
		kaj() {
			count = count + 2;
		}
	);
	
	ghotona_prokash(emitter, "test1");
	ghotona_prokash(emitter, "test2");
	
	count
	`

	result := testEval(input)
	num, ok := result.(*object.Number)
	if !ok || num.Value != 3 {
		t.Errorf("Expected count to be 3 (chained listeners), got: %v", result.Inspect())
	}
}

// TestEventEmitterRealWorldExample tests a real-world usage scenario
func TestEventEmitterRealWorldExample(t *testing.T) {
	input := `
	// Simulate a simple message bus
	dhoro messageBus = ghotona_srishti();
	dhoro messages = [];
	
	// Subscribe to messages
	ghotona_shuno(messageBus, "message", kaj(msg) {
		dhokao(messages, msg);
	});
	
	// Publish messages
	ghotona_prokash(messageBus, "message", "Hello");
	ghotona_prokash(messageBus, "message", "World");
	ghotona_prokash(messageBus, "message", "!");
	
	joro(messages, " ")
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || !strings.Contains(str.Value, "Hello") || !strings.Contains(str.Value, "World") {
		t.Errorf("Expected message with Hello and World, got: %v", result.Inspect())
	}
}
