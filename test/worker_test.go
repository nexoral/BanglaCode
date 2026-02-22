package test

import (
	"BanglaCode/src/object"
	"testing"
	"time"
)

// TestWorkerCreation tests creating a worker
func TestWorkerCreation(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		dekho("Worker running");
	});
	worker
	`

	result := testEval(input)
	if result.Type() != object.WORKER_OBJ {
		t.Errorf("Expected Worker object, got %s", result.Type())
	}

	worker := result.(*object.Worker)
	worker.Mu.RLock()
	isRunning := worker.IsRunning
	worker.Mu.RUnlock()

	if !isRunning {
		t.Errorf("Expected worker to be running")
	}
}

// TestWorkerWithData tests creating worker with initial data
func TestWorkerWithData(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		dekho(kaj_kormi_tothya);
	}, 42);
	worker
	`

	result := testEval(input)
	if result.Type() != object.WORKER_OBJ {
		t.Errorf("Expected Worker object, got %s", result.Type())
	}
}

// TestWorkerTerminate tests terminating a worker
func TestWorkerTerminate(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		dekho("Worker started");
	});
	
	kaj_kormi_bondho(worker);
	
	// Give time for worker to terminate
	process_ghum(100);
	
	worker
	`

	result := testEval(input)
	worker, ok := result.(*object.Worker)
	if !ok {
		t.Errorf("Expected Worker object, got %s", result.Type())
		return
	}

	// Wait a bit for termination
	time.Sleep(200 * time.Millisecond)

	worker.Mu.RLock()
	isRunning := worker.IsRunning
	worker.Mu.RUnlock()

	if isRunning {
		t.Errorf("Expected worker to be terminated")
	}
}

// TestWorkerPostMessage tests sending messages to worker
func TestWorkerPostMessage(t *testing.T) {
	input := `
	dhoro received = [];
	
	dhoro worker = kaj_kormi_srishti(kaj() {
		// Worker receives messages
	});
	
	// Send message to worker
	kaj_kormi_pathao(worker, "Hello Worker");
	kaj_kormi_pathao(worker, 42);
	
	// Give time for processing
	process_ghum(100);
	
	kaj_kormi_bondho(worker);
	
	"sent"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "sent" {
		t.Errorf("Expected 'sent', got %v", result.Inspect())
	}
}

// TestWorkerOnMessage tests setting message handler
func TestWorkerOnMessage(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		dekho("Worker running");
	});
	
	// Set up message handler
	kaj_kormi_shuno(worker, kaj(data) {
		dekho("Received:", data);
	});
	
	// Give time for handler setup
	process_ghum(50);
	
	kaj_kormi_bondho(worker);
	
	"handler_set"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "handler_set" {
		t.Errorf("Expected 'handler_set', got %v", result.Inspect())
	}
}

// TestWorkerBidirectional tests sending messages to worker
func TestWorkerBidirectional(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		dekho("Worker initialized");
	});
	
	// Send messages to worker
	kaj_kormi_pathao(worker, "Message 1");
	kaj_kormi_pathao(worker, "Message 2");
	
	// Give time for processing
	process_ghum(200);
	
	kaj_kormi_bondho(worker);
	
	"messages_sent"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "messages_sent" {
		t.Errorf("Expected 'messages_sent', got %v", result.Inspect())
	}
}

// TestWorkerComputation tests worker performing computation
func TestWorkerComputation(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		// Perform computation
		dhoro sum = 0;
		ghuriye (dhoro i = 1; i <= 10; i = i + 1) {
			sum = sum + i;
		}
		dekho("Sum:", sum);
	});
	
	// Wait for computation
	process_ghum(300);
	
	kaj_kormi_bondho(worker);
	
	"computed"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "computed" {
		t.Errorf("Expected 'computed', got %v", result.Inspect())
	}
}

// TestMultipleWorkers tests creating multiple workers
func TestMultipleWorkers(t *testing.T) {
	input := `
	dhoro worker1 = kaj_kormi_srishti(kaj() {
		dekho("Worker 1 running");
	});
	
	dhoro worker2 = kaj_kormi_srishti(kaj() {
		dekho("Worker 2 running");
	});
	
	// Wait for workers to start
	process_ghum(200);
	
	kaj_kormi_bondho(worker1);
	kaj_kormi_bondho(worker2);
	
	"completed"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "completed" {
		t.Errorf("Expected 'completed', got %v", result.Inspect())
	}
}

// TestWorkerWithComplexData tests passing complex initial data
func TestWorkerWithComplexData(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		// Access initial data
		dhoro config = kaj_kormi_tothya;
		dekho("Count:", config["count"]);
	}, {"count": 5});
	
	process_ghum(200);
	
	kaj_kormi_bondho(worker);
	
	"complex_data_passed"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "complex_data_passed" {
		t.Errorf("Expected 'complex_data_passed', got %v", result.Inspect())
	}
}

// TestWorkerErrorHandling tests error cases
func TestWorkerErrorHandling(t *testing.T) {
	// Test sending message to non-worker
	input1 := `kaj_kormi_pathao("not a worker", "data");`
	result1 := testEval(input1)
	if result1.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for non-worker, got %s", result1.Type())
	}

	// Test creating worker without function
	input2 := `kaj_kormi_srishti("not a function");`
	result2 := testEval(input2)
	if result2.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for non-function, got %s", result2.Type())
	}

	// Test terminating non-worker
	input3 := `kaj_kormi_bondho(42);`
	result3 := testEval(input3)
	if result3.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for non-worker termination, got %s", result3.Type())
	}
}

// TestWorkerLongRunning tests worker that runs for extended time
func TestWorkerLongRunning(t *testing.T) {
	input := `
	dhoro worker = kaj_kormi_srishti(kaj() {
		// Perform multiple operations over time
		ghuriye (dhoro i = 0; i < 3; i = i + 1) {
			dekho("Iteration", i);
			process_ghum(100);
		}
	});
	
	// Wait for all operations
	process_ghum(500);
	
	kaj_kormi_bondho(worker);
	
	"long_running_completed"
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok || str.Value != "long_running_completed" {
		t.Errorf("Expected 'long_running_completed', got %v", result.Inspect())
	}
}
