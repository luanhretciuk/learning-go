# Profiling

## What is this topic?

Go provides profiling tools via pprof. This project covers CPU and memory profiling.

## CPU Profiling

```go
import _ "net/http/pprof"
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

## Memory Profiling

```go
import "runtime/pprof"
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)
f.Close()
```

## Analyzing Profiles

```bash
go tool pprof cpu.prof
go tool pprof mem.prof
```

## Key Concepts

- **pprof**: Profiling package
- **CPU profiling**: Find bottlenecks
- **Memory profiling**: Find memory leaks
- **Analysis**: Use pprof tool
