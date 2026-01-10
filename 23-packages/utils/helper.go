package utils

// Add is an exported function (capitalized)
func Add(a, b int) int {
	return a + b
}

// subtract is unexported (lowercase)
func subtract(a, b int) int {
	return a - b
}

func init() {
	// This runs when the package is imported
}
