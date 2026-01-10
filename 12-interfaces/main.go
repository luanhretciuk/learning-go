package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Interfaces                                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔌 INTERFACES")
	fmt.Println("   Define behavior through method signatures")
	fmt.Println("   type Shape interface {")
	fmt.Println("       Area() float64")
	fmt.Println("       Perimeter() float64")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("✨ IMPLICIT IMPLEMENTATION")
	fmt.Println("   Types implement interfaces implicitly")
	fmt.Println("   • No explicit declaration needed")
	fmt.Println("   • Just implement the methods")
	fmt.Println()
	fmt.Println("📦 EMPTY INTERFACE")
	fmt.Println("   interface{} accepts any type")
	fmt.Println("   • Used for generic-like behavior")
	fmt.Println()
	fmt.Println("🔍 TYPE ASSERTIONS")
	fmt.Println("   Extract concrete type from interface")
	fmt.Println("   circle := shape.(Circle)")
	fmt.Println()
}

func exampleImplicitImplementation() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Implicit Implementation")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var shape1 Shape = Circle{Radius: 5}
	var shape2 Shape = Rectangle{Width: 4, Height: 6}
	
	fmt.Println("Code:")
	fmt.Println("  var shape1 Shape = Circle{Radius: 5}")
	fmt.Println("  var shape2 Shape = Rectangle{Width: 4, Height: 6}")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  shape1 (Circle) - Area: %.2f, Perimeter: %.2f\n",
		shape1.Area(), shape1.Perimeter())
	fmt.Printf("  shape2 (Rectangle) - Area: %.2f, Perimeter: %.2f\n",
		shape2.Area(), shape2.Perimeter())
	fmt.Println()
}

func exampleEmptyInterface() {
	fmt.Println("Example: Empty Interface")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var i interface{} = 42
	var s interface{} = "hello"
	var b interface{} = true
	
	fmt.Println("Code:")
	fmt.Println("  var i interface{} = 42")
	fmt.Println("  var s interface{} = \"hello\"")
	fmt.Println("  var b interface{} = true")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  i: %v (type: %T)\n", i, i)
	fmt.Printf("  s: %v (type: %T)\n", s, s)
	fmt.Printf("  b: %v (type: %T)\n", b, b)
	fmt.Println()
}

func exampleTypeAssertion() {
	fmt.Println("Example: Type Assertion")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var shape Shape = Circle{Radius: 5}
	
	// Type assertion
	circle, ok := shape.(Circle)
	
	fmt.Println("Code:")
	fmt.Println("  var shape Shape = Circle{Radius: 5}")
	fmt.Println("  circle, ok := shape.(Circle)")
	fmt.Println()
	fmt.Println("Output:")
	if ok {
		fmt.Printf("  Assertion successful: %+v\n", circle)
		fmt.Printf("  Circle radius: %.2f\n", circle.Radius)
	} else {
		fmt.Println("  Assertion failed")
	}
	fmt.Println()
}

func examplePolymorphism() {
	fmt.Println("Example: Polymorphism")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 2},
	}
	
	fmt.Println("Code:")
	fmt.Println("  shapes := []Shape{")
	fmt.Println("      Circle{Radius: 3},")
	fmt.Println("      Rectangle{Width: 4, Height: 5},")
	fmt.Println("      Circle{Radius: 2},")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	for i, shape := range shapes {
		fmt.Printf("  Shape %d - Area: %.2f, Perimeter: %.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleImplicitImplementation()
	exampleEmptyInterface()
	exampleTypeAssertion()
	examplePolymorphism()
}
