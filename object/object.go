package object

import "strconv"

// Type is a type of objects.
type Type string

const (
	// INTEGER represents a type of integers.
	INTEGER Type = "INTEGER"
	// BOOLEAN represents a type of booleans.
	BOOLEAN = "BOOLEAN"
	// NULL represents a type of null.
	NULL = "NULL"
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
	return INTEGER
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
	return BOOLEAN
}

// Inspect returns a string representation of the Boolean.
func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

// Null represents the absence of any value.
type Null struct{}

// Type returns the type of the Null.
func (n *Null) Type() Type {
	return NULL
}

// Inspect returns a string representation of the Null.
func (n *Null) Inspect() string {
	return "null"
}
