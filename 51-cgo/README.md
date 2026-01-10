# CGO

## What is this topic?

CGO enables calling C code from Go. This project covers CGO basics and interoperability.

## CGO Example

```go
/*
#include <stdio.h>
void hello() {
    printf("Hello from C\n");
}
*/
import "C"

func main() {
    C.hello()
}
```

## Key Concepts

- **CGO**: C interoperability
- **import "C"**: Enable CGO
- **C types**: Use C types in Go
- **Performance**: CGO has overhead
