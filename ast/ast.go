package ast

import "github.com/skatsuta/monkey-interpreter/token"

// Node represents an AST node.
type Node interface {
	TokenLiteral() string
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
