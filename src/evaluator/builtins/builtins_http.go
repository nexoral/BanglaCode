package builtins

import (
	"BanglaCode/src/object"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	// HTTP Server - server_chalu (সার্ভার চালু - start server)
	Builtins["server_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("first argument to `server_chalu` must be NUMBER (port), got %s", args[0].Type())
			}
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("second argument to `server_chalu` must be FUNCTION (handler), got %s", args[1].Type())
			}

			port := int(args[0].(*object.Number).Value)
			handler := args[1].(*object.Function)

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				reqMap := &object.Map{Pairs: make(map[string]object.Object)}
				reqMap.Pairs["method"] = &object.String{Value: r.Method}
				reqMap.Pairs["path"] = &object.String{Value: r.URL.Path}
				reqMap.Pairs["query"] = &object.String{Value: r.URL.RawQuery}

				headersMap := &object.Map{Pairs: make(map[string]object.Object)}
				for k, v := range r.Header {
					if len(v) > 0 {
						headersMap.Pairs[k] = &object.String{Value: v[0]}
					}
				}
				reqMap.Pairs["headers"] = headersMap

				body, _ := io.ReadAll(r.Body)
				reqMap.Pairs["body"] = &object.String{Value: string(body)}

				resMap := &object.Map{Pairs: make(map[string]object.Object)}
				resMap.Pairs["status"] = &object.Number{Value: 200}
				resMap.Pairs["body"] = &object.String{Value: ""}
				resMap.Pairs["headers"] = &object.Map{Pairs: make(map[string]object.Object)}

				var result object.Object
				if EvalFunc != nil {
					result = EvalFunc(handler, []object.Object{reqMap, resMap})
				}

				if statusObj, ok := resMap.Pairs["status"]; ok {
					if status, ok := statusObj.(*object.Number); ok {
						w.WriteHeader(int(status.Value))
					}
				}

				if headersObj, ok := resMap.Pairs["headers"]; ok {
					if headers, ok := headersObj.(*object.Map); ok {
						for k, v := range headers.Pairs {
							w.Header().Set(k, v.Inspect())
						}
					}
				}

				if bodyObj, ok := resMap.Pairs["body"]; ok {
					fmt.Fprint(w, bodyObj.Inspect())
				} else if result != nil && result != object.NULL {
					fmt.Fprint(w, result.Inspect())
				}
			})

			fmt.Printf("🚀 Server cholche http://localhost:%d e\n", port)
			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			if err != nil {
				return newError("server error: %s", err.Error())
			}
			return object.NULL
		},
	}

	// HTTP GET - anun (আনুন - fetch/bring)
	Builtins["anun"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `anun` must be STRING, got %s", args[0].Type())
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
	}

	// Async HTTP GET - anun_async (আনুন_async)
	Builtins["anun_async"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `anun_async` must be STRING, got %s", args[0].Type())
			}

			url := args[0].(*object.String).Value
			promise := object.CreatePromise()

			go func() {
				resp, err := http.Get(url)
				if err != nil {
					object.RejectPromise(promise, newError("HTTP error: %s", err.Error()))
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					object.RejectPromise(promise, newError("error reading response: %s", err.Error()))
					return
				}

				result := &object.Map{Pairs: make(map[string]object.Object)}
				result.Pairs["status"] = &object.Number{Value: float64(resp.StatusCode)}
				result.Pairs["body"] = &object.String{Value: string(body)}

				object.ResolvePromise(promise, result)
			}()

			return promise
		},
	}

	// JSON Parse - json_poro (JSON পড়ো - read JSON)
	Builtins["json_poro"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `json_poro` must be STRING, got %s", args[0].Type())
			}
			jsonStr := args[0].(*object.String).Value
			return parseJSON(jsonStr)
		},
	}

	// JSON Stringify - json_banao (JSON বানাও - make JSON)
	Builtins["json_banao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: stringifyJSON(args[0])}
		},
	}

	// Simple HTTP response helper - uttor (উত্তর - reply/response)
	Builtins["uttor"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 4 {
				return newError("wrong number of arguments. got=%d, want=2-4 (res, body, [status], [contentType])", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `uttor` must be response MAP, got %s", args[0].Type())
			}
			resMap := args[0].(*object.Map)

			// Set body
			resMap.Pairs["body"] = args[1]

			// Set status (optional, default 200)
			if len(args) >= 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `uttor` must be NUMBER (status), got %s", args[2].Type())
				}
				resMap.Pairs["status"] = args[2]
			}

			// Set content-type (optional)
			if len(args) >= 4 {
				if args[3].Type() != object.STRING_OBJ {
					return newError("fourth argument to `uttor` must be STRING (contentType), got %s", args[3].Type())
				}
				if headersObj, ok := resMap.Pairs["headers"]; ok {
					if headers, ok := headersObj.(*object.Map); ok {
						headers.Pairs["Content-Type"] = args[3]
					}
				}
			}

			return resMap
		},
	}

	// JSON response helper - json_uttor (JSON উত্তর - JSON reply)
	Builtins["json_uttor"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2-3 (res, data, [status])", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `json_uttor` must be response MAP, got %s", args[0].Type())
			}
			resMap := args[0].(*object.Map)

			// Convert data to JSON string
			jsonStr := stringifyJSON(args[1])
			resMap.Pairs["body"] = &object.String{Value: jsonStr}

			// Set status (optional, default 200)
			if len(args) >= 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `json_uttor` must be NUMBER (status), got %s", args[2].Type())
				}
				resMap.Pairs["status"] = args[2]
			}

			// Set content-type to JSON
			if headersObj, ok := resMap.Pairs["headers"]; ok {
				if headers, ok := headersObj.(*object.Map); ok {
					headers.Pairs["Content-Type"] = &object.String{Value: "application/json; charset=utf-8"}
				}
			}

			return resMap
		},
	}
}

// parseJSON converts a JSON string to BanglaCode objects
func parseJSON(jsonStr string) object.Object {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return newError("JSON parse error: %s", err.Error())
	}
	return JsonToObject(data)
}

// JsonToObject recursively converts Go values to BanglaCode objects
func JsonToObject(data interface{}) object.Object {
	switch v := data.(type) {
	case nil:
		return object.NULL
	case bool:
		if v {
			return object.TRUE
		}
		return object.FALSE
	case float64:
		return &object.Number{Value: v}
	case string:
		return &object.String{Value: v}
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, item := range v {
			elements[i] = JsonToObject(item)
		}
		return &object.Array{Elements: elements}
	case map[string]interface{}:
		pairs := make(map[string]object.Object)
		for key, val := range v {
			pairs[key] = JsonToObject(val)
		}
		return &object.Map{Pairs: pairs}
	default:
		return newError("unsupported JSON type")
	}
}

// stringifyJSON converts a BanglaCode object to JSON string
func stringifyJSON(obj object.Object) string {
	data := objectToJSON(obj)
	bytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

// objectToJSON recursively converts BanglaCode objects to Go values
func objectToJSON(obj object.Object) interface{} {
	switch v := obj.(type) {
	case *object.Null:
		return nil
	case *object.Boolean:
		return v.Value
	case *object.Number:
		return v.Value
	case *object.String:
		return v.Value
	case *object.Array:
		arr := make([]interface{}, len(v.Elements))
		for i, el := range v.Elements {
			arr[i] = objectToJSON(el)
		}
		return arr
	case *object.Map:
		m := make(map[string]interface{})
		for key, val := range v.Pairs {
			m[key] = objectToJSON(val)
		}
		return m
	default:
		return obj.Inspect()
	}
}
