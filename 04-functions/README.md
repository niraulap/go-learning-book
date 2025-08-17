# Go Functions - Chapter 4 🐹

Welcome to Chapter 4! Now that you can write code and control program flow, let's learn how to organize your code into reusable functions. Functions are the building blocks of organized, maintainable programs.

## What You'll Learn

- How to create and use functions in Go
- Functions with parameters and return values
- Multiple return values (Go's special feature!)
- How to organize code into reusable blocks
- Function scope and visibility

## What is a Function?

A function is a block of code that performs a specific task. Think of it as a recipe - you give it ingredients (parameters), it does some work, and gives you back a result (return value).

## Basic Function Syntax

```go
func functionName() {
    // code goes here
}

// Call the function
functionName()
```

## Your First Functions

```go
package main

import "fmt"

func main() {
    sayHello()  // Call the function
}

func sayHello() {
    fmt.Println("Hello from a function!")
}
```

## What Each Part Does

### `func sayHello()`
- `func` - tells Go you're creating a function
- `sayHello` - the name of your function
- `()` - empty parentheses mean no parameters
- `{ }` - curly braces contain the function code

### `sayHello()`
- This calls (runs) the function
- The function executes and prints "Hello from a function!"

## Functions with Parameters

Parameters are like inputs to your function:

```go
func greetPerson(name string) {
    fmt.Printf("Hello, %s! Welcome to Go!\n", name)
}

// Call with different names
greetPerson("Alice")
greetPerson("Bob")
```

### Multiple Parameters
```go
func printInfo(name string, age int, city string) {
    fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
}

printInfo("Alice", 25, "New York")
```

## Functions with Return Values

Functions can give you back data:

```go
func calculateSum(a int, b int) int {
    return a + b
}

// Use the return value
result := calculateSum(5, 3)
fmt.Printf("Sum: %d\n", result)
```

### Different Return Types
```go
func createGreeting(name string) string {
    return fmt.Sprintf("Hello, %s! Welcome to Go!", name)
}

func checkAge(age int) bool {
    return age >= 18
}

// Use the returned values
message := createGreeting("Alice")
isAdult := checkAge(18)
```

## Multiple Return Values (Go's Special Feature!)

This is what makes Go unique! Functions can return multiple values:

```go
func divideNumbers(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Get both values
quotient, remainder := divideNumbers(17, 5)
fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)
```

### Multiple Values of Different Types
```go
func getPersonInfo(name string, age int, isStudent bool) (string, int, bool) {
    return name, age, isStudent
}

// Get all three values
name, age, isStudent := getPersonInfo("Alice", 20, true)
```

## Success/Failure Pattern (Important Go Concept!)

The success/failure pattern is a powerful way to handle operations that might fail. Instead of crashing your program, functions tell you whether they succeeded and give you the result.

### Why Use Success/Failure Pattern?

In real programs, operations can fail for many reasons:
- **Math operations**: Division by zero
- **File operations**: File doesn't exist
- **Network operations**: Connection failed
- **User input**: Invalid data entered

### Basic Success/Failure Pattern

```go
func divideNumbers(a int, b int) (bool, int) {
    if b == 0 {
        return false, 0  // Failed: can't divide by zero
    }
    result := a / b
    return true, result  // Success: here's the result
}

// Use it
success, result := divideNumbers(10, 2)
if success {
    fmt.Printf("Division successful: %d\n", result)
} else {
    fmt.Printf("Division failed!\n")
}
```

### Advanced Success/Failure Pattern

```go
func performOperation(operation string, a int, b int) (bool, int) {
    switch operation {
    case "add":
        return true, a + b      // Success: return the sum
    case "multiply":
        return true, a * b      // Success: return the product
    case "divide":
        if b != 0 {
            return true, a / b  // Success: return the quotient
        }
        return false, 0         // Failed: division by zero
    default:
        return false, 0         // Failed: unknown operation
    }
}

// Use it with error handling
success, result := performOperation("divide", 10, 0)
if success {
    fmt.Printf("Operation successful: %d\n", result)
} else {
    fmt.Printf("Operation failed!\n")
    // Maybe try a different approach or use a default value
}
```

### Success/Failure with Error Messages

```go
func validatePassword(password string) (bool, string) {
    if len(password) < 8 {
        return false, "Password too short (need 8+ characters)"
    }
    if !strings.ContainsAny(password, "0123456789") {
        return false, "Password needs at least one number"
    }
    return true, "Password is valid!"
}

// Use it
success, message := validatePassword("abc123")
if success {
    fmt.Println("Password accepted!")
} else {
    fmt.Printf("Password rejected: %s\n", message)
}
```

### Success/Failure with Multiple Data

```go
func getUser(userID int) (bool, string, int) {
    if userID < 0 {
        return false, "Invalid user ID", 0
    }
    if userID == 0 {
        return false, "User not found", 0
    }
    // User found!
    return true, "Alice", 25
}

// Use it
found, name, age := getUser(123)
if found {
    fmt.Printf("User: %s, Age: %d\n", name, age)
} else {
    fmt.Printf("Error: %s\n", name)  // name contains error message
}
```

### Key Points About Success/Failure Pattern

1. **Always check success first** - Don't use the result if success is false
2. **The result value might be meaningless** if success is false
3. **This pattern prevents crashes** - Your program handles errors gracefully
4. **It's very common in Go** - You'll see this pattern everywhere
5. **Better than exceptions** - Errors don't crash your program

### When to Use Success/Failure Pattern

- **Operations that might fail** (file operations, network calls)
- **Input validation** (checking if user data is valid)
- **Resource operations** (database connections, file access)
- **Any time you need to handle errors gracefully**

## Function Benefits

### 1. **Reusability**
- Write once, use many times
- No need to copy and paste code

### 2. **Organization**
- Break large programs into small, manageable pieces
- Each function has one clear purpose

### 3. **Maintainability**
- Fix bugs in one place
- Easy to update and improve

### 4. **Testing**
- Test individual functions
- Easier to find and fix problems

### 5. **Error Handling**
- Use success/failure patterns to handle errors gracefully
- Prevent program crashes
- Provide better user experience

## How to Run Your Program

1. Open your terminal
2. Go to the functions folder: `cd 04-functions`
3. Run the program: `go run main.go`

## Try It Yourself!

### Exercise 1: Create a Simple Function
Add a new function that prints your name:

```go
func printMyName() {
    fmt.Println("My name is [Your Name]")
}

// Call it in main()
printMyName()
```

### Exercise 2: Function with Parameters
Create a function that takes a number and prints it doubled:

```go
func doubleNumber(num int) {
    result := num * 2
    fmt.Printf("%d doubled is %d\n", num, result)
}

doubleNumber(5)  // Should print: 5 doubled is 10
```

### Exercise 3: Function with Return Value
Create a function that calculates the area of a rectangle:

```go
func calculateArea(length int, width int) int {
    return length * width
}

area := calculateArea(5, 3)
fmt.Printf("Area: %d\n", area)
```

### Exercise 4: Multiple Return Values
Create a function that returns both the sum and product of two numbers:

```go
func sumAndProduct(a int, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

sum, product := sumAndProduct(4, 6)
fmt.Printf("Sum: %d, Product: %d\n", sum, product)
```

### Exercise 5: Success/Failure Pattern
Create a function that validates an age and returns success/failure with a message:

```go
func validateAge(age int) (bool, string) {
    if age < 0 {
        return false, "Age cannot be negative"
    }
    if age > 150 {
        return false, "Age seems unrealistic"
    }
    if age < 18 {
        return false, "Too young for this operation"
    }
    return true, "Age is valid"
}

// Test it
success, message := validateAge(25)
if success {
    fmt.Println("Age accepted!")
} else {
    fmt.Printf("Age rejected: %s\n", message)
}
```

## Common Mistakes to Avoid

### 1. **Missing Function Calls**
```go
func sayHello() {
    fmt.Println("Hello!")
}
// Don't forget to call it!
sayHello()  // This line is needed!
```

### 2. **Wrong Parameter Types**
```go
func greetPerson(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

greetPerson(42)  // ❌ Wrong! 42 is int, not string
greetPerson("42") // ✅ Correct! "42" is string
```

### 3. **Not Using Return Values**
```go
func calculateSum(a int, b int) int {
    return a + b
}

calculateSum(5, 3)  // ❌ Return value is ignored
result := calculateSum(5, 3)  // ✅ Return value is stored
```

### 4. **Wrong Number of Return Values**
```go
func divideNumbers(a int, b int) (int, int) {
    return a / b  // ❌ Only returning one value
    // Should return two values
}
```

### 5. **Not Checking Success in Success/Failure Pattern**
```go
func divideNumbers(a int, b int) (bool, int) {
    if b == 0 {
        return false, 0
    }
    return true, a / b
}

success, result := divideNumbers(10, 0)
fmt.Printf("Result: %d\n", result)  // ❌ Using result without checking success!
// Should be:
if success {
    fmt.Printf("Result: %d\n", result)
} else {
    fmt.Println("Division failed!")
}
```

## Key Takeaways

1. **Functions organize code** - Break large programs into small pieces
2. **Parameters are inputs** - Give functions data to work with
3. **Return values are outputs** - Functions give you back results
4. **Multiple returns are powerful** - Go's unique feature for better code
5. **Functions are reusable** - Write once, use many times
6. **Success/failure patterns prevent crashes** - Handle errors gracefully
7. **Always check success before using results** - Protect your program from errors

## Next Steps

After this chapter, you'll learn about:
- Arrays and slices (collections of data)
- Maps (key-value data structures)
- Structs and methods (custom data types)
- Interfaces (Go's way of polymorphism)

## Practice Projects

### Project 1: Calculator Functions
Create functions for basic math operations (add, subtract, multiply, divide) and use them to build a simple calculator. Use success/failure patterns to handle division by zero.

### Project 2: Personal Information Manager
Create functions to manage and display personal information (name, age, hobbies, etc.). Use success/failure patterns to validate input data.

### Project 3: Grade Calculator
Create functions to calculate grades, averages, and provide grade feedback. Use success/failure patterns to handle invalid grades.

### Project 4: File Operations Simulator
Create functions that simulate file operations (read, write, delete) and use success/failure patterns to handle various error conditions.

---

**Great job! You now understand how to organize your Go code into functions and handle errors gracefully! 🎉**

Functions are the foundation of good programming, and success/failure patterns make your programs robust and user-friendly. Practice creating them, experiment with different parameters and return values, and you'll be writing professional Go code in no time! 