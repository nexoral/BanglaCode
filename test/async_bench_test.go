package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"testing"
)

// Helper function to evaluate BanglaCode
func evalBanglaCode(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return evaluator.Eval(program, env)
}

// Benchmark async function creation
func BenchmarkAsyncFunctionCreation(b *testing.B) {
	input := `proyash kaj test() { ferao 42; }`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark async function execution
func BenchmarkAsyncFunctionExecution(b *testing.B) {
	input := `
		proyash kaj test() {
			ferao 42;
		}
		opekha test();
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark ghumaao (async sleep)
func BenchmarkGhumaao(b *testing.B) {
	input := `opekha ghumaao(1);` // 1ms sleep

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark sob_proyash with 3 promises (concurrent wait)
func BenchmarkPromiseAll_3Promises(b *testing.B) {
	input := `
		dhoro p1 = ghumaao(10);
		dhoro p2 = ghumaao(10);
		dhoro p3 = ghumaao(10);
		opekha sob_proyash([p1, p2, p3]);
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark sob_proyash with 10 promises (concurrent wait)
func BenchmarkPromiseAll_10Promises(b *testing.B) {
	input := `
		dhoro promises = [];
		ghuriye (dhoro i = 0; i < 10; i = i + 1) {
			dhokao(promises, ghumaao(5));
		}
		opekha sob_proyash(promises);
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark promise creation overhead
func BenchmarkPromiseCreation(b *testing.B) {
	input := `ghumaao(0);` // Create promise but don't await

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark nested async function calls
func BenchmarkNestedAsyncCalls(b *testing.B) {
	input := `
		proyash kaj inner() {
			ferao 42;
		}
		proyash kaj outer() {
			dhoro result = opekha inner();
			ferao result;
		}
		opekha outer();
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark async file operations (if test file exists)
func BenchmarkAsyncFileRead(b *testing.B) {
	// Create a small test file first
	input := `
		lekho("test_async_bench.txt", "Hello, BanglaCode!");
		opekha poro_async("test_async_bench.txt");
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}

// Benchmark multiple concurrent async operations
func BenchmarkConcurrentAsyncOps(b *testing.B) {
	input := `
		proyash kaj task1() {
			opekha ghumaao(1);
			ferao 1;
		}
		proyash kaj task2() {
			opekha ghumaao(1);
			ferao 2;
		}
		proyash kaj task3() {
			opekha ghumaao(1);
			ferao 3;
		}

		dhoro p1 = task1();
		dhoro p2 = task2();
		dhoro p3 = task3();
		opekha sob_proyash([p1, p2, p3]);
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalBanglaCode(input)
	}
}
