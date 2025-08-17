package main

import "fmt"

func main() {
	fmt.Println("🐹 Go Control Structures Activity Book 🐹")
	fmt.Println("=" * 50)
	
	// ============================================================================
	// SECTION 1: IF STATEMENTS - Making Decisions in Go
	// ============================================================================
	section1_IfStatements()

	// ============================================================================
	// SECTION 2: FOR LOOPS - Repeating Actions in Go
	// ============================================================================
	section2_ForLoops()

	// ============================================================================
	// SECTION 3: RANGE LOOPS - Go's Special Iteration Power
	// ============================================================================
	section3_RangeLoops()

	// ============================================================================
	// SECTION 4: SWITCH STATEMENTS - Multiple Choice Decisions
	// ============================================================================
	section4_SwitchStatements()

	// ============================================================================
	// SECTION 5: DEFER STATEMENTS - Go's Unique Cleanup Feature
	// ============================================================================
	section5_DeferStatements()

	// ============================================================================
	// BONUS SECTION: Interactive Practice Examples
	// ============================================================================
	bonusSection_PracticeExamples()

	fmt.Println("\n🎉 Control Structures Activity Book Complete!")
}

// ============================================================================
// SECTION 1: IF STATEMENTS
// ============================================================================
func section1_IfStatements() {
	fmt.Println("\n SECTION 1: IF STATEMENTS")
	fmt.Println("-" * 30)

	// Example 1: Basic IF Statement
	fmt.Println("Basic IF Statement:")
	age := 18
	if age >= 18 {
		fmt.Println("✅ You are an adult!")
	}

	// Example 2: IF-ELSE Statement
	fmt.Println("\nIF-ELSE Statement:")
	age = 16
	if age >= 18 {
		fmt.Println("✅ You can vote!")
	} else {
		fmt.Println("❌ You cannot vote yet!")
	}

	// Example 3: IF-ELSEIF-ELSE  Chain
	fmt.Println("\nIF-ELSE IF-ELSE Chain:")
	score := 85
	if score >= 90 {
		fmt.Println("🎯 Grade: A (Excellent!)")
	} else if score >= 80 {
		fmt.Println("👍 Grade: B (Good job!)")
	} else if score >= 70 {
		fmt.Println("📚 Grade: C (Keep studying!)")
	} else {
		fmt.Println("⚠️ Grade: D (Need improvement!)")
	}

	// Example 4: IF with Initialization (Go's Cool Feature!)
	fmt.Println("\nIF with Initialization:")
	if testScore := getTestScore(); testScore >= 90 {
		fmt.Printf("🏆 Excellent test score: %d\n", testScore)
	} else if testScore >= 70 {
		fmt.Printf("📖 Good test score: %d\n", testScore)
	} else {
		fmt.Printf("📚 Study more! Test score: %d\n", testScore)
	}
	// Note: 'testScore' variable only exists inside the if block
}

// ============================================================================
// SECTION 2: FOR LOOPS
// ============================================================================
func section2_ForLoops() {

	// Example 1: Traditional FOR Loop
	fmt.Println("Traditional FOR Loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// Example 2: While-Style Loop : Go does not have while loop, so we use for loop like this 
	fmt.Println("\nWhile-Style Loop:")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// Example 3: Infinite Loop with Break
	fmt.Println("\nInfinite Loop with Break:")
	attempts := 0
	for {
		attempts++
		if attempts > 3 {
			fmt.Println("🛑 Max attempts reached! Stopping loop.")
			break
		}
		fmt.Printf("🔄 Attempt %d\n", attempts)
	}

	// Example 4: Loop Control Keywords
	fmt.Println("\nUsing break to exit early:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("🛑 Found 5! Breaking loop at iteration %d\n", i)
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\nUsing continue to skip iterations:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Printf("⏭️ Skipping iteration %d\n", i)
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// ============================================================================
// SECTION 3: RANGE LOOPS
// ============================================================================
func section3_RangeLoops() {
	fmt.Println("\n📚 SECTION 3: RANGE LOOPS")
	fmt.Println("-" * 30)

	numbers := []int{10, 20, 30, 40, 50}

	// Example 1: Range Over Slice (Index and Value)
	fmt.Println("Range with index and value:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example 2: Range Over Slice (Value Only)
	fmt.Println("\nRange with value only:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// Example 3: Range Over Slice (Index Only)
	fmt.Println("\nRange with index only:")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}

	// Example 4: Range Over Map
	fmt.Println("\nRange over map:")
	person := map[string]string{
		"name": "Alice",
		"city": "New York",
		"job":  "Developer",
		"age":  "25",
	}
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Example 5: Range Over String
	fmt.Println("\nRange over string:")
	word := "Go"
	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c (Unicode: %d)\n", index, char, char)
	}
}

// ============================================================================
// SECTION 4: SWITCH STATEMENTS
// ============================================================================
func section4_SwitchStatements() {
	fmt.Println("\n📚 SECTION 4: SWITCH STATEMENTS")
	fmt.Println("-" * 30)

	// Example 1: Basic Switch Statement
	fmt.Println("Basic Switch Statement:")
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("📅 Start of work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("💼 Middle of work week")
	case "Friday":
		fmt.Println("🎉 TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("🏖️ Weekend!")
	default:
		fmt.Println("❓ Unknown day")
	}

	// Example 2: Switch with Expression
	fmt.Println("\nSwitch with Expression:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("🏆 Grade: A")
	case score >= 80:
		fmt.Println("👍 Grade: B")
	case score >= 70:
		fmt.Println("📚 Grade: C")
	case score >= 60:
		fmt.Println("⚠️ Grade: D")
	default:
		fmt.Println("❌ Grade: F")
	}

	// Example 3: Switch with Fallthrough
	fmt.Println("\nSwitch with Fallthrough:")
	grade := 85
	switch {
	case grade >= 90:
		fmt.Print("A")
		fallthrough
	case grade >= 80:
		fmt.Print("B")
		fallthrough
	case grade >= 70:
		fmt.Print("C")
		fallthrough
	default:
		fmt.Print("D")
	}
	fmt.Println(" grade")
}

// ============================================================================
// SECTION 5: DEFER STATEMENTS
// ============================================================================
func section5_DeferStatements() {
	fmt.Println("\n📚 SECTION 5: DEFER STATEMENTS")
	fmt.Println("-" * 30)

	// Example 1: Basic Defer Statement
	fmt.Println("Basic Defer Statement:")
	defer fmt.Println("🚪 This runs LAST (when function exits)")
	defer fmt.Println("🚪 This runs SECOND to last")
	defer fmt.Println("🚪 This runs THIRD to last")
	
	fmt.Println("Function body executing...")
	fmt.Println("About to exit function...")

	// Example 2: Practical Defer Examples
	fmt.Println("\nPractical Defer Examples:")
	
	// Simulate file operations
	fmt.Println("📁 Opening file 'data.txt'...")
	defer fmt.Println("🗂️ Closing file 'data.txt' (deferred cleanup)")
	
	fmt.Println("📖 Reading file content...")
	fmt.Println("📖 File operations complete!")

	// Simulate database operations
	fmt.Println("\n🗄️ Connecting to database...")
	defer fmt.Println("🔌 Closing database connection (deferred)")
	
	fmt.Println("💾 Starting transaction...")
	defer fmt.Println("✅ Committing transaction (deferred)")
	
	fmt.Println("📊 Executing queries...")
	fmt.Println("📊 Database operations complete!")
}

// ============================================================================
// BONUS SECTION: Interactive Practice Examples
// ============================================================================
func bonusSection_PracticeExamples() {
	fmt.Println("\n🎁 BONUS SECTION: Combined Examples")
	fmt.Println("-" * 30)

	// Example 1: Number Guessing Game Logic
	fmt.Println("Number Guessing Game Logic:")
	targetNumber := 7
	guesses := []int{3, 8, 7, 5}
	
	for i, guess := range guesses {
		fmt.Printf("Guess %d: %d\n", i+1, guess)
		
		if guess == targetNumber {
			fmt.Println("🎉 Correct! You found the number!")
			break
		} else if guess < targetNumber {
			fmt.Println("📈 Too low! Try a higher number.")
		} else {
			fmt.Println("📉 Too high! Try a lower number.")
		}
	}

	// Example 2: Grade Analysis with Range
	fmt.Println("\nGrade Analysis with Range:")
	studentScores := []int{95, 87, 92, 78, 100, 85, 90}
	
	var excellentCount, goodCount, needsImprovement int
	
	for _, score := range studentScores {
		switch {
		case score >= 90:
			excellentCount++
		case score >= 80:
			goodCount++
		default:
			needsImprovement++
		}
	}
	
	fmt.Printf("📊 Analysis: Excellent: %d, Good: %d, Needs Improvement: %d\n", 
		excellentCount, goodCount, needsImprovement)

	// Example 3: Resource Management Simulation
	fmt.Println("\nResource Management Simulation:")
	fmt.Println("🚀 Starting complex operation...")
	
	defer fmt.Println("🗂️ Closing file 'config.txt'")
	defer fmt.Println("🔌 Disconnecting from web service")
	defer fmt.Println("🔌 Closing database connection")
	
	fmt.Println("⚙️ Performing operations...")
	fmt.Println("✅ Operations complete! Cleanup will happen automatically.")
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// getTestScore returns a test score for demonstration purposes
func getTestScore() int {
	return 95
} 