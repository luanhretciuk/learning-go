package main

import (
	"errors"
	"fmt"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Error Handling                            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("❌ ERROR INTERFACE")
	fmt.Println("   type error interface {")
	fmt.Println("       Error() string")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("↩️  RETURNING ERRORS")
	fmt.Println("   Functions return errors as last value:")
	fmt.Println("   func divide(a, b float64) (float64, error)")
	fmt.Println()
	fmt.Println("✅ CHECKING ERRORS")
	fmt.Println("   Always check errors:")
	fmt.Println("   result, err := divide(10, 0)")
	fmt.Println("   if err != nil { ... }")
	fmt.Println()
	fmt.Println("📦 ERROR WRAPPING")
	fmt.Println("   Add context with fmt.Errorf:")
	fmt.Println("   return fmt.Errorf(\"failed: %w\", err)")
	fmt.Println()
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func exampleBasicError() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Error Handling")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	result1, err1 := divide(10, 2)
	result2, err2 := divide(10, 0)
	
	fmt.Println("Code:")
	fmt.Println("  result1, err1 := divide(10, 2)")
	fmt.Println("  result2, err2 := divide(10, 0)")
	fmt.Println()
	fmt.Println("Output:")
	if err1 != nil {
		fmt.Printf("  Error: %v\n", err1)
	} else {
		fmt.Printf("  result1: %.2f\n", result1)
	}
	if err2 != nil {
		fmt.Printf("  Error: %v\n", err2)
	} else {
		fmt.Printf("  result2: %.2f\n", result2)
	}
	fmt.Println()
}

func exampleErrorWrapping() {
	fmt.Println("Example: Error Wrapping")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	process := func() error {
		if err := divide(10, 0); err != nil {
			return fmt.Errorf("failed to process division: %w", err)
		}
		return nil
	}
	
	err := process()
	
	fmt.Println("Code:")
	fmt.Println("  if err := divide(10, 0); err != nil {")
	fmt.Println("      return fmt.Errorf(\"failed: %w\", err)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	}
	fmt.Println()
}

var ErrNotFound = errors.New("not found")

func exampleCustomErrors() {
	fmt.Println("Example: Custom Errors")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	type MyError struct {
		Code    int
		Message string
	}
	
	myError := func(code int, msg string) *MyError {
		return &MyError{Code: code, Message: msg}
	}
	
	var err error = myError(404, "Resource not found")
	
	fmt.Println("Code:")
	fmt.Println("  type MyError struct {")
	fmt.Println("      Code    int")
	fmt.Println("      Message string")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Error: %v\n", err)
	
	if myErr, ok := err.(*MyError); ok {
		fmt.Printf("  Code: %d, Message: %s\n", myErr.Code, myErr.Message)
	}
	fmt.Println()
}

func exampleErrorIs() {
	fmt.Println("Example: errors.Is")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	err := ErrNotFound
	
	fmt.Println("Code:")
	fmt.Println("  var ErrNotFound = errors.New(\"not found\")")
	fmt.Println("  errors.Is(err, ErrNotFound)")
	fmt.Println()
	fmt.Println("Output:")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("  Error is ErrNotFound")
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleBasicError()
	exampleErrorWrapping()
	exampleCustomErrors()
	exampleErrorIs()
}
