package builtins

import (
	"BanglaCode/src/object"
	"os"
)

func init() {
	// Read file - poro (পড়ো - read)
	Builtins["poro"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `poro` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content, err := os.ReadFile(path)
			if err != nil {
				return newError("error reading file: %s", err.Error())
			}
			return &object.String{Value: string(content)}
		},
	}

	// Write file - lekho (লেখো - write)
	Builtins["lekho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `lekho` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content := args[1].Inspect()
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				return newError("error writing file: %s", err.Error())
			}
			return object.TRUE
		},
	}

	// Async file read - poro_async (পড়ো_async)
	Builtins["poro_async"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `poro_async` must be STRING, got %s", args[0].Type())
			}

			path := args[0].(*object.String).Value
			promise := object.CreatePromise()

			go func() {
				content, err := os.ReadFile(path)
				if err != nil {
					object.RejectPromise(promise, newError("error reading file: %s", err.Error()))
					return
				}
				object.ResolvePromise(promise, &object.String{Value: string(content)})
			}()

			return promise
		},
	}

	// Async file write - lekho_async (লেখো_async)
	Builtins["lekho_async"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `lekho_async` must be STRING, got %s", args[0].Type())
			}

			path := args[0].(*object.String).Value
			content := args[1].Inspect()
			promise := object.CreatePromise()

			go func() {
				err := os.WriteFile(path, []byte(content), 0644)
				if err != nil {
					object.RejectPromise(promise, newError("error writing file: %s", err.Error()))
					return
				}
				object.ResolvePromise(promise, object.TRUE)
			}()

			return promise
		},
	}
}
