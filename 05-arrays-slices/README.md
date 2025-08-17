# Go Arrays and Slices - Chapter 5 🐹

Welcome to Chapter 5! Now that you can create functions and control program flow, let's learn how to work with collections of data. Arrays and slices are fundamental to Go programming and will open up many new possibilities.

## What You'll Learn

- How to create and use arrays (fixed-size collections)
- How to work with slices (dynamic collections - Go's superpower!)
- Range loops with collections (building on Chapter 3)
- Modifying and manipulating collections
- Functions that work with arrays and slices
- When to use arrays vs. slices

## What are Collections?

Collections are groups of related data stored together. Think of them like:
- **Arrays**: A fixed row of boxes (you can't add or remove boxes)
- **Slices**: A flexible row of boxes (you can add, remove, and resize)

## Arrays (Fixed-Size Collections)

Arrays have a fixed size that you must specify when creating them.

### Creating Arrays

```go
// Method 1: Declare then set values
var numbers [5]int
numbers[0] = 10
numbers[1] = 20
numbers[2] = 30
numbers[3] = 40
numbers[4] = 50

// Method 2: Array literal (create and initialize at once)
colors := [3]string{"red", "green", "blue"}

// Method 3: Size inference (let Go count for you)
grades := [...]int{95, 87, 92, 78, 100}
```

### Accessing Array Elements

```go
// Access by index (starting from 0)
firstGrade := grades[0]    // 95
lastGrade := grades[len(grades)-1]  // 100

// Get array length
arrayLength := len(grades)  // 5
```

### Array Limitations

- **Fixed size** - Cannot grow or shrink after creation
- **Must know size** at declaration time
- **Less flexible** than slices
- **Good for** when you know exactly how many items you need

## Slices (Dynamic Collections - Go's Superpower!)

Slices are dynamic collections that can grow and shrink. They're much more flexible than arrays and are used more often in Go.

### Creating Slices

```go
// Method 1: Slice literal
fruits := []string{"apple", "banana", "orange"}

// Method 2: Empty slice
var emptySlice []int

// Method 3: Using make (with length and capacity)
numbers := make([]int, 3, 5)  // length 3, capacity 5
```

### Adding Elements (append)

```go
fruits := []string{"apple", "banana"}

// Add one element
fruits = append(fruits, "orange")

// Add multiple elements
fruits = append(fruits, "grape", "mango")
```

**⚠️ Important: Why `fruits = append(fruits, "orange")`?**

This is one of the most important concepts in Go slices! The `append()` function can do two different things:

1. **Modify the original slice** (if there's enough capacity)
2. **Create a completely new slice** (if capacity is exceeded)

Go requires you to assign the result because **you can't know in advance** which will happen!

```go
fruits := make([]string, 0, 2)  // length: 0, capacity: 2

// First append - fits in capacity, modifies original slice
fruits = append(fruits, "apple")    // length: 1, capacity: 2

// Second append - still fits, modifies original slice  
fruits = append(fruits, "banana")   // length: 2, capacity: 2

// Third append - EXCEEDS capacity! Creates NEW slice
fruits = append(fruits, "orange")   // length: 3, capacity: 4 (doubled!)

// Fourth append - fits in new capacity, modifies current slice
fruits = append(fruits, "grape")    // length: 4, capacity: 4
```

**Why not just `append(fruits, "orange")`?**

```go
// ❌ WRONG - This could lose the new slice!
append(fruits, "orange")  // Result not assigned!

// ✅ CORRECT - Always assign the result
fruits = append(fruits, "orange")  // Guaranteed to have updated slice
```

**Key Points:**
- **`append()` always returns a slice** - either the original (modified) or a new one
- **Capacity determines behavior** - if exceeded, new memory is allocated
- **Assignment is mandatory** - Go's way of ensuring you always have the correct slice
- **This design is intentional** - provides memory efficiency and safety

### Slicing Operations

Slicing lets you get parts of a slice:

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Get first 3 elements
firstThree := numbers[:3]    // [1, 2, 3]

// Get last 3 elements
lastThree := numbers[len(numbers)-3:]  // [8, 9, 10]

// Get middle elements (index 2 to 5)
middle := numbers[2:6]       // [3, 4, 5, 6]

// Get from index 1 to end
fromIndex1 := numbers[1:]    // [2, 3, 4, 5, 6, 7, 8, 9, 10]
```

### Slice Capacity and Growth

Slices have both length and capacity:
- **Length**: How many elements are currently in the slice
- **Capacity**: How many elements the slice can hold before growing

```go
slice := make([]int, 0, 3)  // length 0, capacity 3
fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

// Add elements
slice = append(slice, 1)     // length 1, capacity 3
slice = append(slice, 2)     // length 2, capacity 3
slice = append(slice, 3)     // length 3, capacity 3
slice = append(slice, 4)     // length 4, capacity 6 (doubled!)
```

## Working with Collections

### Range Loops (Building on Chapter 3)

You already know range loops! Now use them with collections:

```go
scores := []int{95, 87, 92, 78, 100}

// Get both index and value
for i, score := range scores {
    fmt.Printf("Index %d: Score %d\n", i, score)
}

// Get only values
for _, score := range scores {
    if score >= 90 {
        fmt.Printf("High score: %d\n", score)
    }
}

// Get only indices
for i := range scores {
    fmt.Printf("Position %d\n", i)
}
```

### Modifying Slices

```go
numbers := []int{1, 2, 3, 4, 5}

// Change an element
numbers[2] = 30

// Remove an element (index 2)
numbers = append(numbers[:2], numbers[3:]...)

// Insert an element at index 2
numbers = append(numbers[:2], append([]int{25}, numbers[2:]...)...)
```

### Copying Slices

```go
original := []int{10, 20, 30, 40, 50}
copied := make([]int, len(original))
copy(copied, original)

// Now original and copied are independent
original[0] = 100  // copied[0] is still 10
```

## Functions with Collections

### Functions that Take Slices as Parameters

```go
func doubleSlice(numbers []int) {
    for i := range numbers {
        numbers[i] = numbers[i] * 2
    }
}

// Use it
data := []int{10, 20, 30}
doubleSlice(data)  // data is now [20, 40, 60]
```

### Functions that Return New Slices

```go
func doubleSliceReturn(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * 2
    }
    return result
}

// Use it
original := []int{1, 2, 3}
doubled := doubleSliceReturn(original)  // original unchanged, doubled is [2, 4, 6]
```

### Common Collection Functions

```go
// Find maximum value
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

// Filter elements
func filterHighScores(scores []int, threshold int) []int {
    var result []int
    for _, score := range scores {
        if score >= threshold {
            result = append(result, score)
        }
    }
    return result
}
```

## Arrays vs. Slices: When to Use Which?

### Use Arrays When:
- You know the exact size needed
- Size will never change
- You want to ensure data integrity
- Working with fixed-size data structures

### Use Slices When:
- Size might change
- You don't know the size in advance
- You need flexibility
- Most common use case in Go

## How to Run Your Program

1. Open your terminal
2. Go to the arrays-slices folder: `cd 05-arrays-slices`
3. Run the program: `go run main.go`

## Try It Yourself!

### Exercise 1: Create and Modify Arrays
Create an array of 5 numbers and modify the middle element:

```go
numbers := [5]int{10, 20, 30, 40, 50}
numbers[2] = 35
fmt.Printf("Modified array: %v\n", numbers)
```

### Exercise 2: Work with Slices
Create a slice of fruits and add/remove elements:

```go
fruits := []string{"apple", "banana"}
fruits = append(fruits, "orange")
fruits = append(fruits, "grape", "mango")
fmt.Printf("Fruits: %v\n", fruits)
```

### Exercise 3: Slice Operations
Practice slicing operations:

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
firstHalf := numbers[:5]
secondHalf := numbers[5:]
fmt.Printf("First half: %v\n", firstHalf)
fmt.Printf("Second half: %v\n", secondHalf)
```

### Exercise 4: Functions with Collections
Create a function that finds the average of a slice:

```go
func findAverage(numbers []int) float64 {
    if len(numbers) == 0 {
        return 0
    }
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return float64(sum) / float64(len(numbers))
}

// Test it
scores := []int{85, 90, 78, 92, 88}
average := findAverage(scores)
fmt.Printf("Average score: %.2f\n", average)
```

### Exercise 5: Filter and Transform
Create functions to filter and transform data:

```go
func filterEvenNumbers(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

func squareNumbers(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * num
    }
    return result
}

// Test them
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := filterEvenNumbers(data)
squared := squareNumbers(evens)
fmt.Printf("Original: %v\n", data)
fmt.Printf("Even numbers: %v\n", evens)
fmt.Printf("Squared evens: %v\n", squared)
```

## Common Mistakes to Avoid

### 1. **Array Index Out of Bounds**
```go
numbers := [3]int{1, 2, 3}
value := numbers[5]  // ❌ Panic! Index 5 doesn't exist
```

### 2. **Forgetting to Assign append Result**
```go
fruits := []string{"apple"}
fruits = append(fruits, "banana")  // ✅ Correct
append(fruits, "banana")           // ❌ Wrong! Result not assigned
```

**Why this matters:**
- `append()` can create a new slice if capacity is exceeded
- Without assignment, you might lose the new slice
- Go requires assignment to ensure you always have the correct slice
- This is a fundamental Go design principle for safety

### 3. **Modifying Slice Without Copying**
```go
original := []int{1, 2, 3}
reference := original        // ❌ This is a reference, not a copy
reference[0] = 100          // Changes both original and reference

// ✅ Correct way to copy
copied := make([]int, len(original))
copy(copied, original)
copied[0] = 100             // Only copied changes
```

### 4. **Ignoring Slice Capacity**
```go
slice := make([]int, 0, 3)
for i := 0; i < 10; i++ {
    slice = append(slice, i)  // Capacity will grow automatically
    // But this can be inefficient for large slices
}
```

## Key Takeaways

1. **Arrays are fixed-size** - Good when you know the exact size needed
2. **Slices are dynamic** - Can grow and shrink as needed
3. **Range loops work great** with collections
4. **append() adds elements** to slices
5. **Slicing operations** let you get parts of collections
6. **Functions can take and return** collections
7. **Copy slices** when you need independent copies
8. **Slices are more common** in Go than arrays

## Next Steps

After this chapter, you'll learn about:
- Maps (key-value collections)
- Structs and methods (custom data types)
- Interfaces (Go's way of polymorphism)
- Error handling (proper error management)

## Practice Projects

### Project 1: Grade Book
Create a program that stores student grades in slices, calculates averages, finds highest/lowest scores, and filters students by performance.

### Project 2: Shopping Cart
Build a shopping cart system using slices to store items, quantities, and prices. Include functions to add/remove items and calculate totals.

### Project 3: Number Analyzer
Create functions that analyze number collections: find min/max, calculate statistics, filter by criteria, and transform data.

### Project 4: Text Processor
Build a text processing system that works with slices of strings: find words, count occurrences, filter by length, and transform text.

---

**Excellent work! You now understand how to work with collections of data in Go! 🎉**

Arrays and slices are fundamental to Go programming. Practice creating them, manipulating them, and writing functions that work with them. You'll use these concepts in almost every Go program you write! 