package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Variables & Types                         ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📝 VARIABLE DECLARATION")
	fmt.Println("   Three ways to declare:")
	fmt.Println("   1. var name string = \"Go\"")
	fmt.Println("   2. name := \"Go\"  (short declaration, most common)")
	fmt.Println("   3. var x, y int = 1, 2  (multiple)")
	fmt.Println()
	fmt.Println("🔢 BASIC TYPES")
	fmt.Println("   • Integers: int, int8, int16, int32, int64")
	fmt.Println("   • Unsigned: uint, uint8, uint16, uint32, uint64")
	fmt.Println("   • Floating: float32, float64")
	fmt.Println("   • Complex: complex64, complex128")
	fmt.Println("   • Other: bool, string, byte, rune")
	fmt.Println()
	fmt.Println("0️⃣  ZERO VALUES")
	fmt.Println("   Variables automatically initialized:")
	fmt.Println("   • Numbers: 0")
	fmt.Println("   • Booleans: false")
	fmt.Println("   • Strings: \"\" (empty)")
	fmt.Println("   • Pointers/slices/maps: nil")
	fmt.Println()
	fmt.Println("🔒 CONSTANTS")
	fmt.Println("   Declared with const:")
	fmt.Println("   const Pi = 3.14159")
	fmt.Println("   const (")
	fmt.Println("       StatusOK = 200")
	fmt.Println("       StatusNotFound = 404")
	fmt.Println("   )")
	fmt.Println()
	fmt.Println("🔢 IOTA")
	fmt.Println("   Special constant generator:")
	fmt.Println("   const (")
	fmt.Println("       Sunday = iota  // 0")
	fmt.Println("       Monday         // 1")
	fmt.Println("       Tuesday        // 2")
	fmt.Println("   )")
	fmt.Println()
}

func exampleVariableDeclaration() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Variable Declaration")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	// Method 1: var keyword
	var name string = "Go"
	var age int = 13
	
	// Method 2: Short declaration (most common)
	language := "Golang"
	version := 1.21
	
	// Method 3: Multiple variables
	var x, y int = 10, 20
	a, b := 30, 40
	
	fmt.Println("Code:")
	fmt.Println("  var name string = \"Go\"")
	fmt.Println("  var age int = 13")
	fmt.Println("  language := \"Golang\"")
	fmt.Println("  version := 1.21")
	fmt.Println("  var x, y int = 10, 20")
	fmt.Println("  a, b := 30, 40")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  name: %s, age: %d\n", name, age)
	fmt.Printf("  language: %s, version: %.2f\n", language, version)
	fmt.Printf("  x: %d, y: %d\n", x, y)
	fmt.Printf("  a: %d, b: %d\n", a, b)
	fmt.Println()
}

func exampleZeroValues() {
	fmt.Println("Example: Zero Values")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	
	fmt.Println("Code:")
	fmt.Println("  var i int")
	fmt.Println("  var f float64")
	fmt.Println("  var b bool")
	fmt.Println("  var s string")
	fmt.Println("  var p *int")
	fmt.Println()
	fmt.Println("What happens:")
	fmt.Println("  Variables are automatically initialized to zero values")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  int zero value: %d\n", i)
	fmt.Printf("  float64 zero value: %.2f\n", f)
	fmt.Printf("  bool zero value: %t\n", b)
	fmt.Printf("  string zero value: \"%s\" (empty)\n", s)
	fmt.Printf("  pointer zero value: %v (nil)\n", p)
	fmt.Println()
}

func exampleConstants() {
	fmt.Println("Example: Constants")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	const Pi = 3.14159
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)
	
	fmt.Println("Code:")
	fmt.Println("  const Pi = 3.14159")
	fmt.Println("  const (")
	fmt.Println("      StatusOK = 200")
	fmt.Println("      StatusNotFound = 404")
	fmt.Println("      StatusError = 500")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("What happens:")
	fmt.Println("  Constants are immutable values")
	fmt.Println("  Cannot be changed after declaration")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Pi: %.5f\n", Pi)
	fmt.Printf("  StatusOK: %d\n", StatusOK)
	fmt.Printf("  StatusNotFound: %d\n", StatusNotFound)
	fmt.Printf("  StatusError: %d\n", StatusError)
	fmt.Println()
}

func exampleIota() {
	fmt.Println("Example: iota Constant Generator")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	const (
		Sunday = iota // 0
		Monday        // 1
		Tuesday       // 2
		Wednesday     // 3
		Thursday      // 4
		Friday        // 5
		Saturday      // 6
	)
	
	fmt.Println("Code:")
	fmt.Println("  const (")
	fmt.Println("      Sunday = iota  // 0")
	fmt.Println("      Monday         // 1")
	fmt.Println("      Tuesday        // 2")
	fmt.Println("      ...")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("What happens:")
	fmt.Println("  iota starts at 0 and increments for each constant")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Sunday: %d\n", Sunday)
	fmt.Printf("  Monday: %d\n", Monday)
	fmt.Printf("  Tuesday: %d\n", Tuesday)
	fmt.Printf("  Wednesday: %d\n", Wednesday)
	fmt.Printf("  Thursday: %d\n", Thursday)
	fmt.Printf("  Friday: %d\n", Friday)
	fmt.Printf("  Saturday: %d\n", Saturday)
	fmt.Println()
}

func exampleTypeConversion() {
	fmt.Println("Example: Type Conversion")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)
	
	fmt.Println("Code:")
	fmt.Println("  var i int = 42")
	fmt.Println("  var f float64 = float64(i)")
	fmt.Println("  var u uint = uint(i)")
	fmt.Println()
	fmt.Println("What happens:")
	fmt.Println("  Go requires explicit type conversion")
	fmt.Println("  Cannot implicitly convert types")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  int: %d\n", i)
	fmt.Printf("  float64: %.2f\n", f)
	fmt.Printf("  uint: %d\n", u)
	fmt.Println()
}

func main() {
	printTheory()
	exampleVariableDeclaration()
	exampleZeroValues()
	exampleConstants()
	exampleIota()
	exampleTypeConversion()
}
