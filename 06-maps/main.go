package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Maps                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🗺️  MAPS")
	fmt.Println("   Go's built-in hash table implementation")
	fmt.Println("   • Key-value pairs")
	fmt.Println("   • Fast lookups")
	fmt.Println("   • Dynamic size")
	fmt.Println()
	fmt.Println("📝 MAP DECLARATION")
	fmt.Println("   var m map[string]int        // nil map")
	fmt.Println("   m := make(map[string]int)   // empty map")
	fmt.Println("   m := map[string]int{        // map literal")
	fmt.Println("       \"one\": 1,")
	fmt.Println("       \"two\": 2,")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("🔍 ACCESSING VALUES")
	fmt.Println("   value := m[\"key\"]           // get value")
	fmt.Println("   value, ok := m[\"key\"]      // get value and check existence")
	fmt.Println()
	fmt.Println("✏️  SETTING VALUES")
	fmt.Println("   m[\"key\"] = value")
	fmt.Println()
	fmt.Println("🗑️  DELETING VALUES")
	fmt.Println("   delete(m, \"key\")")
	fmt.Println()
	fmt.Println("🔄 ITERATING MAPS")
	fmt.Println("   for key, value := range m { ... }")
	fmt.Println()
}

func exampleMapCreation() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Map Creation")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var m1 map[string]int
	m2 := make(map[string]int)
	m3 := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	
	fmt.Println("Code:")
	fmt.Println("  var m1 map[string]int  // nil map")
	fmt.Println("  m2 := make(map[string]int)  // empty map")
	fmt.Println("  m3 := map[string]int{")
	fmt.Println("      \"one\": 1,")
	fmt.Println("      \"two\": 2,")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  m1: %v, nil: %t\n", m1, m1 == nil)
	fmt.Printf("  m2: %v, nil: %t\n", m2, m2 == nil)
	fmt.Printf("  m3: %v\n", m3)
	fmt.Println()
}

func exampleMapOperations() {
	fmt.Println("Example: Map Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	
	fmt.Println("Code:")
	fmt.Println("  m := make(map[string]int)")
	fmt.Println("  m[\"one\"] = 1")
	fmt.Println("  m[\"two\"] = 2")
	fmt.Println("  m[\"three\"] = 3")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  m: %v\n", m)
	fmt.Println()
	
	// Accessing
	value1 := m["one"]
	value2, ok := m["four"]
	
	fmt.Println("Code:")
	fmt.Println("  value1 := m[\"one\"]")
	fmt.Println("  value2, ok := m[\"four\"]")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  m[\"one\"]: %d\n", value1)
	fmt.Printf("  m[\"four\"]: %d, exists: %t\n", value2, ok)
	fmt.Println()
	
	// Deleting
	delete(m, "two")
	fmt.Println("Code:")
	fmt.Println("  delete(m, \"two\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  m after delete: %v\n", m)
	fmt.Println()
}

func exampleMapIteration() {
	fmt.Println("Example: Map Iteration")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	m := map[string]int{
		"apple": 5,
		"banana": 3,
		"cherry": 8,
	}
	
	fmt.Println("Code:")
	fmt.Println("  for key, value := range m {")
	fmt.Println("      fmt.Println(key, value)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  (Note: iteration order is random)")
	for key, value := range m {
		fmt.Printf("  %s: %d\n", key, value)
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleMapCreation()
	exampleMapOperations()
	exampleMapIteration()
}
