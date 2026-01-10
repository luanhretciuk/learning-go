package main

import (
	"fmt"
	"sync/atomic"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Sync Atomic                               ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("⚛️  ATOMIC OPERATIONS")
	fmt.Println("   Low-level atomic memory operations")
	fmt.Println("   atomic.AddInt64(&counter, 1)")
	fmt.Println("   atomic.LoadInt64(&counter)")
	fmt.Println("   atomic.StoreInt64(&counter, 100)")
	fmt.Println()
	fmt.Println("🔄 COMPARE AND SWAP")
	fmt.Println("   Conditional atomic update")
	fmt.Println("   atomic.CompareAndSwapInt64(&counter, old, new)")
	fmt.Println()
	fmt.Println("📦 ATOMIC TYPES")
	fmt.Println("   Convenient wrapper types (Go 1.19+)")
	fmt.Println("   var counter atomic.Int64")
	fmt.Println()
}

func exampleAtomicOperations() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Atomic Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var counter int64
	
	atomic.AddInt64(&counter, 1)
	atomic.AddInt64(&counter, 1)
	value := atomic.LoadInt64(&counter)
	
	fmt.Println("Code:")
	fmt.Println("  var counter int64")
	fmt.Println("  atomic.AddInt64(&counter, 1)")
	fmt.Println("  value := atomic.LoadInt64(&counter)")
	fmt.Println()
	fmt.Printf("Output:\n")
	fmt.Printf("  Counter value: %d\n", value)
	fmt.Println()
}

func exampleCompareAndSwap() {
	fmt.Println("Example: Compare and Swap")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var counter int64
	atomic.StoreInt64(&counter, 10)
	
	old := atomic.LoadInt64(&counter)
	new := old + 1
	
	fmt.Println("Code:")
	fmt.Println("  old := atomic.LoadInt64(&counter)")
	fmt.Println("  new := old + 1")
	fmt.Println("  atomic.CompareAndSwapInt64(&counter, old, new)")
	fmt.Println()
	
	swapped := atomic.CompareAndSwapInt64(&counter, old, new)
	fmt.Println("Output:")
	fmt.Printf("  Swapped: %t\n", swapped)
	fmt.Printf("  Counter: %d\n", atomic.LoadInt64(&counter))
	fmt.Println()
}

func exampleAtomicTypes() {
	fmt.Println("Example: Atomic Types (Go 1.19+)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var counter atomic.Int64
	counter.Add(1)
	counter.Add(1)
	value := counter.Load()
	
	fmt.Println("Code:")
	fmt.Println("  var counter atomic.Int64")
	fmt.Println("  counter.Add(1)")
	fmt.Println("  value := counter.Load()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Counter value: %d\n", value)
	fmt.Println()
}

func main() {
	printTheory()
	exampleAtomicOperations()
	exampleCompareAndSwap()
	exampleAtomicTypes()
}
