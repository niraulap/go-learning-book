package main

import "fmt"

func main() {
	fmt.Println("🐹 Go Functions - Chapter 4 🐹")
	fmt.Println("=" * 50)

	// ============================================================================
	// SECTION 1: Basic Functions
	// ============================================================================
	section1_BasicFunctions()

	// ============================================================================
	// SECTION 2: Functions with Parameters
	// ============================================================================
	section2_FunctionsWithParameters()

	// ============================================================================
	// SECTION 3: Functions with Return Values
	// ============================================================================
	section3_FunctionsWithReturns()

	// ============================================================================
	// SECTION 4: Multiple Return Values (Go's Special Feature!)
	// ============================================================================
	section4_MultipleReturns()

	fmt.Println("\n🎉 Chapter 4 Complete! You understand Go functions!")
}

// ============================================================================
// SECTION 1: Basic Functions
// ============================================================================
func section1_BasicFunctions() {
	fmt.Println("\n📚 SECTION 1: Basic Functions")
	fmt.Println("-" * 30)

	// Call our first function
	sayHello()
	
	// Call it multiple times
	sayHello()
	sayHello()
	
	// Call another function
	printSeparator()
}

// sayHello is a simple function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello from a function!")
}

// printSeparator prints a line to separate output
func printSeparator() {
	fmt.Println("---")
}

// ============================================================================
// SECTION 2: Functions with Parameters
// ============================================================================
func section2_FunctionsWithParameters() {
	fmt.Println("\n📚 SECTION 2: Functions with Parameters")
	fmt.Println("-" * 30)

	// Call functions with different parameters
	greetPerson("Alice")
	greetPerson("Bob")
	greetPerson("Charlie")
	
	// Function with multiple parameters
	printInfo("Alice", 25, "New York")
	printInfo("Bob", 30, "London")
	
	// Function with numbers
	addNumbers(5, 3)
	addNumbers(10, 20)
	addNumbers(100, 200)
}

// greetPerson takes a name parameter and prints a greeting
func greetPerson(name string) {
	fmt.Printf("Hello, %s! Welcome to Go programming!\n", name)
}

// printInfo takes multiple parameters of different types
func printInfo(name string, age int, city string) {
	fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
}

// addNumbers takes two integers and prints their sum
func addNumbers(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

// ============================================================================
// SECTION 3: Functions with Return Values
// ============================================================================
func section3_FunctionsWithReturns() {
	fmt.Println("\n📚 SECTION 3: Functions with Return Values")
	fmt.Println("-" * 30)

	// Call functions and use their return values
	result1 := calculateSum(5, 3)
	fmt.Printf("Sum: %d\n", result1)
	
	result2 := calculateSum(10, 20)
	fmt.Printf("Sum: %d\n", result2)
	
	// Function that returns a string
	message := createGreeting("Alice")
	fmt.Println(message)
	
	// Function that returns a boolean
	isAdult := checkAge(18)
	fmt.Printf("Is 18 an adult? %t\n", isAdult)
	
	isAdult = checkAge(15)
	fmt.Printf("Is 15 an adult? %t\n", isAdult)
}

// calculateSum takes two integers and returns their sum
func calculateSum(a int, b int) int {
	return a + b
}

// createGreeting takes a name and returns a greeting message
func createGreeting(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go!", name)
}

// checkAge takes an age and returns true if adult (18+)
func checkAge(age int) bool {
	return age >= 18
}

// ============================================================================
// SECTION 4: Multiple Return Values (Go's Special Feature!)
// ============================================================================
func section4_MultipleReturns() {
	fmt.Println("\n📚 SECTION 4: Multiple Return Values (Go's Special Feature!)")
	fmt.Println("-" * 30)

	// Call function with multiple return values
	quotient, remainder := divideNumbers(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)
	
	quotient, remainder = divideNumbers(20, 4)
	fmt.Printf("20 ÷ 4 = %d remainder %d\n", quotient, remainder)
	
	// Function that returns multiple values of different types
	name, age, isStudent := getPersonInfo("Alice", 20, true)
	fmt.Printf("Person: %s, Age: %d, Student: %t\n", name, age, isStudent)
	
	// Function that returns success/failure with data
	success, result := performOperation("add", 10, 5)
	if success {
		fmt.Printf("Operation successful: %d\n", result)
	} else {
		fmt.Printf("Operation failed: %s\n", result)
	}
	
	success, result = performOperation("multiply", 6, 7)
	if success {
		fmt.Printf("Operation successful: %d\n", result)
	} else {
		fmt.Printf("Operation failed: %s\n", result)
	}
}

// divideNumbers takes two integers and returns quotient and remainder
func divideNumbers(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// getPersonInfo returns multiple values of different types
func getPersonInfo(name string, age int, isStudent bool) (string, int, bool) {
	return name, age, isStudent
}

// performOperation performs math operations and returns success status and result
func performOperation(operation string, a int, b int) (bool, int) {
	switch operation {
	case "add":
		return true, a + b
	case "subtract":
		return true, a - b
	case "multiply":
		return true, a * b
	case "divide":
		if b != 0 {
			return true, a / b
		}
		return false, 0
	default:
		return false, 0
	}
} 