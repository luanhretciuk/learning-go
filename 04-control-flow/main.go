package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Control Flow                              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔀 IF/ELSE STATEMENTS")
	fmt.Println("   • if condition { ... }")
	fmt.Println("   • if condition { ... } else { ... }")
	fmt.Println("   • if init; condition { ... }  (with initialization)")
	fmt.Println()
	fmt.Println("🔀 SWITCH STATEMENTS")
	fmt.Println("   • switch value { case ... }")
	fmt.Println("   • switch { case condition: ... }  (no expression)")
	fmt.Println("   • fallthrough for explicit fall-through")
	fmt.Println()
	fmt.Println("🔄 FOR LOOPS")
	fmt.Println("   • for init; condition; post { ... }  (traditional)")
	fmt.Println("   • for condition { ... }  (while-style)")
	fmt.Println("   • for { ... }  (infinite)")
	fmt.Println("   • for index, value := range collection { ... }")
	fmt.Println()
	fmt.Println("⏸️  CONTROL STATEMENTS")
	fmt.Println("   • break - exit loop")
	fmt.Println("   • continue - skip to next iteration")
	fmt.Println("   • goto - jump to label (rarely used)")
	fmt.Println()
}

func exampleIfElse() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: If/Else Statements")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	x := 5
	fmt.Println("Code:")
	fmt.Println("  x := 5")
	fmt.Println("  if x > 0 {")
	fmt.Println("      fmt.Println(\"positive\")")
	fmt.Println("  } else if x < 0 {")
	fmt.Println("      fmt.Println(\"negative\")")
	fmt.Println("  } else {")
	fmt.Println("      fmt.Println(\"zero\")")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	if x > 0 {
		fmt.Println("  positive")
	} else if x < 0 {
		fmt.Println("  negative")
	} else {
		fmt.Println("  zero")
	}
	fmt.Println()
}

func exampleSwitch() {
	fmt.Println("Example: Switch Statement")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	day := "Friday"
	fmt.Println("Code:")
	fmt.Println("  day := \"Friday\"")
	fmt.Println("  switch day {")
	fmt.Println("  case \"Monday\":")
	fmt.Println("      fmt.Println(\"Start of week\")")
	fmt.Println("  case \"Friday\":")
	fmt.Println("      fmt.Println(\"End of week\")")
	fmt.Println("  default:")
	fmt.Println("      fmt.Println(\"Mid week\")")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	switch day {
	case "Monday":
		fmt.Println("  Start of week")
	case "Friday":
		fmt.Println("  End of week")
	default:
		fmt.Println("  Mid week")
	}
	fmt.Println()
}

func exampleForLoop() {
	fmt.Println("Example: For Loop")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  for i := 0; i < 5; i++ {")
	fmt.Println("      fmt.Println(i)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d\n", i)
	}
	fmt.Println()
}

func exampleRangeLoop() {
	fmt.Println("Example: Range Loop")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	numbers := []int{10, 20, 30}
	fmt.Println("Code:")
	fmt.Println("  numbers := []int{10, 20, 30}")
	fmt.Println("  for index, value := range numbers {")
	fmt.Println("      fmt.Println(index, value)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	for index, value := range numbers {
		fmt.Printf("  index: %d, value: %d\n", index, value)
	}
	fmt.Println()
}

func exampleBreakContinue() {
	fmt.Println("Example: Break and Continue")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  for i := 0; i < 10; i++ {")
	fmt.Println("      if i == 3 { continue }")
	fmt.Println("      if i == 7 { break }")
	fmt.Println("      fmt.Println(i)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Printf("  %d\n", i)
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleIfElse()
	exampleSwitch()
	exampleForLoop()
	exampleRangeLoop()
	exampleBreakContinue()
}
