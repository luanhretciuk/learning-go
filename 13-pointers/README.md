# Pointers

## What is this topic?

Pointers hold the memory address of a value. This project covers pointer basics, when to use pointers, nil pointers, and pointer operations.

## Pointer Basics

```go
var x int = 42
var p *int = &x  // p points to x
*p = 21          // dereference: change value at address
```

## Address Operator

The `&` operator gets the address:

```go
x := 42
p := &x  // p is a pointer to x
```

## Dereference Operator

The `*` operator dereferences a pointer:

```go
p := &x
value := *p  // get value at address
*p = 10     // set value at address
```

## Nil Pointers

Pointers can be `nil`:

```go
var p *int  // nil pointer
if p != nil {
    *p = 42  // safe to dereference
}
```

## When to Use Pointers

- **Modify function arguments**: Pass by reference
- **Avoid copying large structs**: More efficient
- **Optional values**: `nil` represents absence
- **Method receivers**: Pointer receivers can modify

## Key Concepts

- **Address operator**: `&` gets address
- **Dereference operator**: `*` gets/sets value
- **Nil pointers**: Uninitialized pointers are `nil`
- **Pass by reference**: Modify original values
