package main

import (
	"fmt"
	"math"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: First-Class Functions                     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔧 FUNCTIONS AS VALUES")
	fmt.Println("   Functions can be assigned to variables:")
	fmt.Println("   add := func(a, b int) int { return a + b }")
	fmt.Println()
	fmt.Println("📝 FUNCTION TYPES")
	fmt.Println("   You can define function types:")
	fmt.Println("   type Operation func(int, int) int")
	fmt.Println()
	fmt.Println("🔒 CLOSURES")
	fmt.Println("   Functions can capture variables from enclosing scope:")
	fmt.Println("   func counter() func() int {")
	fmt.Println("       count := 0")
	fmt.Println("       return func() int { count++; return count }")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("🎯 METHOD EXPRESSIONS")
	fmt.Println("   Methods can be called as functions:")
	fmt.Println("   dist := Point.Distance")
	fmt.Println()
}

func exampleFunctionsAsValues() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Functions as Values")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	add := func(a, b int) int {
		return a + b
	}
	
	multiply := func(a, b int) int {
		return a * b
	}
	
	fmt.Println("Code:")
	fmt.Println("  add := func(a, b int) int { return a + b }")
	fmt.Println("  multiply := func(a, b int) int { return a * b }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  add(3, 4) = %d\n", add(3, 4))
	fmt.Printf("  multiply(3, 4) = %d\n", multiply(3, 4))
	fmt.Println()
}

func exampleFunctionTypes() {
	fmt.Println("Example: Function Types")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	type Operation func(int, int) int
	
	var op Operation = func(a, b int) int {
		return a + b
	}
	
	fmt.Println("Code:")
	fmt.Println("  type Operation func(int, int) int")
	fmt.Println("  var op Operation = func(a, b int) int { return a + b }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  op(5, 3) = %d\n", op(5, 3))
	fmt.Println()
}

func exampleClosures() {
	fmt.Println("Example: Closures")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()
	
	fmt.Println("Code:")
	fmt.Println("  counter := func() func() int {")
	fmt.Println("      count := 0")
	fmt.Println("      return func() int {")
	fmt.Println("          count++")
	fmt.Println("          return count")
	fmt.Println("      }")
	fmt.Println("  }()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  counter(): %d\n", counter())
	fmt.Printf("  counter(): %d\n", counter())
	fmt.Printf("  counter(): %d\n", counter())
	fmt.Println()
	fmt.Println("💡 Note: The closure captures and modifies 'count'")
	fmt.Println()
}

type Point struct {
	X, Y float64
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func exampleMethodExpressions() {
	fmt.Println("Example: Method Expressions")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p := Point{3, 4}
	dist := Point.Distance // method expression
	
	fmt.Println("Code:")
	fmt.Println("  p := Point{3, 4}")
	fmt.Println("  dist := Point.Distance  // method expression")
	fmt.Println("  result := dist(p)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  p.Distance(): %.2f\n", p.Distance())
	fmt.Printf("  dist(p): %.2f\n", dist(p))
	fmt.Println()
}

func main() {
	printTheory()
	exampleFunctionsAsValues()
	exampleFunctionTypes()
	exampleClosures()
	exampleMethodExpressions()
}
