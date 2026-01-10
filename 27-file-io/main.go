package main

import (
	"fmt"
	"os"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: File I/O                                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📖 READING FILES")
	fmt.Println("   data, err := os.ReadFile(\"file.txt\")")
	fmt.Println()
	fmt.Println("✏️  WRITING FILES")
	fmt.Println("   err := os.WriteFile(\"file.txt\", data, 0644)")
	fmt.Println()
}

func exampleReadWrite() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Read and Write File")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	filename := "example.txt"
	content := []byte("Hello, Go!")
	
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
		return
	}
	
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return
	}
	
	fmt.Println("Code:")
	fmt.Println("  os.WriteFile(\"example.txt\", content, 0644)")
	fmt.Println("  data, _ := os.ReadFile(\"example.txt\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Written and read: %s\n", string(data))
	
	os.Remove(filename) // cleanup
	fmt.Println()
}

func main() {
	printTheory()
	exampleReadWrite()
}
