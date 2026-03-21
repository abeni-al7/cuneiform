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
	if !p.curTokenIs(lexer.LBRACE) && !p.curTokenIs(lexer.LBRACKET) {
		err := fmt.Errorf("expected top-level JSON object or array, got %q", p.curToken.Type)
		p.errors = append(p.errors, err)
		return nil, err
	}

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
	obj := &ObjectNode{Fields: []ObjectField{}}

	if p.peekTokenIs(lexer.RBRACE) {
		p.nextToken()
		return obj, nil
	}

	for {
		p.nextToken()
		if !p.curTokenIs(lexer.STRING) {
			err := fmt.Errorf("expected object key string, got %q", p.curToken.Type)
			p.errors = append(p.errors, err)
			return nil, err
		}

		keyValue, err := p.parseString()
		if err != nil {
			return nil, err
		}

		key, ok := keyValue.(*StringNode)
		if !ok {
			err := fmt.Errorf("expected object key node to be string")
			p.errors = append(p.errors, err)
			return nil, err
		}

		if !p.expectPeek(lexer.COLON) {
			return nil, p.errors[len(p.errors)-1]
		}

		p.nextToken()
		value, err := p.ParseValue()
		if err != nil {
			return nil, err
		}

		obj.Fields = append(obj.Fields, ObjectField{Key: key, Value: value})

		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken()
			continue
		}

		if p.peekTokenIs(lexer.RBRACE) {
			p.nextToken()
			return obj, nil
		}

		err = fmt.Errorf("expected next token %q or %q, got %q", lexer.COMMA, lexer.RBRACE, p.peekToken.Type)
		p.errors = append(p.errors, err)
		return nil, err
	}
}

func (p *Parser) parseArray() (Value, error) {
	arr := &ArrayNode{Elements: []Value{}}

	if p.peekTokenIs(lexer.RBRACKET) {
		p.nextToken()
		return arr, nil
	}

	for {
		p.nextToken()
		element, err := p.ParseValue()
		if err != nil {
			return nil, err
		}

		arr.Elements = append(arr.Elements, element)

		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken()
			continue
		}

		if p.peekTokenIs(lexer.RBRACKET) {
			p.nextToken()
			return arr, nil
		}

		err = fmt.Errorf("expected next token %q or %q, got %q", lexer.COMMA, lexer.RBRACKET, p.peekToken.Type)
		p.errors = append(p.errors, err)
		return nil, err
	}
}

func (p *Parser) parseString() (Value, error) {
	return &StringNode{Value: p.curToken.Literal}, nil
}

func (p *Parser) parseNumber() (Value, error) {
	return &NumberNode{Raw: p.curToken.Literal}, nil
}

func (p *Parser) parseBoolean() (Value, error) {
	return &BooleanNode{Value: p.curTokenIs(lexer.TRUE)}, nil
}

func (p *Parser) parseNull() (Value, error) {
	return &NullNode{}, nil
}

func (p *Parser) todo(area string) error {
	err := fmt.Errorf("%s is not implemented yet", area)
	p.errors = append(p.errors, err)
	return err
}