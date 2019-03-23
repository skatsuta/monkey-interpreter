package eval

import (
	"github.com/skatsuta/monkey-interpreter/ast"
	"github.com/skatsuta/monkey-interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
