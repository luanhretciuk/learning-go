package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Methods                                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔧 METHODS")
	fmt.Println("   Functions with a special receiver argument")
	fmt.Println("   func (p Point) Distance() float64 { ... }")
	fmt.Println()
	fmt.Println("📋 VALUE RECEIVERS")
	fmt.Println("   func (p Point) Move(dx, dy float64)")
	fmt.Println("   • Operate on a copy")
	fmt.Println("   • Cannot modify original")
	fmt.Println()
	fmt.Println("📍 POINTER RECEIVERS")
	fmt.Println("   func (p *Point) Move(dx, dy float64)")
	fmt.Println("   • Can modify original")
	fmt.Println("   • More efficient for large structs")
	fmt.Println()
	fmt.Println("🔄 METHOD SETS")
	fmt.Println("   Go automatically handles conversion:")
	fmt.Println("   • Value receiver: can be called on value or pointer")
	fmt.Println("   • Pointer receiver: can be called on value or pointer")
	fmt.Println()
}

// Value receiver - operates on copy
func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Value receiver - doesn't modify original
func (p Point) MoveValue(dx, dy float64) Point {
	p.X += dx
	p.Y += dy
	return p
}

// Pointer receiver - modifies original
func (p *Point) MovePointer(dx, dy float64) {
	p.X += dx
	p.Y += dy
}

func exampleValueReceiver() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Value Receiver")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p := Point{3, 4}
	fmt.Println("Code:")
	fmt.Println("  p := Point{3, 4}")
	fmt.Println("  distance := p.Distance()")
	fmt.Println()
	fmt.Printf("Output:\n")
	fmt.Printf("  p: %+v\n", p)
	fmt.Printf("  distance: %.2f\n", p.Distance())
	fmt.Println()
}

func exampleValueVsPointer() {
	fmt.Println("Example: Value Receiver vs Pointer Receiver")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	
	fmt.Println("Code:")
	fmt.Println("  p1 := Point{1, 2}")
	fmt.Println("  p2 := Point{1, 2}")
	fmt.Println("  p1.MoveValue(10, 10)  // value receiver")
	fmt.Println("  p2.MovePointer(10, 10)  // pointer receiver")
	fmt.Println()
	
	p1.MoveValue(10, 10)
	p2.MovePointer(10, 10)
	
	fmt.Println("Output:")
	fmt.Printf("  p1 after MoveValue: %+v (unchanged)\n", p1)
	fmt.Printf("  p2 after MovePointer: %+v (modified)\n", p2)
	fmt.Println()
}

func exampleMethodSets() {
	fmt.Println("Example: Method Sets")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p := Point{3, 4}
	pp := &Point{5, 12}
	
	fmt.Println("Code:")
	fmt.Println("  p := Point{3, 4}")
	fmt.Println("  pp := &Point{5, 12}")
	fmt.Println("  p.Distance()  // value can call value receiver")
	fmt.Println("  pp.Distance()  // pointer can call value receiver")
	fmt.Println()
	
	fmt.Println("Output:")
	fmt.Printf("  p.Distance(): %.2f\n", p.Distance())
	fmt.Printf("  pp.Distance(): %.2f\n", pp.Distance())
	fmt.Println()
	fmt.Println("💡 Go automatically converts between value and pointer")
	fmt.Println()
}

func main() {
	printTheory()
	exampleValueReceiver()
	exampleValueVsPointer()
	exampleMethodSets()
}
