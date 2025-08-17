# Go Control Structures: Complete Beginner's Guide 🐹

Welcome to the world of Go control structures! This guide is designed to take you from complete beginner to confident Go programmer. We'll cover everything step by step with clear examples.

## Table of Contents
1. [What Are Control Structures?](#what-are-control-structures)
2. [IF Statements](#if-statements)
3. [FOR Loops](#for-loops)
4. [RANGE Loop](#range-loop)
5. [SWITCH Statements](#switch-statements)
6. [DEFER Statement](#defer-statement)
7. [Common Mistakes to Avoid](#common-mistakes-to-avoid)
8. [Practice Exercises](#practice-exercises)
9. [Next Steps](#next-steps)

## What Are Control Structures? 🤔

Control structures are the building blocks that let your program make decisions and repeat actions. Think of them as the "brain" of your program that controls the flow of execution.

**In Go, you have:**
- **IF statements** - Make decisions (if this, then that)
- **FOR loops** - Repeat actions multiple times
- **SWITCH statements** - Choose from multiple options
- **RANGE loops** - Go's special way to iterate over collections
- **DEFER statements** - Schedule actions for later

## IF Statements 🔍

### Basic IF Statement

**Syntax:**
```go
if condition {
    // code to run if condition is true
}
```

**Example:**
```go
age := 18
if age >= 18 {
    fmt.Println("You are an adult!")
}
```

**Key Points:**
- ✅ **No parentheses needed** around the condition
- ✅ **Curly braces are required** (even for single lines)
- ✅ **Condition must be a boolean** (true/false)

### IF-ELSE Statement

**Syntax:**
```go
if condition {
    // code if condition is true
} else {
    // code if condition is false
}
```

**Example:**
```go
age := 16
if age >= 18 {
    fmt.Println("You can vote!")
} else {
    fmt.Println("You cannot vote yet!")
}
```

### IF-ELSE IF-ELSE Chain

**Syntax:**
```go
if condition1 {
    // code if condition1 is true
} else if condition2 {
    // code if condition2 is true
} else {
    // code if no conditions are true
}
```

**Example:**
```go
score := 85
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: D")
}
```

### IF with Initialization (Go-Specific!)

This is one of Go's coolest features! You can initialize a variable and check it in one line.

**Syntax:**
```go
if variable := expression; condition {
    // use variable here
}
// variable is NOT available here
```

**Example:**
```go
if score := getScore(); score >= 90 {
    fmt.Printf("Excellent score: %d\n", score)
} else {
    fmt.Printf("Good score: %d\n", score)
}
// score variable is NOT available here!
```

**Why This Is Useful:**
- ✅ **Cleaner code** - no need to declare variable separately
- ✅ **Scope control** - variable only exists where you need it
- ✅ **Performance** - can avoid unnecessary calculations

## FOR Loops 🔄

Go only has one loop construct: `for`. But don't worry - it's incredibly powerful and can do everything other languages do with `while`, `do-while`, and `for` loops!

### Traditional FOR Loop

**Syntax:**
```go
for initialization; condition; increment {
    // code to repeat
}
```

**Example:**
```go
for i := 0; i < 5; i++ {
    fmt.Printf("Count: %d\n", i)
}
// Output: Count: 0, Count: 1, Count: 2, Count: 3, Count: 4
```

**Breakdown:**
- `i := 0` - Initialize counter to 0
- `i < 5` - Keep looping while i is less than 5
- `i++` - Add 1 to i after each iteration

### While-Style Loop

**Syntax:**
```go
for condition {
    // code to repeat
}
```

**Example:**
```go
counter := 0
for counter < 3 {
    fmt.Printf("Counter: %d\n", counter)
    counter++
}
// Output: Counter: 0, Counter: 1, Counter: 2
```

**When to Use:**
- ✅ When you don't know how many iterations you need
- ✅ When the increment is not regular
- ✅ When you need complex loop conditions

### Infinite Loop

**Syntax:**
```go
for {
    // code to repeat forever (or until break)
}
```

**Example:**
```go
attempts := 0
for {
    attempts++
    if attempts > 3 {
        fmt.Println("Max attempts reached!")
        break  // Exit the loop
    }
    fmt.Printf("Attempt %d\n", attempts)
}
// Output: Attempt 1, Attempt 2, Attempt 3, Max attempts reached!
```

**Important:**
- ⚠️ **Always have a way to exit** (like `break` or `return`)
- ⚠️ **Infinite loops can crash your program** if not handled properly

### Loop Control Keywords

**`break`** - Exit the loop immediately
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Exit when i equals 5
    }
    fmt.Printf("%d ", i)
}
// Output: 0 1 2 3 4
```

**`continue`** - Skip to next iteration
```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue  // Skip printing when i equals 2
    }
    fmt.Printf("%d ", i)
}
// Output: 0 1 3 4
```

## RANGE Loop (Go's Superpower!) 🚀

The `range` loop is Go's special way to iterate over collections. It's safer and more convenient than traditional loops!

### Range Over Slice/Array

**Syntax:**
```go
for index, value := range slice {
    // use index and value
}
```

**Example:**
```go
numbers := []int{10, 20, 30, 40, 50}

// Get both index and value
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
// Output:
// Index: 0, Value: 10
// Index: 1, Value: 20
// Index: 2, Value: 30
// Index: 3, Value: 40
// Index: 4, Value: 50
```

**Variations:**

**Only value (ignore index):**
```go
for _, value := range numbers {
    fmt.Printf("Value: %d\n", value)
}
// Output: Value: 10, Value: 20, Value: 30, Value: 40, Value: 50
```
**Note:** `_` means "ignore this value"

**Only index:**
```go
for index := range numbers {
    fmt.Printf("Index: %d\n", index)
}
// Output: Index: 0, Index: 1, Index: 2, Index: 3, Index: 4
```

### Range Over Map

**Syntax:**
```go
for key, value := range map {
    // use key and value
}
```

**Example:**
```go
person := map[string]string{
    "name": "Alice",
    "city": "New York",
    "job":  "Developer",
}

for key, value := range person {
    fmt.Printf("%s: %s\n", key, value)
}
// Output (order may vary):
// name: Alice
// city: New York
// job: Developer
```

**Important:** Map iteration order is **not guaranteed** in Go!

### Range Over String

**Syntax:**
```go
for index, char := range string {
    // use index and character
}
```

**Example:**
```go
word := "Go"
for index, char := range word {
    fmt.Printf("Index: %d, Character: %c (Unicode: %d)\n", index, char, char)
}
// Output:
// Index: 0, Character: G (Unicode: 71)
// Index: 1, Character: o (Unicode: 111)
```

**Why Range is Awesome:**
- ✅ **Safer** - no risk of going out of bounds
- ✅ **Cleaner** - no need to manage loop variables
- ✅ **More readable** - intent is clear
- ✅ **Works with any collection** - slices, maps, strings, channels

## SWITCH Statements 🔀

Go's switch statement is powerful and safer than other languages because it automatically breaks after each case!

### Basic Switch

**Syntax:**
```go
switch variable {
case value1:
    // code for value1
case value2:
    // code for value2
default:
    // code if no cases match
}
```

**Example:**
```go
day := "Friday"
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Tuesday", "Wednesday", "Thursday":
    fmt.Println("Middle of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Unknown day")
}
// Output: TGIF!
```

**Key Features:**
- ✅ **No break needed** - Go automatically breaks after each case
- ✅ **Multiple values per case** - `case "Tuesday", "Wednesday"`
- ✅ **Default case** - handles unmatched values

### Switch with Expression

**Syntax:**
```go
switch {
case condition1:
    // code if condition1 is true
case condition2:
    // code if condition2 is true
default:
    // code if no conditions are true
}
```

**Example:**
```go
score := 85
switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
case score >= 70:
    fmt.Println("Grade: C")
default:
    fmt.Println("Grade: D")
}
// Output: Grade: B
```

### Fallthrough (Go-Specific!)

If you want cases to fall through (like in C), use the `fallthrough` keyword.

**Example:**
```go
grade := 85
switch {
case grade >= 90:
    fmt.Print("A")
    fallthrough  // Continue to next case
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
// Output: BCD grade
```

**When to Use Fallthrough:**
- ⚠️ **Rarely needed** - Go's automatic break is usually better
- ✅ **Useful for specific logic** like grade ranges
- ⚠️ **Can make code harder to read** - use sparingly

## DEFER Statement (Go's Unique Feature!) 🚪

The `defer` statement is one of Go's most unique features. It schedules a function call to run when the current function exits.

### Basic Defer

**Syntax:**
```go
defer functionCall()
```

**Example:**
```go
func example() {
    fmt.Println("Starting function...")
    
    defer fmt.Println("This runs LAST!")
    defer fmt.Println("This runs SECOND to last!")
    
    fmt.Println("Function body executing...")
    fmt.Println("About to exit function...")
}

// Output:
// Starting function...
// Function body executing...
// About to exit function...
// This runs SECOND to last!
// This runs LAST!
```

### Defer Order (LIFO)

Defer statements are executed in **Last In, First Out** order:

```go
func deferOrder() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    
    fmt.Println("Function body")
}

// Output:
// Function body
// Third defer
// Second defer
// First defer
```

### Common Use Cases

**File Operations:**
```go
func readFile() {
    file := openFile("data.txt")
    defer file.Close()  // Always close the file
    
    // ... read file operations ...
    // file.Close() is automatically called here
}
```

**Database Connections:**
```go
func queryDatabase() {
    db := connectToDB()
    defer db.Close()  // Always close the connection
    
    // ... database operations ...
    // db.Close() is automatically called here
}
```

**Why Defer is Awesome:**
- ✅ **Automatic cleanup** - resources are always released
- ✅ **Cleaner code** - no need to remember cleanup in multiple places
- ✅ **Error-safe** - cleanup happens even if function returns early
- ✅ **Multiple defers** - can have several cleanup operations

## Common Mistakes to Avoid ⚠️

### 1. **Missing Curly Braces**
```go
// ❌ Wrong - missing braces
if age >= 18
    fmt.Println("Adult!")

// ✅ Correct
if age >= 18 {
    fmt.Println("Adult!")
}
```

### 2. **Unnecessary Parentheses**
```go
// ❌ Wrong - unnecessary parentheses
if (age >= 18) {
    fmt.Println("Adult!")
}

// ✅ Correct
if age >= 18 {
    fmt.Println("Adult!")
}
```

### 3. **Infinite Loop Without Exit**
```go
// ❌ Dangerous - no way to exit
for {
    fmt.Println("This runs forever!")
}

// ✅ Safe - has exit condition
for {
    if shouldStop {
        break
    }
    fmt.Println("Processing...")
}
```

### 4. **Ignoring Range Variables**
```go
// ❌ Wrong - ignoring both index and value
for range numbers {
    fmt.Println("Processing...")
}

// ✅ Better - at least acknowledge what you're iterating
for _, value := range numbers {
    fmt.Printf("Processing: %d\n", value)
}
```

### 5. **Forgetting Defer Order**
```go
// ❌ Confusing - defer order matters
defer fmt.Println("First")
defer fmt.Println("Second")
defer fmt.Println("Third")

// ✅ Clear - think about the order
defer fmt.Println("Cleanup 1")
defer fmt.Println("Cleanup 2")
defer fmt.Println("Cleanup 3")
```

## Practice Exercises 🏋️

### Exercise 1: Grade Calculator
Create a program that takes a score and prints the grade using if-else statements.

### Exercise 2: Number Guessing Game
Create a simple number guessing game using loops and if statements.

### Exercise 3: FizzBuzz
Print numbers 1-100, but for multiples of 3 print "Fizz", multiples of 5 print "Buzz", and multiples of both print "FizzBuzz".

### Exercise 4: Word Counter
Use range to count characters in a string and print each character with its count.

### Exercise 5: Resource Manager
Create a function that uses defer to ensure resources are properly cleaned up.

## Run the Program

```bash
cd 03-control-structures
go run main.go
```

## Next Steps 🚀

After mastering control structures, you're ready for:

1. **Functions** - Parameters, return values, multiple returns
2. **Arrays and Slices** - Working with collections of data
3. **Maps** - Key-value data structures
4. **Structs and Methods** - Creating custom types
5. **Interfaces** - Go's way of polymorphism
6. **Goroutines and Channels** - Concurrency (Go's superpower!)

## Additional Resources 📚

- [Go Official Tour](https://tour.golang.org/) - Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) - Practical examples
- [Effective Go](https://golang.org/doc/effective_go.html) - Best practices
- [Go Language Specification](https://golang.org/ref/spec) - Complete reference

## Tips for Success 💡

1. **Practice regularly** - Control structures are fundamental
2. **Experiment with variations** - Try different approaches
3. **Read other people's code** - See how they use control structures
4. **Use Go's tools** - `go fmt` to format code, `go vet` to find issues
5. **Don't rush** - Master the basics before moving to advanced topics

---

**Happy coding! Remember: The best way to learn Go is to write lots of Go code! 🎉**

If you have questions or get stuck, don't hesitate to experiment and try different approaches. Go is designed to be simple and readable, so trust your instincts! 