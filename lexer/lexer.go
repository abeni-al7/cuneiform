package lexer

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
}

func NewLexer(input []byte) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.ch == 0 {
		return NewToken(EOF, "")
	}

	// TODO: Replace placeholder tokenization with full JSON lexing logic.
	tok := NewToken(ILLEGAL, string(l.ch))
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.position = l.readPosition
		l.ch = 0
		return
	}

	l.position = l.readPosition
	l.ch = l.input[l.readPosition]
	l.readPosition++

	if l.ch == '\n' {
		l.line++
		l.column = 0
		return
	}

	l.column++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	// TODO: Implement escaped-string handling and termination checks.
	return ""
}

func (l *Lexer) readNumber() string {
	// TODO: Implement integer, fraction, and exponent number lexing.
	return ""
}

func (l *Lexer) readIdentifierOrKeyword() string {
	// TODO: Implement scanning for true/false/null keywords.
	return ""
}