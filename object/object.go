package object

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/skatsuta/monkey-interpreter/ast"
)

// Type is a type of objects.
type Type string

const (
	// IntegerType represents a type of integers.
	IntegerType Type = "Integer"
	// BooleanType represents a type of booleans.
	BooleanType = "Boolean"
	// NullType represents a type of null.
	NullType = "Null"
	// ReturnValueType represents a type of return values.
	ReturnValueType = "ReturnValue"
	// ErrorType represents a type of errors.
	ErrorType = "Error"
	// FunctionType represents a type of functions.
	FunctionType = "Function"
	// StringType represents a type of strings.
	StringType = "String"
)

// Object represents an object of Monkey language.
type Object interface {
	Type() Type
	Inspect() string
}

// Integer represents an integer.
type Integer struct {
	Value int64
}

// Type returns the type of the Integer.
func (i *Integer) Type() Type {
	return IntegerType
}

// Inspect returns a string representation of the Integer.
func (i *Integer) Inspect() string {
	return strconv.FormatInt(i.Value, 10)
}

// Boolean represents a boolean.
type Boolean struct {
	Value bool
}

// Type returns the type of the Boolean.
func (b *Boolean) Type() Type {
	return BooleanType
}

// Inspect returns a string representation of the Boolean.
func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

// Null represents the absence of any value.
type Null struct{}

// Type returns the type of the Null.
func (n *Null) Type() Type {
	return NullType
}

// Inspect returns a string representation of the Null.
func (n *Null) Inspect() string {
	return "null"
}

// ReturnValue represents a return value.
type ReturnValue struct {
	Value Object
}

// Type returns the type of the ReturnValue.
func (rv *ReturnValue) Type() Type {
	return ReturnValueType
}

// Inspect returns a string representation of the ReturnValue.
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error represents an error.
type Error struct {
	Message string
}

// Type returns the type of the Error.
func (e *Error) Type() Type {
	return ErrorType
}

// Inspect returns a string representation of the Error.
func (e *Error) Inspect() string {
	return "Error:" + e.Message
}

// Function represents a function.
type Function struct {
	Parameters []*ast.Ident
	Body       *ast.BlockStatement
	Env        Environment
}

// Type returns the type of the Function.
func (f *Function) Type() Type {
	return FunctionType
}

// Inspect returns a string representation of the Function.
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := make([]string, 0, len(f.Parameters))
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String represents a string.
type String struct {
	Value string
}

// Type returns the type of the String.
func (s *String) Type() Type {
	return StringType
}

// Inspect returns a string representation of the String.
func (s *String) Inspect() string {
	return s.Value
}
