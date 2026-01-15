package main

import (
	"fmt"
	"runtime"
	"time"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Goroutines                               ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🚀 GOROUTINES")
	fmt.Println("   Lightweight threads managed by Go runtime")
	fmt.Println("   • Much cheaper than OS threads (2KB stack vs 1-2MB)")
	fmt.Println("   • Managed by Go scheduler (M:N threading model)")
	fmt.Println("   • Can run concurrently or in parallel")
	fmt.Println("   • Syntax: go func() { ... }()")
	fmt.Println()
	fmt.Println("⚡ HOW GOROUTINES WORK")
	fmt.Println("   • Go runtime maps many goroutines to fewer OS threads")
	fmt.Println("   • Scheduler uses cooperative multitasking")
	fmt.Println("   • Goroutines yield at channel ops, syscalls, function calls")
	fmt.Println("   • No preemption overhead like OS threads")
	fmt.Println()
	fmt.Println("⚙️  GOMAXPROCS")
	fmt.Println("   Controls number of OS threads for goroutines")
	fmt.Println("   • Default: number of CPU cores")
	fmt.Println("   • runtime.GOMAXPROCS(0) - get current value")
	fmt.Println("   • runtime.GOMAXPROCS(n) - set to n")
	fmt.Println("   • runtime.NumCPU() - number of CPU cores")
	fmt.Println()
	fmt.Println("🔄 GOROUTINE LIFECYCLE")
	fmt.Println("   • Created with 'go' keyword")
	fmt.Println("   • Runs independently from caller")
	fmt.Println("   • Main goroutine exits -> program exits")
	fmt.Println("   • Use channels or sync primitives to coordinate")
	fmt.Println()
}

func exampleBasicGoroutine() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Goroutine")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code:")
	fmt.Println("  go func() {")
	fmt.Println("      fmt.Println(\"Hello from goroutine\")")
	fmt.Println("  }()")
	fmt.Println("  fmt.Println(\"Hello from main\")")
	fmt.Println("  time.Sleep(100 * time.Millisecond)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Main goroutine continues immediately after 'go'")
	fmt.Println("  • New goroutine runs concurrently")
	fmt.Println("  • Order of output is non-deterministic")
	fmt.Println("  • Sleep ensures goroutine finishes (not ideal, see WaitGroup)")
	fmt.Println()

	fmt.Println("Output:")
	go func() {
		fmt.Println("  Hello from goroutine")
	}()
	fmt.Println("  Hello from main")
	time.Sleep(100 * time.Millisecond) // Wait for goroutine
	fmt.Println()
}

func exampleMultipleGoroutines() {
	fmt.Println("Example: Multiple Goroutines")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code:")
	fmt.Println("  for i := 0; i < 3; i++ {")
	fmt.Println("      go func(n int) {")
	fmt.Println("          fmt.Println(\"Goroutine\", n)")
	fmt.Println("      }(i)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Each iteration creates a new goroutine")
	fmt.Println("  • All goroutines run concurrently")
	fmt.Println("  • Output order is non-deterministic")
	fmt.Println("  • Note: passing 'i' as parameter avoids closure capture issues")
	fmt.Println()

	fmt.Println("Output:")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("  Goroutine %d\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func exampleGoroutineClosure() {
	fmt.Println("Example: Goroutine Closure Issue")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code (WRONG - closure captures loop variable):")
	fmt.Println("  for i := 0; i < 3; i++ {")
	fmt.Println("      go func() {")
	fmt.Println("          fmt.Println(i)  // All print 3!")
	fmt.Println("      }()")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Code (CORRECT - pass as parameter):")
	fmt.Println("  for i := 0; i < 3; i++ {")
	fmt.Println("      go func(n int) {")
	fmt.Println("          fmt.Println(n)  // Prints 0, 1, 2")
	fmt.Println("      }(i)")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("Output (wrong way):")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("  Wrong: %d\n", i) // All print 3!
		}()
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Output (correct way):")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("  Correct: %d\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func exampleGOMAXPROCS() {
	fmt.Println("Example: GOMAXPROCS")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	procs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()

	fmt.Println("Code:")
	fmt.Println("  runtime.GOMAXPROCS(0)  // get current value")
	fmt.Println("  runtime.NumCPU()       // number of CPU cores")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • GOMAXPROCS sets max OS threads for goroutines")
	fmt.Println("  • Default equals number of CPU cores")
	fmt.Println("  • More threads = more parallelism, but more overhead")
	fmt.Println("  • Usually best to leave at default")
	fmt.Println()

	fmt.Println("Output:")
	fmt.Printf("  GOMAXPROCS: %d\n", procs)
	fmt.Printf("  NumCPU: %d\n", numCPU)
	fmt.Println()
}

func exampleGoroutineLifecycle() {
	fmt.Println("Example: Goroutine Lifecycle")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code:")
	fmt.Println("  go func() {")
	fmt.Println("      for i := 0; i < 3; i++ {")
	fmt.Println("          fmt.Println(\"Working\", i)")
	fmt.Println("          time.Sleep(100 * time.Millisecond)")
	fmt.Println("      }")
	fmt.Println("  }()")
	fmt.Println("  time.Sleep(400 * time.Millisecond)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Goroutine starts immediately")
	fmt.Println("  • Runs independently from main")
	fmt.Println("  • Main must wait or goroutine may not finish")
	fmt.Println("  • Better: use channels or WaitGroup (see sync-primitives)")
	fmt.Println()

	fmt.Println("Output:")
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("  Working %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(400 * time.Millisecond)
	fmt.Println()
}

func exampleConcurrentWork() {
	fmt.Println("Example: Concurrent Work")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	work := func(id int, duration time.Duration) {
		fmt.Printf("  Worker %d: starting\n", id)
		time.Sleep(duration)
		fmt.Printf("  Worker %d: done\n", id)
	}

	fmt.Println("Code:")
	fmt.Println("  go work(1, 200ms)")
	fmt.Println("  go work(2, 100ms)")
	fmt.Println("  go work(3, 150ms)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • All workers start concurrently")
	fmt.Println("  • They run in parallel (if GOMAXPROCS > 1)")
	fmt.Println("  • Total time ≈ longest worker, not sum")
	fmt.Println()

	fmt.Println("Output:")
	go work(1, 200*time.Millisecond)
	go work(2, 100*time.Millisecond)
	go work(3, 150*time.Millisecond)
	time.Sleep(250 * time.Millisecond)
	fmt.Println()
}

func main() {
	printTheory()
	exampleBasicGoroutine()
	exampleMultipleGoroutines()
	exampleGoroutineClosure()
	exampleGOMAXPROCS()
	exampleGoroutineLifecycle()
	exampleConcurrentWork()
}
