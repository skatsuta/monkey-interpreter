package object

// Environment associates values with variable names.
type Environment interface {
	// Get retrieves the value of a variable named by the `name`.
	// If the variable is present in the environment the value is returned and the boolean is true.
	// Otherwise the returned value will be nil and the boolean will be false.
	Get(name string) (Object, bool)

	// Set sets the `val` of a variable named by the `name` and returns the `val` itself.
	Set(name string, val Object) Object
}

// environment implements Environment interface.
// environment is not thread safe, so do not use it in multiple goroutines.
type environment struct {
	store map[string]Object
}

// NewEnvironment returns a new object which implements Environment interface.
func NewEnvironment() Environment {
	return &environment{store: make(map[string]Object)}
}

// Get retrieves the value of a variable named by the `name`.
// If the variable is present in the environment the value is returned and the boolean is true.
// Otherwise the returned value will be nil and the boolean will be false.
func (e *environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set sets the `val` of a variable named by the `name` and returns the `val` itself.
func (e *environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
