# Go Error Handling - Chapter 10 🐹

Welcome to Chapter 10! Now that you understand interfaces and pointers, you're ready to learn about **error handling** - one of Go's most distinctive and powerful features. 

**Don't worry if this feels overwhelming at first!** Error handling is a big topic, but we'll break it down step by step, starting with the absolute basics and building up to advanced patterns. By the end, you'll understand why Go's approach is actually simpler and more reliable than what you might have seen in other languages.

## What You'll Learn

- **What errors are** in Go (hint: they're just values!)
- **Why Go handles errors differently** from other languages
- **How to check for errors** (the `if err != nil` pattern)
- **How to create your own error types** using your struct knowledge
- **How to handle multiple errors** and build error chains
- **Real examples** you can run and experiment with

## Let's Start Simple: What is an Error in Go?

### **Think of Errors Like This:**

Imagine you're asking someone to do a task:
- **Success**: They complete the task and give you the result
- **Failure**: They can't complete the task and tell you why

In Go, functions work the same way:
- **Success**: Function returns the result you want
- **Failure**: Function returns an error explaining what went wrong

### **The Basic Pattern:**

```go
// Every function that can fail returns two things:
result, err := someFunction()

// You ALWAYS check the error first
if err != nil {
    // Something went wrong - handle it!
    fmt.Printf("Error: %v\n", err)
    return err  // or handle it some other way
}

// If we get here, err is nil, so we can safely use result
fmt.Printf("Success! Result: %v\n", result)
```

## Why is Go Different from Other Languages?

### **Other Languages (Java, Python, C#) Use "Exceptions":**

```java
// Java - Exceptions can happen anywhere, anytime
try {
    user = database.getUser(id);  // This might "throw" an exception
    user.updateProfile(data);     // This might also throw an exception
    // If anything throws an exception, execution jumps to catch block
} catch (UserNotFoundException e) {
    // Handle user not found
} catch (DatabaseException e) {
    // Handle database problems
}
// Problems with exceptions:
// 1. You don't know which line might fail
// 2. Execution can jump around unexpectedly
// 3. It's easy to forget to handle errors
```

### **Go Uses "Return Values":**

```go
// Go - Errors are returned, not thrown
user, err := database.GetUser(id)
if err != nil {
    // Handle the error right here
    return fmt.Errorf("failed to get user: %w", err)
}

err = user.UpdateProfile(data)
if err != nil {
    // Handle this error too
    return fmt.Errorf("failed to update profile: %w", err)
}

// Benefits of Go's approach:
// 1. You know exactly where errors can happen
// 2. Execution flows normally (no jumping around)
// 3. You can't forget to handle errors (Go will warn you)
```

## Section 1: Go's Error Philosophy - Let's Break It Down

### **1. Errors are Just Values (Like Strings or Numbers)**

```go
// An error is just a variable that can hold error information
var err error                    // This creates an error variable
fmt.Printf("Error: %v\n", err)  // Prints: <nil> (nil means "no error")

// You can assign errors to variables
err = errors.New("something went wrong")
fmt.Printf("Error: %v\n", err)  // Prints: something went wrong

// You can compare errors
if err != nil {
    fmt.Println("There is an error!")
} else {
    fmt.Println("No error!")
}
```

### **2. What is an Interface? (Important Concept!)**

Before we talk about the `error` interface, let me explain what an **interface** is in Go. This is a fundamental concept!

#### **Think of Interfaces Like This:**

An interface is like a **contract** or **promise**. It says:
> "If you want to be considered an X, you must have these specific abilities"

For example:
- **To be a "Vehicle"**: You must be able to `Start()` and `Stop()`
- **To be a "Reader"**: You must be able to `Read()` data
- **To be an "error"**: You must be able to `Error()` and return a string

#### **The `error` Interface is Super Simple**

```go
// The error interface is a contract that says:
// "To be considered an error, you must have an Error() method that returns a string"
type error interface {
    Error() string
}

// That's it! Any type that has an Error() method is automatically an error
// Let's create a simple error type:
type SimpleError struct {
    Message string
}

// This method makes SimpleError satisfy the error interface
func (e SimpleError) Error() string {
    return e.Message
}

// Now SimpleError is an error! (because it has the Error() method)
var myError error = SimpleError{Message: "my custom error"}
fmt.Printf("My error: %v\n", myError)  // Prints: my custom error
```

#### **Why is This So Cool?**

1. **No explicit declaration needed** - Go automatically knows SimpleError is an error
2. **Any type can be an error** - structs, ints, strings, anything!
3. **Flexible and powerful** - you can create error types with extra data and methods

#### **Let's See More Examples:**

```go
// Even a simple string can be an error!
type StringError string

func (s StringError) Error() string {
    return string(s)
}

// Now a string is an error!
var stringErr error = StringError("something went wrong")
fmt.Printf("String error: %v\n", stringErr)

// Even an integer can be an error!
type CodeError int

func (c CodeError) Error() string {
    return fmt.Sprintf("error code: %d", c)
}

// Now an int is an error!
var codeErr error = CodeError(404)
fmt.Printf("Code error: %v\n", codeErr)
```

### **3. Errors Can Be Stored in Collections (Just Like Other Values)**

```go
// Store errors in a slice
errorList := []error{
    errors.New("error 1"),
    errors.New("error 2"),
    errors.New("error 3"),
}

// Print all errors
for i, err := range errorList {
    fmt.Printf("Error %d: %v\n", i+1, err)
}

// Store errors in a map
errorMap := map[string]error{
    "validation": errors.New("validation failed"),
    "database":   errors.New("database connection failed"),
    "network":    errors.New("network timeout"),
}

// Look up errors by category
if err, exists := errorMap["validation"]; exists {
    fmt.Printf("Validation error: %v\n", err)
}
```

## Section 2: Basic Error Handling - The Foundation

### **Creating Errors (Three Ways)**

```go
// Method 1: errors.New (for simple errors)
err := errors.New("file not found")

// Method 2: fmt.Errorf (for errors with formatting)
filename := "config.txt"
err := fmt.Errorf("failed to read file %s", filename)

// Method 3: Custom error types (we'll learn this in Section 3)
err := &ValidationError{Field: "age", Message: "must be 18+"}
```

### **The Golden Rule: Always Check Errors**

```go
// ❌ WRONG: Don't ignore errors
file, _ := os.Open("file.txt")  // The _ ignores the error!
// What if the file doesn't exist? Your program will crash later!

// ✅ CORRECT: Always check errors
file, err := os.Open("file.txt")
if err != nil {
    // Handle the error
    fmt.Printf("Failed to open file: %v\n", err)
    return err
}
// Now we know the file is open and safe to use
```

### **Real Example: Division Function**

```go
func divideNumbers(a, b float64) (float64, error) {
    if b == 0 {
        // Return an error if dividing by zero
        return 0, errors.New("division by zero")
    }
    // Return the result and nil (no error)
    return a / b, nil
}

// Using the function:
result, err := divideNumbers(10, 2)
if err != nil {
    fmt.Printf("Division failed: %v\n", err)
} else {
    fmt.Printf("Result: %f\n", result)  // Prints: Result: 5.000000
}

result, err = divideNumbers(10, 0)
if err != nil {
    fmt.Printf("Division failed: %v\n", err)  // Prints: Division failed: division by zero
} else {
    fmt.Printf("Result: %f\n", result)
}
```

## Section 3: Error Handling Patterns - Building Good Habits

### **Pattern 1: Early Returns (Happy Path Left-Aligned)**

This is one of the most important patterns in Go error handling. The idea is to handle errors early and return quickly, keeping the "happy path" (success case) aligned to the left.

#### **❌ WRONG: Nested Error Handling**

```go
func processUser(user User) error {
    if user.Name != "" {
        if user.Age >= 18 {
            if user.Email != "" {
                if strings.Contains(user.Email, "@") {
                    // All validations passed
                    fmt.Println("User is valid!")
                    return nil
                } else {
                    return errors.New("invalid email format")
                }
            } else {
                return errors.New("email is required")
            }
        } else {
            return errors.New("user must be at least 18")
        }
    } else {
        return errors.New("name is required")
    }
}
```

#### **✅ CORRECT: Early Returns (Happy Path Left-Aligned)**

```go
func processUser(user User) error {
    // Check name first
    if user.Name == "" {
        return errors.New("name is required")
    }
    
    // Check age
    if user.Age < 18 {
        return errors.New("user must be at least 18")
    }
    
    // Check email
    if user.Email == "" {
        return errors.New("email is required")
    }
    
    // Check email format
    if !strings.Contains(user.Email, "@") {
        return errors.New("invalid email format")
    }
    
    // If we get here, everything is valid
    fmt.Println("User is valid!")
    return nil
}
```

#### **Why Early Returns Are Better:**

1. **Easier to read** - no deep nesting
2. **Easier to debug** - errors happen at the top level
3. **Easier to maintain** - add new validations without changing structure
4. **Happy path is clear** - success case is obvious

### **Pattern 2: Error Propagation - Passing Errors Up the Chain**

When a function calls another function that can fail, you need to decide how to handle the error.

```go
// When a function calls another function that can fail:
func processUser(name string, age int) error {
    // Step 1: Validate the name
    if err := validateName(name); err != nil {
        // Add context to the error and pass it up
        return fmt.Errorf("name validation failed: %w", err)
    }
    
    // Step 2: Validate the age
    if err := validateAge(age); err != nil {
        // Add context to the error and pass it up
        return fmt.Errorf("age validation failed: %w", err)
    }
    
    // If we get here, everything worked
    return nil
}

// Helper functions:
func validateName(name string) error {
    if len(name) < 2 {
        return fmt.Errorf("name too short: %s (minimum 2 characters)", name)
    }
    return nil
}

func validateAge(age int) error {
    if age < 18 {
        return fmt.Errorf("age too young: %d (minimum 18)", age)
    }
    return nil
}

// Using it:
if err := processUser("Alice", 15); err != nil {
    fmt.Printf("Processing failed: %v\n", err)
    // Output: Processing failed: age validation failed: age too young: 15 (minimum 18)
}
```

### **Pattern 3: Multiple Error Handling - Collecting All Problems**

Sometimes you want to collect all the problems and report them together, rather than stopping at the first error.

```go
func validateUser(user User) error {
    var errors []error  // Slice to collect all errors
    
    // Check name
    if err := validateName(user.Name); err != nil {
        errors = append(errors, err)
    }
    
    // Check age
    if err := validateAge(user.Age); err != nil {
        errors = append(errors, err)
    }
    
    // Check email
    if err := validateEmail(user.Email); err != nil {
        errors = append(errors, err)
    }
    
    // If we found any errors, return them all
    if len(errors) > 0 {
        return &AggregatedError{Errors: errors}
    }
    
    // No errors found
    return nil
}

// Using it:
user := User{Name: "A", Age: 15, Email: "invalid-email"}
if err := validateUser(user); err != nil {
    fmt.Printf("Validation failed: %v\n", err)
    // Output: Validation failed: multiple errors (3): name too short: A (minimum 2 characters); age too young: 15 (minimum 18); invalid email format: invalid-email
}
```

## Section 4: Custom Error Types - Now We're Getting Advanced!

This is where your knowledge from Chapter 7 really shines! You can create meaningful error types that carry lots of useful information.

### **Why Create Custom Error Types?**

Think about it: when something goes wrong, you want to know:
- **What** went wrong
- **Where** it went wrong  
- **Why** it went wrong
- **What data** was involved

A simple string like "error occurred" doesn't tell you much. But a custom error type can tell you everything!

### **Let's Build a Validation Error Step by Step**

```go
// Step 1: Define the error struct
type ValidationError struct {
    Field   string      // Which field had the problem
    Message string      // What the problem was
    Value   interface{} // What value caused the problem
}

// Step 2: Make it an error by implementing the Error() method
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s (value: %v)", 
        e.Field, e.Message, e.Value)
}

// Step 3: Add helper methods to make it more useful
func (e *ValidationError) IsValidationError() bool {
    return true
}

func (e *ValidationError) GetField() string {
    return e.Field
}

func (e *ValidationError) GetValue() interface{} {
    return e.Value
}
```

### **Using Your Custom Error Type**

```go
// Create a validation error
err := &ValidationError{
    Field:   "age",
    Message: "must be at least 18",
    Value:   15,
}

// Print the error (calls the Error() method automatically)
fmt.Printf("Error: %v\n", err)
// Output: Error: validation failed for age: must be at least 18 (value: 15)

// Use the helper methods
fmt.Printf("Field with problem: %s\n", err.GetField())  // age
fmt.Printf("Problematic value: %v\n", err.GetValue())   // 15
fmt.Printf("Is this a validation error? %t\n", err.IsValidationError())  // true
```

### **Let's Build More Error Types**

```go
// Database Error - for database problems
type DatabaseError struct {
    Operation string  // What operation failed (SELECT, INSERT, etc.)
    Table     string  // Which table was involved
    Message   string  // What went wrong
    Code      int     // Database error code
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("database error in %s on table %s: %s (code: %d)", 
        e.Operation, e.Table, e.Message, e.Code)
}

func (e *DatabaseError) IsDatabaseError() bool {
    return true
}

func (e *DatabaseError) IsRetryable() bool {
    // Some database errors can be retried
    return e.Code == 1001 || e.Code == 1002
}

// Network Error - for network problems
type NetworkError struct {
    URL     string        // Which URL failed
    Timeout time.Duration // How long we waited
    Message string        // What went wrong
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("network error for %s: %s (timeout: %v)", 
        e.URL, e.Message, e.Timeout)
}

func (e *NetworkError) IsNetworkError() bool {
    return true
}
```

### **How to Check Error Types**

```go
// Method 1: Type assertion (basic way)
if validationErr, ok := err.(*ValidationError); ok {
    fmt.Printf("Validation error on field: %s\n", validationErr.Field)
}

// Method 2: Using errors.As (Go 1.13+ way - recommended!)
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    fmt.Printf("Validation error on field: %s\n", validationErr.Field)
}

// Let's create helper functions for each error type:
func IsValidationError(err error) bool {
    var validationErr *ValidationError
    return errors.As(err, &validationErr)
}

func IsDatabaseError(err error) bool {
    var dbErr *DatabaseError
    return errors.As(err, &dbErr)
}

func IsNetworkError(err error) bool {
    var networkErr *NetworkError
    return errors.As(err, &networkErr)
}

// Now you can use them like this:
if err := processUser(user); err != nil {
    switch {
    case IsValidationError(err):
        fmt.Println("Please fix the validation issues")
    case IsDatabaseError(err):
        fmt.Println("Database operation failed")
    case IsNetworkError(err):
        fmt.Println("Network problem occurred")
    default:
        fmt.Println("Unknown error occurred")
    }
}
```

## Section 5: Advanced Error Patterns - The Cool Stuff!

Now we're getting into the really powerful features! These patterns help you build professional, production-ready applications.

### **Error Wrapping - Adding Context Without Losing Information**

Think of error wrapping like putting a letter in an envelope, then putting that envelope in a bigger envelope, and so on. Each layer adds more context, but you can still get back to the original message.

#### **Why Wrap Errors?**

Imagine this scenario:
1. **Database level**: "connection timeout"
2. **Service level**: "failed to get user data"
3. **API level**: "user profile update failed"

Without wrapping, you'd lose the original database error. With wrapping, you get the full story!

```go
// Let's build an error chain step by step:

// Step 1: Original error (database level)
originalErr := errors.New("connection timeout")
fmt.Printf("Original error: %v\n", originalErr)
// Output: Original error: connection timeout

// Step 2: Wrap it with database context
dbErr := fmt.Errorf("database query failed: %w", originalErr)
fmt.Printf("Database error: %v\n", dbErr)
// Output: Database error: database query failed: connection timeout

// Step 3: Wrap it with service context
serviceErr := fmt.Errorf("failed to get user data: %w", dbErr)
fmt.Printf("Service error: %v\n", serviceErr)
// Output: Service error: failed to get user data: database query failed: connection timeout

// Step 4: Wrap it with API context
apiErr := fmt.Errorf("user profile update failed: %w", serviceErr)
fmt.Printf("API error: %v\n", apiErr)
// Output: API error: user profile update failed: failed to get user data: database query failed: connection timeout
```

#### **Unwrapping Errors - Going Back Through the Chain**

```go
// Now let's go backwards through the chain:
currentErr := apiErr
level := 1

fmt.Println("Error chain (from most recent to original):")
for currentErr != nil {
    fmt.Printf("Level %d: %v\n", level, currentErr)
    currentErr = errors.Unwrap(currentErr)
    level++
}

// Output:
// Level 1: API error: user profile update failed: failed to get user data: database query failed: connection timeout
// Level 2: Service error: failed to get user data: database query failed: connection timeout
// Level 3: Database error: database query failed: connection timeout
// Level 4: Original error: connection timeout
```

#### **Checking Error Types in Wrapped Errors**

```go
// You can check if a wrapped error contains a specific error type
if errors.Is(apiErr, originalErr) {
    fmt.Println("Yes! The API error contains the original connection timeout error")
}

// You can also extract specific error types from the chain
var dbError *DatabaseError
if errors.As(apiErr, &dbError) {
    fmt.Printf("Found database error in chain: %v\n", dbError)
}
```

### **Sentinel Errors - Predefined Error Values**

Sentinel errors are predefined error values that represent specific error conditions. They're useful for checking for known error types.

```go
// Common sentinel errors
var (
    ErrNotFound = errors.New("not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrTimeout = errors.New("timeout")
    ErrPermissionDenied = errors.New("permission denied")
)

// Using sentinel errors
func findUser(id string) (*User, error) {
    if id == "" {
        return nil, ErrInvalidInput
    }
    
    // Simulate user lookup
    if id == "nonexistent" {
        return nil, ErrNotFound
    }
    
    // Simulate permission check
    if id == "admin" {
        return nil, ErrPermissionDenied
    }
    
    return &User{ID: id, Name: "User " + id}, nil
}

// Checking for sentinel errors
user, err := findUser("nonexistent")
if err != nil {
    switch {
    case errors.Is(err, ErrNotFound):
        fmt.Println("User not found")
    case errors.Is(err, ErrInvalidInput):
        fmt.Println("Invalid user ID")
    case errors.Is(err, ErrPermissionDenied):
        fmt.Println("Access denied")
    default:
        fmt.Printf("Unknown error: %v\n", err)
    }
}
```

### **Error Aggregation - Collecting Multiple Errors**

Sometimes you want to collect all the problems and report them together, rather than stopping at the first error.

#### **Why Aggregate Errors?**

Think about a form submission:
- User enters invalid name (too short)
- User enters invalid age (too young)
- User enters invalid email (wrong format)

Instead of telling the user "name is too short" and then after they fix that, "age is too young", you can tell them all the problems at once!

```go
// AggregatedError collects multiple errors
type AggregatedError struct {
    Errors []error
}

func (ae *AggregatedError) Error() string {
    if len(ae.Errors) == 0 {
        return "no errors"
    }
    
    var messages []string
    for _, err := range ae.Errors {
        messages = append(messages, err.Error())
    }
    
    return fmt.Sprintf("multiple errors (%d): %s", 
        len(ae.Errors), strings.Join(messages, "; "))
}

// Helper methods
func (ae *AggregatedError) ErrorCount() int {
    return len(ae.Errors)
}

func (ae *AggregatedError) HasValidationErrors() bool {
    for _, err := range ae.Errors {
        if IsValidationError(err) {
            return true
        }
    }
    return false
}

func (ae *AggregatedError) HasDatabaseErrors() bool {
    for _, err := range ae.Errors {
        if IsDatabaseError(err) {
            return true
        }
    }
    return false
}

// Function to collect all validation errors
func validateUserComprehensive(user User) error {
    var errors []error
    
    // Check name
    if err := validateName(user.Name); err != nil {
        errors = append(errors, err)
    }
    
    // Check age
    if err := validateAge(user.Age); err != nil {
        errors = append(errors, err)
    }
    
    // Check email
    if err := validateEmail(user.Email); err != nil {
        errors = append(errors, err)
    }
    
    // If we found any errors, return them all
    if len(errors) > 0 {
        return &AggregatedError{Errors: errors}
    }
    
    return nil
}

// Using it:
user := User{Name: "A", Age: 15, Email: "invalid-email"}
if err := validateUserComprehensive(user); err != nil {
    if aggErr, ok := err.(*AggregatedError); ok {
        fmt.Printf("Found %d validation errors:\n", aggErr.ErrorCount())
        for i, validationErr := range aggErr.Errors {
            fmt.Printf("  %d. %v\n", i+1, validationErr)
        }
        
        if aggErr.HasValidationErrors() {
            fmt.Println("Please fix the validation issues above")
        }
    }
}
```

## Section 6: Error Handling Best Practices

### **1. Distinguishing Between Expected vs Unexpected Errors**

Not all errors are the same. Some are expected (like validation failures), while others are unexpected (like system failures).

```go
// Expected errors - these are part of normal operation
func validateUserAge(age int) error {
    if age < 18 {
        return &ValidationError{
            Field:   "age",
            Message: "must be at least 18",
            Value:   age,
        }
    }
    return nil
}

// Unexpected errors - these indicate system problems
func connectToDatabase() error {
    // Simulate connection failure
    return &DatabaseError{
        Operation: "CONNECT",
        Table:     "N/A",
        Message:   "connection refused",
        Code:      1001,
    }
}

// Handle them differently
func processUser(user User) error {
    // Expected error - return it to the user
    if err := validateUserAge(user.Age); err != nil {
        return err  // User needs to fix this
    }
    
    // Unexpected error - log it and return generic message
    if err := connectToDatabase(); err != nil {
        // Log the full error for debugging
        log.Printf("Database connection failed: %v", err)
        // Return generic message to user
        return errors.New("service temporarily unavailable")
    }
    
    return nil
}
```

### **2. Logging vs Returning Errors - Don't Mix Responsibilities**

This is a crucial concept: **logging and error handling are separate concerns**.

```go
// ❌ WRONG: Mixing logging and error handling
func processUser(user User) error {
    if user.Age < 18 {
        log.Printf("User %s is too young: %d", user.Name, user.Age)
        return fmt.Errorf("user %s is too young: %d", user.Name, user.Age)
    }
    return nil
}

// ✅ CORRECT: Separate logging from error handling
func processUser(user User) error {
    if user.Age < 18 {
        // Log for debugging/monitoring
        log.Printf("Validation failed: user %s age %d is below minimum", user.Name, user.Age)
        // Return clean error for user
        return &ValidationError{
            Field:   "age",
            Message: "must be at least 18",
            Value:   user.Age,
        }
    }
    return nil
}

// The calling code decides how to handle the error
func handleUserSubmission(user User) {
    if err := processUser(user); err != nil {
        // Log the error for debugging
        log.Printf("User submission failed: %v", err)
        
        // Show appropriate message to user
        if IsValidationError(err) {
            fmt.Println("Please fix the validation issues and try again")
        } else {
            fmt.Println("An unexpected error occurred. Please try again later.")
        }
    } else {
        fmt.Println("User submitted successfully!")
    }
}
```

### **3. Retry Patterns - Don't Give Up Too Easily**

Some errors are temporary and can be fixed by trying again. Network timeouts, database connection issues, and temporary server problems often fall into this category.

#### **Simple Retry Pattern**

```go
// Retry an operation multiple times
func retryOperation(maxAttempts int, operation func() error) bool {
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        fmt.Printf("Attempt %d: ", attempt)
        
        if err := operation(); err != nil {
            fmt.Printf("Failed: %v\n", err)
            if attempt < maxAttempts {
                fmt.Println("  Retrying...")
            }
        } else {
            fmt.Println("Succeeded!")
            return true
        }
    }
    return false
}

// Simulate an unreliable operation (like a network call)
func simulateUnreliableOperation() error {
    // Simulate 70% failure rate
    if time.Now().UnixNano()%10 < 7 {
        return errors.New("operation failed")
    }
    return nil
}

// Use it
success := retryOperation(3, simulateUnreliableOperation)
if success {
    fmt.Println("Operation succeeded after retries")
} else {
    fmt.Println("Operation failed after all retries")
}
```

#### **Smart Retry with Error Type Checking**

```go
// Only retry certain types of errors
func smartRetry(operation func() error) error {
    maxAttempts := 3
    
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        err := operation()
        if err == nil {
            return nil  // Success!
        }
        
        // Check if this error is retryable
        if isRetryableError(err) {
            if attempt < maxAttempts {
                fmt.Printf("Retryable error, attempting %d of %d...\n", attempt+1, maxAttempts)
                time.Sleep(time.Duration(attempt) * time.Second)  // Wait longer each time
                continue
            }
        }
        
        // Non-retryable error or max attempts reached
        return err
    }
    
    return errors.New("max retry attempts reached")
}

func isRetryableError(err error) bool {
    // Network errors are usually retryable
    if IsNetworkError(err) {
        return true
    }
    
    // Some database errors are retryable
    if IsDatabaseError(err) {
        if dbErr, ok := err.(*DatabaseError); ok {
            return dbErr.IsRetryable()
        }
    }
    
    return false
}
```

### **4. Fallback Patterns - Plan B, C, and D**

When your primary method fails, have backup plans ready!

#### **Data Source Fallback**

```go
// Try multiple data sources with fallbacks
func getDataWithFallback() (string, error) {
    // Try primary source first
    if data, err := getDataFromPrimary(); err == nil {
        return data, nil
    }
    
    // Primary failed, try secondary
    if data, err := getDataFromSecondary(); err == nil {
        return data, nil
    }
    
    // Secondary failed, try cache
    if data, err := getDataFromCache(); err == nil {
        return data, nil
    }
    
    // All sources failed
    return "", errors.New("all data sources failed")
}

// Simulate different data sources
func getDataFromPrimary() (string, error) {
    return "", errors.New("primary source unavailable")
}

func getDataFromSecondary() (string, error) {
    return "", errors.New("secondary source unavailable")
}

func getDataFromCache() (string, error) {
    return "cached data", nil  // Cache works!
}

// Using it:
data, err := getDataWithFallback()
if err != nil {
    fmt.Printf("All data sources failed: %v\n", err)
} else {
    fmt.Printf("Data retrieved: %s\n", data)
}
```

## Section 7: Panics vs Errors - When to Use Each

### **What are Panics?**

Panics are Go's way of handling **unrecoverable situations**. They're like emergency stops that crash your program.

```go
// Panic example
func divideByZero(a, b int) int {
    if b == 0 {
        panic("division by zero")  // This will crash the program!
    }
    return a / b
}

// Using it (this will crash)
// result := divideByZero(10, 0)  // PANIC!
```

### **When to Use Panics vs Errors**

#### **❌ DON'T Use Panics For:**
- **Expected failures** (file not found, validation errors)
- **User input problems** (invalid email, wrong password)
- **Network timeouts** (connection issues, slow responses)
- **Business logic failures** (insufficient funds, access denied)

#### **✅ DO Use Panics For:**
- **Programming errors** (accessing nil pointer, index out of bounds)
- **Unrecoverable situations** (out of memory, corrupted data)
- **Initialization failures** (can't start the application)
- **Assertion failures** (invariant violations)

### **Recover for Handling Panics Safely (Rare Cases)**

The `recover` function can catch panics and prevent your program from crashing. However, this should be used very sparingly.

```go
// Example of using recover (advanced topic)
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            // Convert panic to error
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    
    result = a / b
    return result, nil
}

// Using it
result, err := safeDivide(10, 0)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Result: %d\n", result)
}
```

### **Better Approach: Use Errors Instead of Panics**

```go
// ✅ BETTER: Use errors for expected failures
func divideSafely(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Using it
result, err := divideSafely(10, 0)
if err != nil {
    fmt.Printf("Division failed: %v\n", err)
} else {
    fmt.Printf("Result: %d\n", result)
}
```

## How to Run Your Program

1. Open your terminal
2. Go to the error-handling folder: `cd 10-error-handling`
3. Run the program: `go run main.go`

## Try It Yourself!

### Exercise 1: Create a File Error Type
Create a custom error type for file operations:

```go
type FileError struct {
    Filename string
    Operation string
    Message   string
    Code      int
}

func (e *FileError) Error() string {
    return fmt.Sprintf("file error in %s for %s: %s (code: %d)", 
        e.Operation, e.Filename, e.Message, e.Code)
}

func (e *FileError) IsPermissionError() bool {
    return e.Code == 13
}

func (e *FileError) IsNotFoundError() bool {
    return e.Code == 2
}
```

### Exercise 2: Build a Validation System
Create a comprehensive validation system:

```go
type ValidationResult struct {
    IsValid bool
    Errors  []*ValidationError
}

func (vr *ValidationResult) AddError(field, message string, value interface{}) {
    vr.Errors = append(vr.Errors, &ValidationError{
        Field:   field,
        Message: message,
        Value:   value,
    })
    vr.IsValid = false
}

func (vr *ValidationResult) GetErrorsForField(field string) []*ValidationError {
    var fieldErrors []*ValidationError
    for _, err := range vr.Errors {
        if err.Field == field {
            fieldErrors = append(fieldErrors, err)
        }
    }
    return fieldErrors
}

// Use it
func validateUser(user User) *ValidationResult {
    result := &ValidationResult{IsValid: true}
    
    if len(user.Name) < 2 {
        result.AddError("name", "must be at least 2 characters", user.Name)
    }
    
    if user.Age < 18 {
        result.AddError("age", "must be at least 18", user.Age)
    }
    
    if !strings.Contains(user.Email, "@") {
        result.AddError("email", "invalid email format", user.Email)
    }
    
    return result
}
```

### Exercise 3: Implement Error Recovery
Create a system that can recover from different types of errors:

```go
type ErrorRecovery struct {
    MaxRetries int
    Backoff    time.Duration
}

func (er *ErrorRecovery) Recover(operation func() error) error {
    for attempt := 1; attempt <= er.MaxRetries; attempt++ {
        if err := operation(); err != nil {
            if attempt == er.MaxRetries {
                return fmt.Errorf("operation failed after %d attempts: %w", er.MaxRetries, err)
            }
            
            // Wait before retrying
            time.Sleep(er.Backoff * time.Duration(attempt))
            continue
        }
        return nil
    }
    return errors.New("unexpected error in recovery")
}

// Use it
recovery := &ErrorRecovery{MaxRetries: 3, Backoff: time.Second}
err := recovery.Recover(func() error {
    return simulateUnreliableOperation()
})
```

## Common Mistakes to Avoid

### 1. **Ignoring Errors**
```go
// ❌ WRONG: Ignoring errors
file, _ := os.Open("file.txt")  // Error ignored!
defer file.Close()

// ✅ CORRECT: Handle errors
file, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer file.Close()
```

### 2. **Not Wrapping Errors**
```go
// ❌ WRONG: Losing context
if err := validateUser(user); err != nil {
    return err  // Original error without context
}

// ✅ CORRECT: Wrap with context
if err := validateUser(user); err != nil {
    return fmt.Errorf("user validation failed: %w", err)
}
```

### 3. **Checking Error Types Incorrectly**
```go
// ❌ WRONG: Type assertion without checking
if validationErr, ok := err.(*ValidationError); ok {
    // This might panic if err is nil
}

// ✅ CORRECT: Use errors.As
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    // Safe type checking
}
```

### 4. **Using Panics for Expected Errors**
```go
// ❌ WRONG: Panic for expected failure
func validateAge(age int) {
    if age < 18 {
        panic("age too young")  // Don't do this!
    }
}

// ✅ CORRECT: Return error for expected failure
func validateAge(age int) error {
    if age < 18 {
        return errors.New("age too young")
    }
    return nil
}
```

### 5. **Mixing Logging and Error Handling**
```go
// ❌ WRONG: Logging in validation function
func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        log.Printf("Invalid email: %s", email)  // Don't log here!
        return errors.New("invalid email format")
    }
    return nil
}

// ✅ CORRECT: Separate logging from validation
func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return errors.New("invalid email format")
    }
    return nil
}

// Log in the calling code
if err := validateEmail(user.Email); err != nil {
    log.Printf("Email validation failed for user %s: %v", user.Name, err)
    return err
}
```

## Key Takeaways

1. **Errors are values** - Store, pass, and return them like any other value
2. **Always check errors** - Never ignore them, always handle them explicitly
3. **Use early returns** - Keep the happy path left-aligned
4. **Wrap errors with context** - Use `fmt.Errorf` with `%w` to preserve original errors
5. **Create meaningful error types** - Use structs and methods for rich error information
6. **Use error checking functions** - `errors.Is`, `errors.As`, `errors.Unwrap`
7. **Separate concerns** - Don't mix logging with error handling
8. **Use panics sparingly** - Only for unrecoverable situations
9. **Design for failure** - Build systems that can recover from errors gracefully
10. **Return errors, don't panic** - Panics are for unrecoverable situations

## Next Steps

After this chapter, you'll be ready for:
- **Goroutines and channels** (Go's concurrency features)
- **Testing** (writing tests for your error handling)
- **Packages and modules** (organizing your code)
- **Web development** (handling HTTP errors)
- **Production deployment** (logging and monitoring errors)

---

**Excellent work! You now understand Go's powerful error handling system! 🎉**

Error handling is what makes Go programs robust and production-ready. You've learned how to create meaningful error types, handle failures gracefully, and build systems that can recover from problems.

Practice creating custom error types, implementing error handling patterns, and building resilient applications. This knowledge will make you a much better Go programmer and help you write code that others can trust in production! 🚀 