# Code Documentation

## Project Layout

- main.go: CLI entrypoint and process exit behavior
- main_test.go: fixture-driven acceptance tests
- lexer/token.go: token type definitions and token constructor
- lexer/lexer.go: lexical scanner implementation
- parser/ast.go: AST node definitions
- parser/parser.go: recursive descent parser and syntax validation
- tests/: JSON fixture sets

## main.go

## main

- Delegates to run with process arguments minus executable name.
- Exits process with run return code.

## run(args, stderr)

Responsibilities:
- checks command usage constraints
- reads file content
- calls validatePlaceholder
- prints errors to stderr and returns non-zero code on failure

Behavior:
- no args: exit 1 with usage
- more than one arg: exit 1 with usage
- file read failure: exit 1
- parse failure: exit 1 with parse error
- success: exit 0

## validatePlaceholder(data)

- Constructs lexer and parser.
- Calls parser.Parse.
- Returns parse error if any.

## lexer/token.go

## TokenType

Defines lexical categories used by parser:
- structural: [, ], {, }, :, ,
- literal/value: STRING, NUMBER, TRUE, FALSE, NULL
- control: EOF, Illegal

## Token

- Type: token kind
- Literal: source slice represented by token

## lexer/lexer.go

## Lexer struct

Core fields:
- input: full byte slice
- position: current byte index
- readPosition: next byte index
- ch: current byte
- line, column: position counters (currently tracked, not yet surfaced in errors)

## NewLexer(input)

- initializes state and primes first character with readChar.

## NextToken()

Main scanner function:
1. skip whitespace
2. detect keywords (true/false/null)
3. detect number literals
4. match punctuation and strings
5. emit Illegal for unknown patterns

## readChar()

- advances scanner one byte
- updates line/column accounting
- sets ch to zero-byte sentinel at EOF

## peekChar()

- reads lookahead byte without consuming it

## skipWhitespace()

- consumes space, tab, LF, CR

## readString()

- scans string between double quotes
- validates escape syntax
- rejects invalid escapes
- rejects unescaped control characters
- rejects unterminated strings

Allowed escapes:
- \"
- \\
- \/
- \b
- \f
- \n
- \r
- \t
- \uXXXX (4 hex digits)

## readNumber()

- validates JSON-like number forms:
  - optional leading minus
  - integer component
  - optional fraction
  - optional exponent
- rejects invalid leading zeros and malformed exponent/fraction

## readIdentifierOrKeyword()

- consumes alphabetic run for keyword matching

## helpers

- isDigit
- isLetter
- isHexDigit
- isValidEscapeChar

## parser/ast.go

Defines AST interfaces and concrete node types.

## Interfaces

- Node: Kind()
- Value: Node plus marker method valueNode()

## Node types

- ObjectNode with ObjectField list
- ArrayNode with element list
- StringNode
- NumberNode
- BooleanNode
- NullNode

Each node includes Position field for future source mapping.

## parser/parser.go

## Parser struct

Core parser state:
- l: lexer instance
- curToken and peekToken: current and one-token lookahead
- errors: collected parser errors
- depth: current nesting depth
- maxDepth: nesting limit

## NewParser(l)

- initializes parser
- sets defaultMaxNestingDepth
- primes token window with nextToken twice

## Parse()

Document entrypoint:
1. enforces top-level value must be object or array
2. parses value via ParseValue
3. enforces EOF after parsed document

## ParseValue()

Dispatches by token type to:
- parseObject
- parseArray
- parseString
- parseNumber
- parseBoolean
- parseNull

Returns error on unexpected token.

## expectPeek(tokenType)

- validates lookahead token type
- advances when expected
- records parser error otherwise

## enterComposite(kind) and leaveComposite()

- increments/decrements nesting depth
- enforces maximum nesting constraint
- records error when depth exceeds maxDepth

## parseObject()

Grammar shape:
- object -> { }
- object -> { string : value (, string : value)* }

Behavior:
- supports empty object
- enforces string keys and colon separator
- parses values recursively
- enforces comma-separated fields
- returns error for malformed structure

## parseArray()

Grammar shape:
- array -> [ ]
- array -> [ value (, value)* ]

Behavior:
- supports empty array
- parses elements recursively
- enforces comma-separated elements
- returns error for malformed structure

## parseString/parseNumber/parseBoolean/parseNull

Primitive node constructors from current token literal/type.

## Errors()

- returns copy of internal parser error slice.

## main_test.go

Contains acceptance tests that exercise the full CLI and parser path against fixture corpora.

Coverage includes:
- argument and file handling behavior
- step1 to step4 fixtures
- tests/test pass and fail corpus

Assertions generally verify:
- expected exit code
- whether stderr is present

## Tests Folder

- tests/step1 to tests/step4: progressive parser capability fixtures
- tests/test: broader JSON corpus with pass and fail cases

## Implementation Notes

1. Parser currently builds AST nodes even though primary runtime use is validation.
2. Position fields are placeholders for future richer diagnostics.
3. Depth and lexical checks are intentionally strict to satisfy robustness-focused fixtures.

## Suggested Future Code Improvements

1. Rename validatePlaceholder in main.go to reflect actual parser validation role.
2. Add parser and lexer unit tests for isolated behavior beyond integration fixtures.
3. Add structured error type with line/column for better diagnostics.
4. Introduce parser options struct for configurable maxDepth and strictness flags.
