package main

import (
	"fmt"
	"reflect"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Reflection                                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔍 REFLECTION")
	fmt.Println("   Runtime type and value inspection")
	fmt.Println("   t := reflect.TypeOf(x)")
	fmt.Println("   v := reflect.ValueOf(x)")
	fmt.Println()
}

func exampleReflection() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Type and Value Reflection")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	x := 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	
	fmt.Println("Code:")
	fmt.Println("  t := reflect.TypeOf(x)")
	fmt.Println("  v := reflect.ValueOf(x)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Type: %v\n", t)
	fmt.Printf("  Value: %v\n", v)
	fmt.Println()
}

func main() {
	printTheory()
	exampleReflection()
}
