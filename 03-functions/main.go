package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Functions                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔧 FUNCTION SYNTAX")
	fmt.Println("   func name(params) returnType { ... }")
	fmt.Println("   • func = keyword to define function")
	fmt.Println("   • name = function identifier")
	fmt.Println("   • params = input parameters")
	fmt.Println("   • returnType = output type(s)")
	fmt.Println("   • {} = function body")
	fmt.Println()
	fmt.Println("↩️  MULTIPLE RETURN VALUES")
	fmt.Println("   Go functions can return multiple values:")
	fmt.Println("   func divide(a, b float64) (float64, error)")
	fmt.Println("   • Common pattern: (result, error)")
	fmt.Println("   • Enables explicit error handling")
	fmt.Println()
	fmt.Println("🏷️  NAMED RETURNS")
	fmt.Println("   You can name return values:")
	fmt.Println("   func calc(x, y int) (sum int, product int)")
	fmt.Println("   • Variables are automatically declared")
	fmt.Println("   • Can use 'naked return'")
	fmt.Println()
	fmt.Println("📦 VARIADIC FUNCTIONS")
	fmt.Println("   Functions with variable arguments:")
	fmt.Println("   func sum(numbers ...int) int")
	fmt.Println("   • ... = variadic parameter")
	fmt.Println("   • Treated as slice inside function")
	fmt.Println()
}

func add(a int, b int) int {
	return a + b
}

func exampleBasicFunction() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Function")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  func add(a int, b int) int {")
	fmt.Println("      return a + b")
	fmt.Println("  }")
	fmt.Println()
	result := add(5, 3)
	fmt.Println("Output:")
	fmt.Printf("  add(5, 3) = %d\n", result)
	fmt.Println()
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func exampleMultipleReturns() {
	fmt.Println("Example: Multiple Return Values")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  func divide(a, b float64) (float64, error) {")
	fmt.Println("      if b == 0 {")
	fmt.Println("          return 0, fmt.Errorf(\"division by zero\")")
	fmt.Println("      }")
	fmt.Println("      return a / b, nil")
	fmt.Println("  }")
	fmt.Println()
	result1, err1 := divide(10, 2)
	result2, err2 := divide(10, 0)
	fmt.Println("Output:")
	fmt.Printf("  divide(10, 2) = %.2f, error: %v\n", result1, err1)
	fmt.Printf("  divide(10, 0) = %.2f, error: %v\n", result2, err2)
	fmt.Println()
}

func calculate(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return // naked return
}

func exampleNamedReturns() {
	fmt.Println("Example: Named Returns")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  func calculate(x, y int) (sum int, product int) {")
	fmt.Println("      sum = x + y")
	fmt.Println("      product = x * y")
	fmt.Println("      return  // naked return")
	fmt.Println("  }")
	fmt.Println()
	sum, product := calculate(4, 5)
	fmt.Println("Output:")
	fmt.Printf("  calculate(4, 5) = sum: %d, product: %d\n", sum, product)
	fmt.Println()
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func exampleVariadicFunction() {
	fmt.Println("Example: Variadic Function")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  func sum(numbers ...int) int {")
	fmt.Println("      total := 0")
	fmt.Println("      for _, n := range numbers {")
	fmt.Println("          total += n")
	fmt.Println("      }")
	fmt.Println("      return total")
	fmt.Println("  }")
	fmt.Println()
	result1 := sum(1, 2, 3)
	result2 := sum(1, 2, 3, 4, 5)
	fmt.Println("Output:")
	fmt.Printf("  sum(1, 2, 3) = %d\n", result1)
	fmt.Printf("  sum(1, 2, 3, 4, 5) = %d\n", result2)
	fmt.Println()
}

func main() {
	printTheory()
	exampleBasicFunction()
	exampleMultipleReturns()
	exampleNamedReturns()
	exampleVariadicFunction()
}
