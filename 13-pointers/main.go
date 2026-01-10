package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Pointers                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📍 POINTERS")
	fmt.Println("   Hold the memory address of a value")
	fmt.Println("   var x int = 42")
	fmt.Println("   var p *int = &x  // p points to x")
	fmt.Println()
	fmt.Println("🔗 ADDRESS OPERATOR")
	fmt.Println("   & gets the address of a variable")
	fmt.Println("   p := &x  // p is a pointer to x")
	fmt.Println()
	fmt.Println("⭐ DEREFERENCE OPERATOR")
	fmt.Println("   * gets or sets the value at address")
	fmt.Println("   value := *p  // get value")
	fmt.Println("   *p = 10      // set value")
	fmt.Println()
	fmt.Println("🚫 NIL POINTERS")
	fmt.Println("   Uninitialized pointers are nil")
	fmt.Println("   var p *int  // nil")
	fmt.Println("   Always check for nil before dereferencing")
	fmt.Println()
}

func exampleBasicPointers() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Pointers")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	x := 42
	p := &x
	
	fmt.Println("Code:")
	fmt.Println("  x := 42")
	fmt.Println("  p := &x")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  x: %d\n", x)
	fmt.Printf("  p: %p (address of x)\n", p)
	fmt.Printf("  *p: %d (value at address)\n", *p)
	fmt.Println()
}

func exampleModifyThroughPointer() {
	fmt.Println("Example: Modify Through Pointer")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	x := 10
	p := &x
	
	fmt.Println("Code:")
	fmt.Println("  x := 10")
	fmt.Println("  p := &x")
	fmt.Println("  *p = 20")
	fmt.Println()
	
	*p = 20
	
	fmt.Println("Output:")
	fmt.Printf("  After *p = 20:\n")
	fmt.Printf("  x: %d (modified through pointer)\n", x)
	fmt.Printf("  *p: %d\n", *p)
	fmt.Println()
}

func exampleNilPointers() {
	fmt.Println("Example: Nil Pointers")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var p *int
	
	fmt.Println("Code:")
	fmt.Println("  var p *int")
	fmt.Println("  if p != nil {")
	fmt.Println("      *p = 42")
	fmt.Println("  }")
	fmt.Println()
	
	fmt.Println("Output:")
	if p != nil {
		fmt.Printf("  p is not nil: %d\n", *p)
	} else {
		fmt.Println("  p is nil (uninitialized)")
	}
	fmt.Println()
}

func examplePassByReference() {
	fmt.Println("Example: Pass by Reference")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	modifyValue := func(x *int) {
		*x = 100
	}
	
	x := 10
	fmt.Println("Code:")
	fmt.Println("  x := 10")
	fmt.Println("  modifyValue(&x)")
	fmt.Println()
	
	modifyValue(&x)
	
	fmt.Println("Output:")
	fmt.Printf("  x after modifyValue: %d (modified)\n", x)
	fmt.Println()
}

func main() {
	printTheory()
	exampleBasicPointers()
	exampleModifyThroughPointer()
	exampleNilPointers()
	examplePassByReference()
}
