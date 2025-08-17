package main

import "fmt"

func main() {
	fmt.Println("🐹 Go Arrays and Slices - Chapter 5 🐹")
	fmt.Println("==================================================")

	// ============================================================================
	// SECTION 1: Arrays (Fixed-Size Collections)
	// ============================================================================
	section1_Arrays()

	// ============================================================================
	// SECTION 2: Slices (Dynamic Collections - Go's Superpower!)
	// ============================================================================
	section2_Slices()

	// ============================================================================
	// SECTION 3: Working with Collections
	// ============================================================================
	section3_WorkingWithCollections()

	// ============================================================================
	// SECTION 4: Functions with Collections
	// ============================================================================
	// Create a scores slice that will be used across sections
	scores := []int{95, 87, 92, 78, 100, 85, 90}
	section4_FunctionsWithCollections(scores)

	fmt.Println("\n🎉 Chapter 5 Complete! You understand Go arrays and slices!")
}

// ============================================================================
// SECTION 1: Arrays (Fixed-Size Collections)
// ============================================================================
func section1_Arrays() {
	fmt.Println("\n📚 SECTION 1: Arrays (Fixed-Size Collections)")
	fmt.Println("------------------------------")

	// Declaring arrays with explicit size
	var numbers [5]int
	fmt.Printf("Empty array: %v\n", numbers)
	
	// Setting values in array
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Printf("Array with values: %v\n", numbers)
	
	// Array literal (declare and initialize at once)
	colors := [3]string{"red", "green", "blue"}
	fmt.Printf("Colors array: %v\n", colors)
	
	// Array with size inference
	grades := [...]int{95, 87, 92, 78, 100}
	fmt.Printf("Grades array: %v\n", grades)
	fmt.Printf("Array length: %d\n", len(grades))
	
	// Accessing array elements
	fmt.Printf("First grade: %d\n", grades[0])
	fmt.Printf("Last grade: %d\n", grades[len(grades)-1])
	
	// Array limitations
	fmt.Println("\nArray limitations:")
	fmt.Println("- Fixed size (cannot grow or shrink)")
	fmt.Println("- Must know size at declaration")
	fmt.Println("- Less flexible than slices")
}

// ============================================================================
// SECTION 2: Slices (Dynamic Collections - Go's Superpower!)
// ============================================================================
func section2_Slices() {
	fmt.Println("\n📚 SECTION 2: Slices (Dynamic Collections - Go's Superpower!)")
	fmt.Println("------------------------------")

	// Creating slices
	var emptySlice []int
	fmt.Printf("Empty slice: %v (length: %d, capacity: %d)\n", emptySlice, len(emptySlice), cap(emptySlice))
	
	// Slice literal
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruits slice: %v (length: %d)\n", fruits, len(fruits))
	
	// Creating slice with make
	numbers := make([]int, 3, 5) // length 3, capacity 5
	fmt.Printf("Slice with make: %v (length: %d, capacity: %d)\n", numbers, len(numbers), cap(numbers))
	
	// Adding elements to slice (append)
	fmt.Println("\nAdding elements to slice:")
	fruits = append(fruits, "grape")
	fmt.Printf("After adding grape: %v\n", fruits)
	
	fruits = append(fruits, "mango", "kiwi")
	fmt.Printf("After adding multiple: %v\n", fruits)
	
	// Slicing (getting parts of a slice)
	fmt.Println("\nSlicing operations:")
	fmt.Printf("Original: %v\n", fruits)
	fmt.Printf("First 3: %v\n", fruits[:3])
	fmt.Printf("Last 3: %v\n", fruits[len(fruits)-3:])
	fmt.Printf("Middle 3: %v\n", fruits[1:4])
	
	// Slice capacity and growth
	fmt.Println("\nSlice capacity demonstration:")
	slice := make([]int, 0, 3)
	fmt.Printf("Initial: length=%d, capacity=%d\n", len(slice), cap(slice))
	
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
		fmt.Printf("After adding %d: length=%d, capacity=%d\n", i, len(slice), cap(slice))
	}
}

// ============================================================================
// SECTION 3: Working with Collections
// ============================================================================
func section3_WorkingWithCollections() {
	fmt.Println("\n📚 SECTION 3: Working with Collections")
	fmt.Println("------------------------------")

	// Range loops with slices (building on Chapter 3)
	fmt.Println("Range loops with slices:")
	// Use the same scores data for consistency
	scores := []int{95, 87, 92, 78, 100, 85, 90}
	
	fmt.Println("All scores:")
	for i, score := range scores {
		fmt.Printf("  Index %d: %d\n", i, score)
	}
	
	fmt.Println("\nHigh scores only:")
	for _, score := range scores {
		if score >= 90 {
			fmt.Printf("  High score: %d\n", score)
		}
	}
	
	// Modifying slices
	fmt.Println("\nModifying slices:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", numbers)
	
	// Change an element
	numbers[2] = 30
	fmt.Printf("After changing index 2: %v\n", numbers)
	
	// Remove an element (using append and slicing)
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Printf("After removing index 2: %v\n", numbers)
	
	// Insert an element
	numbers = append(numbers[:2], append([]int{25}, numbers[2:]...)...)
	fmt.Printf("After inserting 25 at index 2: %v\n", numbers)
	
	// Copying slices
	fmt.Println("\nCopying slices:")
	original := []int{10, 20, 30, 40, 50}
	copied := make([]int, len(original))
	copy(copied, original)
	
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copied: %v\n", copied)
	
	// Modify original to show they're independent
	original[0] = 100
	fmt.Printf("After modifying original: %v\n", original)
	fmt.Printf("Copied unchanged: %v\n", copied)
}

// ============================================================================
// SECTION 4: Functions with Collections
// ============================================================================
func section4_FunctionsWithCollections(scores []int) {
	fmt.Println("\n📚 SECTION 4: Functions with Collections")
	fmt.Println("------------------------------")

	// Function that takes a slice as parameter
	data := []int{10, 20, 30, 40, 50}
	fmt.Printf("Original data: %v\n", data)
	
	// Call function to double all numbers
	doubleSlice(data)
	fmt.Printf("After doubling: %v\n", data)
	
	// Function that returns a new slice
	original := []int{1, 2, 3, 4, 5}
	doubled := doubleSliceReturn(original)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Doubled: %v\n", doubled)
	
	// Function that finds maximum value
	maxValue := findMax(scores)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Maximum score: %d\n", maxValue)
	
	// Function that filters slice
	highScores := filterHighScores(scores, 85)
	fmt.Printf("All scores: %v\n", scores)
	fmt.Printf("High scores (85+): %v\n", highScores)
	
	// Function that combines slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	combined := combineSlices(slice1, slice2)
	fmt.Printf("Slice 1: %v\n", slice1)
	fmt.Printf("Slice 2: %v\n", slice2)
	fmt.Printf("Combined: %v\n", combined)
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// doubleSlice doubles all numbers in a slice (modifies original)
func doubleSlice(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 2
	}
}

// doubleSliceReturn returns a new slice with doubled values
func doubleSliceReturn(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num * 2
	}
	return result
}

// findMax finds the maximum value in a slice
func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// filterHighScores returns scores above a threshold
func filterHighScores(scores []int, threshold int) []int {
	var result []int
	for _, score := range scores {
		if score >= threshold {
			result = append(result, score)
		}
	}
	return result
}

// combineSlices combines two slices into one
func combineSlices(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))
	result = append(result, slice1...)
	result = append(result, slice2...)
	return result
} 