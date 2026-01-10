# Select Statement

## What is this topic?

The `select` statement lets a goroutine wait on multiple channel operations. This project covers select syntax, default case, and timeout patterns.

## Basic Select

```go
select {
case msg1 := <-ch1:
    fmt.Println("received", msg1)
case msg2 := <-ch2:
    fmt.Println("received", msg2)
}
```

## Default Case

Non-blocking select:

```go
select {
case msg := <-ch:
    fmt.Println("received", msg)
default:
    fmt.Println("no message")
}
```

## Timeout Pattern

```go
select {
case msg := <-ch:
    fmt.Println("received", msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

## Key Concepts

- **Multiple channels**: Wait on multiple operations
- **Default case**: Non-blocking select
- **Timeout**: Use time.After for timeouts
- **Random selection**: If multiple cases ready, one is randomly chosen
