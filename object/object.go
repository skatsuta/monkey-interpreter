package object

import (
	"bytes"
	"hash/fnv"
	"strconv"
	"strings"

	"github.com/skatsuta/monkey-interpreter/ast"
)

// Type is a type of objects.
type Type string

const (
	// IntegerType represents a type of integers.
	IntegerType Type = "Integer"
	// FloatType represents a type of floating point numbers.
	FloatType = "Float"
	// BooleanType represents a type of booleans.
	BooleanType = "Boolean"
	// NilType represents a type of nil.
	NilType = "Nil"
	// ReturnValueType represents a type of return values.
	ReturnValueType = "ReturnValue"
	// ErrorType represents a type of errors.
	ErrorType = "Error"
	// FunctionType represents a type of functions.
	FunctionType = "Function"
	// StringType represents a type of strings.
	StringType = "String"
	// BuiltinType represents a type of builtin functions.
	BuiltinType = "Builtin"
	// ArrayType represents a type of arrays.
	ArrayType = "Array"
	// HashType represents a type of hashes.
	HashType = "Hash"
)

// Object represents an object of Monkey language.
type Object interface {
	Type() Type
	Inspect() string
}

// HashKey represents a key of a hash.
type HashKey struct {
	Type  Type
	Value uint64
}

// Hashable is the interface that is able to become a hash key.
type Hashable interface {
	HashKey() HashKey
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

// HashKey returns a hash key object for i.
func (i *Integer) HashKey() HashKey {
	return HashKey{
		Type:  i.Type(),
		Value: uint64(i.Value),
	}
}

// Float represents an integer.
type Float struct {
	Value float64
}

// Type returns the type of f.
func (f *Float) Type() Type {
	return FloatType
}

// Inspect returns a string representation of f.
func (f *Float) Inspect() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

// HashKey returns a hash key object for f.
func (f *Float) HashKey() HashKey {
	s := strconv.FormatFloat(f.Value, 'f', -1, 64)
	h := fnv.New64a()
	h.Write([]byte(s))

	return HashKey{
		Type:  f.Type(),
		Value: h.Sum64(),
	}
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

// HashKey returns a hash key object for b.
func (b *Boolean) HashKey() HashKey {
	key := HashKey{Type: b.Type()}
	if b.Value {
		key.Value = 1
	}
	return key
}

// Nil represents the absence of any value.
type Nil struct{}

// Type returns the type of the Nil.
func (n *Nil) Type() Type {
	return NilType
}

// Inspect returns a string representation of the Nil.
func (n *Nil) Inspect() string {
	return "nil"
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
	return "Error: " + e.Message
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

// HashKey returns a hash key object for s.
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}

// BuiltinFunction represents a function signature of builtin functions.
type BuiltinFunction func(args ...Object) Object

// Builtin represents a builtin function.
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the type of the Builtin.
func (b *Builtin) Type() Type {
	return BuiltinType
}

// Inspect returns a string representation of the Builtin.
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// Array represents an array.
type Array struct {
	Elements []Object
}

// Type returns the type of the Array.
func (*Array) Type() Type {
	return ArrayType
}

// Inspect returns a string representation of the Array.
func (a *Array) Inspect() string {
	if a == nil {
		return ""
	}

	elements := make([]string, 0, len(a.Elements))
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// HashPair represents a key-value pair in a hash.
type HashPair struct {
	Key   Object
	Value Object
}

// Hash represents a hash.
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type returns the type of the Hash.
func (*Hash) Type() Type {
	return HashType
}

// Inspect returns a string representation of the Hash.
func (h *Hash) Inspect() string {
	if h == nil {
		return ""
	}

	pairs := make([]string, 0, len(h.Pairs))
	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+": "+pair.Value.Inspect())
	}

	var out bytes.Buffer
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
