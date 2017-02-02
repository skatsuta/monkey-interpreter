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
>> let b = a > 3;
>> let c = a * 99;
>> if (b) { 10 } else { 1 };
10
>> let d = if (c > a) { 99 } else { 100 };
>> d
99
>> d * c * a;
245025
```
