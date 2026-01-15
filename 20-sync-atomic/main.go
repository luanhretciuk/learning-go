package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
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
	fmt.Println("   • Operations are indivisible (all-or-nothing)")
	fmt.Println("   • No intermediate state visible to other goroutines")
	fmt.Println("   • Faster than mutex for simple operations")
	fmt.Println("   • Hardware-level guarantees")
	fmt.Println()
	fmt.Println("📋 BASIC OPERATIONS")
	fmt.Println("   atomic.AddInt64(&counter, 1)    // add")
	fmt.Println("   atomic.LoadInt64(&counter)      // read")
	fmt.Println("   atomic.StoreInt64(&counter, 100) // write")
	fmt.Println("   atomic.SwapInt64(&counter, 50)   // swap")
	fmt.Println()
	fmt.Println("🔄 COMPARE AND SWAP (CAS)")
	fmt.Println("   Conditional atomic update")
	fmt.Println("   • atomic.CompareAndSwapInt64(&var, old, new)")
	fmt.Println("   • Updates only if current value == old")
	fmt.Println("   • Returns true if swapped, false otherwise")
	fmt.Println("   • Foundation for lock-free algorithms")
	fmt.Println()
	fmt.Println("📦 ATOMIC TYPES (Go 1.19+)")
	fmt.Println("   Convenient wrapper types")
	fmt.Println("   • atomic.Int64, atomic.Uint64, atomic.Bool")
	fmt.Println("   • atomic.Pointer[T] for type-safe pointers")
	fmt.Println("   • Methods: Add(), Load(), Store(), Swap()")
	fmt.Println()
	fmt.Println("⚖️  ATOMIC vs MUTEX")
	fmt.Println("   Use Atomic for:")
	fmt.Println("   • Simple operations (increment, read, write)")
	fmt.Println("   • Single variable updates")
	fmt.Println("   • Performance-critical code")
	fmt.Println()
	fmt.Println("   Use Mutex for:")
	fmt.Println("   • Complex operations (multiple statements)")
	fmt.Println("   • Multiple variables together")
	fmt.Println("   • Read-write patterns (use RWMutex)")
	fmt.Println()
}

func exampleAtomicOperations() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Atomic Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var counter int64

	fmt.Println("Code:")
	fmt.Println("  var counter int64")
	fmt.Println("  atomic.AddInt64(&counter, 1)")
	fmt.Println("  atomic.AddInt64(&counter, 1)")
	fmt.Println("  value := atomic.LoadInt64(&counter)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • AddInt64 atomically adds to counter")
	fmt.Println("  • LoadInt64 atomically reads counter")
	fmt.Println("  • No race conditions - thread-safe")
	fmt.Println("  • Much faster than mutex for simple ops")
	fmt.Println()

	atomic.AddInt64(&counter, 1)
	atomic.AddInt64(&counter, 1)
	value := atomic.LoadInt64(&counter)

	fmt.Println("Output:")
	fmt.Printf("  Counter value: %d\n", value)
	fmt.Println()
}

func exampleAtomicConcurrent() {
	fmt.Println("Example: Atomic Operations - Concurrent Access")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var counter int64

	fmt.Println("Code:")
	fmt.Println("  var counter int64")
	fmt.Println("  for i := 0; i < 10; i++ {")
	fmt.Println("      go func() {")
	fmt.Println("          atomic.AddInt64(&counter, 1)")
	fmt.Println("      }()")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • 10 goroutines increment concurrently")
	fmt.Println("  • Atomic operations prevent race conditions")
	fmt.Println("  • Final value is always 10 (correct)")
	fmt.Println()

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Output:")
	fmt.Printf("  Counter: %d (correct with atomic)\n", atomic.LoadInt64(&counter))
	fmt.Println()
}

func exampleCompareAndSwap() {
	fmt.Println("Example: Compare and Swap (CAS)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var counter int64
	atomic.StoreInt64(&counter, 10)

	fmt.Println("Code:")
	fmt.Println("  atomic.StoreInt64(&counter, 10)")
	fmt.Println("  old := atomic.LoadInt64(&counter)  // 10")
	fmt.Println("  new := old + 1                      // 11")
	fmt.Println("  swapped := atomic.CompareAndSwapInt64(&counter, old, new)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • CAS updates only if value hasn't changed")
	fmt.Println("  • If counter == old, set to new and return true")
	fmt.Println("  • If counter != old, do nothing and return false")
	fmt.Println("  • Useful for optimistic locking")
	fmt.Println()

	old := atomic.LoadInt64(&counter)
	new := old + 1

	fmt.Println("Output:")
	swapped := atomic.CompareAndSwapInt64(&counter, old, new)
	fmt.Printf("  Swapped: %t\n", swapped)
	fmt.Printf("  Counter: %d\n", atomic.LoadInt64(&counter))

	// Try again (counter is now 11, not 10)
	old = 10
	new = 20
	swapped = atomic.CompareAndSwapInt64(&counter, old, new)
	fmt.Printf("  Second attempt (old=10, counter=11): swapped=%t\n", swapped)
	fmt.Printf("  Counter still: %d\n", atomic.LoadInt64(&counter))
	fmt.Println()
}

func exampleAtomicTypes() {
	fmt.Println("Example: Atomic Types (Go 1.19+)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var counter atomic.Int64

	fmt.Println("Code:")
	fmt.Println("  var counter atomic.Int64")
	fmt.Println("  counter.Add(1)")
	fmt.Println("  counter.Add(1)")
	fmt.Println("  value := counter.Load()")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Modern API (Go 1.19+)")
	fmt.Println("  • No need for pointers")
	fmt.Println("  • Type-safe and cleaner")
	fmt.Println("  • Same performance as pointer-based API")
	fmt.Println()

	counter.Add(1)
	counter.Add(1)
	value := counter.Load()

	fmt.Println("Output:")
	fmt.Printf("  Counter value: %d\n", value)
	fmt.Println()
}

func exampleAtomicVsMutex() {
	fmt.Println("Example: Atomic vs Mutex Performance")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var atomicCounter int64
	var mutexCounter int64
	var mu sync.Mutex

	iterations := 100000

	// Atomic version
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	atomicTime := time.Since(start)

	// Mutex version
	start = time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				mu.Lock()
				mutexCounter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)

	fmt.Println("Code:")
	fmt.Println("  // Atomic: atomic.AddInt64(&counter, 1)")
	fmt.Println("  // Mutex:  mu.Lock(); counter++; mu.Unlock()")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Atomic operations are faster for simple ops")
	fmt.Println("  • Mutex has overhead (lock/unlock)")
	fmt.Println("  • Use atomic when possible")
	fmt.Println()

	fmt.Println("Output:")
	fmt.Printf("  Atomic counter: %d (time: %v)\n", atomicCounter, atomicTime)
	fmt.Printf("  Mutex counter:  %d (time: %v)\n", mutexCounter, mutexTime)
	fmt.Printf("  Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
	fmt.Println()
}

func exampleAtomicBool() {
	fmt.Println("Example: Atomic Bool")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var flag atomic.Bool

	fmt.Println("Code:")
	fmt.Println("  var flag atomic.Bool")
	fmt.Println("  flag.Store(true)")
	fmt.Println("  value := flag.Load()")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Thread-safe boolean flag")
	fmt.Println("  • Common pattern for shutdown signals")
	fmt.Println("  • No need for mutex")
	fmt.Println()

	flag.Store(true)
	value := flag.Load()

	fmt.Println("Output:")
	fmt.Printf("  Flag value: %t\n", value)
	fmt.Println()
}

func exampleCASLoop() {
	fmt.Println("Example: CAS in Loop (Lock-Free Increment)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var counter int64

	increment := func() {
		for {
			old := atomic.LoadInt64(&counter)
			new := old + 1
			if atomic.CompareAndSwapInt64(&counter, old, new) {
				return // Success!
			}
			// Failed, retry (another goroutine changed value)
		}
	}

	fmt.Println("Code:")
	fmt.Println("  for {")
	fmt.Println("      old := atomic.LoadInt64(&counter)")
	fmt.Println("      new := old + 1")
	fmt.Println("      if atomic.CompareAndSwapInt64(&counter, old, new) {")
	fmt.Println("          return  // success")
	fmt.Println("      }")
	fmt.Println("      // retry if CAS failed")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Lock-free algorithm using CAS")
	fmt.Println("  • Retries if another goroutine modified value")
	fmt.Println("  • No mutex needed")
	fmt.Println("  • Foundation for complex lock-free data structures")
	fmt.Println()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()

	fmt.Println("Output:")
	fmt.Printf("  Counter: %d (lock-free increment)\n", atomic.LoadInt64(&counter))
	fmt.Println()
}

func main() {
	printTheory()
	exampleAtomicOperations()
	exampleAtomicConcurrent()
	exampleCompareAndSwap()
	exampleAtomicTypes()
	exampleAtomicVsMutex()
	exampleAtomicBool()
	exampleCASLoop()
}
