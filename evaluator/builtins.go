package evaluator

import (
	"BanglaCode/object"
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// EvalFunc is a function pointer for evaluating AST nodes (set by evaluator.go to avoid circular dependency)
var EvalFunc func(handler *object.Function, args []object.Object) object.Object

var builtins = map[string]*object.Builtin{
	"dekho": {
		Fn: func(args ...object.Object) object.Object {
			for i, arg := range args {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(arg.Inspect())
			}
			fmt.Println()
			return object.NULL
		},
	},
	"length": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Number{Value: float64(len(arg.Value))}
			case *object.Array:
				return &object.Number{Value: float64(len(arg.Elements))}
			default:
				return newError("argument to `length` not supported, got %s", args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			arr.Elements = append(arr.Elements, args[1])
			return arr
		},
	},
	"pop": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				lastElement := arr.Elements[length-1]
				arr.Elements = arr.Elements[:length-1]
				return lastElement
			}
			return object.NULL
		},
	},
	"keys": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to `keys` must be MAP, got %s", args[0].Type())
			}

			mapObj := args[0].(*object.Map)
			keys := make([]object.Object, 0, len(mapObj.Pairs))
			for key := range mapObj.Pairs {
				keys = append(keys, &object.String{Value: key})
			}
			return &object.Array{Elements: keys}
		},
	},
	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: string(args[0].Type())}
		},
	},
	"string": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: args[0].Inspect()}
		},
	},
	"number": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Number:
				return arg
			case *object.String:
				var num float64
				_, err := fmt.Sscanf(arg.Value, "%f", &num)
				if err != nil {
					return newError("cannot convert string to number: %s", arg.Value)
				}
				return &object.Number{Value: num}
			case *object.Boolean:
				if arg.Value {
					return &object.Number{Value: 1}
				}
				return &object.Number{Value: 0}
			default:
				return newError("cannot convert %s to number", arg.Type())
			}
		},
	},
	// Math functions
	"sqrt": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `sqrt` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Sqrt(num)}
		},
	},
	"pow": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ || args[1].Type() != object.NUMBER_OBJ {
				return newError("arguments to `pow` must be NUMBERs")
			}
			base := args[0].(*object.Number).Value
			exp := args[1].(*object.Number).Value
			return &object.Number{Value: math.Pow(base, exp)}
		},
	},
	"floor": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `floor` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Floor(num)}
		},
	},
	"ceil": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `ceil` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Ceil(num)}
		},
	},
	"round": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `round` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Round(num)}
		},
	},
	"abs": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `abs` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Abs(num)}
		},
	},
	"min": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			minVal := math.Inf(1)
			for _, arg := range args {
				if arg.Type() != object.NUMBER_OBJ {
					return newError("all arguments to `min` must be NUMBERs")
				}
				val := arg.(*object.Number).Value
				if val < minVal {
					minVal = val
				}
			}
			return &object.Number{Value: minVal}
		},
	},
	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			maxVal := math.Inf(-1)
			for _, arg := range args {
				if arg.Type() != object.NUMBER_OBJ {
					return newError("all arguments to `max` must be NUMBERs")
				}
				val := arg.(*object.Number).Value
				if val > maxVal {
					maxVal = val
				}
			}
			return &object.Number{Value: maxVal}
		},
	},
	// String functions
	"upper": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `upper` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToUpper(str)}
		},
	},
	"lower": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `lower` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToLower(str)}
		},
	},
	"split": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("arguments to `split` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			sep := args[1].(*object.String).Value
			parts := strings.Split(str, sep)
			elements := make([]object.Object, len(parts))
			for i, p := range parts {
				elements[i] = &object.String{Value: p}
			}
			return &object.Array{Elements: elements}
		},
	},
	"join": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `join` must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `join` must be STRING, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			sep := args[1].(*object.String).Value
			parts := make([]string, len(arr.Elements))
			for i, el := range arr.Elements {
				parts[i] = el.Inspect()
			}
			return &object.String{Value: strings.Join(parts, sep)}
		},
	},
	"trim": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `trim` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.TrimSpace(str)}
		},
	},
	"indexOf": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("arguments to `indexOf` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			substr := args[1].(*object.String).Value
			return &object.Number{Value: float64(strings.Index(str, substr))}
		},
	},
	"substring": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `substring` must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `substring` must be NUMBER, got %s", args[1].Type())
			}
			str := args[0].(*object.String).Value
			start := int(args[1].(*object.Number).Value)
			end := len(str)
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `substring` must be NUMBER, got %s", args[2].Type())
				}
				end = int(args[2].(*object.Number).Value)
			}
			if start < 0 {
				start = 0
			}
			if end > len(str) {
				end = len(str)
			}
			if start > end {
				return &object.String{Value: ""}
			}
			return &object.String{Value: str[start:end]}
		},
	},
	"replace": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
				return newError("all arguments to `replace` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			old := args[1].(*object.String).Value
			new := args[2].(*object.String).Value
			return &object.String{Value: strings.ReplaceAll(str, old, new)}
		},
	},
	// Array functions
	"slice": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `slice` must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `slice` must be NUMBER, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			start := int(args[1].(*object.Number).Value)
			end := len(arr.Elements)
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `slice` must be NUMBER, got %s", args[2].Type())
				}
				end = int(args[2].(*object.Number).Value)
			}
			if start < 0 {
				start = 0
			}
			if end > len(arr.Elements) {
				end = len(arr.Elements)
			}
			if start > end {
				return &object.Array{Elements: []object.Object{}}
			}
			newElements := make([]object.Object, end-start)
			copy(newElements, arr.Elements[start:end])
			return &object.Array{Elements: newElements}
		},
	},
	"reverse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `reverse` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length)
			for i := 0; i < length; i++ {
				newElements[i] = arr.Elements[length-1-i]
			}
			return &object.Array{Elements: newElements}
		},
	},
	"includes": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `includes` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			target := args[1]
			for _, el := range arr.Elements {
				if objectsEqual(el, target) {
					return object.TRUE
				}
			}
			return object.FALSE
		},
	},
	"sort": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `sort` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			newElements := make([]object.Object, len(arr.Elements))
			copy(newElements, arr.Elements)
			sort.Slice(newElements, func(i, j int) bool {
				// Sort numbers numerically
				if newElements[i].Type() == object.NUMBER_OBJ && newElements[j].Type() == object.NUMBER_OBJ {
					return newElements[i].(*object.Number).Value < newElements[j].(*object.Number).Value
				}
				// Sort strings alphabetically
				return newElements[i].Inspect() < newElements[j].Inspect()
			})
			return &object.Array{Elements: newElements}
		},
	},
	// Utility functions
	"time": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Number{Value: float64(time.Now().UnixMilli())}
		},
	},
	"random": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Number{Value: rand.Float64()}
		},
	},
	"input": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 0 {
				fmt.Print(args[0].Inspect())
			}
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return &object.String{Value: ""}
				}
				return newError("error reading input: %s", err.Error())
			}
			return &object.String{Value: strings.TrimSpace(text)}
		},
	},
	"readFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `readFile` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content, err := os.ReadFile(path)
			if err != nil {
				return newError("error reading file: %s", err.Error())
			}
			return &object.String{Value: string(content)}
		},
	},
	"writeFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `writeFile` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content := args[1].Inspect()
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				return newError("error writing file: %s", err.Error())
			}
			return object.TRUE
		},
	},
	"exit": {
		Fn: func(args ...object.Object) object.Object {
			code := 0
			if len(args) > 0 && args[0].Type() == object.NUMBER_OBJ {
				code = int(args[0].(*object.Number).Value)
			}
			os.Exit(code)
			return object.NULL
		},
	},
	"sleep": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `sleep` must be NUMBER, got %s", args[0].Type())
			}
			ms := int64(args[0].(*object.Number).Value)
			time.Sleep(time.Duration(ms) * time.Millisecond)
			return object.NULL
		},
	},
	// HTTP Server functions
	"http_server": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("first argument to `http_server` must be NUMBER (port), got %s", args[0].Type())
			}
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("second argument to `http_server` must be FUNCTION (handler), got %s", args[1].Type())
			}

			port := int(args[0].(*object.Number).Value)
			handler := args[1].(*object.Function)

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				// Create request object
				reqMap := &object.Map{Pairs: make(map[string]object.Object)}
				reqMap.Pairs["method"] = &object.String{Value: r.Method}
				reqMap.Pairs["path"] = &object.String{Value: r.URL.Path}
				reqMap.Pairs["query"] = &object.String{Value: r.URL.RawQuery}

				// Headers
				headersMap := &object.Map{Pairs: make(map[string]object.Object)}
				for k, v := range r.Header {
					if len(v) > 0 {
						headersMap.Pairs[k] = &object.String{Value: v[0]}
					}
				}
				reqMap.Pairs["headers"] = headersMap

				// Body
				body, _ := io.ReadAll(r.Body)
				reqMap.Pairs["body"] = &object.String{Value: string(body)}

				// Create response object
				resMap := &object.Map{Pairs: make(map[string]object.Object)}
				resMap.Pairs["status"] = &object.Number{Value: 200}
				resMap.Pairs["body"] = &object.String{Value: ""}
				resMap.Pairs["headers"] = &object.Map{Pairs: make(map[string]object.Object)}

				// Call handler using EvalFunc
				var result object.Object
				if EvalFunc != nil {
					result = EvalFunc(handler, []object.Object{reqMap, resMap})
				}

				// Get updated response values
				if statusObj, ok := resMap.Pairs["status"]; ok {
					if status, ok := statusObj.(*object.Number); ok {
						w.WriteHeader(int(status.Value))
					}
				}

				// Set headers
				if headersObj, ok := resMap.Pairs["headers"]; ok {
					if headers, ok := headersObj.(*object.Map); ok {
						for k, v := range headers.Pairs {
							w.Header().Set(k, v.Inspect())
						}
					}
				}

				// Write body
				if bodyObj, ok := resMap.Pairs["body"]; ok {
					fmt.Fprint(w, bodyObj.Inspect())
				} else if result != nil && result != object.NULL {
					fmt.Fprint(w, result.Inspect())
				}
			})

			fmt.Printf("🚀 Server running at http://localhost:%d\n", port)
			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			if err != nil {
				return newError("server error: %s", err.Error())
			}
			return object.NULL
		},
	},
	"http_get": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `http_get` must be STRING, got %s", args[0].Type())
			}
			url := args[0].(*object.String).Value

			resp, err := http.Get(url)
			if err != nil {
				return newError("HTTP error: %s", err.Error())
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return newError("error reading response: %s", err.Error())
			}

			result := &object.Map{Pairs: make(map[string]object.Object)}
			result.Pairs["status"] = &object.Number{Value: float64(resp.StatusCode)}
			result.Pairs["body"] = &object.String{Value: string(body)}

			return result
		},
	},
}
