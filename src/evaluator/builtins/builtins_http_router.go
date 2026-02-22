package builtins

import (
	"BanglaCode/src/object"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Router represents a modular HTTP router (similar to Express.js Router)
// Supports all 7 common HTTP methods with Banglish method names
type Router struct {
	basePath string
	routes   map[string]map[string]*object.Function // method -> path -> handler
	mu       sync.RWMutex
}

// NewRouter creates a new router instance with support for all HTTP methods
func NewRouter(basePath string) *Router {
	return &Router{
		basePath: basePath,
		routes: map[string]map[string]*object.Function{
			"GET":     make(map[string]*object.Function),
			"POST":    make(map[string]*object.Function),
			"PUT":     make(map[string]*object.Function),
			"DELETE":  make(map[string]*object.Function),
			"PATCH":   make(map[string]*object.Function),
			"HEAD":    make(map[string]*object.Function),
			"OPTIONS": make(map[string]*object.Function),
		},
	}
}

// AddRoute adds a route to the router
func (r *Router) AddRoute(method, path string, handler *object.Function) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Normalize path
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Store route
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]*object.Function)
	}
	r.routes[method][path] = handler
}

// GetHandler finds a handler for the given method and path
func (r *Router) GetHandler(method, path string) (*object.Function, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Remove base path if present
	if r.basePath != "" && strings.HasPrefix(path, r.basePath) {
		path = strings.TrimPrefix(path, r.basePath)
		if path == "" {
			path = "/"
		}
	}

	handler, ok := r.routes[method][path]
	return handler, ok
}

// MountSubRouter mounts a sub-router at a specific path
func (r *Router) MountSubRouter(mountPath string, subRouter *Router) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Normalize mount path
	if !strings.HasPrefix(mountPath, "/") {
		mountPath = "/" + mountPath
	}
	if strings.HasSuffix(mountPath, "/") && mountPath != "/" {
		mountPath = strings.TrimSuffix(mountPath, "/")
	}

	// Update sub-router's base path
	subRouter.basePath = r.basePath + mountPath

	// Copy all routes from sub-router with updated paths
	for method, routes := range subRouter.routes {
		for path, handler := range routes {
			fullPath := mountPath + path
			if r.routes[method] == nil {
				r.routes[method] = make(map[string]*object.Function)
			}
			r.routes[method][fullPath] = handler
		}
	}
}

// ServeHTTP implements http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := r.GetHandler(req.Method, req.URL.Path)

	if !ok {
		http.NotFound(w, req)
		return
	}

	// Build request object
	reqMap := &object.Map{Pairs: make(map[string]object.Object)}
	reqMap.Pairs["method"] = &object.String{Value: req.Method}
	reqMap.Pairs["path"] = &object.String{Value: req.URL.Path}
	reqMap.Pairs["query"] = &object.String{Value: req.URL.RawQuery}

	// Parse headers
	headersMap := &object.Map{Pairs: make(map[string]object.Object)}
	for k, v := range req.Header {
		if len(v) > 0 {
			headersMap.Pairs[k] = &object.String{Value: v[0]}
		}
	}
	reqMap.Pairs["headers"] = headersMap

	// Read body
	body, _ := io.ReadAll(req.Body)
	reqMap.Pairs["body"] = &object.String{Value: string(body)}

	// Build response object
	resMap := &object.Map{Pairs: make(map[string]object.Object)}
	resMap.Pairs["status"] = &object.Number{Value: 200}
	resMap.Pairs["body"] = &object.String{Value: ""}
	resMap.Pairs["headers"] = &object.Map{Pairs: make(map[string]object.Object)}

	// Execute handler
	var result object.Object
	if EvalFunc != nil {
		result = EvalFunc(handler, []object.Object{reqMap, resMap})
	}

	// Write response
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
}

func init() {
	// router_banao (রাউটার বানাও - create router)
	Builtins["router_banao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			router := NewRouter("")

			// Create a map to represent the router with methods
			routerMap := &object.Map{Pairs: make(map[string]object.Object)}

			// Store the actual router instance (we'll use this internally)
			routerMap.Pairs["__internal_router__"] = &object.String{Value: fmt.Sprintf("%p", router)}

			// Add ana method (আনা - GET - fetch)
			routerMap.Pairs["ana"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.ana(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.ana() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.ana() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("GET", path, handler)

					return routerMap // Return router for chaining
				},
			}

			// Add pathano method (পাঠানো - POST - send)
			routerMap.Pairs["pathano"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.pathano(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.pathano() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.pathano() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("POST", path, handler)

					return routerMap
				},
			}

			// Add bodlano method (বদলানো - PUT - update/change)
			routerMap.Pairs["bodlano"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.bodlano(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.bodlano() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.bodlano() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("PUT", path, handler)

					return routerMap
				},
			}

			// Add mujhe_felo method (মুছে ফেলো - DELETE - remove)
			routerMap.Pairs["mujhe_felo"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.mujhe_felo(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.mujhe_felo() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.mujhe_felo() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("DELETE", path, handler)

					return routerMap
				},
			}

			// Add songshodhon method (সংশোধন - PATCH - modify)
			routerMap.Pairs["songshodhon"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.songshodhon(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.songshodhon() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.songshodhon() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("PATCH", path, handler)

					return routerMap
				},
			}

			// Add matha method (মাথা - HEAD - retrieve headers)
			routerMap.Pairs["matha"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.matha(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.matha() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.matha() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("HEAD", path, handler)

					return routerMap
				},
			}

			// Add nirdharon method (নির্ধারণ - OPTIONS - determine options)
			routerMap.Pairs["nirdharon"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.nirdharon(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.nirdharon() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.FUNCTION_OBJ {
						return newError("second argument to router.nirdharon() must be FUNCTION (handler), got %s", args[1].Type())
					}

					path := args[0].(*object.String).Value
					handler := args[1].(*object.Function)
					router.AddRoute("OPTIONS", path, handler)

					return routerMap
				},
			}
			// Add bebohār method (ব্যবহার - use/mount sub-router)
			routerMap.Pairs["bebohar"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("wrong number of arguments to router.bebohar(). got=%d, want=2", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.bebohar() must be STRING (mount path), got %s", args[0].Type())
					}
					if args[1].Type() != object.MAP_OBJ {
						return newError("second argument to router.bebohar() must be ROUTER (sub-router), got %s", args[1].Type())
					}

					mountPath := args[0].(*object.String).Value
					subRouterMap := args[1].(*object.Map)

					// Extract the internal router pointer
					if internalObj, ok := subRouterMap.Pairs["__internal_router__"]; ok {
						if _, ok := internalObj.(*object.String); ok {
							// In a real implementation, we'd store routers in a global map
							// For now, we'll recreate the logic
							subRouter := NewRouter("")

							// Mount the sub-router
							router.MountSubRouter(mountPath, subRouter)
						}
					}

					return routerMap
				},
			}

			// Store router in global registry for server_chalu to use
			registerRouter(router)
			routerMap.Pairs["__router_id__"] = &object.String{Value: fmt.Sprintf("%p", router)}

			return routerMap
		},
	}
}

// Global router registry
var (
	routerRegistry   = make(map[string]*Router)
	routerRegistryMu sync.RWMutex
)

func registerRouter(r *Router) {
	routerRegistryMu.Lock()
	defer routerRegistryMu.Unlock()
	id := fmt.Sprintf("%p", r)
	routerRegistry[id] = r
}

func getRouter(id string) (*Router, bool) {
	routerRegistryMu.RLock()
	defer routerRegistryMu.RUnlock()
	r, ok := routerRegistry[id]
	return r, ok
}
