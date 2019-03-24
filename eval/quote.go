package eval

import (
	"github.com/skatsuta/monkey-interpreter/ast"
	"github.com/skatsuta/monkey-interpreter/object"
)

const (
	// FuncNameQuote is a name for quote function.
	FuncNameQuote = "quote"
	// FuncNameUnquote is a name for unquote function.
	FuncNameUnquote = "unquote"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
