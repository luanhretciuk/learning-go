# Fmt Formatting

## What is this topic?

The fmt package provides formatted I/O. This project covers Printf, Sprintf, Fprintf, format verbs, and custom formatting.

## Print Functions

```go
fmt.Print("hello")
fmt.Println("hello")
fmt.Printf("Value: %d\n", 42)
```

## Format Verbs

- `%v` - default format
- `%d` - integer
- `%f` - float
- `%s` - string
- `%t` - boolean
- `%p` - pointer
- `%T` - type

## Sprintf

Format to string:

```go
s := fmt.Sprintf("Value: %d", 42)
```

## Fprintf

Format to writer:

```go
fmt.Fprintf(os.Stdout, "Value: %d", 42)
```

## Custom Formatting

Implement `Stringer` interface:

```go
func (p Point) String() string {
    return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}
```

## Key Concepts

- **Format verbs**: Control output format
- **Sprintf**: Format to string
- **Fprintf**: Format to writer
- **Stringer**: Custom string representation
