# Defer, Panic, Recover

## What is this topic?

Defer, panic, and recover are Go's mechanisms for control flow and error handling. This project covers defer execution, panic mechanism, and recover from panic.

## Defer

`defer` schedules a function call to run after the surrounding function returns:

```go
func example() {
    defer fmt.Println("cleanup")
    fmt.Println("work")
    // "cleanup" prints after function returns
}
```

## Defer Order

Deferred calls execute in LIFO (Last In, First Out) order:

```go
defer fmt.Println("first")
defer fmt.Println("second")
// prints: second, first
```

## Panic

`panic` stops normal execution and begins panicking:

```go
panic("something went wrong")
```

## Recover

`recover` regains control of a panicking goroutine:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

## Common Patterns

- **Resource cleanup**: Close files, unlock mutexes
- **Error recovery**: Catch panics in critical sections
- **Logging**: Log function entry/exit

## Key Concepts

- **Defer**: Execute after function returns
- **Panic**: Stop normal execution
- **Recover**: Catch panics
- **LIFO order**: Deferred calls execute in reverse order
