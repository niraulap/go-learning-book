# Go Maps - Chapter 6 🐹

Welcome to Chapter 6! Now that you understand arrays and slices (indexed collections), let's learn about **maps** - Go's powerful key-value collections. Maps are like dictionaries where you look up information by name instead of by position.

## What You'll Learn

- How to create and use maps (key-value collections)
- Go's unique "comma ok" idiom for checking key existence
- Range loops with maps (building on Chapters 3 and 5)
- Functions that work with maps
- When to use maps vs. slices vs. arrays
- Real-world applications like databases and configuration

## What are Maps?

Maps are collections where data is organized by **keys** instead of **indices**. Think of them like:
- **Phone book**: Name → Phone number
- **Dictionary**: Word → Definition
- **Shopping cart**: Product → Quantity
- **User database**: Username → User information

## Maps vs. Arrays/Slices

| Feature | Arrays | Slices | Maps |
|---------|--------|--------|------|
| **Access by** | Index (0, 1, 2...) | Index (0, 1, 2...) | Key ("name", "age"...) |
| **Size** | Fixed | Dynamic | Dynamic |
| **Order** | Guaranteed | Guaranteed | **Not guaranteed** |
| **Use case** | Fixed-size lists | Dynamic lists | Lookup tables |

## Creating Maps

### Method 1: Map Literal (Most Common)

```go
// Create and initialize at once
studentGrades := map[string]int{
    "Alice":   95,
    "Bob":     87,
    "Charlie": 92,
    "Diana":   78,
    "Eve":     100,
}
```

### Method 2: Using make (Empty Map)

```go
// Create empty map
shoppingCart := make(map[string]int)

// Add items later
shoppingCart["apple"] = 5
shoppingCart["banana"] = 3
```

### Method 3: Declare then Initialize

```go
// Declare first
var config map[string]string

// Initialize later
config = map[string]string{
    "database": "postgres",
    "port":     "5432",
    "host":     "localhost",
}
```

## Accessing and Setting Values

```go
// Access values
aliceGrade := studentGrades["Alice"]  // 95
bobGrade := studentGrades["Bob"]      // 87

// Set values
studentGrades["Frank"] = 88
studentGrades["Alice"] = 96  // Update existing value
```

## The "Comma Ok" Idiom (Go's Special Feature!)

This is one of Go's most elegant features for checking if a key exists:

```go
// Basic access (returns zero value if key doesn't exist)
grade := studentGrades["Frank"]  // Returns 0 (int zero value)

// Comma ok idiom - check if key exists
if grade, exists := studentGrades["Frank"]; exists {
    fmt.Printf("Frank's grade: %d\n", grade)
} else {
    fmt.Println("Frank not found in records")
}
```

**Why this is brilliant:**
- **One line** does two things: gets value AND checks existence
- **No separate lookup** needed
- **Go-specific** - you won't see this in many other languages
- **Eliminates errors** from missing keys

## Map Operations

### Checking if Keys Exist

```go
// Method 1: Comma ok idiom (recommended)
if value, exists := myMap["key"]; exists {
    fmt.Printf("Key exists: %v\n", value)
} else {
    fmt.Println("Key not found")
}

// Method 2: Check against zero value (less reliable)
if myMap["key"] != 0 {  // Only works for non-zero values!
    fmt.Printf("Key exists: %v\n", myMap["key"])
}
```

### Deleting Keys

```go
// Delete a key
delete(studentGrades, "Bob")

// Safe to delete non-existent keys
delete(studentGrades, "Nonexistent")  // No error, no panic
```

### Getting Map Length

```go
studentCount := len(studentGrades)  // Number of key-value pairs
```

## Working with Maps

### Range Loops (Building on Chapters 3 and 5)

You already know range loops! Now use them with maps:

```go
bookRatings := map[string]int{
    "Go Programming":    5,
    "Python Basics":     4,
    "JavaScript Guide":  3,
}

// Get both key and value
for book, rating := range bookRatings {
    fmt.Printf("%s: %d stars\n", book, rating)
}

// Get only keys
for book := range bookRatings {
    fmt.Printf("Book: %s\n", book)
}

// Get only values (less common)
for _, rating := range bookRatings {
    fmt.Printf("Rating: %d\n", rating)
}
```

### Filtering Maps

```go
// Find high-rated books
for book, rating := range bookRatings {
    if rating >= 4 {
        fmt.Printf("%s: %d stars (recommended)\n", book, rating)
    }
}
```

### Modifying Map Values

```go
// Boost low ratings
for book, rating := range bookRatings {
    if rating < 3 {
        bookRatings[book] = rating + 1
    }
}
```

## Important: Map Copying

**⚠️ Critical Concept: Maps are reference types!**

```go
originalMap := map[string]int{"a": 1, "b": 2, "c": 3}

// ❌ WRONG: This creates a reference, not a copy!
referenceMap := originalMap
referenceMap["a"] = 100
// Now BOTH maps have "a": 100!

// ✅ CORRECT: Proper copying
copiedMap := make(map[string]int)
for key, value := range originalMap {
    copiedMap[key] = value
}
copiedMap["a"] = 200  // Only copiedMap changes
```

## Functions with Maps

### Functions that Take Maps as Parameters

```go
func printStudentInfo(students map[string]map[string]interface{}, name string) {
    if student, exists := students[name]; exists {
        fmt.Printf("%s: Age %v, Grade %v\n", 
            name, student["age"], student["grade"])
    } else {
        fmt.Printf("Student %s not found\n", name)
    }
}

// Use it
students := map[string]map[string]interface{}{
    "Alice": {"age": 20, "grade": "A"},
}
printStudentInfo(students, "Alice")
```

### Functions that Return New Maps

```go
func getActiveStudents(students map[string]map[string]interface{}) map[string]map[string]interface{} {
    active := make(map[string]map[string]interface{})
    for name, info := range students {
        if activeStatus, ok := info["active"].(bool); ok && activeStatus {
            active[name] = info
        }
    }
    return active
}

// Use it
activeStudents := getActiveStudents(students)
```

### Functions that Modify Maps

```go
func addStudent(students map[string]map[string]interface{}, name string, age int, grade string) {
    students[name] = map[string]interface{}{
        "age":   age,
        "grade": grade,
    }
}

// Use it
addStudent(students, "Diana", 21, "B")
```

## Real-World Map Applications

### 1. **Configuration Management**
```go
config := map[string]string{
    "database": "postgres",
    "port":     "5432",
    "host":     "localhost",
    "username": "admin",
}
```

### 2. **User Management System**
```go
users := map[string]map[string]interface{}{
    "alice": {
        "name":     "Alice Johnson",
        "email":    "alice@example.com",
        "age":      25,
        "isActive": true,
    },
    "bob": {
        "name":     "Bob Smith",
        "email":    "bob@example.com",
        "age":      30,
        "isActive": false,
    },
}
```

### 3. **Shopping Cart**
```go
cart := map[string]int{
    "apple":     5,
    "banana":    3,
    "orange":    2,
    "milk":      1,
}
```

### 4. **Word Frequency Counter**
```go
wordCount := map[string]int{
    "hello": 3,
    "world": 2,
    "go":    5,
}
```

## How to Run Your Program

1. Open your terminal
2. Go to the maps folder: `cd 06-maps`
3. Run the program: `go run main.go`

## Try It Yourself!

### Exercise 1: Create and Use a Map
Create a map of your favorite movies and their ratings:

```go
movies := map[string]int{
    "The Matrix":     5,
    "Inception":      4,
    "Interstellar":   5,
}

// Add a new movie
movies["The Dark Knight"] = 5

// Check if a movie exists
if rating, exists := movies["Inception"]; exists {
    fmt.Printf("Inception rating: %d\n", rating)
}
```

### Exercise 2: Build a Simple Phone Book
Create a phone book system:

```go
phoneBook := make(map[string]string)

// Add contacts
phoneBook["Alice"] = "555-0101"
phoneBook["Bob"] = "555-0102"
phoneBook["Charlie"] = "555-0103"

// Look up a number
if number, exists := phoneBook["Alice"]; exists {
    fmt.Printf("Alice's number: %s\n", number)
} else {
    fmt.Println("Alice not found in phone book")
}
```

### Exercise 3: Grade Book with Maps
Build on your arrays/slices knowledge to create a grade book:

```go
gradeBook := map[string][]int{
    "Alice":   {95, 87, 92, 78, 100},
    "Bob":     {88, 76, 94, 82, 89},
    "Charlie": {91, 85, 88, 90, 87},
}

// Calculate average for a student
func getAverage(grades []int) float64 {
    if len(grades) == 0 {
        return 0
    }
    sum := 0
    for _, grade := range grades {
        sum += grade
    }
    return float64(sum) / float64(len(grades))
}

// Use it
aliceAverage := getAverage(gradeBook["Alice"])
fmt.Printf("Alice's average: %.2f\n", aliceAverage)
```

### Exercise 4: Inventory Management
Create an inventory system:

```go
inventory := map[string]map[string]interface{}{
    "laptop": {
        "quantity": 10,
        "price":    999.99,
        "category": "electronics",
    },
    "mouse": {
        "quantity": 25,
        "price":    29.99,
        "category": "accessories",
    },
}

// Function to check low stock
func checkLowStock(inv map[string]map[string]interface{}, threshold int) []string {
    var lowStock []string
    for item, info := range inv {
        if quantity, ok := info["quantity"].(int); ok && quantity < threshold {
            lowStock = append(lowStock, item)
        }
    }
    return lowStock
}

// Use it
lowStockItems := checkLowStock(inventory, 15)
fmt.Printf("Low stock items: %v\n", lowStockItems)
```

### Exercise 5: Word Frequency Analyzer
Build a text analysis tool:

```go
func analyzeText(text string) map[string]int {
    words := strings.Fields(strings.ToLower(text))
    wordCount := make(map[string]int)
    
    for _, word := range words {
        // Remove punctuation
        word = strings.Trim(word, ".,!?;:")
        if word != "" {
            wordCount[word]++
        }
    }
    return wordCount
}

// Test it
text := "Hello world! Hello Go! Go is amazing!"
frequencies := analyzeText(text)
fmt.Printf("Word frequencies: %v\n", frequencies)
```

## Common Mistakes to Avoid

### 1. **Forgetting the Comma Ok Idiom**
```go
// ❌ WRONG: This might give you a zero value for a valid key!
value := myMap["key"]
if value != 0 {  // What if the actual value is 0?
    // This logic is flawed!
}

// ✅ CORRECT: Use comma ok idiom
if value, exists := myMap["key"]; exists {
    // Now you know the key definitely exists
}
```

### 2. **Assuming Map Order**
```go
// ❌ WRONG: Maps don't guarantee order!
for key, value := range myMap {
    fmt.Printf("%s: %v\n", key, value)  // Order might change each run!
}

// ✅ CORRECT: If you need order, use slices or sort the keys
keys := make([]string, 0, len(myMap))
for key := range myMap {
    keys = append(keys, key)
}
sort.Strings(keys)  // Now iterate in sorted order
for _, key := range keys {
    fmt.Printf("%s: %v\n", key, myMap[key])
}
```

### 3. **Modifying Maps During Iteration**
```go
// ❌ WRONG: Modifying map during iteration can cause issues
for key, value := range myMap {
    if value < 0 {
        delete(myMap, key)  // Can cause problems!
    }
}

// ✅ CORRECT: Collect keys to delete, then delete after iteration
var keysToDelete []string
for key, value := range myMap {
    if value < 0 {
        keysToDelete = append(keysToDelete, key)
    }
}
for _, key := range keysToDelete {
    delete(myMap, key)
}
```

### 4. **Not Checking Key Existence**
```go
// ❌ WRONG: Accessing non-existent keys returns zero value
userAge := userAges["nonexistent"]  // Returns 0
if userAge > 18 {  // This will be true for 0!
    fmt.Println("User is adult")  // Wrong!
}

// ✅ CORRECT: Always check if key exists
if userAge, exists := userAges["nonexistent"]; exists {
    if userAge > 18 {
        fmt.Println("User is adult")
    }
} else {
    fmt.Println("User not found")
}
```

## Key Takeaways

1. **Maps are key-value collections** - Access data by name, not position
2. **Use comma ok idiom** - Always check if keys exist
3. **Maps are reference types** - Copy carefully when you need independent data
4. **Order is not guaranteed** - Don't rely on map iteration order
5. **Zero values exist** - Missing keys return appropriate zero values
6. **Great for lookups** - Perfect for databases, configuration, and caches
7. **Range loops work** - Iterate through keys, values, or both
8. **Functions love maps** - Pass maps as parameters and return them

## When to Use Maps vs. Slices

### **Use Maps When:**
- You need to **look up data by name/key**
- **Order doesn't matter**
- You have **unique identifiers** for your data
- Building **databases, caches, or lookup tables**

### **Use Slices When:**
- You need **ordered collections**
- You're **iterating through all items**
- You need **index-based access**
- Building **lists, queues, or stacks**

## Next Steps

After this chapter, you'll learn about:
- **Structs and methods** (custom data types with behavior)
- **Interfaces** (Go's way of polymorphism)
- **Error handling** (building robust programs)
- **Goroutines and channels** (Go's concurrency features)

## Practice Projects

### Project 1: Student Management System
Build a complete system that stores student information, grades, and attendance using maps and slices together.

### Project 2: Configuration Manager
Create a system that loads, validates, and manages application configuration from different sources.

### Project 3: Simple Database
Build a basic in-memory database with CRUD operations (Create, Read, Update, Delete).

### Project 4: Cache System
Implement a simple caching system that stores frequently accessed data with expiration times.

---

**Excellent work! You now understand Go's three main collection types: arrays, slices, and maps! 🎉**

Maps are incredibly powerful and are used extensively in real Go programs. They're perfect for building databases, configuration systems, caches, and any application where you need to look up data by name rather than position.

Practice creating maps, using the comma ok idiom, and building functions that work with maps. You'll find yourself using them constantly in real-world Go development! 🚀 