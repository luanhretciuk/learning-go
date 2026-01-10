# Context Package

## What is this topic?

The context package provides cancellation, timeouts, and request-scoped values. This project covers context creation, cancellation, timeouts, and values.

## Context Creation

```go
ctx := context.Background()  // root context
ctx := context.TODO()        // placeholder
```

## With Cancellation

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

## With Timeout

```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()
```

## With Deadline

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(parent, deadline)
defer cancel()
```

## With Values

```go
ctx := context.WithValue(parent, "key", "value")
value := ctx.Value("key")
```

## Checking Context

```go
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // continue
}
```

## Key Concepts

- **Cancellation**: Propagate cancellation signals
- **Timeouts**: Automatic cancellation after timeout
- **Values**: Request-scoped values
- **Propagation**: Context flows through call chain
