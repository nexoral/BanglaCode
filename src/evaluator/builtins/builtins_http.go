package builtins

import (
	"BanglaCode/src/object"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func init() {
	// server_chalu (সার্ভার চালু - start server)
	// Accepts a Router (MAP with __router_id__) or a plain function handler.
	Builtins["server_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("first argument to `server_chalu` must be NUMBER (port), got %s", args[0].Type())
			}
			port := int(args[0].(*object.Number).Value)

			// Router mode
			if args[1].Type() == object.MAP_OBJ {
				routerMap := args[1].(*object.Map)
				ridObj, ok := routerMap.Pairs["__router_id__"]
				if !ok {
					return newError("second argument to `server_chalu` is not a valid router")
				}
				rid, ok := ridObj.(*object.String)
				if !ok {
					return newError("invalid router ID in `server_chalu`")
				}
				router, found := getRouter(rid.Value)
				if !found {
					return newError("router not found — was it created with router_banao()?")
				}
				fmt.Printf("🚀 Server cholche http://localhost:%d e (Router mode)\n", port)
				if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
					return newError("server error: %s", err.Error())
				}
				return object.NULL
			}

			// Function-based mode (backward compatible)
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("second argument to `server_chalu` must be FUNCTION or ROUTER, got %s", args[1].Type())
			}
			handler := args[1].(*object.Function)
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				body, _ := io.ReadAll(r.Body)
				reqMap := buildRequestMap(r, body, nil)
				resMap := buildResponseMap()
				if EvalFunc != nil {
					EvalFunc(handler, []object.Object{reqMap, resMap})
				}
				writeHTTPResponse(w, resMap, false)
			})
			fmt.Printf("🚀 Server cholche http://localhost:%d e\n", port)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
				return newError("server error: %s", err.Error())
			}
			return object.NULL
		},
	}

	// anun (আনুন - HTTP client, backward compatible + extended with options)
	// anun(url)                         → GET
	// anun(url, {method, body, headers}) → any method
	Builtins["anun"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("wrong number of arguments. got=%d, want=1-2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `anun` must be STRING (url), got %s", args[0].Type())
			}
			resp, err := doHTTPRequest(args)
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

	// anun_async (আনুন async - async HTTP client)
	Builtins["anun_async"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("wrong number of arguments. got=%d, want=1-2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `anun_async` must be STRING (url), got %s", args[0].Type())
			}
			promise := object.CreatePromise()
			// Capture args slice for goroutine
			capturedArgs := args
			go func() {
				resp, err := doHTTPRequest(capturedArgs)
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

	// json_poro (JSON পড়ো - parse JSON string)
	Builtins["json_poro"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `json_poro` must be STRING, got %s", args[0].Type())
			}
			return parseJSON(args[0].(*object.String).Value)
		},
	}

	// json_banao (JSON বানাও - stringify to JSON)
	Builtins["json_banao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: stringifyJSON(args[0])}
		},
	}

	// uttor (উত্তর - set response body, status, content-type)
	Builtins["uttor"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 4 {
				return newError("wrong number of arguments. got=%d, want=2-4", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `uttor` must be response MAP, got %s", args[0].Type())
			}
			resMap := args[0].(*object.Map)
			resMap.Pairs["body"] = args[1]
			if len(args) >= 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `uttor` must be NUMBER (status), got %s", args[2].Type())
				}
				resMap.Pairs["status"] = args[2]
			}
			if len(args) >= 4 {
				if args[3].Type() != object.STRING_OBJ {
					return newError("fourth argument to `uttor` must be STRING (contentType), got %s", args[3].Type())
				}
				if h, ok := resMap.Pairs["headers"].(*object.Map); ok {
					h.Pairs["Content-Type"] = args[3]
				}
			}
			return resMap
		},
	}

	// json_uttor (JSON উত্তর - send JSON response)
	Builtins["json_uttor"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2-3", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `json_uttor` must be response MAP, got %s", args[0].Type())
			}
			resMap := args[0].(*object.Map)
			resMap.Pairs["body"] = &object.String{Value: stringifyJSON(args[1])}
			if len(args) >= 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `json_uttor` must be NUMBER (status), got %s", args[2].Type())
				}
				resMap.Pairs["status"] = args[2]
			}
			if h, ok := resMap.Pairs["headers"].(*object.Map); ok {
				h.Pairs["Content-Type"] = &object.String{Value: "application/json; charset=utf-8"}
			}
			return resMap
		},
	}
}

// doHTTPRequest builds and executes an HTTP request from BanglaCode args.
// args[0] = url STRING
// args[1] = options MAP (optional): method, body, headers
func doHTTPRequest(args []object.Object) (*http.Response, error) {
	rawURL := args[0].(*object.String).Value
	method := "GET"
	bodyStr := ""
	extraHeaders := map[string]string{}

	if len(args) == 2 && args[1].Type() == object.MAP_OBJ {
		opts := args[1].(*object.Map)
		if m, ok := opts.Pairs["method"].(*object.String); ok {
			method = strings.ToUpper(m.Value)
		}
		if b, ok := opts.Pairs["body"].(*object.String); ok {
			bodyStr = b.Value
		}
		if h, ok := opts.Pairs["headers"].(*object.Map); ok {
			for k, v := range h.Pairs {
				if vs, ok := v.(*object.String); ok {
					extraHeaders[k] = vs.Value
				}
			}
		}
	}

	var bodyReader io.Reader
	if bodyStr != "" {
		bodyReader = strings.NewReader(bodyStr)
	}

	req, err := http.NewRequest(method, rawURL, bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	return client.Do(req)
}

// parseJSON converts a JSON string to BanglaCode objects.
func parseJSON(jsonStr string) object.Object {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return newError("JSON parse error: %s", err.Error())
	}
	return JsonToObject(data)
}

// JsonToObject recursively converts Go values to BanglaCode objects.
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
		pairs := make(map[string]object.Object, len(v))
		for key, val := range v {
			pairs[key] = JsonToObject(val)
		}
		return &object.Map{Pairs: pairs}
	default:
		return newError("unsupported JSON type: %T", data)
	}
}

// stringifyJSON converts a BanglaCode object to a JSON string.
func stringifyJSON(obj object.Object) string {
	data := objectToJSON(obj)
	bytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

// objectToJSON recursively converts BanglaCode objects to Go values for JSON marshalling.
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
		m := make(map[string]interface{}, len(v.Pairs))
		for key, val := range v.Pairs {
			m[key] = objectToJSON(val)
		}
		return m
	default:
		return obj.Inspect()
	}
}
