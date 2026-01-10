# Control Flow

## What is this topic?

Control flow statements determine the order in which statements are executed. This project covers if/else, switch, for loops, and control statements like break/continue.

## If/Else Statements

### Basic If

```go
if x > 0 {
    fmt.Println("positive")
}
```

### If-Else

```go
if x > 0 {
    fmt.Println("positive")
} else {
    fmt.Println("non-positive")
}
```

### If-Else If-Else

```go
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}
```

### If with Initialization

```go
if err := doSomething(); err != nil {
    fmt.Println("error:", err)
}
```

## Switch Statements

### Basic Switch

```go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
default:
    fmt.Println("Mid week")
}
```

### Switch without Expression

```go
switch {
case x > 0:
    fmt.Println("positive")
case x < 0:
    fmt.Println("negative")
default:
    fmt.Println("zero")
}
```

### Switch with Fallthrough

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("one or two")
}
```

## For Loops

### Traditional For Loop

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### While-Style Loop

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

### Infinite Loop

```go
for {
    // do something
    if condition {
        break
    }
}
```

### Range Loop

```go
for index, value := range slice {
    fmt.Println(index, value)
}
```

## Control Statements

- **break**: Exit loop immediately
- **continue**: Skip to next iteration
- **goto**: Jump to label (rarely used)

## Key Concepts

- **if/else**: Conditional execution
- **switch**: Multi-way branching
- **for**: Loop constructs (3 forms)
- **range**: Iterate over collections
- **break/continue**: Loop control
