package token

// Type is a token type.
type Type string

const (
	// ILLEGAL is a token type for illegal tokens.
	ILLEGAL Type = "ILLEGAL"
	// EOF is a token type that represents end of file.
	EOF = "EOF"

	// IDENT is a token type for identifiers.
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT is a token type for integers.
	INT = "INT"

	// ASSIGN is a token type for assignment operators.
	ASSIGN = "="
	// PLUS is a token type for addition.
	PLUS = "+"

	// COMMA is a token type for commas.
	COMMA = ","
	// SEMICOLON is a token type for semicolons.
	SEMICOLON = ";"

	// LPAREN is a token type for left parentheses.
	LPAREN = "("
	// RPAREN is a token type for right parentheses.
	RPAREN = ")"
	// LBRACE is a token type for left braces.
	LBRACE = "{"
	// RBRACE is a token type for right braces.
	RBRACE = "}"

	// FUNCTION is a token type for functions.
	FUNCTION = "FUNCTION"
	// LET is a token type for lets.
	LET = "LET"
)

// Token represents a token which has a token type and literal.
type Token struct {
	Type    Type
	Literal string
}

// Language keywords
var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks the language keywords to see whether the given identifier is a keyword.
// If it is, it returns the keyword's Type constant. If it isn't, it just gets back IDENT.
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
