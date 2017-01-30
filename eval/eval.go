package eval

import (
	"github.com/skatsuta/monkey-interpreter/ast"
	"github.com/skatsuta/monkey-interpreter/object"
)

var (
	// NULL represents a value of null reference.
	NULL = &object.Null{}
	// TRUE represents a value of true literals.
	TRUE = &object.Boolean{Value: true}
	// FALSE represents a value of false literals.
	FALSE = &object.Boolean{Value: false}
)

// Eval evaluates the given node and returns an evaluated object.
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	default:
		return nil
	}
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
