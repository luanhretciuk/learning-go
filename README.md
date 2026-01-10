# 📚 Complete Go Language Learning Repository

This repository contains a comprehensive collection of Go projects to learn the Go programming language from scratch, with detailed theoretical explanations and practical examples. 

## 🎯 Objective

Master the Go programming language through practical examples with detailed theoretical explanations that are displayed during program execution. This repository covers all fundamental and advanced concepts of Go.

## 📋 Project Index

### Fundamentals (01-06)

1. **[01-hello-world](01-hello-world/)** - Basic Go program structure
2. **[02-variables-types](02-variables-types/)** - Variables, types, constants, zero values
3. **[03-functions](03-functions/)** - Functions, parameters, returns, variadic functions
4. **[04-control-flow](04-control-flow/)** - if/else, switch, for loops, break/continue
5. **[05-arrays-slices](05-arrays-slices/)** - Arrays and slices
6. **[06-maps](06-maps/)** - Maps (hash tables)

### Strings & Type System (07-09)

7. **[07-strings-runes](07-strings-runes/)** - Strings, runes, UTF-8 encoding
8. **[08-type-conversions](08-type-conversions/)** - Type conversions, type switches
9. **[09-first-class-functions](09-first-class-functions/)** - Functions as values, closures

### Object-Oriented Concepts (10-13)

10. **[10-structs](10-structs/)** - Structs and custom types
11. **[11-methods](11-methods/)** - Methods on types
12. **[12-interfaces](12-interfaces/)** - Interfaces and polymorphism
13. **[13-pointers](13-pointers/)** - Pointers and references

### Error Handling & Control Flow (14-15)

14. **[14-error-handling](14-error-handling/)** - Error handling patterns
15. **[15-defer-panic-recover](15-defer-panic-recover/)** - Defer, panic, recover

### Concurrency (16-22)

16. **[16-goroutines](16-goroutines/)** - Concurrency: goroutines
17. **[17-channels](17-channels/)** - Channels for communication
18. **[18-select](18-select/)** - Select statement
19. **[19-sync-primitives](19-sync-primitives/)** - Mutex, RWMutex, WaitGroup, Once, Cond
20. **[20-sync-atomic](20-sync-atomic/)** - Atomic operations
21. **[21-sync-map](21-sync-map/)** - sync.Map for concurrent maps
22. **[22-context](22-context/)** - Context package

### Standard Library - Core (23-28)

23. **[23-packages](23-packages/)** - Package organization
24. **[24-fmt-formatting](24-fmt-formatting/)** - fmt package: Printf, Sprintf, etc.
25. **[25-json](25-json/)** - JSON encoding/decoding
26. **[26-encoding](26-encoding/)** - Other encodings: gob, xml, base64, hex
27. **[27-file-io](27-file-io/)** - File operations
28. **[28-bytes-strconv](28-bytes-strconv/)** - bytes and strconv packages

### Standard Library - Utilities (29-32)

29. **[29-path-filepath](29-path-filepath/)** - Path manipulation
30. **[30-time](30-time/)** - time package: Time, Duration, formatting
31. **[31-regexp](31-regexp/)** - Regular expressions
32. **[32-sort](32-sort/)** - Sorting and custom sorting

### Standard Library - Networking & I/O (33-40)

33. **[33-http-server](33-http-server/)** - HTTP server basics
34. **[34-net-tcp-udp](34-net-tcp-udp/)** - TCP/UDP networking
35. **[35-url](35-url/)** - URL parsing and manipulation
36. **[36-flag](36-flag/)** - Command-line flags
37. **[37-log](37-log/)** - Logging with log package
38. **[38-os-exec](38-os-exec/)** - Executing external commands
39. **[39-os-signal](39-os-signal/)** - Signal handling
40. **[40-templates](40-templates/)** - text/template and html/template

### Testing & Debugging (41-43)

41. **[41-testing](41-testing/)** - Unit testing, benchmarks, examples
42. **[42-race-detector](42-race-detector/)** - Race detector and concurrency debugging
43. **[43-profiling](43-profiling/)** - Profiling with pprof

### Advanced Language Features (44-50)

44. **[44-unsafe](44-unsafe/)** - unsafe package
45. **[45-go-embed](45-go-embed/)** - go:embed directive
46. **[46-go-generate](46-go-generate/)** - go:generate and code generation
47. **[47-build-tags](47-build-tags/)** - Build tags and constraints
48. **[48-workspaces](48-workspaces/)** - Multi-module workspaces
49. **[49-reflection](49-reflection/)** - Reflection package
50. **[50-generics](50-generics/)** - Generics (Go 1.18+)

### System Integration & Patterns (51-52)

51. **[51-cgo](51-cgo/)** - CGO: Interoperability with C
52. **[52-advanced-patterns](52-advanced-patterns/)** - Common Go patterns and idioms

## 🚀 How to Use

### Run all projects

```bash
make all
```

### Run a specific project

```bash
make 01-hello-world
make 25-json
# etc...
```

### Show help

```bash
make help
```

## 📖 Project Structure

Each project contains:

- **README.md**: Complete theoretical explanation of the topic
- **main.go**: Code with practical examples that prints formatted theory when executed

## 🎓 Recommended Study Order

Projects are numbered in the recommended study order, starting with fundamentals and progressing to more advanced topics. Each project builds on knowledge from previous ones.

## 📝 Requirements

- Go installed (version 1.21 or higher)
- Make (to run via Makefile)

## 🔧 Installation

```bash
# Check if Go is installed
go version

# If not, install Go from https://go.dev/dl/
```

## 📚 Additional Resources

- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Go Documentation](https://pkg.go.dev/std)

---

**Happy studying! 🚀**
# learning-go
