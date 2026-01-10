# Generics

## What is this topic?

Generics (Go 1.18+) provide type parameters. This project covers generic functions and types.

## Generic Function

```go
func Print[T any](v T) {
    fmt.Println(v)
}
```

## Generic Type

```go
type Stack[T any] struct {
    items []T
}
```

## Constraints

```go
func Add[T int | float64](a, b T) T {
    return a + b
}
```

## Key Concepts

- **Type parameters**: Generic types
- **Constraints**: Limit allowed types
- **any**: Any type constraint
- **comparable**: Comparable types
