# Monkey interpreter

Monkey programming language interpreter designed in [_Writing An Interpreter In Go_](https://interpreterbook.com/).


## Installation

```sh
$ go get -v -u github.com/skatsuta/monkey-interpreter/...
```


## Getting started

```sh
$ monkey-interpreter
This is the Monkey programming language!
Feel free to type in commands
>> let a = 5;
>> let b = a * 99;
>> let c = if (b > a) { 99 } else { 100 };
>> c
99
>> c * b * a;
245025
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
