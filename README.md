# Monkey interpreter

Monkey programming language interpreter designed in [_Writing An Interpreter In Go_](https://interpreterbook.com/).


## Usage

```sh
$ go get -v -u github.com/skatsuta/monkey-interpreter/...
...

$ $GOPATH/bin/monkey-interpreter
This is the Monkey programming language!
Feel free to type in commands
>> 
```


## Getting started with Monkey

### Variable bindings and arithmetic expressions

```sh
>> let a = 10;
>> let b = a * 2;
>> (a + b) / 2 - 3;
12
```

### If expressions

```sh
>> let a = 10;
>> let b = a * 2;
>> let c = if (b > a) { 99 } else { 100 };
>> c
99
```

### Functions and closures

```sh
>> let multiply = fn(x, y) { x * y };
>> multiply(50 / 2, 1 * 2)
50
>> fn(x) { x + 10 }(10)
20
>> let newAdder = fn(x) { fn(y) { x + y }; };
>> let addTwo = newAdder(2);
>> addTwo(3);
5
>> let sub = fn(a, b) { a - b };
>> let applyFunc = fn(a, b, func) { func(a, b) };
>> applyFunc(10, 2, sub);
8
```

### Strings

```sh
>> let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } };
>> let hello = makeGreeter("Hello");
>> hello("skatsuta");
Hello skatsuta!
```

### Builtin functions

```sh
>> len("hello");
5
>> len("∑");
3
```
