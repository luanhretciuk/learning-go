# Reflection

## What is this topic?

The reflect package provides runtime reflection. This project covers type and value reflection.

## Type Reflection

```go
t := reflect.TypeOf(x)
```

## Value Reflection

```go
v := reflect.ValueOf(x)
```

## Key Concepts

- **TypeOf**: Get type information
- **ValueOf**: Get value information
- **Use cases**: Generic-like behavior, serialization
- **Performance**: Reflection is slower
