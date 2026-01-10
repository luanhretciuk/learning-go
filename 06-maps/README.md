# Maps

## What is this topic?

Maps are Go's built-in hash table implementation. This project covers map creation, access, deletion, iteration, and zero values.

## Map Declaration

### Creating Maps

```go
var m map[string]int        // nil map
m := make(map[string]int)   // empty map
m := map[string]int{        // map literal
    "one": 1,
    "two": 2,
}
```

## Map Operations

### Accessing Values

```go
value := m["key"]           // get value
value, ok := m["key"]      // get value and check existence
```

### Setting Values

```go
m["key"] = value           // set value
```

### Deleting Values

```go
delete(m, "key")           // delete key-value pair
```

### Checking Existence

```go
value, ok := m["key"]
if ok {
    // key exists
}
```

## Iterating Maps

```go
for key, value := range m {
    fmt.Println(key, value)
}
```

## Map Characteristics

- **Zero value**: `nil` (cannot add to nil map)
- **Key types**: Must be comparable (numbers, strings, arrays, structs with comparable fields)
- **Value types**: Any type
- **Order**: Iteration order is random (not insertion order)

## Key Concepts

- **Map declaration**: `make()` or literal
- **Access**: Returns value and existence flag
- **Deletion**: `delete()` function
- **Iteration**: `range` keyword
- **Zero value**: `nil` (must initialize before use)
