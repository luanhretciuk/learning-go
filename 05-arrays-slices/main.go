package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Arrays & Slices                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📦 ARRAYS")
	fmt.Println("   Fixed-size sequences:")
	fmt.Println("   var arr [5]int  // array of 5 integers")
	fmt.Println("   • Fixed size (cannot grow)")
	fmt.Println("   • Value type (copied when assigned)")
	fmt.Println("   • Size is part of the type")
	fmt.Println()
	fmt.Println("🔪 SLICES")
	fmt.Println("   Dynamic arrays:")
	fmt.Println("   s := []int{1, 2, 3}")
	fmt.Println("   • Dynamic size (can grow)")
	fmt.Println("   • Reference type")
	fmt.Println("   • Built on top of arrays")
	fmt.Println()
	fmt.Println("🔧 SLICE INTERNALS")
	fmt.Println("   A slice has three components:")
	fmt.Println("   • Pointer: to underlying array")
	fmt.Println("   • Length: number of elements")
	fmt.Println("   • Capacity: max elements without reallocation")
	fmt.Println()
	fmt.Println("➕ APPEND")
	fmt.Println("   s = append(s, 4, 5)")
	fmt.Println("   • Adds elements to slice")
	fmt.Println("   • May reallocate if capacity exceeded")
	fmt.Println()
	fmt.Println("✂️  SLICING")
	fmt.Println("   s[1:3]  // from index 1 to 3 (exclusive)")
	fmt.Println("   s[:3]   // from start to 3")
	fmt.Println("   s[2:]   // from 2 to end")
	fmt.Println()
}

func exampleArrays() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Arrays")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}
	
	fmt.Println("Code:")
	fmt.Println("  var arr1 [5]int")
	fmt.Println("  arr2 := [5]int{1, 2, 3}")
	fmt.Println("  arr3 := [...]int{1, 2, 3, 4}")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  arr1: %v (zero values)\n", arr1)
	fmt.Printf("  arr2: %v\n", arr2)
	fmt.Printf("  arr3: %v (size inferred)\n", arr3)
	fmt.Println()
}

func exampleSlices() {
	fmt.Println("Example: Slices")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := make([]int, 5)
	s4 := make([]int, 5, 10)
	
	fmt.Println("Code:")
	fmt.Println("  var s1 []int  // nil slice")
	fmt.Println("  s2 := []int{1, 2, 3}")
	fmt.Println("  s3 := make([]int, 5)")
	fmt.Println("  s4 := make([]int, 5, 10)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  s1: %v, len: %d, cap: %d, nil: %t\n", s1, len(s1), cap(s1), s1 == nil)
	fmt.Printf("  s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	fmt.Printf("  s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	fmt.Printf("  s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
	fmt.Println()
}

func exampleAppend() {
	fmt.Println("Example: Append")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s := []int{1, 2, 3}
	fmt.Println("Code:")
	fmt.Println("  s := []int{1, 2, 3}")
	fmt.Println("  s = append(s, 4, 5)")
	fmt.Println()
	fmt.Printf("  Before: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	s = append(s, 4, 5)
	fmt.Printf("  After:  %v, len: %d, cap: %d\n", s, len(s), cap(s))
	fmt.Println()
}

func exampleSlicing() {
	fmt.Println("Example: Slicing")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s := []int{0, 1, 2, 3, 4}
	fmt.Println("Code:")
	fmt.Println("  s := []int{0, 1, 2, 3, 4}")
	fmt.Println("  s[1:3]  // from index 1 to 3 (exclusive)")
	fmt.Println("  s[:3]   // from start to 3")
	fmt.Println("  s[2:]   // from 2 to end")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  s:      %v\n", s)
	fmt.Printf("  s[1:3]: %v\n", s[1:3])
	fmt.Printf("  s[:3]:  %v\n", s[:3])
	fmt.Printf("  s[2:]:  %v\n", s[2:])
	fmt.Println()
}

func main() {
	printTheory()
	exampleArrays()
	exampleSlices()
	exampleAppend()
	exampleSlicing()
}
