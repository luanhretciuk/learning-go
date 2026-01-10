# Packages

## What is this topic?

Packages organize Go code. This project covers package organization, imports, exports, init functions, and package initialization order.

## Package Declaration

```go
package main  // executable
package utils // library
```

## Exports

- **Exported**: Capitalized names (visible outside package)
- **Unexported**: Lowercase names (package-private)

```go
func PublicFunction() {}  // exported
func privateFunction() {} // unexported
```

## Imports

```go
import "fmt"
import (
    "fmt"
    "os"
)
```

## Init Functions

```go
func init() {
    // initialization code
}
```

## Package Initialization

1. Import dependencies
2. Initialize package-level variables
3. Run init functions (in order)

## Key Concepts

- **Package organization**: Group related code
- **Exports**: Capitalization determines visibility
- **Init functions**: Package initialization
- **Import cycle**: Packages cannot import each other
