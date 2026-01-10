package main

import (
	"fmt"
	"sort"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Sort                                     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔢 SORTING")
	fmt.Println("   sort.Ints(slice)")
	fmt.Println("   sort.Strings(slice)")
	fmt.Println()
	fmt.Println("🎯 CUSTOM SORTING")
	fmt.Println("   sort.Slice(slice, func(i, j int) bool { ... })")
	fmt.Println()
}

func exampleBasicSort() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Sort")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ints := []int{3, 1, 4, 1, 5}
	strings := []string{"c", "a", "b"}
	
	fmt.Println("Code:")
	fmt.Println("  sort.Ints(ints)")
	fmt.Println("  sort.Strings(strings)")
	fmt.Println()
	
	sort.Ints(ints)
	sort.Strings(strings)
	
	fmt.Println("Output:")
	fmt.Printf("  Sorted ints: %v\n", ints)
	fmt.Printf("  Sorted strings: %v\n", strings)
	fmt.Println()
}

type Person struct {
	Name string
	Age  int
}

func exampleCustomSort() {
	fmt.Println("Example: Custom Sort")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	
	fmt.Println("Code:")
	fmt.Println("  sort.Slice(people, func(i, j int) bool {")
	fmt.Println("      return people[i].Age < people[j].Age")
	fmt.Println("  })")
	fmt.Println()
	
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	
	fmt.Println("Output:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleBasicSort()
	exampleCustomSort()
}
