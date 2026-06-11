package builtins

import (
	"BanglaCode/src/object"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func init() {
	// cors_chharpao (ছাড়পাও - allow cross-origin requests)
	// cors_chharpao(app)
	// cors_chharpao(app, {"origin": "https://example.com", "methods": "GET,POST", "headers": "...", "maxAge": "86400"})
	Builtins["cors_chharpao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("wrong number of arguments. got=%d, want=1-2 (app, [options])", len(args))
			}
			router, err := extractRouter("cors_chharpao", args[0])
			if err != nil {
				return err
			}

			opts := CORSOptions{
				Origin:  "*",
				Methods: "GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS",
				Headers: "Content-Type,Authorization,X-Requested-With",
				MaxAge:  "86400",
			}
			if len(args) == 2 && args[1].Type() == object.MAP_OBJ {
				m := args[1].(*object.Map)
				if v, ok := m.Pairs["origin"].(*object.String); ok {
					opts.Origin = v.Value
				}
				if v, ok := m.Pairs["methods"].(*object.String); ok {
					opts.Methods = v.Value
				}
				if v, ok := m.Pairs["headers"].(*object.String); ok {
					opts.Headers = v.Value
				}
				if v, ok := m.Pairs["maxAge"].(*object.String); ok {
					opts.MaxAge = v.Value
				}
			}

			router.mu.Lock()
			router.corsEnabled = true
			router.corsOptions = opts
			router.mu.Unlock()
			return args[0]
		},
	}

	// file_dao (ফাইল দাও - serve static files from directory)
	// file_dao(app, "/public", "./static_dir")
	Builtins["file_dao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3 (app, urlPrefix, dirPath)", len(args))
			}
			router, err := extractRouter("file_dao", args[0])
			if err != nil {
				return err
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `file_dao` must be STRING (url prefix), got %s", args[1].Type())
			}
			if args[2].Type() != object.STRING_OBJ {
				return newError("third argument to `file_dao` must be STRING (directory path), got %s", args[2].Type())
			}

			urlPrefix := args[1].(*object.String).Value
			dirPath := args[2].(*object.String).Value

			if !strings.HasPrefix(urlPrefix, "/") {
				urlPrefix = "/" + urlPrefix
			}
			fileHandler := http.StripPrefix(urlPrefix, http.FileServer(http.Dir(dirPath)))

			router.mu.Lock()
			router.fileRoutes = append(router.fileRoutes, FileRoute{prefix: urlPrefix, handler: fileHandler})
			router.mu.Unlock()
			return args[0]
		},
	}

	// ghurao (ঘোরাও - redirect to another URL)
	// ghurao(res, "/login")         → 302 Found
	// ghurao(res, "/permanent", 301) → 301 Moved Permanently
	Builtins["ghurao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2-3 (res, url, [status])", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `ghurao` must be response MAP, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `ghurao` must be STRING (url), got %s", args[1].Type())
			}
			resMap := args[0].(*object.Map)
			redirectURL := args[1].(*object.String).Value
			status := 302
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `ghurao` must be NUMBER (status), got %s", args[2].Type())
				}
				status = int(args[2].(*object.Number).Value)
			}
			resMap.Pairs["status"] = &object.Number{Value: float64(status)}
			if h, ok := resMap.Pairs["headers"].(*object.Map); ok {
				h.Pairs["Location"] = &object.String{Value: redirectURL}
			}
			return resMap
		},
	}

	// kuki_rakho (কুকি রাখো - set a cookie on the response)
	// kuki_rakho(res, "name", "value")
	// kuki_rakho(res, "name", "value", {"httpOnly": sotti, "secure": sotti, "maxAge": 3600, "path": "/", "sameSite": "Lax"})
	Builtins["kuki_rakho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 3 || len(args) > 4 {
				return newError("wrong number of arguments. got=%d, want=3-4 (res, name, value, [opts])", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `kuki_rakho` must be response MAP, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `kuki_rakho` must be STRING (name), got %s", args[1].Type())
			}
			if args[2].Type() != object.STRING_OBJ {
				return newError("third argument to `kuki_rakho` must be STRING (value), got %s", args[2].Type())
			}

			resMap := args[0].(*object.Map)
			name := args[1].(*object.String).Value
			value := args[2].(*object.String).Value

			cookieStr := fmt.Sprintf("%s=%s", name, value)

			if len(args) == 4 && args[3].Type() == object.MAP_OBJ {
				opts := args[3].(*object.Map)
				if v, ok := opts.Pairs["path"].(*object.String); ok {
					cookieStr += "; Path=" + v.Value
				} else {
					cookieStr += "; Path=/"
				}
				if v, ok := opts.Pairs["maxAge"].(*object.Number); ok {
					cookieStr += fmt.Sprintf("; Max-Age=%d", int(v.Value))
				}
				if v, ok := opts.Pairs["sameSite"].(*object.String); ok {
					cookieStr += "; SameSite=" + v.Value
				}
				if v, ok := opts.Pairs["httpOnly"].(*object.Boolean); ok && v.Value {
					cookieStr += "; HttpOnly"
				}
				if v, ok := opts.Pairs["secure"].(*object.Boolean); ok && v.Value {
					cookieStr += "; Secure"
				}
			} else {
				cookieStr += "; Path=/"
			}

			if h, ok := resMap.Pairs["headers"].(*object.Map); ok {
				// Append to existing Set-Cookie or set new one
				existing, hasExisting := h.Pairs["Set-Cookie"].(*object.String)
				if hasExisting && existing.Value != "" {
					h.Pairs["Set-Cookie"] = &object.String{Value: existing.Value + "\r\nSet-Cookie: " + cookieStr}
				} else {
					h.Pairs["Set-Cookie"] = &object.String{Value: cookieStr}
				}
			}
			return resMap
		},
	}

	// html_uttor (HTML উত্তর - serve an HTML file as response)
	// html_uttor(res, "./views/index.html")
	Builtins["html_uttor"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2 (res, filepath)", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `html_uttor` must be response MAP, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `html_uttor` must be STRING (filepath), got %s", args[1].Type())
			}
			resMap := args[0].(*object.Map)
			filepath := args[1].(*object.String).Value
			content, err := os.ReadFile(filepath)
			if err != nil {
				return newError("html_uttor: could not read file '%s': %s", filepath, err.Error())
			}
			resMap.Pairs["body"] = &object.String{Value: string(content)}
			if h, ok := resMap.Pairs["headers"].(*object.Map); ok {
				h.Pairs["Content-Type"] = &object.String{Value: "text/html; charset=utf-8"}
			}
			return resMap
		},
	}

	// log_chalu (লগ চালু - enable request logging on the router)
	// log_chalu(app)
	Builtins["log_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1 (app)", len(args))
			}
			router, err := extractRouter("log_chalu", args[0])
			if err != nil {
				return err
			}
			router.mu.Lock()
			router.logEnabled = true
			router.mu.Unlock()
			return args[0]
		},
	}
}

// extractRouter retrieves the *Router from a BanglaCode router map.
func extractRouter(fn string, arg object.Object) (*Router, object.Object) {
	if arg.Type() != object.MAP_OBJ {
		return nil, newError("first argument to `%s` must be ROUTER (from router_banao()), got %s", fn, arg.Type())
	}
	routerMap := arg.(*object.Map)
	ridObj, ok := routerMap.Pairs["__router_id__"]
	if !ok {
		return nil, newError("first argument to `%s` is not a valid router", fn)
	}
	rid, ok := ridObj.(*object.String)
	if !ok {
		return nil, newError("invalid router ID in `%s`", fn)
	}
	router, found := getRouter(rid.Value)
	if !found {
		return nil, newError("`%s`: router not found — was it created with router_banao()?", fn)
	}
	return router, nil
}
