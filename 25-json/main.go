package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
	City string `json:"city,omitempty"`
}

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: JSON                                      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📤 MARSHAL")
	fmt.Println("   Convert Go value to JSON")
	fmt.Println("   data, err := json.Marshal(person)")
	fmt.Println()
	fmt.Println("📥 UNMARSHAL")
	fmt.Println("   Convert JSON to Go value")
	fmt.Println("   err := json.Unmarshal(data, &person)")
	fmt.Println()
	fmt.Println("🏷️  STRUCT TAGS")
	fmt.Println("   Control JSON field names")
	fmt.Println("   Name string `json:\"name\"`")
	fmt.Println("   Age  int    `json:\"age,omitempty\"`")
	fmt.Println()
}

func exampleMarshal() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Marshal")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	person := Person{Name: "Alice", Age: 30}
	data, err := json.Marshal(person)
	
	fmt.Println("Code:")
	fmt.Println("  person := Person{Name: \"Alice\", Age: 30}")
	fmt.Println("  data, err := json.Marshal(person)")
	fmt.Println()
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Println("Output:")
	fmt.Printf("  JSON: %s\n", string(data))
	fmt.Println()
}

func exampleUnmarshal() {
	fmt.Println("Example: Unmarshal")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	jsonData := `{"name":"Bob","age":25}`
	var person Person
	
	err := json.Unmarshal([]byte(jsonData), &person)
	
	fmt.Println("Code:")
	fmt.Println("  jsonData := `{\"name\":\"Bob\",\"age\":25}`")
	fmt.Println("  err := json.Unmarshal([]byte(jsonData), &person)")
	fmt.Println()
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Println("Output:")
	fmt.Printf("  Person: %+v\n", person)
	fmt.Println()
}

func exampleStructTags() {
	fmt.Println("Example: Struct Tags")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	person := Person{Name: "Charlie", Age: 0} // Age is zero value
	data, _ := json.Marshal(person)
	
	fmt.Println("Code:")
	fmt.Println("  type Person struct {")
	fmt.Println("      Name string `json:\"name\"`")
	fmt.Println("      Age  int    `json:\"age,omitempty\"`")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  JSON (with omitempty): %s\n", string(data))
	fmt.Println("  Note: Age is omitted because it's zero value")
	fmt.Println()
}

func main() {
	printTheory()
	exampleMarshal()
	exampleUnmarshal()
	exampleStructTags()
}
