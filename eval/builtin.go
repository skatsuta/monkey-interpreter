package eval

import "github.com/skatsuta/monkey-interpreter/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if l := len(args); l != 1 {
				return newError("wrong number of arguments. want=1, got=%d", l)
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
}
