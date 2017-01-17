package ast

import (
	"bytes"

	"github.com/skatsuta/monkey-interpreter/token"
)

// Node represents an AST node.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement.
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression.
type Expression interface {
	Node
	expressionNode()
}

// Program is a top-level AST node of a program.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the first token literal of a program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a let statement.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Ident
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns a token literal of let statement.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Ident represents an identifier.
type Ident struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Ident) expressionNode() {}

// TokenLiteral returns a token literal of an identifier.
func (i *Ident) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Ident) String() string {
	return i.Value
}

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns a token literal of return statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents an expression statement.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns a token literal of expression statement.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
