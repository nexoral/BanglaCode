package builtins

import (
	"BanglaCode/src/object"
	"math"
)

func init() {
	// Square root - borgomul (বর্গমূল)
	Builtins["borgomul"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `borgomul` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Sqrt(num)}
		},
	}

	// Power - ghat (ঘাত)
	Builtins["ghat"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ || args[1].Type() != object.NUMBER_OBJ {
				return newError("arguments to `ghat` must be NUMBERs")
			}
			base := args[0].(*object.Number).Value
			exp := args[1].(*object.Number).Value
			return &object.Number{Value: math.Pow(base, exp)}
		},
	}

	// Floor - niche (নিচে - down)
	Builtins["niche"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `niche` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Floor(num)}
		},
	}

	// Ceil - upore (উপরে - up)
	Builtins["upore"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `upore` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Ceil(num)}
		},
	}

	// Round - kache (কাছে - near)
	Builtins["kache"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `kache` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Round(num)}
		},
	}

	// Absolute - niratek (নিরপেক্ষ - absolute)
	Builtins["niratek"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `niratek` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Abs(num)}
		},
	}

	// Min - choto (ছোট - small)
	Builtins["choto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			minVal := math.Inf(1)
			for _, arg := range args {
				if arg.Type() != object.NUMBER_OBJ {
					return newError("all arguments to `choto` must be NUMBERs")
				}
				val := arg.(*object.Number).Value
				if val < minVal {
					minVal = val
				}
			}
			return &object.Number{Value: minVal}
		},
	}

	// Max - boro (বড় - big)
	Builtins["boro"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			maxVal := math.Inf(-1)
			for _, arg := range args {
				if arg.Type() != object.NUMBER_OBJ {
					return newError("all arguments to `boro` must be NUMBERs")
				}
				val := arg.(*object.Number).Value
				if val > maxVal {
					maxVal = val
				}
			}
			return &object.Number{Value: maxVal}
		},
	}
}
