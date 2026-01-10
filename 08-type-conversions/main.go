package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Type Conversions                         ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔄 TYPE CONVERSIONS")
	fmt.Println("   Go requires explicit type conversion:")
	fmt.Println("   var i int = 42")
	fmt.Println("   var f float64 = float64(i)")
	fmt.Println("   • Cannot implicitly convert types")
	fmt.Println()
	fmt.Println("🔍 TYPE ASSERTIONS")
	fmt.Println("   Used with interfaces to extract concrete types:")
	fmt.Println("   s := i.(string)        // assertion")
	fmt.Println("   s, ok := i.(string)    // safe assertion")
	fmt.Println()
	fmt.Println("🔀 TYPE SWITCHES")
	fmt.Println("   Switch on the type of an interface value:")
	fmt.Println("   switch v := i.(type) {")
	fmt.Println("   case int: ...")
	fmt.Println("   case string: ...")
	fmt.Println("   }")
	fmt.Println()
}

func exampleTypeConversion() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Type Conversion")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)
	var b byte = byte(i)
	
	fmt.Println("Code:")
	fmt.Println("  var i int = 42")
	fmt.Println("  var f float64 = float64(i)")
	fmt.Println("  var u uint = uint(i)")
	fmt.Println("  var b byte = byte(i)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  int: %d\n", i)
	fmt.Printf("  float64: %.2f\n", f)
	fmt.Printf("  uint: %d\n", u)
	fmt.Printf("  byte: %d\n", b)
	fmt.Println()
}

func exampleTypeAssertion() {
	fmt.Println("Example: Type Assertion")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var i interface{} = "hello"
	
	// Unsafe assertion
	s1 := i.(string)
	
	// Safe assertion
	s2, ok := i.(string)
	_, ok2 := i.(int)
	
	fmt.Println("Code:")
	fmt.Println("  var i interface{} = \"hello\"")
	fmt.Println("  s1 := i.(string)  // unsafe")
	fmt.Println("  s2, ok := i.(string)  // safe")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  s1: %s\n", s1)
	fmt.Printf("  s2: %s, ok: %t\n", s2, ok)
	fmt.Printf("  int assertion ok: %t\n", ok2)
	fmt.Println()
}

func exampleTypeSwitch() {
	fmt.Println("Example: Type Switch")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var i interface{} = 42
	
	fmt.Println("Code:")
	fmt.Println("  var i interface{} = 42")
	fmt.Println("  switch v := i.(type) {")
	fmt.Println("  case int:")
	fmt.Println("      fmt.Println(\"int:\", v)")
	fmt.Println("  case string:")
	fmt.Println("      fmt.Println(\"string:\", v)")
	fmt.Println("  default:")
	fmt.Println("      fmt.Println(\"unknown\")")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	switch v := i.(type) {
	case int:
		fmt.Printf("  int: %d\n", v)
	case string:
		fmt.Printf("  string: %s\n", v)
	default:
		fmt.Printf("  unknown: %v\n", v)
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleTypeConversion()
	exampleTypeAssertion()
	exampleTypeSwitch()
}
