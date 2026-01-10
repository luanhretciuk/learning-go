package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Profiling                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📊 PROFILING")
	fmt.Println("   import _ \"net/http/pprof\"")
	fmt.Println("   go tool pprof http://localhost:6060/debug/pprof/profile")
	fmt.Println()
}

func main() {
	printTheory()
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Profiling Setup")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  import _ \"net/http/pprof\"")
	fmt.Println("  go func() {")
	fmt.Println("      log.Println(http.ListenAndServe(\"localhost:6060\", nil))")
	fmt.Println("  }()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  Profiling server would start on :6060")
	fmt.Println("  (Not started in this example)")
	fmt.Println()
}
