package builtins

import (
	"BanglaCode/src/object"
	"fmt"
)

func init() {
	// router_banao (রাউটার বানাও - create router)
	Builtins["router_banao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			router := NewRouter("")
			registerRouter(router)

			routerMap := &object.Map{Pairs: make(map[string]object.Object)}
			routerMap.Pairs["__router_id__"] = &object.String{Value: fmt.Sprintf("%p", router)}

			// ana (আনা - GET - fetch)
			routerMap.Pairs["ana"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.ana", args); err != nil {
						return err
					}
					router.AddRoute("GET", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// pathano (পাঠানো - POST - send)
			routerMap.Pairs["pathano"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.pathano", args); err != nil {
						return err
					}
					router.AddRoute("POST", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// bodlano (বদলানো - PUT - update)
			routerMap.Pairs["bodlano"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.bodlano", args); err != nil {
						return err
					}
					router.AddRoute("PUT", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// mujhe_felo (মুছে ফেলো - DELETE - remove)
			routerMap.Pairs["mujhe_felo"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.mujhe_felo", args); err != nil {
						return err
					}
					router.AddRoute("DELETE", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// songshodhon (সংশোধন - PATCH - modify)
			routerMap.Pairs["songshodhon"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.songshodhon", args); err != nil {
						return err
					}
					router.AddRoute("PATCH", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// matha (মাথা - HEAD)
			routerMap.Pairs["matha"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.matha", args); err != nil {
						return err
					}
					router.AddRoute("HEAD", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// nirdharon (নির্ধারণ - OPTIONS)
			routerMap.Pairs["nirdharon"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if err := requireRoute("router.nirdharon", args); err != nil {
						return err
					}
					router.AddRoute("OPTIONS", args[0].(*object.String).Value, args[1])
					return routerMap
				},
			}

			// majhe (মাঝে - middleware intercept - agorao = next)
			routerMap.Pairs["majhe"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 1 {
						return newError("router.majhe() takes exactly 1 argument (handler), got %d", len(args))
					}
					if args[0].Type() != object.FUNCTION_OBJ && args[0].Type() != object.BUILTIN_OBJ {
						return newError("argument to router.majhe() must be FUNCTION, got %s", args[0].Type())
					}
					router.AddMiddleware(args[0])
					return routerMap
				},
			}

			// bebohar (ব্যবহার - mount sub-router) — FIXED: looks up actual sub-router
			routerMap.Pairs["bebohar"] = &object.Builtin{
				Fn: func(args ...object.Object) object.Object {
					if len(args) != 2 {
						return newError("router.bebohar() takes 2 arguments (mountPath, subRouter), got %d", len(args))
					}
					if args[0].Type() != object.STRING_OBJ {
						return newError("first argument to router.bebohar() must be STRING (path), got %s", args[0].Type())
					}
					if args[1].Type() != object.MAP_OBJ {
						return newError("second argument to router.bebohar() must be ROUTER, got %s", args[1].Type())
					}
					mountPath := args[0].(*object.String).Value
					subRouterMap := args[1].(*object.Map)

					ridObj, ok := subRouterMap.Pairs["__router_id__"]
					if !ok {
						return newError("invalid router object passed to bebohar()")
					}
					rid, ok := ridObj.(*object.String)
					if !ok {
						return newError("invalid router ID in bebohar()")
					}
					subRouter, found := getRouter(rid.Value)
					if !found {
						return newError("sub-router not found in registry")
					}
					router.MountSubRouter(mountPath, subRouter)
					return routerMap
				},
			}

			return routerMap
		},
	}
}

// requireRoute validates the 2-argument (path STRING, handler FUNCTION) signature
// used by all HTTP method registrations.
func requireRoute(name string, args []object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments to %s(). got=%d, want=2", name, len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("first argument to %s() must be STRING (path), got %s", name, args[0].Type())
	}
	if args[1].Type() != object.FUNCTION_OBJ && args[1].Type() != object.BUILTIN_OBJ {
		return newError("second argument to %s() must be FUNCTION (handler), got %s", name, args[1].Type())
	}
	return nil
}
