package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Structs                                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🏗️  STRUCTS")
	fmt.Println("   Composite data types that group related fields")
	fmt.Println("   type Person struct {")
	fmt.Println("       Name string")
	fmt.Println("       Age  int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("🔧 INITIALIZATION")
	fmt.Println("   • Positional: Person{\"Alice\", 30}")
	fmt.Println("   • Named: Person{Name: \"Bob\", Age: 25}")
	fmt.Println("   • Zero value: var p Person")
	fmt.Println()
}

func exampleStructDefinition() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Struct Definition and Initialization")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p1 := Person{"Alice", 30}
	p2 := Person{Name: "Bob", Age: 25}
	p3 := Person{Name: "Charlie"}
	var p4 Person
	
	fmt.Println("Code:")
	fmt.Println("  p1 := Person{\"Alice\", 30}")
	fmt.Println("  p2 := Person{Name: \"Bob\", Age: 25}")
	fmt.Println("  p3 := Person{Name: \"Charlie\"}")
	fmt.Println("  var p4 Person")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  p1: %+v\n", p1)
	fmt.Printf("  p2: %+v\n", p2)
	fmt.Printf("  p3: %+v\n", p3)
	fmt.Printf("  p4: %+v (zero value)\n", p4)
	fmt.Println()
}

type Employee struct {
	Person
	Salary float64
}

func exampleEmbeddedStructs() {
	fmt.Println("Example: Embedded Structs")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	e := Employee{
		Person: Person{Name: "David", Age: 35},
		Salary: 50000,
	}
	
	fmt.Println("Code:")
	fmt.Println("  type Employee struct {")
	fmt.Println("      Person")
	fmt.Println("      Salary float64")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  e: %+v\n", e)
	fmt.Printf("  e.Name: %s (promoted field)\n", e.Name)
	fmt.Println()
}

func main() {
	printTheory()
	exampleStructDefinition()
	exampleEmbeddedStructs()
}
