package main

import (
	"flag"
	"fmt"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Flag                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🚩 COMMAND-LINE FLAGS")
	fmt.Println("   var name = flag.String(\"name\", \"default\", \"usage\")")
	fmt.Println("   flag.Parse()")
	fmt.Println()
}

func exampleFlag() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Flag Parsing")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	name := flag.String("name", "World", "Name to greet")
	age := flag.Int("age", 0, "Age")
	
	fmt.Println("Code:")
	fmt.Println("  name := flag.String(\"name\", \"World\", \"usage\")")
	fmt.Println("  age := flag.Int(\"age\", 0, \"usage\")")
	fmt.Println("  flag.Parse()")
	fmt.Println()
	
	flag.Parse()
	
	fmt.Println("Output:")
	fmt.Printf("  Name: %s\n", *name)
	fmt.Printf("  Age: %d\n", *age)
	fmt.Println("  (Run with: go run main.go -name=Alice -age=30)")
	fmt.Println()
}

func main() {
	printTheory()
	exampleFlag()
}
