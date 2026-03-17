package lexer

type TokenType string

const (
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"
	EOF TokenType = "EOF"
	ILLEGAL TokenType = "Illegal"
)

type Token struct {
	Type TokenType
	Literal string
}