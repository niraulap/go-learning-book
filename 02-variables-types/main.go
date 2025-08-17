package main

import "fmt"

func main() {
	fmt.Println("🐹 Go Variables and Types - Chapter 2 🐹")
	fmt.Println("=" * 50)

	// ============================================================================
	// SECTION 1: Basic Variable Declaration
	// ============================================================================
	section1_BasicVariables()

	// ============================================================================
	// SECTION 2: Different Data Types
	// ============================================================================
	section2_DataTypes()

	// ============================================================================
	// SECTION 3: Type Inference and Conversion
	// ============================================================================
	section3_TypeInference()

	fmt.Println("\n🎉 Chapter 2 Complete! You understand Go variables and types!")
}

// ============================================================================
// SECTION 1: Basic Variable Declaration
// ============================================================================
func section1_BasicVariables() {
	fmt.Println("\n📚 SECTION 1: Basic Variable Declaration")
	fmt.Println("-" * 30)

	// Method 1: var keyword with explicit type
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.6
	var isStudent bool = true

	fmt.Printf("Name: %s (type: %T)\n", name, name)
	fmt.Printf("Age: %d (type: %T)\n", age, age)
	fmt.Printf("Height: %.1f (type: %T)\n", height, height)
	fmt.Printf("Is Student: %t (type: %T)\n", isStudent, isStudent)

	// Method 2: var keyword with type inference
	var city = "New York"
	var population = 8336817
	var temperature = 72.5

	fmt.Printf("\nCity: %s (type: %T)\n", city, city)
	fmt.Printf("Population: %d (type: %T)\n", population, population)
	fmt.Printf("Temperature: %.1f (type: %T)\n", temperature, temperature)

	// Method 3: Short variable declaration (:=)
	country := "Canada"
	area := 9984670
	currency := "CAD"

	fmt.Printf("\nCountry: %s (type: %T)\n", country, country)
	fmt.Printf("Area: %d (type: %T)\n", area, area)
	fmt.Printf("Currency: %s (type: %T)\n", currency, currency)
}

// ============================================================================
// SECTION 2: Different Data Types
// ============================================================================
func section2_DataTypes() {
	fmt.Println("\n📚 SECTION 2: Different Data Types")
	fmt.Println("-" * 30)

	// Integer types
	var smallInt int8 = 127
	var bigInt int64 = 9223372036854775807
	var unsignedInt uint = 255

	fmt.Printf("Small int: %d (type: %T)\n", smallInt, smallInt)
	fmt.Printf("Big int: %d (type: %T)\n", bigInt, bigInt)
	fmt.Printf("Unsigned int: %d (type: %T)\n", unsignedInt, unsignedInt)

	// Float types
	var smallFloat float32 = 3.14159
	var bigFloat float64 = 3.141592653589793

	fmt.Printf("\nSmall float: %f (type: %T)\n", smallFloat, smallFloat)
	fmt.Printf("Big float: %.15f (type: %T)\n", bigFloat, bigFloat)

	// String and boolean
	var message string = "Hello, Go!"
	var isActive bool = false

	fmt.Printf("\nMessage: %s (type: %T)\n", message, message)
	fmt.Printf("Is Active: %t (type: %T)\n", isActive, isActive)

	// Zero values (default values)
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("\nZero values:\n")
	fmt.Printf("Default int: %d\n", defaultInt)
	fmt.Printf("Default float: %f\n", defaultFloat)
	fmt.Printf("Default string: %q\n", defaultString)
	fmt.Printf("Default bool: %t\n", defaultBool)
}

// ============================================================================
// SECTION 3: Type Inference and Conversion
// ============================================================================
func section3_TypeInference() {
	fmt.Println("\n📚 SECTION 3: Type Inference and Conversion")
	fmt.Println("-" * 30)

	// Type inference examples
	text := "Hello"           // Go infers string
	number := 42              // Go infers int
	decimal := 3.14           // Go infers float64
	flag := true              // Go infers bool

	fmt.Printf("Text: %s (type: %T)\n", text, text)
	fmt.Printf("Number: %d (type: %T)\n", number, number)
	fmt.Printf("Decimal: %.2f (type: %T)\n", decimal, decimal)
	fmt.Printf("Flag: %t (type: %T)\n", flag, flag)

	// Type conversion
	var intValue int = 42
	var floatValue float64 = float64(intValue)
	var stringValue string = fmt.Sprintf("%d", intValue)

	fmt.Printf("\nType conversion:\n")
	fmt.Printf("Original: %d (type: %T)\n", intValue, intValue)
	fmt.Printf("To float: %.1f (type: %T)\n", floatValue, floatValue)
	fmt.Printf("To string: %s (type: %T)\n", stringValue, stringValue)

	// Constants
	const PI = 3.14159
	const MAX_SIZE = 1000
	const APP_NAME = "GoPractice"

	fmt.Printf("\nConstants:\n")
	fmt.Printf("PI: %.5f\n", PI)
	fmt.Printf("MAX_SIZE: %d\n", MAX_SIZE)
	fmt.Printf("APP_NAME: %s\n", APP_NAME)
} 