package builtins

import (
	"BanglaCode/src/object"
	"time"
)

func init() {
	// goti_shima (গতি সীমা - per-IP sliding-window rate limiter)
	// goti_shima(app, maxRequests, windowSeconds)
	// Example: goti_shima(app, 100, 60)  → max 100 req per 60 seconds per IP
	Builtins["goti_shima"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3 (app, maxReq, windowSec)", len(args))
			}
			router, err := extractRouter("goti_shima", args[0])
			if err != nil {
				return err
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `goti_shima` must be NUMBER (max requests), got %s", args[1].Type())
			}
			if args[2].Type() != object.NUMBER_OBJ {
				return newError("third argument to `goti_shima` must be NUMBER (window seconds), got %s", args[2].Type())
			}
			maxReq := int(args[1].(*object.Number).Value)
			windowSec := int(args[2].(*object.Number).Value)
			if maxReq <= 0 || windowSec <= 0 {
				return newError("goti_shima: maxReq and windowSec must be positive numbers")
			}
			router.mu.Lock()
			router.rateLimiter = &RateLimiter{
				max:    maxReq,
				window: time.Duration(windowSec) * time.Second,
				hits:   make(map[string][]time.Time),
			}
			router.mu.Unlock()
			return args[0]
		},
	}

	// sankochon_chalu (সংকোচন চালু - enable gzip compression)
	// sankochon_chalu(app)
	// Compresses responses when client sends Accept-Encoding: gzip
	Builtins["sankochon_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1 (app)", len(args))
			}
			router, err := extractRouter("sankochon_chalu", args[0])
			if err != nil {
				return err
			}
			router.mu.Lock()
			router.gzipEnabled = true
			router.mu.Unlock()
			return args[0]
		},
	}

	// somoy_shima (সময় সীমা - request timeout)
	// somoy_shima(app, seconds)
	// Example: somoy_shima(app, 30)  → 30-second handler timeout
	Builtins["somoy_shima"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2 (app, seconds)", len(args))
			}
			router, err := extractRouter("somoy_shima", args[0])
			if err != nil {
				return err
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `somoy_shima` must be NUMBER (seconds), got %s", args[1].Type())
			}
			secs := args[1].(*object.Number).Value
			if secs <= 0 {
				return newError("somoy_shima: seconds must be a positive number")
			}
			router.mu.Lock()
			router.timeout = time.Duration(secs * float64(time.Second))
			router.mu.Unlock()
			return args[0]
		},
	}

	// akaar_shima (আকার সীমা - request body size limit)
	// akaar_shima(app, bytes)
	// Example: akaar_shima(app, 1048576)  → 1 MB body limit
	Builtins["akaar_shima"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2 (app, bytes)", len(args))
			}
			router, err := extractRouter("akaar_shima", args[0])
			if err != nil {
				return err
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `akaar_shima` must be NUMBER (bytes), got %s", args[1].Type())
			}
			bytes := int64(args[1].(*object.Number).Value)
			if bytes <= 0 {
				return newError("akaar_shima: bytes must be a positive number")
			}
			router.mu.Lock()
			router.maxBodyBytes = bytes
			router.mu.Unlock()
			return args[0]
		},
	}

	// bhul_sambhalo (ভুল সামলাও - error middleware handler)
	// bhul_sambhalo(app, kaj(err, req, res) { ... })
	// Called when a route handler returns an ERROR object
	Builtins["bhul_sambhalo"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2 (app, handler)", len(args))
			}
			router, err := extractRouter("bhul_sambhalo", args[0])
			if err != nil {
				return err
			}
			if args[1].Type() != object.FUNCTION_OBJ && args[1].Type() != object.BUILTIN_OBJ {
				return newError("second argument to `bhul_sambhalo` must be FUNCTION (handler), got %s", args[1].Type())
			}
			router.mu.Lock()
			router.errorHandler = args[1]
			router.mu.Unlock()
			return args[0]
		},
	}
}
