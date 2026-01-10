# Sync Atomic

## What is this topic?

The sync/atomic package provides low-level atomic memory operations. This project covers atomic operations, atomic types, and compare-and-swap.

## Atomic Operations

```go
var counter int64
atomic.AddInt64(&counter, 1)
value := atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 100)
```

## Compare and Swap

```go
old := atomic.LoadInt64(&counter)
new := old + 1
if atomic.CompareAndSwapInt64(&counter, old, new) {
    // swap succeeded
}
```

## Atomic Types

```go
var counter atomic.Int64
counter.Add(1)
value := counter.Load()
counter.Store(100)
```

## Key Concepts

- **Atomic operations**: Thread-safe operations
- **Compare-and-swap**: Conditional atomic update
- **Atomic types**: Convenient wrapper types (Go 1.19+)
- **Low-level**: Use when sync primitives aren't enough
