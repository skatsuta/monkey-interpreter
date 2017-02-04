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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if l := len(args); l != 1 {
				return newError("wrong number of arguments. want=1, got=%d", l)
			}

			if typ := args[0].Type(); typ != object.ArrayType {
				return newError("argument to `first` must be Array, got %s", typ)
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) == 0 {
				return NilValue
			}
			return arr.Elements[0]
		},
	},

	"last": {
		Fn: func(args ...object.Object) object.Object {
			if l := len(args); l != 1 {
				return newError("wrong number of arguments. want=1, got=%d", l)
			}

			if typ := args[0].Type(); typ != object.ArrayType {
				return newError("argument to `last` must be Array, got %s", typ)
			}

			arr := args[0].(*object.Array)
			l := len(arr.Elements)
			if l == 0 {
				return NilValue
			}
			return arr.Elements[l-1]
		},
	},
}
