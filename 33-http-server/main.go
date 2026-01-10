package main

import (
	"fmt"
	"net/http"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: HTTP Server                               ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🌐 HTTP SERVER")
	fmt.Println("   http.HandleFunc(\"/\", handler)")
	fmt.Println("   http.ListenAndServe(\":8080\", nil)")
	fmt.Println()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func exampleHTTPServer() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: HTTP Server")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  http.HandleFunc(\"/\", handler)")
	fmt.Println("  http.ListenAndServe(\":8080\", nil)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  Server would start on :8080")
	fmt.Println("  (Not started in this example)")
	fmt.Println()
}

func main() {
	printTheory()
	exampleHTTPServer()
}
