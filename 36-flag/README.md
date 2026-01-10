# Flag

## What is this topic?

The flag package provides command-line flag parsing. This project covers defining and parsing command-line flags.

## Define Flags

```go
var name = flag.String("name", "default", "usage")
var age = flag.Int("age", 0, "usage")
```

## Parse Flags

```go
flag.Parse()
```

## Key Concepts

- **String/Int/Bool**: Define flag types
- **Parse**: Parse command-line arguments
- **Usage**: Display help message
