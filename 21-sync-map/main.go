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
	fmt.Println("   Concurrent map safe for concurrent use")
	fmt.Println("   var m sync.Map")
	fmt.Println()
	fmt.Println("📝 OPERATIONS")
	fmt.Println("   m.Store(\"key\", \"value\")")
	fmt.Println("   value, ok := m.Load(\"key\")")
	fmt.Println("   m.Delete(\"key\")")
	fmt.Println()
	fmt.Println("⚠️  WHEN TO USE")
	fmt.Println("   • Multiple goroutines read/write")
	fmt.Println("   • Entry-specific locks would be complex")
	fmt.Println("   • Write-heavy or disjoint key sets")
	fmt.Println()
}

func exampleSyncMap() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Sync Map Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var m sync.Map
	
	m.Store("one", 1)
	m.Store("two", 2)
	m.Store("three", 3)
	
	fmt.Println("Code:")
	fmt.Println("  var m sync.Map")
	fmt.Println("  m.Store(\"one\", 1)")
	fmt.Println("  value, ok := m.Load(\"one\")")
	fmt.Println()
	
	value, ok := m.Load("one")
	fmt.Println("Output:")
	if ok {
		fmt.Printf("  Loaded: %v\n", value)
	}
	
	m.Delete("two")
	_, exists := m.Load("two")
	fmt.Printf("  After delete, key exists: %t\n", exists)
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
	fmt.Println("      return true")
	fmt.Println("  })")
	fmt.Println()
	fmt.Println("Output:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})
	fmt.Println()
}

func main() {
	printTheory()
	exampleSyncMap()
	exampleRange()
}
