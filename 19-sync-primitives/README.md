# Sync Primitives

## What is this topic?

The sync package provides synchronization primitives. This project covers Mutex, RWMutex, WaitGroup, Once, and Cond.

## Mutex

Mutual exclusion lock:

```go
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

## RWMutex

Read-write mutex (multiple readers, single writer):

```go
var rw sync.RWMutex
rw.RLock()  // read lock
rw.RUnlock()
rw.Lock()   // write lock
rw.Unlock()
```

## WaitGroup

Wait for goroutines to finish:

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

## Once

Execute function exactly once:

```go
var once sync.Once
once.Do(func() {
    // initialization
})
```

## Cond

Condition variable for signaling:

```go
cond := sync.NewCond(&mu)
cond.Wait()
cond.Signal()
```

## Key Concepts

- **Mutex**: Mutual exclusion
- **RWMutex**: Multiple readers, single writer
- **WaitGroup**: Wait for goroutines
- **Once**: Execute once
- **Cond**: Condition variables
