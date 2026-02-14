package builtins

import (
	"BanglaCode/src/object"
	"sync"
	"time"
)

func init() {
	// ghumaao (ঘুমাও) - async sleep
	Builtins["ghumaao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `ghumaao` must be NUMBER, got %s", args[0].Type())
			}

			ms := int64(args[0].(*object.Number).Value)
			promise := object.CreatePromise()

			go func() {
				time.Sleep(time.Duration(ms) * time.Millisecond)
				object.ResolvePromise(promise, object.NULL)
			}()

			return promise
		},
	}

	// sob_proyash (সব_প্রয়াস) - Promise.all
	// FIXED: Now waits for all promises CONCURRENTLY instead of sequentially
	Builtins["sob_proyash"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `sob_proyash` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			promises := make([]*object.Promise, len(arr.Elements))

			// Validate all elements are promises
			for i, el := range arr.Elements {
				p, ok := el.(*object.Promise)
				if !ok {
					return newError("all elements must be promises, got %s at index %d", el.Type(), i)
				}
				promises[i] = p
			}

			resultPromise := object.CreatePromise()

			go func() {
				results := make([]object.Object, len(promises))
				var wg sync.WaitGroup
				var mu sync.Mutex
				var firstError object.Object

				// Wait for all promises concurrently
				for i, p := range promises {
					wg.Add(1)
					go func(idx int, promise *object.Promise) {
						defer wg.Done()

						// Wait for this promise to complete
						select {
						case res := <-promise.ResultChan:
							mu.Lock()
							results[idx] = res
							mu.Unlock()
						case err := <-promise.ErrorChan:
							mu.Lock()
							if firstError == nil {
								firstError = err
							}
							mu.Unlock()
						}
					}(i, p)
				}

				// Wait for all goroutines to complete
				wg.Wait()

				// If any promise rejected, reject the result promise
				if firstError != nil {
					object.RejectPromise(resultPromise, firstError)
					return
				}

				// All promises resolved successfully
				object.ResolvePromise(resultPromise, &object.Array{Elements: results})
			}()

			return resultPromise
		},
	}
}
