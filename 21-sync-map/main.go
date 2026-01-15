package main

import (
	"fmt"
	"sync"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Sync Map                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🗺️  SYNC.MAP")
	fmt.Println("   Thread-safe map for concurrent use")
	fmt.Println("   • No need for external mutex")
	fmt.Println("   • Optimized for specific use cases")
	fmt.Println("   • Different API than regular map")
	fmt.Println()
	fmt.Println("📝 OPERATIONS")
	fmt.Println("   m.Store(key, value)              // write")
	fmt.Println("   value, ok := m.Load(key)        // read")
	fmt.Println("   value, ok := m.LoadOrStore(key, value)  // read or write")
	fmt.Println("   m.Delete(key)                  // delete")
	fmt.Println("   m.Range(func(key, value) bool)  // iterate")
	fmt.Println()
	fmt.Println("⚠️  WHEN TO USE SYNC.MAP")
	fmt.Println("   ✅ Good for:")
	fmt.Println("   • Multiple goroutines read/write different keys")
	fmt.Println("   • Write-heavy workloads")
	fmt.Println("   • Entry-specific locks would be complex")
	fmt.Println("   • Disjoint key sets per goroutine")
	fmt.Println()
	fmt.Println("   ❌ Not good for:")
	fmt.Println("   • All goroutines access same keys")
	fmt.Println("   • Read-heavy with few writes")
	fmt.Println("   • Need type safety (uses interface{})")
	fmt.Println("   • Simple use cases (regular map + mutex is simpler)")
	fmt.Println()
	fmt.Println("⚖️  SYNC.MAP vs MAP + MUTEX")
	fmt.Println("   sync.Map:")
	fmt.Println("   • Optimized for write-heavy, disjoint keys")
	fmt.Println("   • No type safety (interface{})")
	fmt.Println("   • Different API")
	fmt.Println()
	fmt.Println("   map + mutex:")
	fmt.Println("   • Simpler, more familiar")
	fmt.Println("   • Type-safe")
	fmt.Println("   • Better for read-heavy workloads")
	fmt.Println()
}

func exampleSyncMap() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Sync Map Basic Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var m sync.Map

	fmt.Println("Code:")
	fmt.Println("  var m sync.Map")
	fmt.Println("  m.Store(\"one\", 1)")
	fmt.Println("  m.Store(\"two\", 2)")
	fmt.Println("  value, ok := m.Load(\"one\")")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Store() writes a key-value pair")
	fmt.Println("  • Load() reads a value (returns ok=false if not found)")
	fmt.Println("  • Thread-safe without external synchronization")
	fmt.Println()

	m.Store("one", 1)
	m.Store("two", 2)
	m.Store("three", 3)

	value, ok := m.Load("one")
	fmt.Println("Output:")
	if ok {
		fmt.Printf("  Loaded 'one': %v\n", value)
	}

	m.Delete("two")
	_, exists := m.Load("two")
	fmt.Printf("  After delete 'two', exists: %t\n", exists)
	fmt.Println()
}

func exampleLoadOrStore() {
	fmt.Println("Example: LoadOrStore - Atomic Read-Modify-Write")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var m sync.Map

	fmt.Println("Code:")
	fmt.Println("  value, loaded := m.LoadOrStore(\"key\", \"default\")")
	fmt.Println("  // loaded=true if key existed, false if stored new value")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Atomically loads value if key exists")
	fmt.Println("  • Stores new value if key doesn't exist")
	fmt.Println("  • Returns loaded=true if value was already there")
	fmt.Println("  • Useful for initialization patterns")
	fmt.Println()

	value, loaded := m.LoadOrStore("config", "default")
	fmt.Println("Output:")
	fmt.Printf("  First call: value=%v, loaded=%t\n", value, loaded)

	value, loaded = m.LoadOrStore("config", "new")
	fmt.Printf("  Second call: value=%v, loaded=%t\n", value, loaded)
	fmt.Println()
}

func exampleRange() {
	fmt.Println("Example: Range over Sync Map")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var m sync.Map
	m.Store("a", 1)
	m.Store("b", 2)
	m.Store("c", 3)

	fmt.Println("Code:")
	fmt.Println("  m.Range(func(key, value interface{}) bool {")
	fmt.Println("      fmt.Println(key, value)")
	fmt.Println("      return true  // continue iteration")
	fmt.Println("  })")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Range() iterates over all key-value pairs")
	fmt.Println("  • Function receives key and value (both interface{})")
	fmt.Println("  • Return false to stop iteration")
	fmt.Println("  • Snapshot may not be consistent (concurrent modifications)")
	fmt.Println()

	fmt.Println("Output:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true // continue
	})
	fmt.Println()
}

func exampleConcurrentAccess() {
	fmt.Println("Example: Concurrent Access - Different Keys")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var m sync.Map
	var wg sync.WaitGroup

	fmt.Println("Code:")
	fmt.Println("  // Each goroutine writes to different keys")
	fmt.Println("  go func() { m.Store(\"key1\", 1) }()")
	fmt.Println("  go func() { m.Store(\"key2\", 2) }()")
	fmt.Println("  go func() { m.Store(\"key3\", 3) }()")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • sync.Map optimized for disjoint key access")
	fmt.Println("  • Each goroutine works with different keys")
	fmt.Println("  • Minimal contention between goroutines")
	fmt.Println("  • This is the ideal use case for sync.Map")
	fmt.Println()

	fmt.Println("Output:")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			m.Store(key, n*10)
			fmt.Printf("  Stored %s: %d\n", key, n*10)
		}(i)
	}
	wg.Wait()

	fmt.Println("  Final map contents:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("    %v: %v\n", key, value)
		return true
	})
	fmt.Println()
}

func exampleSyncMapVsMutex() {
	fmt.Println("Example: sync.Map vs map + Mutex")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("sync.Map (this example):")
	fmt.Println("  var m sync.Map")
	fmt.Println("  m.Store(\"key\", \"value\")")
	fmt.Println("  value, ok := m.Load(\"key\")")
	fmt.Println("  • No type safety (interface{})")
	fmt.Println("  • Optimized for write-heavy, disjoint keys")
	fmt.Println("  • Different API")
	fmt.Println()

	fmt.Println("map + Mutex (alternative):")
	fmt.Println("  var mu sync.Mutex")
	fmt.Println("  m := make(map[string]int)")
	fmt.Println("  mu.Lock()")
	fmt.Println("  m[\"key\"] = 1")
	fmt.Println("  mu.Unlock()")
	fmt.Println("  • Type-safe")
	fmt.Println("  • Better for read-heavy workloads")
	fmt.Println("  • More familiar API")
	fmt.Println()

	fmt.Println("When to use each:")
	fmt.Println("  • Use sync.Map: write-heavy, different keys per goroutine")
	fmt.Println("  • Use map+mutex: read-heavy, type safety needed, simple cases")
	fmt.Println()
}

func exampleTypeAssertion() {
	fmt.Println("Example: Type Assertion with sync.Map")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var m sync.Map

	m.Store("name", "Alice")
	m.Store("age", 30)
	m.Store("active", true)

	fmt.Println("Code:")
	fmt.Println("  value, ok := m.Load(\"name\")")
	fmt.Println("  if ok {")
	fmt.Println("      str := value.(string)  // type assertion")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • sync.Map stores interface{} values")
	fmt.Println("  • Need type assertion to get concrete type")
	fmt.Println("  • Type assertion can panic if wrong type")
	fmt.Println("  • Consider using map+mutex for type safety")
	fmt.Println()

	fmt.Println("Output:")
	if value, ok := m.Load("name"); ok {
		if str, ok := value.(string); ok {
			fmt.Printf("  name: %s\n", str)
		}
	}

	if value, ok := m.Load("age"); ok {
		if age, ok := value.(int); ok {
			fmt.Printf("  age: %d\n", age)
		}
	}

	if value, ok := m.Load("active"); ok {
		if active, ok := value.(bool); ok {
			fmt.Printf("  active: %t\n", active)
		}
	}
	fmt.Println()
}

func exampleWriteHeavy() {
	fmt.Println("Example: Write-Heavy Workload (sync.Map strength)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	var m sync.Map
	var wg sync.WaitGroup

	fmt.Println("Code:")
	fmt.Println("  // Many goroutines writing different keys")
	fmt.Println("  for i := 0; i < 100; i++ {")
	fmt.Println("      go func(n int) {")
	fmt.Println("          m.Store(fmt.Sprintf(\"key%d\", n), n)")
	fmt.Println("      }(i)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • sync.Map excels at write-heavy workloads")
	fmt.Println("  • Each goroutine writes to different keys")
	fmt.Println("  • Minimal contention")
	fmt.Println("  • Better performance than map+mutex in this case")
	fmt.Println()

	fmt.Println("Output (writing 20 keys concurrently):")
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			m.Store(key, n*10)
		}(i)
	}
	wg.Wait()

	count := 0
	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Printf("  Total keys stored: %d\n", count)
	fmt.Println()
}

func main() {
	printTheory()
	exampleSyncMap()
	exampleLoadOrStore()
	exampleRange()
	exampleConcurrentAccess()
	exampleSyncMapVsMutex()
	exampleTypeAssertion()
	exampleWriteHeavy()
}
