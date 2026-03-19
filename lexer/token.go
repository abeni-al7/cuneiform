package lexer

type TokenType string

const (
	LBRACKET TokenType = "["
	LBRACE TokenType = "{"
	RBRACKET TokenType = "]"
	RBRACE TokenType = "}"
	COLON TokenType = ":"
	COMMA TokenType = ","

	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"
	TRUE   TokenType = "TRUE"
	FALSE  TokenType = "FALSE"
	NULL   TokenType = "NULL"

	EOF TokenType = "EOF"
	ILLEGAL TokenType = "Illegal"
)

type Token struct {
	Type TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}