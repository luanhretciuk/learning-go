# Hello World

## What is this topic?

This project introduces the basic structure of a Go program. You'll learn about the fundamental building blocks that every Go program needs.

## Go Program Structure

### Package Declaration

Every Go file starts with a `package` declaration:

```go
package main
```

- `package main` is special - it tells Go this is an executable program
- Other packages (like `package fmt`) are libraries that can be imported
- The package name determines what can be accessed from other files

### The main Function

Every executable Go program must have a `main` function:

```go
func main() {
    // Your code here
}
```

- `func` is the keyword to define a function
- `main` is the entry point - Go starts executing here
- `()` means this function takes no parameters
- `{}` contains the function body

### Imports

To use code from other packages, you import them:

```go
import "fmt"
```

- `import` brings in functionality from other packages
- `"fmt"` is the format package - used for printing
- Multiple imports can be grouped:
  ```go
  import (
      "fmt"
      "os"
  )
  ```

### Printing Output

The `fmt` package provides printing functions:

- `fmt.Println()` - prints and adds a newline
- `fmt.Print()` - prints without a newline
- `fmt.Printf()` - formatted printing (like C's printf)

## Key Concepts

- **Package main**: Makes a program executable
- **func main()**: Entry point of the program
- **Imports**: Bring in functionality from other packages
- **fmt package**: Standard library for formatted I/O

## Running Go Programs

### Using `go run`

```bash
go run main.go
```

- Compiles and runs in one step
- Good for development
- Creates temporary executable

### Using `go build`

```bash
go build main.go
./main  # or main.exe on Windows
```

- Compiles to a binary file
- Creates a permanent executable
- Good for distribution

## Example Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

This simple program:
1. Declares it's in the `main` package (executable)
2. Imports `fmt` for printing
3. Defines the `main` function (entry point)
4. Prints "Hello, World!" to the console
