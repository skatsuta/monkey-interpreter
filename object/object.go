package object

import "strconv"

// Type is a type of objects.
type Type string

const (
	// IntegerType represents a type of integers.
	IntegerType Type = "Integer"
	// BooleanType represents a type of booleans.
	BooleanType = "Boolean"
	// NullType represents a type of null.
	NullType = "Null"
	// ReturnValueType represents return values.
	ReturnValueType = "ReturnValue"
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

// Type return the type of the Integer.
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
