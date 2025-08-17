package main

import "fmt"

func main() {
	fmt.Println("🐹 Go Maps - Chapter 6 🐹")
	fmt.Println("==================================================")

	// ============================================================================
	// SECTION 1: Creating and Using Maps
	// ============================================================================
	section1_CreatingMaps()

	// ============================================================================
	// SECTION 2: Map Operations
	// ============================================================================
	section2_MapOperations()

	// ============================================================================
	// SECTION 3: Working with Maps
	// ============================================================================
	section3_WorkingWithMaps()

	// ============================================================================
	// SECTION 4: Functions with Maps
	// ============================================================================
	section4_FunctionsWithMaps()

	fmt.Println("\n🎉 Chapter 6 Complete! You understand Go maps!")
}

// ============================================================================
// SECTION 1: Creating and Using Maps
// ============================================================================
func section1_CreatingMaps() {
	fmt.Println("\n📚 SECTION 1: Creating and Using Maps")
	fmt.Println("----------------------------------------")

	// Method 1: Map literal (create and initialize at once)
	fmt.Println("Method 1: Map literal")
	studentGrades := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
		"Diana":   78,
		"Eve":     100,
	}
	fmt.Printf("Student grades: %v\n", studentGrades)

	// Method 2: Using make (empty map)
	fmt.Println("\nMethod 2: Using make")
	shoppingCart := make(map[string]int)
	fmt.Printf("Empty shopping cart: %v\n", shoppingCart)

	// Method 3: Declare then initialize
	fmt.Println("\nMethod 3: Declare then initialize")
	var config map[string]string
	config = map[string]string{
		"database": "postgres",
		"port":     "5432",
		"host":     "localhost",
	}
	fmt.Printf("Configuration: %v\n", config)

	// Accessing map values
	fmt.Println("\nAccessing map values:")
	fmt.Printf("Alice's grade: %d\n", studentGrades["Alice"])
	fmt.Printf("Bob's grade: %d\n", studentGrades["Bob"])

	// Setting map values
	fmt.Println("\nSetting map values:")
	shoppingCart["apple"] = 5
	shoppingCart["banana"] = 3
	shoppingCart["orange"] = 2
	fmt.Printf("Shopping cart after adding items: %v\n", shoppingCart)

	// Zero values for missing keys
	fmt.Println("\nZero values for missing keys:")
	fmt.Printf("Frank's grade (doesn't exist): %d\n", studentGrades["Frank"])
	fmt.Printf("Is Frank in the map? %t\n", studentGrades["Frank"] != 0)

	// Different value types
	fmt.Println("\nDifferent value types:")
	userInfo := map[string]interface{}{
		"name":     "Alice",
		"age":      25,
		"isActive": true,
		"height":   5.6,
	}
	fmt.Printf("User info: %v\n", userInfo)
}

// ============================================================================
// SECTION 2: Map Operations
// ============================================================================
func section2_MapOperations() {
	fmt.Println("\n📚 SECTION 2: Map Operations")
	fmt.Println("------------------------------")

	// Creating a map to work with
	inventory := map[string]int{
		"laptop":   10,
		"mouse":    25,
		"keyboard": 15,
		"monitor":  8,
	}
	fmt.Printf("Initial inventory: %v\n", inventory)

	// Checking if a key exists
	fmt.Println("\nChecking if keys exist:")
	
	// Method 1: Check if value is zero value
	if inventory["laptop"] != 0 {
		fmt.Printf("Laptop exists: %d units\n", inventory["laptop"])
	} else {
		fmt.Println("Laptop not found")
	}

	// Method 2: Using comma ok idiom (Go's special feature!)
	fmt.Println("\nUsing comma ok idiom:")
	if quantity, exists := inventory["laptop"]; exists {
		fmt.Printf("Laptop exists: %d units\n", quantity)
	} else {
		fmt.Println("Laptop not found")
	}

	if quantity, exists := inventory["headphones"]; exists {
		fmt.Printf("Headphones exist: %d units\n", quantity)
	} else {
		fmt.Println("Headphones not found")
	}

	// Deleting keys
	fmt.Println("\nDeleting keys:")
	delete(inventory, "mouse")
	fmt.Printf("After deleting mouse: %v\n", inventory)

	// Trying to delete non-existent key (safe)
	delete(inventory, "nonexistent")
	fmt.Printf("After trying to delete nonexistent: %v\n", inventory)

	// Getting map length
	fmt.Printf("\nMap length: %d\n", len(inventory))

	// Updating existing values
	fmt.Println("\nUpdating existing values:")
	inventory["laptop"] = 12
	fmt.Printf("After updating laptop quantity: %v\n", inventory)

	// Adding new keys
	fmt.Println("\nAdding new keys:")
	inventory["speakers"] = 5
	fmt.Printf("After adding speakers: %v\n", inventory)
}

// ============================================================================
// SECTION 3: Working with Maps
// ============================================================================
func section3_WorkingWithMaps() {
	fmt.Println("\n📚 SECTION 3: Working with Maps")
	fmt.Println("--------------------------------")

	// Creating a map to work with
	bookRatings := map[string]int{
		"Go Programming":    5,
		"Python Basics":     4,
		"JavaScript Guide":  3,
		"Rust Tutorial":     5,
		"Java Reference":    2,
	}

	// Range loops over maps (building on Chapter 3 and 5)
	fmt.Println("Range loops over maps:")
	
	// Get both key and value
	fmt.Println("\nAll books and ratings:")
	for book, rating := range bookRatings {
		fmt.Printf("  %s: %d stars\n", book, rating)
	}

	// Get only keys
	fmt.Println("\nJust the book titles:")
	for book := range bookRatings {
		fmt.Printf("  %s\n", book)
	}

	// Get only values (less common but possible)
	fmt.Println("\nJust the ratings:")
	for _, rating := range bookRatings {
		fmt.Printf("  %d stars\n", rating)
	}

	// Filtering maps
	fmt.Println("\nFiltering maps (high-rated books):")
	for book, rating := range bookRatings {
		if rating >= 4 {
			fmt.Printf("  %s: %d stars (recommended)\n", book, rating)
		}
	}

	// Modifying map values
	fmt.Println("\nModifying map values:")
	for book, rating := range bookRatings {
		if rating < 3 {
			bookRatings[book] = rating + 1 // Give low-rated books a boost
		}
	}
	fmt.Printf("After boosting low ratings: %v\n", bookRatings)

	// Copying maps (important concept!)
	fmt.Println("\nCopying maps:")
	originalMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Original map: %v\n", originalMap)

	// Method 1: Direct assignment (creates reference, not copy!)
	referenceMap := originalMap
	referenceMap["a"] = 100
	fmt.Printf("After modifying reference: %v\n", referenceMap)
	fmt.Printf("Original also changed: %v\n", originalMap)

	// Method 2: Proper copying
	originalMap = map[string]int{"a": 1, "b": 2, "c": 3} // Reset
	copiedMap := make(map[string]int)
	for key, value := range originalMap {
		copiedMap[key] = value
	}
	
	copiedMap["a"] = 200
	fmt.Printf("After modifying copy: %v\n", copiedMap)
	fmt.Printf("Original unchanged: %v\n", originalMap)
}

// ============================================================================
// SECTION 4: Functions with Maps
// ============================================================================
func section4_FunctionsWithMaps() {
	fmt.Println("\n📚 SECTION 4: Functions with Maps")
	fmt.Println("-----------------------------------")

	// Sample data
	students := map[string]map[string]interface{}{
		"Alice": {
			"age":    20,
			"grade":  "A",
			"city":   "New York",
			"active": true,
		},
		"Bob": {
			"age":    22,
			"grade":  "B",
			"city":   "London",
			"active": false,
		},
		"Charlie": {
			"age":    19,
			"grade":  "A",
			"city":   "Paris",
			"active": true,
		},
	}

	// Function that takes a map as parameter
	fmt.Println("Function that takes a map as parameter:")
	printStudentInfo(students, "Alice")
	printStudentInfo(students, "Bob")
	printStudentInfo(students, "Frank") // Doesn't exist

	// Function that returns a new map
	fmt.Println("\nFunction that returns a new map:")
	activeStudents := getActiveStudents(students)
	fmt.Printf("Active students: %v\n", activeStudents)

	// Function that modifies a map
	fmt.Println("\nFunction that modifies a map:")
	fmt.Printf("Before modification: %v\n", students)
	addStudent(students, "Diana", 21, "B", "Berlin", true)
	fmt.Printf("After adding Diana: %v\n", students)

	// Function that finds students by criteria
	fmt.Println("\nFunction that finds students by criteria:")
	gradeAStudents := findStudentsByGrade(students, "A")
	fmt.Printf("Students with grade A: %v\n", gradeAStudents)

	// Function that counts students by city
	fmt.Println("\nFunction that counts students by city:")
	cityCounts := countStudentsByCity(students)
	fmt.Printf("Students by city: %v\n", cityCounts)

	// Function that returns success/failure with data (building on Chapter 4)
	fmt.Println("\nFunction with success/failure pattern:")
	success, studentData := getStudentDetails(students, "Alice")
	if success {
		fmt.Printf("Student found: %v\n", studentData)
	} else {
		fmt.Println("Student not found")
	}

	success, studentData = getStudentDetails(students, "Frank")
	if success {
		fmt.Printf("Student found: %v\n", studentData)
	} else {
		fmt.Println("Student not found")
	}
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// printStudentInfo prints information about a specific student
func printStudentInfo(students map[string]map[string]interface{}, name string) {
	if student, exists := students[name]; exists {
		fmt.Printf("  %s: Age %v, Grade %v, City %v, Active %v\n",
			name, student["age"], student["grade"], student["city"], student["active"])
	} else {
		fmt.Printf("  Student %s not found\n", name)
	}
}

// getActiveStudents returns a map of only active students
func getActiveStudents(students map[string]map[string]interface{}) map[string]map[string]interface{} {
	active := make(map[string]map[string]interface{})
	for name, info := range students {
		if activeStatus, ok := info["active"].(bool); ok && activeStatus {
			active[name] = info
		}
	}
	return active
}

// addStudent adds a new student to the map
func addStudent(students map[string]map[string]interface{}, name string, age int, grade string, city string, active bool) {
	students[name] = map[string]interface{}{
		"age":    age,
		"grade":  grade,
		"city":   city,
		"active": active,
	}
}

// findStudentsByGrade returns students with a specific grade
func findStudentsByGrade(students map[string]map[string]interface{}, grade string) []string {
	var result []string
	for name, info := range students {
		if studentGrade, ok := info["grade"].(string); ok && studentGrade == grade {
			result = append(result, name)
		}
	}
	return result
}

// countStudentsByCity returns a map counting students by city
func countStudentsByCity(students map[string]map[string]interface{}) map[string]int {
	cityCounts := make(map[string]int)
	for _, info := range students {
		if city, ok := info["city"].(string); ok {
			cityCounts[city]++
		}
	}
	return cityCounts
}

// getStudentDetails returns success status and student data (success/failure pattern)
func getStudentDetails(students map[string]map[string]interface{}, name string) (bool, map[string]interface{}) {
	if student, exists := students[name]; exists {
		return true, student
	}
	return false, nil
} 