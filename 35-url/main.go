package main

import (
	"fmt"
	"net/url"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: URL                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔗 URL PARSING")
	fmt.Println("   u, _ := url.Parse(\"https://example.com/path?key=value\")")
	fmt.Println()
}

func exampleURLParse() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: URL Parsing")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	u, _ := url.Parse("https://example.com/path?key=value")
	
	fmt.Println("Code:")
	fmt.Println("  u, _ := url.Parse(\"https://example.com/path?key=value\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Scheme: %s\n", u.Scheme)
	fmt.Printf("  Host: %s\n", u.Host)
	fmt.Printf("  Path: %s\n", u.Path)
	fmt.Printf("  Query: %s\n", u.RawQuery)
	fmt.Println()
}

func main() {
	printTheory()
	exampleURLParse()
}
