package builtins

import (
	"BanglaCode/src/object"
	"sync"
	"time"
)

var (
	timerMu     sync.Mutex
	nextTimerID float64 = 1
	timeouts            = map[int]chan struct{}{}
	intervals           = map[int]chan struct{}{}
)

func init() {
	registerSetTimeout()
	registerSetInterval()
	registerClearTimeout()
	registerClearInterval()
}

func registerSetTimeout() {
	Builtins["setTimeout"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		cb, cbArgs, ms, errObj := parseTimerArgs("setTimeout", args)
		if errObj != nil {
			return errObj
		}

		id, stopCh := newTimerID(true)
		go func() {
			select {
			case <-time.After(time.Duration(ms) * time.Millisecond):
				EvalFunc(cb, cbArgs)
			case <-stopCh:
			}
			timerMu.Lock()
			delete(timeouts, id)
			timerMu.Unlock()
		}()
		return &object.Number{Value: float64(id)}
	}}
}

func registerSetInterval() {
	Builtins["setInterval"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		cb, cbArgs, ms, errObj := parseTimerArgs("setInterval", args)
		if errObj != nil {
			return errObj
		}
		if ms <= 0 {
			ms = 1
		}

		id, stopCh := newTimerID(false)
		go func() {
			ticker := time.NewTicker(time.Duration(ms) * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					EvalFunc(cb, cbArgs)
				case <-stopCh:
					timerMu.Lock()
					delete(intervals, id)
					timerMu.Unlock()
					return
				}
			}
		}()
		return &object.Number{Value: float64(id)}
	}}
}

func registerClearTimeout() {
	Builtins["clearTimeout"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clearTimer(args, true)
	}}
}

func registerClearInterval() {
	Builtins["clearInterval"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return clearTimer(args, false)
	}}
}

func parseTimerArgs(name string, args []object.Object) (*object.Function, []object.Object, int64, *object.Error) {
	if len(args) < 2 {
		return nil, nil, 0, newError("wrong number of arguments. got=%d, want>=2", len(args))
	}
	if args[0].Type() != object.FUNCTION_OBJ {
		return nil, nil, 0, newError("first argument to `%s` must be FUNCTION, got %s", name, args[0].Type())
	}
	if args[1].Type() != object.NUMBER_OBJ {
		return nil, nil, 0, newError("second argument to `%s` must be NUMBER delay(ms), got %s", name, args[1].Type())
	}
	cb := args[0].(*object.Function)
	ms := int64(args[1].(*object.Number).Value)
	cbArgs := []object.Object{}
	if len(args) > 2 {
		cbArgs = append(cbArgs, args[2:]...)
	}
	return cb, cbArgs, ms, nil
}

func newTimerID(timeout bool) (int, chan struct{}) {
	timerMu.Lock()
	defer timerMu.Unlock()

	id := int(nextTimerID)
	nextTimerID++
	stopCh := make(chan struct{}, 1)
	if timeout {
		timeouts[id] = stopCh
	} else {
		intervals[id] = stopCh
	}
	return id, stopCh
}

func clearTimer(args []object.Object, timeout bool) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.NUMBER_OBJ {
		return newError("argument must be NUMBER timer id, got %s", args[0].Type())
	}
	id := int(args[0].(*object.Number).Value)

	timerMu.Lock()
	defer timerMu.Unlock()

	if timeout {
		if ch, ok := timeouts[id]; ok {
			select {
			case ch <- struct{}{}:
			default:
			}
			delete(timeouts, id)
		}
	} else {
		if ch, ok := intervals[id]; ok {
			select {
			case ch <- struct{}{}:
			default:
			}
			delete(intervals, id)
		}
	}
	return object.NULL
}
