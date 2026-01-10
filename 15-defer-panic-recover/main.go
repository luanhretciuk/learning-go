package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Defer, Panic, Recover                     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("⏸️  DEFER")
	fmt.Println("   Schedules function call after surrounding function returns")
	fmt.Println("   defer fmt.Println(\"cleanup\")")
	fmt.Println()
	fmt.Println("📋 DEFER ORDER")
	fmt.Println("   Deferred calls execute in LIFO order")
	fmt.Println("   defer fmt.Println(\"first\")")
	fmt.Println("   defer fmt.Println(\"second\")")
	fmt.Println("   // prints: second, first")
	fmt.Println()
	fmt.Println("🚨 PANIC")
	fmt.Println("   Stops normal execution and begins panicking")
	fmt.Println("   panic(\"something went wrong\")")
	fmt.Println()
	fmt.Println("🔄 RECOVER")
	fmt.Println("   Regains control of a panicking goroutine")
	fmt.Println("   if r := recover(); r != nil { ... }")
	fmt.Println()
}

func exampleDefer() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Defer")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  fmt.Println(\"start\")")
	fmt.Println("  defer fmt.Println(\"deferred\")")
	fmt.Println("  fmt.Println(\"end\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  start")
	defer fmt.Println("  deferred")
	fmt.Println("  end")
	fmt.Println()
}

func exampleDeferOrder() {
	fmt.Println("Example: Defer Order (LIFO)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  defer fmt.Println(\"first\")")
	fmt.Println("  defer fmt.Println(\"second\")")
	fmt.Println("  defer fmt.Println(\"third\")")
	fmt.Println()
	fmt.Println("Output:")
	defer fmt.Println("  first")
	defer fmt.Println("  second")
	defer fmt.Println("  third")
	fmt.Println("  (deferred calls will execute in reverse order)")
	fmt.Println()
}

func examplePanic() {
	fmt.Println("Example: Panic")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  func mayPanic() {")
	fmt.Println("      panic(\"something went wrong\")")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  (This would panic if called)")
	fmt.Println("  💡 Panic stops normal execution")
	fmt.Println()
}

func exampleRecover() {
	fmt.Println("Example: Recover")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	safeCall := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Recovered from panic: %v\n", r)
			}
		}()
		fmt.Println("  Calling function that may panic...")
		panic("something went wrong")
	}
	
	fmt.Println("Code:")
	fmt.Println("  defer func() {")
	fmt.Println("      if r := recover(); r != nil {")
	fmt.Println("          fmt.Println(\"Recovered:\", r)")
	fmt.Println("      }")
	fmt.Println("  }()")
	fmt.Println()
	fmt.Println("Output:")
	safeCall()
	fmt.Println()
}

func exampleDeferCleanup() {
	fmt.Println("Example: Defer for Cleanup")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	resource := "file"
	
	fmt.Println("Code:")
	fmt.Println("  resource := \"file\"")
	fmt.Println("  defer fmt.Println(\"Closing\", resource)")
	fmt.Println("  fmt.Println(\"Using\", resource)")
	fmt.Println()
	
	defer fmt.Println("  Closing", resource)
	fmt.Println("Output:")
	fmt.Println("  Using", resource)
	fmt.Println()
}

func main() {
	printTheory()
	exampleDefer()
	exampleDeferOrder()
	examplePanic()
	exampleRecover()
	exampleDeferCleanup()
}
