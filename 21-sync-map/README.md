# Sync Map

## What is this topic?

sync.Map is a concurrent map that is safe for concurrent use without additional locking. This project covers when and how to use sync.Map.

## Sync Map

```go
var m sync.Map
m.Store("key", "value")
value, ok := m.Load("key")
m.Delete("key")
```

## When to Use

- Multiple goroutines read, write, and overwrite entries
- Entry-specific locks would be complex
- Map is write-heavy or has disjoint key sets per goroutine

## When NOT to Use

- Most cases - regular map with mutex is often better
- When you can partition keys across goroutines
- When writes are infrequent

## Key Concepts

- **Concurrent-safe**: No additional locking needed
- **Specialized use**: Not a drop-in replacement for map
- **Performance**: Optimized for specific concurrent patterns
