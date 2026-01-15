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
	fmt.Println("🔒 MUTEX (Mutual Exclusion Lock)")
	fmt.Println("   Ensures only one goroutine accesses shared resource")
	fmt.Println("   • mu.Lock() - acquire lock (blocks if locked)")
	fmt.Println("   • mu.Unlock() - release lock")
	fmt.Println("   • Protects critical sections")
	fmt.Println("   • Always use defer mu.Unlock() to prevent deadlocks")
	fmt.Println()
	fmt.Println("📖 RWMUTEX (Read-Write Mutex)")
	fmt.Println("   Allows multiple readers OR one writer")
	fmt.Println("   • rw.RLock() - acquire read lock (multiple allowed)")
	fmt.Println("   • rw.RUnlock() - release read lock")
	fmt.Println("   • rw.Lock() - acquire write lock (exclusive)")
	fmt.Println("   • rw.Unlock() - release write lock")
	fmt.Println("   • Better than Mutex when reads >> writes")
	fmt.Println()
	fmt.Println("⏳ WAITGROUP")
	fmt.Println("   Waits for a group of goroutines to finish")
	fmt.Println("   • wg.Add(n) - add n to counter")
	fmt.Println("   • wg.Done() - decrement counter by 1")
	fmt.Println("   • wg.Wait() - block until counter is 0")
	fmt.Println("   • Use defer wg.Done() for safety")
	fmt.Println()
	fmt.Println("1️⃣  ONCE")
	fmt.Println("   Executes a function exactly once")
	fmt.Println("   • once.Do(func() { ... })")
	fmt.Println("   • Thread-safe initialization")
	fmt.Println("   • Useful for singletons, lazy init")
	fmt.Println()
}

func exampleMutex() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Mutex - Protecting Shared Counter")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var mu sync.Mutex
	counter := 0

	fmt.Println("Code:")
	fmt.Println("  var mu sync.Mutex")
	fmt.Println("  counter := 0")
	fmt.Println("  go func() {")
	fmt.Println("      mu.Lock()")
	fmt.Println("      defer mu.Unlock()  // Always unlock")
	fmt.Println("      counter++")
	fmt.Println("  }()")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Without mutex: race condition (wrong results)")
	fmt.Println("  • With mutex: only one goroutine increments at a time")
	fmt.Println("  • defer ensures unlock even if panic occurs")
	fmt.Println()

	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Output:")
	fmt.Printf("  Counter: %d (correct with mutex)\n", counter)
	fmt.Println()
}

func exampleMutexWithoutLock() {
	fmt.Println("Example: Race Condition Without Mutex")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	counter := 0

	fmt.Println("Code (WRONG - no mutex):")
	fmt.Println("  counter := 0")
	fmt.Println("  go func() { counter++ }()  // Race condition!")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Multiple goroutines modify counter simultaneously")
	fmt.Println("  • Read-modify-write is not atomic")
	fmt.Println("  • Result is unpredictable (data race)")
	fmt.Println("  • Use mutex or atomic operations to fix")
	fmt.Println()

	for i := 0; i < 5; i++ {
		go func() {
			counter++ // Race condition!
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Output:")
	fmt.Printf("  Counter: %d (may be wrong due to race condition)\n", counter)
	fmt.Println("  Note: Run with 'go run -race' to detect race conditions")
	fmt.Println()
}

func exampleRWMutex() {
	fmt.Println("Example: RWMutex - Multiple Readers")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var rw sync.RWMutex
	data := make(map[string]int)
	data["count"] = 0

	read := func(id int) {
		rw.RLock()
		defer rw.RUnlock()
		value := data["count"]
		fmt.Printf("  Reader %d: %d\n", id, value)
	}

	write := func(value int) {
		rw.Lock()
		defer rw.Unlock()
		data["count"] = value
		fmt.Printf("  Writer: set to %d\n", value)
	}

	fmt.Println("Code:")
	fmt.Println("  var rw sync.RWMutex")
	fmt.Println("  rw.RLock()  // multiple readers can lock")
	fmt.Println("  rw.Lock()   // only one writer can lock")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Multiple goroutines can read simultaneously")
	fmt.Println("  • Only one goroutine can write (exclusive)")
	fmt.Println("  • Better performance than Mutex for read-heavy workloads")
	fmt.Println()

	fmt.Println("Output:")
	// Multiple readers
	for i := 0; i < 3; i++ {
		go read(i)
	}
	time.Sleep(50 * time.Millisecond)

	// One writer
	write(10)
	time.Sleep(50 * time.Millisecond)

	// More readers
	for i := 3; i < 5; i++ {
		go read(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func exampleWaitGroup() {
	fmt.Println("Example: WaitGroup - Waiting for Goroutines")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var wg sync.WaitGroup

	fmt.Println("Code:")
	fmt.Println("  var wg sync.WaitGroup")
	fmt.Println("  for i := 0; i < 3; i++ {")
	fmt.Println("      wg.Add(1)  // increment counter")
	fmt.Println("      go func(n int) {")
	fmt.Println("          defer wg.Done()  // decrement when done")
	fmt.Println("          // do work")
	fmt.Println("      }(i)")
	fmt.Println("  }")
	fmt.Println("  wg.Wait()  // block until counter is 0")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Add(1) increments the counter")
	fmt.Println("  • Done() decrements the counter")
	fmt.Println("  • Wait() blocks until counter reaches 0")
	fmt.Println("  • Better than time.Sleep() for coordination")
	fmt.Println()

	fmt.Println("Output:")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("  Goroutine %d: starting\n", n)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("  Goroutine %d: finished\n", n)
		}(i)
	}

	fmt.Println("  Main: waiting for goroutines...")
	wg.Wait()
	fmt.Println("  Main: all goroutines finished")
	fmt.Println()
}

func exampleOnce() {
	fmt.Println("Example: Once - Thread-Safe Initialization")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var once sync.Once
	initCount := 0

	initFunc := func() {
		initCount++
		fmt.Printf("  Initialization executed (count: %d)\n", initCount)
	}

	fmt.Println("Code:")
	fmt.Println("  var once sync.Once")
	fmt.Println("  initFunc := func() { ... }")
	fmt.Println("  once.Do(initFunc)  // executes only once")
	fmt.Println("  once.Do(initFunc)  // does nothing")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Do() executes function exactly once")
	fmt.Println("  • Thread-safe: multiple goroutines can call Do()")
	fmt.Println("  • Only first call executes the function")
	fmt.Println("  • Useful for lazy initialization, singletons")
	fmt.Println()

	fmt.Println("Output:")
	for i := 0; i < 5; i++ {
		go once.Do(initFunc)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  Total calls to Do(): 5")
	fmt.Printf("  Initialization executed: %d time(s)\n", initCount)
	fmt.Println()
}

func exampleMutexDeadlock() {
	fmt.Println("Example: Mutex Deadlock Prevention")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code (WRONG - can deadlock):")
	fmt.Println("  mu.Lock()")
	fmt.Println("  // ... code that might panic ...")
	fmt.Println("  mu.Unlock()  // never reached if panic!")
	fmt.Println()
	fmt.Println("Code (CORRECT - use defer):")
	fmt.Println("  mu.Lock()")
	fmt.Println("  defer mu.Unlock()  // always executes")
	fmt.Println("  // ... code that might panic ...")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Without defer: panic prevents Unlock()")
	fmt.Println("  • Locked mutex blocks all other goroutines")
	fmt.Println("  • Program deadlocks (hangs forever)")
	fmt.Println("  • Always use defer mu.Unlock()")
	fmt.Println()
}

func exampleWaitGroupPattern() {
	fmt.Println("Example: WaitGroup - Worker Pool Pattern")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var wg sync.WaitGroup
	jobs := []string{"job1", "job2", "job3", "job4", "job5"}

	processJob := func(job string) {
		defer wg.Done()
		fmt.Printf("  Processing %s\n", job)
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("  Completed %s\n", job)
	}

	fmt.Println("Code:")
	fmt.Println("  var wg sync.WaitGroup")
	fmt.Println("  for _, job := range jobs {")
	fmt.Println("      wg.Add(1)")
	fmt.Println("      go processJob(job)")
	fmt.Println("  }")
	fmt.Println("  wg.Wait()  // wait for all jobs")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Process multiple jobs concurrently")
	fmt.Println("  • WaitGroup ensures all jobs complete")
	fmt.Println("  • Main goroutine waits for all workers")
	fmt.Println()

	fmt.Println("Output:")
	for _, job := range jobs {
		wg.Add(1)
		go processJob(job)
	}

	fmt.Println("  Main: waiting for all jobs...")
	wg.Wait()
	fmt.Println("  Main: all jobs completed")
	fmt.Println()
}

func main() {
	printTheory()
	exampleMutex()
	exampleMutexWithoutLock()
	exampleRWMutex()
	exampleWaitGroup()
	exampleOnce()
	exampleMutexDeadlock()
	exampleWaitGroupPattern()
}
