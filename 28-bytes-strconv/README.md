# Bytes & Strconv

## What is this topic?

The bytes and strconv packages provide byte slice manipulation and string conversions. This project covers common operations.

## Bytes Package

```go
import "bytes"
var buf bytes.Buffer
buf.WriteString("hello")
result := buf.String()
```

## Strconv Package

```go
import "strconv"
s := strconv.Itoa(42)        // int to string
i, _ := strconv.Atoi("42")   // string to int
f, _ := strconv.ParseFloat("3.14", 64)
```

## Key Concepts

- **bytes.Buffer**: Efficient string building
- **strconv**: String conversions
- **Parse functions**: Convert strings to numbers
- **Format functions**: Convert numbers to strings
