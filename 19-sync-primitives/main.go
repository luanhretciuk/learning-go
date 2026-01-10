package main

import (
	"fmt"
	"sync"
	"time"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Sync Primitives                           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔒 MUTEX")
	fmt.Println("   Mutual exclusion lock")
	fmt.Println("   mu.Lock() ... mu.Unlock()")
	fmt.Println()
	fmt.Println("📖 RWMUTEX")
	fmt.Println("   Read-write mutex")
	fmt.Println("   rw.RLock() / rw.Lock()")
	fmt.Println()
	fmt.Println("⏳ WAITGROUP")
	fmt.Println("   Wait for goroutines to finish")
	fmt.Println("   wg.Add(1) ... wg.Done() ... wg.Wait()")
	fmt.Println()
	fmt.Println("1️⃣  ONCE")
	fmt.Println("   Execute function exactly once")
	fmt.Println("   once.Do(func() { ... })")
	fmt.Println()
}

func exampleMutex() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Mutex")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var mu sync.Mutex
	counter := 0
	
	for i := 0; i < 3; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("Code:")
	fmt.Println("  var mu sync.Mutex")
	fmt.Println("  mu.Lock()")
	fmt.Println("  counter++")
	fmt.Println("  mu.Unlock()")
	fmt.Println()
	fmt.Printf("Output:\n")
	fmt.Printf("  Counter: %d\n", counter)
	fmt.Println()
}

func exampleWaitGroup() {
	fmt.Println("Example: WaitGroup")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("  Goroutine %d finished\n", n)
		}(i)
	}
	
	fmt.Println("Code:")
	fmt.Println("  var wg sync.WaitGroup")
	fmt.Println("  wg.Add(1)")
	fmt.Println("  go func() { defer wg.Done(); ... }()")
	fmt.Println("  wg.Wait()")
	fmt.Println()
	fmt.Println("Output:")
	wg.Wait()
	fmt.Println("  All goroutines finished")
	fmt.Println()
}

func exampleOnce() {
	fmt.Println("Example: Once")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var once sync.Once
	initFunc := func() {
		fmt.Println("  Initialization (executed once)")
	}
	
	fmt.Println("Code:")
	fmt.Println("  var once sync.Once")
	fmt.Println("  once.Do(initFunc)")
	fmt.Println()
	fmt.Println("Output:")
	for i := 0; i < 3; i++ {
		once.Do(initFunc)
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleMutex()
	exampleWaitGroup()
	exampleOnce()
}
