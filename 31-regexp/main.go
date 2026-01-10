package main

import (
	"fmt"
	"regexp"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Regexp                                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔍 REGULAR EXPRESSIONS")
	fmt.Println("   Pattern matching")
	fmt.Println("   re := regexp.MustCompile(`\\d+`)")
	fmt.Println()
}

func exampleRegexp() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Regexp")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	re := regexp.MustCompile(`\d+`)
	text := "abc123def456"
	
	fmt.Println("Code:")
	fmt.Println("  re := regexp.MustCompile(`\\d+`)")
	fmt.Println("  matches := re.FindAllString(text, -1)")
	fmt.Println()
	
	matches := re.FindAllString(text, -1)
	fmt.Println("Output:")
	fmt.Printf("  Text: %s\n", text)
	fmt.Printf("  Matches: %v\n", matches)
	fmt.Println()
}

func main() {
	printTheory()
	exampleRegexp()
}
