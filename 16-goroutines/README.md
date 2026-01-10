# Goroutines

## What is this topic?

Goroutines are lightweight threads managed by the Go runtime. This project covers creating goroutines, goroutine lifecycle, and GOMAXPROCS.

## Creating Goroutines

Use the `go` keyword to start a goroutine:

```go
go func() {
    fmt.Println("Hello from goroutine")
}()
```

## Goroutine Lifecycle

- Goroutines are lightweight (small stack)
- Managed by Go runtime scheduler
- Can run concurrently or in parallel
- Main goroutine must stay alive

## GOMAXPROCS

Controls the number of OS threads:

```go
runtime.GOMAXPROCS(4)  // use 4 threads
```

## Synchronization

Goroutines need synchronization mechanisms:
- Channels
- sync package primitives
- WaitGroups

## Key Concepts

- **go keyword**: Start new goroutine
- **Lightweight**: Small memory footprint
- **Concurrent execution**: Can run simultaneously
- **Synchronization**: Need mechanisms to coordinate
