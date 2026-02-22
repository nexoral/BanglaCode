package builtins

import (
	"BanglaCode/src/object"
	"io"
	"os"
	"path/filepath"
	"time"
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

	// Append to file - file_jog (ফাইল জোগ - add/append)
	Builtins["file_jog"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `file_jog` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content := args[1].Inspect()

			// Open file in append mode, create if doesn't exist
			f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return newError("error opening file for append: %s", err.Error())
			}
			defer f.Close()

			if _, err := f.WriteString(content); err != nil {
				return newError("error appending to file: %s", err.Error())
			}
			return object.TRUE
		},
	}

	// Delete file - file_mochho (ফাইল মোছো - delete/erase)
	Builtins["file_mochho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `file_mochho` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			err := os.Remove(path)
			if err != nil {
				return newError("error deleting file: %s", err.Error())
			}
			return object.TRUE
		},
	}

	// Copy file - file_nokol (ফাইল নকল - copy/duplicate)
	Builtins["file_nokol"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `file_nokol` must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `file_nokol` must be STRING, got %s", args[1].Type())
			}

			src := args[0].(*object.String).Value
			dst := args[1].(*object.String).Value

			// Open source file
			sourceFile, err := os.Open(src)
			if err != nil {
				return newError("error opening source file: %s", err.Error())
			}
			defer sourceFile.Close()

			// Create destination file
			destFile, err := os.Create(dst)
			if err != nil {
				return newError("error creating destination file: %s", err.Error())
			}
			defer destFile.Close()

			// Copy content
			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				return newError("error copying file: %s", err.Error())
			}

			return object.TRUE
		},
	}

	// Delete folder - folder_mochho (ফোল্ডার মোছো - delete folder)
	Builtins["folder_mochho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("wrong number of arguments. got=%d, want=1 or 2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `folder_mochho` must be STRING, got %s", args[0].Type())
			}

			path := args[0].(*object.String).Value
			recursive := false

			// Check for recursive flag (second argument)
			if len(args) == 2 {
				if args[1].Type() != object.BOOLEAN_OBJ {
					return newError("second argument to `folder_mochho` must be BOOLEAN, got %s", args[1].Type())
				}
				recursive = args[1].(*object.Boolean).Value
			}

			var err error
			if recursive {
				// Remove directory and all contents
				err = os.RemoveAll(path)
			} else {
				// Remove empty directory only
				err = os.Remove(path)
			}

			if err != nil {
				return newError("error deleting folder: %s", err.Error())
			}
			return object.TRUE
		},
	}

	// Watch file for changes - file_dekhun (ফাইল দেখুন - watch file)
	Builtins["file_dekhun"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `file_dekhun` must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("second argument to `file_dekhun` must be FUNCTION, got %s", args[1].Type())
			}

			path := args[0].(*object.String).Value
			callback := args[1].(*object.Function)

			// Get initial file info
			initialInfo, err := os.Stat(path)
			if err != nil {
				return newError("error getting file info: %s", err.Error())
			}
			lastModTime := initialInfo.ModTime()

			// Create watcher object
			watcher := &object.Map{
				Pairs: make(map[string]object.Object),
			}
			watcher.Pairs["path"] = &object.String{Value: path}
			watcher.Pairs["active"] = object.TRUE

			// Start watching in goroutine
			go func() {
				ticker := time.NewTicker(1 * time.Second)
				defer ticker.Stop()

				for range ticker.C {
					// Check if watcher is still active
					if active, ok := watcher.Pairs["active"]; ok {
						if b, ok := active.(*object.Boolean); ok && !b.Value {
							return // Stop watching
						}
					}

					// Check for file changes
					info, err := os.Stat(path)
					if err != nil {
						continue // File might have been deleted
					}

					if info.ModTime().After(lastModTime) {
						lastModTime = info.ModTime()

						// Call callback with event type and filename
						if EvalFunc != nil {
							EvalFunc(callback, []object.Object{
								&object.String{Value: "change"},
								&object.String{Value: filepath.Base(path)},
							})
						}
					}
				}
			}()

			return watcher
		},
	}

	// Stop file watching - file_dekhun_bondho (ফাইল দেখুন বন্ধ - stop watching)
	Builtins["file_dekhun_bondho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to `file_dekhun_bondho` must be MAP (watcher), got %s", args[0].Type())
			}

			watcher := args[0].(*object.Map)
			watcher.Pairs["active"] = object.FALSE
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

// Helper function to extend environment for callback
func extendFunctionEnvForCallback(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for i, param := range fn.Parameters {
		if i < len(args) {
			env.Set(param.Value, args[i])
		}
	}
	return env
}
