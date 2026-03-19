package parser

import (
	"fmt"

	"github.com/abeni-al7/cuneiform/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
	errors    []error
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() (Value, error) {
	value, err := p.ParseValue()
	if err != nil {
		return nil, err
	}

	if !p.peekTokenIs(lexer.EOF) {
		err := fmt.Errorf("expected end of input, got %q", p.peekToken.Type)
		p.errors = append(p.errors, err)
		return nil, err
	}

	return value, nil
}

func (p *Parser) ParseValue() (Value, error) {
	switch p.curToken.Type {
	case lexer.LBRACE:
		return p.parseObject()
	case lexer.LBRACKET:
		return p.parseArray()
	case lexer.STRING:
		return p.parseString()
	case lexer.NUMBER:
		return p.parseNumber()
	case lexer.TRUE, lexer.FALSE:
		return p.parseBoolean()
	case lexer.NULL:
		return p.parseNull()
	default:
		err := fmt.Errorf("unexpected token %q", p.curToken.Type)
		p.errors = append(p.errors, err)
		return nil, err
	}
}

func (p *Parser) Errors() []error {
	out := make([]error, len(p.errors))
	copy(out, p.errors)
	return out
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(tokenType lexer.TokenType) bool {
	return p.curToken.Type == tokenType
}

func (p *Parser) peekTokenIs(tokenType lexer.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) expectPeek(tokenType lexer.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}

	err := fmt.Errorf("expected next token %q, got %q", tokenType, p.peekToken.Type)
	p.errors = append(p.errors, err)
	return false
}

func (p *Parser) parseObject() (Value, error) {
	if !p.expectPeek(lexer.RBRACE) {
		return nil, p.errors[len(p.errors)-1]
	}

	return &ObjectNode{Fields: []ObjectField{}}, nil
}

func (p *Parser) parseArray() (Value, error) {
	return nil, p.todo("parseArray")
}

func (p *Parser) parseString() (Value, error) {
	return nil, p.todo("parseString")
}

func (p *Parser) parseNumber() (Value, error) {
	return nil, p.todo("parseNumber")
}

func (p *Parser) parseBoolean() (Value, error) {
	return nil, p.todo("parseBoolean")
}

func (p *Parser) parseNull() (Value, error) {
	return nil, p.todo("parseNull")
}

func (p *Parser) todo(area string) error {
	err := fmt.Errorf("%s is not implemented yet", area)
	p.errors = append(p.errors, err)
	return err
}