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
	fmt.Println("   go func() { ... }()")
	fmt.Println()
	fmt.Println("⚡ CHARACTERISTICS")
	fmt.Println("   • Lightweight (small stack)")
	fmt.Println("   • Managed by Go scheduler")
	fmt.Println("   • Can run concurrently or in parallel")
	fmt.Println()
	fmt.Println("⚙️  GOMAXPROCS")
	fmt.Println("   Controls number of OS threads")
	fmt.Println("   runtime.GOMAXPROCS(4)")
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
	fmt.Println("Output:")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("  Goroutine %d\n", n)
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
	fmt.Println("  runtime.NumCPU()       // number of CPUs")
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

func main() {
	printTheory()
	exampleBasicGoroutine()
	exampleMultipleGoroutines()
	exampleGOMAXPROCS()
	exampleGoroutineLifecycle()
}
