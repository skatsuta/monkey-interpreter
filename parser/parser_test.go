package parser

import (
	"fmt"
	"testing"

	"github.com/skatsuta/monkey-interpreter/ast"
	"github.com/skatsuta/monkey-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	length := len(program.Statements)
	if length != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", length)
	}
	checkParserErrors(t, p)

	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdent) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	length := len(errors)
	if length == 0 {
		return
	}

	t.Errorf("parser has %d errors", length)
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	length := len(program.Statements)
	if length != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", length)
	}
	checkParserErrors(t, p)

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStmt. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program has not enough statements. got=%d", l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	testIdent(t, stmt.Expression, input)
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program has not enough statements. got=%d", l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	testIntegerLiteral(t, stmt.Expression, 5)
}

func TestParsingPrefixExpressions(t *testing.T) {
	tests := []struct {
		input        string
		operator     string
		integerValue interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		length := len(program.Statements)
		if length != 1 {
			t.Fatalf("program has not enough statements. got=%d", length)
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp not *ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Errorf("exp.operator is not %s. got=%s", tt.operator, exp.Operator)
		}

		testLiteralExpression(t, exp.Right, tt.integerValue)
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) {
	i, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
	}

	if i.Value != value {
		t.Errorf("i.Value not %d. got=%d", value, i.Value)
	}

	if i.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("i.TokenLiteral() not %d. got=%s", value, i.TokenLiteral())
	}
}

func testIdent(t *testing.T, expr ast.Expression, value string) {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		t.Errorf("expr not *ast.Ident. got=%T", expr)
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral() not %s. got=%s", value, ident.TokenLiteral())
	}
}

func testLiteralExpression(t *testing.T, expr ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, expr, int64(v))
	case int64:
		testIntegerLiteral(t, expr, v)
	case string:
		testIdent(t, expr, v)
	case bool:
		testBooleanLiteral(t, expr, v)
	default:
		t.Errorf("type of expr not handled. got=%T", expr)
	}
}

func testInfixExpression(t *testing.T, expr ast.Expression, left interface{}, operator string,
	right interface{}) {
	op, ok := expr.(*ast.InfixExpression)
	if !ok {
		t.Errorf("expr is not ast.OperatorExpression. got=%T(%s)", expr, expr)
	}

	testLiteralExpression(t, op.Left, left)

	if op.Operator != operator {
		t.Errorf("expr.Operator is not %q. got=%q", operator, op.Operator)
	}

	testLiteralExpression(t, op.Right, right)
}

func TestParsingInfixExpressions(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range tests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkParserErrors(t, p)

		l := len(program.Statements)
		if l != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, l)
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("expr not *ast.InfixExpression. got=%T", stmt.Expression)
		}

		testLiteralExpression(t, expr.Left, tt.leftValue)

		if expr.Operator != tt.operator {
			t.Errorf("expr.Operator is not %q. got=%s", tt.operator, expr.Operator)
		}

		testLiteralExpression(t, expr.Right, tt.rightValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
		{"a + add(b * c) + d", "((a + add((b * c))) + d)"},
		{"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))", "add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))"},
		{"add(a + b + c * d / f + g)", "add((((a + b) + ((c * d) / f)) + g))"},
	}

	for _, tt := range tests {
		p := New(lexer.New(tt.input))

		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func testBooleanLiteral(t *testing.T, expr ast.Expression, value bool) {
	b, ok := expr.(*ast.Boolean)
	if !ok {
		t.Errorf("b not *ast.Boolean. got=%T", expr)
	}
	if b.Value != value {
		t.Errorf("b.Value not %t. got=%t", value, b.Value)
	}
	if b.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("b.TokenLiteral() not %t. got=%s", value, b.TokenLiteral())
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkParserErrors(t, p)

		l := len(program.Statements)
		if l != 1 {
			t.Fatalf("program has not enough statements. got=%d", l)
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		testBooleanLiteral(t, stmt.Expression, tt.expected)
	}
}

func TestIfExpression(t *testing.T) {
	input := "if (x < y) { x }"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d", 1, l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	expr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Errorf("stmt.Expression is not ast.Expression. got=%T", stmt.Expression)
	}

	testInfixExpression(t, expr.Condition, "x", "<", "y")

	l = len(expr.Consequence.Statements)
	if l != 1 {
		t.Errorf("consequence is not %d statements. got=%d\n", 1, l)
	}

	cons, ok := expr.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not *ast.ExpressionStatement. got=%T",
			expr.Consequence.Statements[0])
	}

	testIdent(t, cons.Expression, "x")

	if expr.Alternative != nil {
		t.Errorf("expr.Alternative.Statements was not nil. got=%+v", expr.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := "if (x < y) { x } else { y }"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d", 1, l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	expr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Errorf("stmt.Expression is not ast.Expression. got=%T", stmt.Expression)
	}

	testInfixExpression(t, expr.Condition, "x", "<", "y")

	l = len(expr.Consequence.Statements)
	if l != 1 {
		t.Errorf("consequence is not %d statements. got=%d\n", 1, l)
	}

	cons, ok := expr.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not *ast.ExpressionStatement. got=%T",
			expr.Consequence.Statements[0])
	}

	testIdent(t, cons.Expression, "x")

	alt, ok := expr.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not *ast.ExpressionStatement. got=%T",
			expr.Alternative.Statements[0])
	}

	testIdent(t, alt.Expression, "y")
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := "fn(x, y) { x + y; }"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d", 1, l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	f, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Errorf("stmt.Expression is not *ast.FunctionLiteral. got=%T", stmt.Expression)
	}

	l = len(f.Parameters)
	if l != 2 {
		t.Fatalf("function literal parameters wrong. want=%d, got=%d\n", 2, l)
	}

	testLiteralExpression(t, f.Parameters[0], "x")
	testLiteralExpression(t, f.Parameters[1], "y")

	l = len(f.Body.Statements)
	if l != 1 {
		t.Fatalf("f.Body.Statements has not %d statements. got=%d\n", 1, l)
	}

	bodyStmt, ok := f.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("f.Body.Statements[0] is not *ast.ExpressionStatement. got=%T", f.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"fn() {}", []string{}},
		{"fn(x) {};", []string{"x"}},
		{"fn(x, y, z) {};", []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		f := stmt.Expression.(*ast.FunctionLiteral)

		if len(f.Parameters) != len(tt.expected) {
			t.Errorf("length parameters wrong. want=%d, got=%d", len(tt.expected), len(f.Parameters))
		}

		for i, ident := range tt.expected {
			testLiteralExpression(t, f.Parameters[i], ident)
		}
	}
}

func TestCallFunctionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	expr, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Errorf("stmt.Expression is not *ast.CallExpression. got=%T", stmt.Expression)
	}

	testIdent(t, expr.Function, "add")

	l = len(expr.Arguments)
	if l != 3 {
		t.Fatalf("wrong length of arguments. got=%d\n", l)
	}

	testLiteralExpression(t, expr.Arguments[0], 1)
	testInfixExpression(t, expr.Arguments[1], 2, "*", 3)
	testInfixExpression(t, expr.Arguments[2], 4, "+", 5)
}
