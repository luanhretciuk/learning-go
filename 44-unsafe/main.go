package main

import (
	"fmt"
	"unsafe"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Unsafe                                    ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("⚠️  UNSAFE PACKAGE")
	fmt.Println("   Low-level operations")
	fmt.Println("   Use with extreme caution")
	fmt.Println()
}

func exampleUnsafe() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Unsafe Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var x int = 42
	p := unsafe.Pointer(&x)
	size := unsafe.Sizeof(x)
	
	fmt.Println("Code:")
	fmt.Println("  var x int = 42")
	fmt.Println("  p := unsafe.Pointer(&x)")
	fmt.Println("  size := unsafe.Sizeof(x)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Value: %d\n", x)
	fmt.Printf("  Pointer: %p\n", p)
	fmt.Printf("  Size: %d bytes\n", size)
	fmt.Println()
	fmt.Println("⚠️  Warning: Use unsafe only when absolutely necessary")
	fmt.Println()
}

func main() {
	printTheory()
	exampleUnsafe()
}
