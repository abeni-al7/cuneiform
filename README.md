# Cuneiform

Cuneiform is a JSON validator CLI written in Go. It parses a JSON file and returns a success or failure exit code.

## Build

Compile the project into a binary named cuneiform:

```bash
go build -o cuneiform .
```

You can also run without compiling:

```bash
go run . path/to/file.json
```

## Usage

Run the compiled binary with exactly one JSON file path:

```bash
./cuneiform path/to/file.json
```

Examples:

```bash
./cuneiform tests/step4/valid2.json
./cuneiform tests/test/fail18.json
```

## Exit Codes

- 0: input is valid for the parser rules implemented in this project.
- 1: input is invalid, unreadable, or command usage is incorrect.

## Current Features

- Validates JSON input from a file path.
- Requires top-level payload to be an object or an array.
- Supports object parsing with string keys and value parsing for:
	- strings
	- numbers
	- booleans
	- null
	- nested objects
	- nested arrays
- Validates JSON string escape sequences and rejects illegal escapes (for example \x).
- Rejects invalid JSON forms in the fixture corpus, including malformed structure and over-nesting.
- Enforces a maximum nesting depth to guard against excessively deep payloads.

## What This Project Taught Me

- How to design a lexer and parser in small, testable steps.
- How to evolve CLI behavior from plumbing to real validation.
- Why fixture-driven testing is useful for parser correctness.
- How strict JSON details matter in practice (escape rules, token boundaries, structural limits).
- How to use defensive parsing techniques, like depth limits, to improve robustness.

## Contributing

Contributions are welcome.

Suggested workflow:

1. Fork the repository and create a feature branch.
2. Add or update tests first for the behavior you are changing.
3. Implement the smallest change needed to make tests pass.
4. Run the full test suite:

```bash
go test ./...
```

5. Open a pull request with a clear description of the change and the related fixtures/tests.

Please keep changes focused, preserve existing style, and avoid unrelated refactors in the same PR.