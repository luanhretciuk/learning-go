# Go Generate

## What is this topic?

The go:generate directive runs code generation tools. This project covers using go generate.

## Generate Directive

```go
//go:generate stringer -type=Status
type Status int
```

## Running Generate

```bash
go generate ./...
```

## Key Concepts

- **go:generate**: Code generation directive
- **stringer**: Generate String() methods
- **Custom tools**: Run any command
