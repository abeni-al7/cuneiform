# Cuneiform

## CLI (Plumbing Phase)

Current CLI scope is argument and file-path plumbing only.

### Usage

```bash
go run . <path-to-json-file>
```

### Input Contract

- Exactly one positional file path is required.
- Multiple file paths are not supported yet.

### Exit Codes (Current Phase)

- `1` when no file path is provided.
- `1` when more than one path is provided.
- `1` when the file cannot be read.
- `1` for readable files because JSON validation is not implemented yet.

Real valid/invalid (`0`/`1`) behavior will be wired once parser validation is implemented.