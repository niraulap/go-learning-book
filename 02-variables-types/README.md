# Go Variables and Types - Chapter 2 🐹

Welcome to Chapter 2! Now that you can print text, let's learn how to store and work with different types of data in Go.

## What You'll Learn

- How to declare variables in Go
- Different data types and when to use them
- How Go automatically figures out types
- How to convert between types
- What constants are and how to use them

## Your First Variables

```go
package main

import "fmt"

func main() {
    var name string = "Alice"
    var age int = 25
    var height float64 = 5.6
    var isStudent bool = true
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

## What Each Part Does

### `var name string = "Alice"`
- `var` - tells Go you want to create a variable
- `name` - the name of your variable
- `string` - the type of data it will hold
- `"Alice"` - the actual value

### `var age int = 25`
- `int` means integer (whole numbers)
- Can store: 1, 2, 3, -1, -2, 0, etc.

### `var height float64 = 5.6`
- `float64` means decimal number
- Can store: 3.14, 2.5, 0.0, etc.

### `var isStudent bool = true`
- `bool` means boolean (true or false)
- Can store: `true` or `false`

## Three Ways to Declare Variables

### Method 1: Explicit Type
```go
var name string = "Alice"
var age int = 25
```
- You tell Go exactly what type you want
- Most clear and explicit
- Good for beginners

### Method 2: Type Inference
```go
var city = "New York"
var population = 8336817
```
- Go figures out the type automatically
- Cleaner syntax
- Still uses `var` keyword

### Method 3: Short Declaration
```go
country := "Canada"
area := 9984670
```
- Use `:=` instead of `var`
- Go figures out the type
- Can only use inside functions
- Most common in Go code

## Data Types in Go

### Integer Types
- `int` - regular integer (most common)
- `int8` - small integer (-128 to 127)
- `int64` - big integer
- `uint` - positive integers only

### Float Types
- `float32` - decimal number (less precise)
- `float64` - decimal number (more precise, default)

### Other Types
- `string` - text like "Hello"
- `bool` - true or false

## Type Inference

Go is smart! It can figure out what type you want:

```go
text := "Hello"     // Go knows this is string
number := 42        // Go knows this is int
decimal := 3.14     // Go knows this is float64
flag := true        // Go knows this is bool
```

## Type Conversion

Sometimes you need to change one type to another:

```go
var intValue int = 42
var floatValue float64 = float64(intValue)  // Convert int to float
var stringValue string = fmt.Sprintf("%d", intValue)  // Convert int to string
```

## Constants

Constants are values that never change:

```go
const PI = 3.14159
const MAX_SIZE = 1000
const APP_NAME = "GoPractice"
```

- Use `const` instead of `var`
- Value must be known when writing the code
- Convention: Use UPPER_CASE for constants

## Zero Values

When you declare a variable without giving it a value, Go gives it a default:

```go
var defaultInt int        // Gets 0
var defaultFloat float64  // Gets 0.0
var defaultString string  // Gets "" (empty string)
var defaultBool bool      // Gets false
```

## How to Run Your Program

1. Open your terminal
2. Go to the variables-types folder: `cd 02-variables-types`
3. Run the program: `go run main.go`

## Try It Yourself!

1. Change the values in the variables
2. Add new variables with different types
3. Try type conversion between different types
4. Create your own constants

## Common Mistakes to Avoid

1. **Wrong Type Names** - Use `int` not `integer`, `bool` not `boolean`
2. **Missing Quotes** - Strings need quotes: `"Hello"` not `Hello`
3. **Type Mismatch** - Can't assign a string to an int variable
4. **Forgetting :=** - Use `:=` for short declaration, not `=`

## Key Takeaways

1. **Variables store data** - Think of them as labeled boxes
2. **Go has types** - Different types for different kinds of data
3. **Type inference is smart** - Go can guess what you want
4. **Constants never change** - Use them for values that stay the same
5. **Zero values exist** - Variables get sensible defaults

## Next Steps

After this chapter, you'll learn about:
- Control structures (if statements, loops)
- Functions (reusable blocks of code)
- Arrays and slices (collections of data)

---

**Great job! You now understand how to store and work with data in Go! 🎉** 