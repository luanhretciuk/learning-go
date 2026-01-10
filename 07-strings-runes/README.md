# Strings & Runes

## What is this topic?

Strings and runes are fundamental for text processing in Go. This project covers string operations, runes, UTF-8 encoding, and the strings package.

## Strings

Strings in Go are sequences of bytes (UTF-8 encoded):

```go
s := "Hello, 世界"
len(s)  // byte length, not character count
```

## Runes

A rune is an alias for `int32` and represents a Unicode code point:

```go
r := 'A'        // rune literal
r := rune(65)   // same as 'A'
```

## UTF-8 Encoding

Go strings are UTF-8 encoded by default:
- Each character can be 1-4 bytes
- `len(string)` returns byte length
- Use `[]rune` for character count

## Strings Package

Common operations from `strings` package:

```go
strings.Contains(s, substr)
strings.HasPrefix(s, prefix)
strings.HasSuffix(s, suffix)
strings.Index(s, substr)
strings.Replace(s, old, new, n)
strings.Split(s, sep)
strings.Join(slice, sep)
strings.ToUpper(s)
strings.ToLower(s)
strings.TrimSpace(s)
```

## String Builder

For efficient string concatenation:

```go
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" World")
result := builder.String()
```

## Key Concepts

- **Strings**: UTF-8 encoded byte sequences
- **Runes**: Unicode code points
- **Byte length vs character count**: Important distinction
- **strings package**: Rich string manipulation functions
- **strings.Builder**: Efficient string building
