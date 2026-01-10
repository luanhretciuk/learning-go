# Error Handling

## What is this topic?

Go uses explicit error handling. This project covers the error interface, error patterns, error wrapping, and custom errors.

## Error Interface

The `error` type is an interface:

```go
type error interface {
    Error() string
}
```

## Returning Errors

Functions return errors as the last return value:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Checking Errors

Always check errors:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
```

## Error Wrapping

Wrap errors to add context:

```go
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

## Error Checking Functions

```go
errors.Is(err, target)  // check if error is target
errors.As(err, target)  // check if error is type
```

## Custom Errors

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

## Key Concepts

- **Error interface**: `Error() string`
- **Explicit handling**: Always check errors
- **Error wrapping**: Add context with `%w`
- **Custom errors**: Define your own error types
