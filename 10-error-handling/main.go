package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// User represents a user in the system
type User struct {
	ID   string
	Name string
	Age  int
	Email string
}

// ValidationError represents validation failures
type ValidationError struct {
	Field   string      // Which field had the problem
	Message string      // What the problem was
	Value   interface{} // What value caused the problem
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s (value: %v)", 
		e.Field, e.Message, e.Value)
}

func (e *ValidationError) IsValidationError() bool {
	return true
}

func (e *ValidationError) GetField() string {
	return e.Field
}

func (e *ValidationError) GetValue() interface{} {
	return e.Value
}

// DatabaseError represents database operation failures
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

// NetworkError represents network operation failures
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

// Helper functions for error type checking
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

// Simple error types for demonstration
type SimpleError struct {
	Message string
}

func (e SimpleError) Error() string {
	return e.Message
}

type StringError string

func (s StringError) Error() string {
	return string(s)
}

type CodeError int

func (c CodeError) Error() string {
	return fmt.Sprintf("error code: %d", c)
}

func main() {
	fmt.Println("🐹 Go Error Handling - Chapter 10")
	fmt.Println("==================================")
	
	// Section 1: Go's Error Philosophy
	section1_ErrorPhilosophy()
	
	// Section 2: Basic Error Handling
	section2_BasicErrorHandling()
	
	// Section 3: Error Handling Patterns
	section3_ErrorHandlingPatterns()
	
	// Section 4: Custom Error Types
	section4_CustomErrorTypes()
	
	// Section 5: Advanced Error Patterns
	section5_AdvancedErrorPatterns()
	
	// Section 6: Error Handling Best Practices
	section6_ErrorHandlingBestPractices()
	
	// Section 7: Panics vs Errors
	section7_PanicsVsErrors()
	
	fmt.Println("\n🎉 Error handling chapter completed!")
}

// Section 1: Go's Error Philosophy - Let's Break It Down
func section1_ErrorPhilosophy() {
	fmt.Println("\n📚 Section 1: Go's Error Philosophy")
	fmt.Println("------------------------------------")
	
	// 1. Errors are Just Values (Like Strings or Numbers)
	fmt.Println("\n1. Errors are Just Values:")
	var err error
	fmt.Printf("Error: %v\n", err)  // <nil> (nil means "no error")
	
	err = errors.New("something went wrong")
	fmt.Printf("Error: %v\n", err)  // something went wrong
	
	if err != nil {
		fmt.Println("There is an error!")
	} else {
		fmt.Println("No error!")
	}
	
	// 2. What is an Interface? (Important Concept!)
	fmt.Println("\n2. Understanding the error Interface:")
	
	// Simple error type
	myError := SimpleError{Message: "my custom error"}
	fmt.Printf("My error: %v\n", myError)
	
	// String error type
	stringErr := StringError("something went wrong")
	fmt.Printf("String error: %v\n", stringErr)
	
	// Code error type
	codeErr := CodeError(404)
	fmt.Printf("Code error: %v\n", codeErr)
	
	// 3. Errors Can Be Stored in Collections
	fmt.Println("\n3. Errors in Collections:")
	
	errorList := []error{
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	}
	
	for i, err := range errorList {
		fmt.Printf("Error %d: %v\n", i+1, err)
	}
	
	errorMap := map[string]error{
		"validation": errors.New("validation failed"),
		"database":   errors.New("database connection failed"),
		"network":    errors.New("network timeout"),
	}
	
	if err, exists := errorMap["validation"]; exists {
		fmt.Printf("Validation error: %v\n", err)
	}
}

// Section 2: Basic Error Handling - The Foundation
func section2_BasicErrorHandling() {
	fmt.Println("\n🔧 Section 2: Basic Error Handling")
	fmt.Println("-----------------------------------")
	
	// Creating Errors (Three Ways)
	fmt.Println("\n1. Creating Errors:")
	
	// Method 1: errors.New
	err1 := errors.New("file not found")
	fmt.Printf("Method 1: %v\n", err1)
	
	// Method 2: fmt.Errorf
	filename := "config.txt"
	err2 := fmt.Errorf("failed to read file %s", filename)
	fmt.Printf("Method 2: %v\n", err2)
	
	// Method 3: Custom error types
	err3 := &ValidationError{Field: "age", Message: "must be 18+", Value: 15}
	fmt.Printf("Method 3: %v\n", err3)
	
	// The Golden Rule: Always Check Errors
	fmt.Println("\n2. Always Check Errors:")
	
	// Real Example: Division Function
	result, err := divideNumbers(10, 2)
	if err != nil {
		fmt.Printf("Division failed: %v\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
	
	result, err = divideNumbers(10, 0)
	if err != nil {
		fmt.Printf("Division failed: %v\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
}

// Division function that demonstrates error handling
func divideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Section 3: Error Handling Patterns - Building Good Habits
func section3_ErrorHandlingPatterns() {
	fmt.Println("\n🎯 Section 3: Error Handling Patterns")
	fmt.Println("--------------------------------------")
	
	// Pattern 1: Early Returns (Happy Path Left-Aligned)
	fmt.Println("\n1. Early Returns (Happy Path Left-Aligned):")
	
	user := User{Name: "Alice", Age: 15, Email: "alice@example.com"}
	if err := processUser(user); err != nil {
		fmt.Printf("User processing failed: %v\n", err)
	}
	
	// Pattern 2: Error Propagation - Passing Errors Up the Chain
	fmt.Println("\n2. Error Propagation:")
	
	if err := processUserWithContext(user); err != nil {
		fmt.Printf("Processing failed: %v\n", err)
	}
	
	// Pattern 3: Multiple Error Handling - Collecting All Problems
	fmt.Println("\n3. Multiple Error Handling:")
	
	if err := validateUserComprehensive(user); err != nil {
		if aggErr, ok := err.(*AggregatedError); ok {
			fmt.Printf("Found %d validation errors:\n", aggErr.ErrorCount())
			for i, validationErr := range aggErr.Errors {
				fmt.Printf("  %d. %v\n", i+1, validationErr)
			}
		}
	}
}

// Pattern 1: Early Returns (Happy Path Left-Aligned)
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

// Pattern 2: Error Propagation
func processUserWithContext(user User) error {
	// Step 1: Validate the name
	if err := validateName(user.Name); err != nil {
		return fmt.Errorf("name validation failed: %w", err)
	}
	
	// Step 2: Validate the age
	if err := validateAge(user.Age); err != nil {
		return fmt.Errorf("age validation failed: %w", err)
	}
	
	return nil
}

// Helper functions for validation
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

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email format: %s", email)
	}
	return nil
}

// Pattern 3: Multiple Error Handling
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

// Section 4: Custom Error Types - Now We're Getting Advanced!
func section4_CustomErrorTypes() {
	fmt.Println("\n🚀 Section 4: Custom Error Types")
	fmt.Println("---------------------------------")
	
	// Why Create Custom Error Types?
	fmt.Println("\n1. Custom Error Types with Rich Information:")
	
	validationErr := &ValidationError{
		Field:   "age",
		Message: "must be at least 18",
		Value:   15,
	}
	
	fmt.Printf("Error: %v\n", validationErr)
	fmt.Printf("Field with problem: %s\n", validationErr.GetField())
	fmt.Printf("Problematic value: %v\n", validationErr.GetValue())
	fmt.Printf("Is this a validation error? %t\n", validationErr.IsValidationError())
	
	// More Error Types
	fmt.Println("\n2. Database and Network Errors:")
	
	dbErr := &DatabaseError{
		Operation: "SELECT",
		Table:     "users",
		Message:   "connection timeout",
		Code:      1001,
	}
	
	fmt.Printf("Database error: %v\n", dbErr)
	fmt.Printf("Is retryable? %t\n", dbErr.IsRetryable())
	
	networkErr := &NetworkError{
		URL:     "https://api.example.com",
		Timeout: 30 * time.Second,
		Message: "connection refused",
	}
	
	fmt.Printf("Network error: %v\n", networkErr)
	
	// How to Check Error Types
	fmt.Println("\n3. Checking Error Types:")
	
	// Method 1: Type assertion (this will fail since dbErr is a DatabaseError)
	fmt.Println("Type assertion example: dbErr is a DatabaseError, not a ValidationError")
	
	// Let's try with a validation error instead
	valErr := &ValidationError{Field: "age", Message: "too young", Value: 15}
	fmt.Printf("Validation error on field: %s\n", valErr.Field)
	
	// Method 2: Using errors.As (recommended!)
	var validationErr2 *ValidationError
	if errors.As(dbErr, &validationErr2) {
		fmt.Printf("Validation error on field: %s\n", validationErr2.Field)
	} else {
		fmt.Println("Not a validation error")
	}
	
	// Using helper functions
	fmt.Printf("Is validation error? %t\n", IsValidationError(dbErr))
	fmt.Printf("Is database error? %t\n", IsDatabaseError(dbErr))
	fmt.Printf("Is network error? %t\n", IsNetworkError(dbErr))
}

// Section 5: Advanced Error Patterns - The Cool Stuff!
func section5_AdvancedErrorPatterns() {
	fmt.Println("\n🌟 Section 5: Advanced Error Patterns")
	fmt.Println("--------------------------------------")
	
	// Error Wrapping - Adding Context Without Losing Information
	fmt.Println("\n1. Error Wrapping:")
	
	// Build an error chain step by step
	originalErr := errors.New("connection timeout")
	fmt.Printf("Original error: %v\n", originalErr)
	
	dbErr := fmt.Errorf("database query failed: %w", originalErr)
	fmt.Printf("Database error: %v\n", dbErr)
	
	serviceErr := fmt.Errorf("failed to get user data: %w", dbErr)
	fmt.Printf("Service error: %v\n", serviceErr)
	
	apiErr := fmt.Errorf("user profile update failed: %w", serviceErr)
	fmt.Printf("API error: %v\n", apiErr)
	
	// Unwrapping Errors - Going Back Through the Chain
	fmt.Println("\n2. Unwrapping Errors:")
	
	currentErr := apiErr
	level := 1
	
	fmt.Println("Error chain (from most recent to original):")
	for currentErr != nil {
		fmt.Printf("Level %d: %v\n", level, currentErr)
		currentErr = errors.Unwrap(currentErr)
		level++
	}
	
	// Checking Error Types in Wrapped Errors
	fmt.Println("\n3. Checking Wrapped Errors:")
	
	if errors.Is(apiErr, originalErr) {
		fmt.Println("Yes! The API error contains the original connection timeout error")
	}
	
	var dbError *DatabaseError
	if errors.As(apiErr, &dbError) {
		fmt.Printf("Found database error in chain: %v\n", dbError)
	} else {
		fmt.Println("No database error found in chain")
	}
	
	// Sentinel Errors - Predefined Error Values
	fmt.Println("\n4. Sentinel Errors:")
	
	// Common sentinel errors
	var (
		ErrNotFound = errors.New("not found")
		ErrInvalidInput = errors.New("invalid input")
		ErrPermissionDenied = errors.New("permission denied")
	)
	
	// Using sentinel errors
	_, err := findUser("nonexistent")
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
	
	// Error Aggregation - Collecting Multiple Errors
	fmt.Println("\n5. Error Aggregation:")
	
	user2 := User{Name: "A", Age: 15, Email: "invalid-email"}
	if err := validateUserComprehensive(user2); err != nil {
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
}

// Function to demonstrate sentinel errors
func findUser(id string) (*User, error) {
	if id == "" {
		return nil, errors.New("invalid input")
	}
	
	// Simulate user lookup
	if id == "nonexistent" {
		return nil, errors.New("not found")
	}
	
	// Simulate permission check
	if id == "admin" {
		return nil, errors.New("permission denied")
	}
	
	return &User{ID: id, Name: "User " + id}, nil
}

// Section 6: Error Handling Best Practices
func section6_ErrorHandlingBestPractices() {
	fmt.Println("\n📋 Section 6: Error Handling Best Practices")
	fmt.Println("---------------------------------------------")
	
	// 1. Distinguishing Between Expected vs Unexpected Errors
	fmt.Println("\n1. Expected vs Unexpected Errors:")
	
	user := User{Name: "Alice", Age: 15, Email: "alice@example.com"}
	if err := processUserWithErrorTypes(user); err != nil {
		fmt.Printf("Processing failed: %v\n", err)
	}
	
	// 2. Logging vs Returning Errors - Don't Mix Responsibilities
	fmt.Println("\n2. Separating Logging from Error Handling:")
	
	handleUserSubmission(user)
	
	// 3. Retry Patterns - Don't Give Up Too Easily
	fmt.Println("\n3. Retry Patterns:")
	
	success := retryOperation(3, simulateUnreliableOperation)
	if success {
		fmt.Println("Operation succeeded after retries")
	} else {
		fmt.Println("Operation failed after all retries")
	}
	
	// Smart retry with error type checking
	err := smartRetry(simulateUnreliableOperation)
	if err != nil {
		fmt.Printf("Smart retry failed: %v\n", err)
	} else {
		fmt.Println("Smart retry succeeded!")
	}
	
	// 4. Fallback Patterns - Plan B, C, and D
	fmt.Println("\n4. Fallback Patterns:")
	
	data, err := getDataWithFallback()
	if err != nil {
		fmt.Printf("All data sources failed: %v\n", err)
	} else {
		fmt.Printf("Data retrieved: %s\n", data)
	}
}

// Expected vs Unexpected Errors
func processUserWithErrorTypes(user User) error {
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

func connectToDatabase() error {
	// Simulate connection failure
	return &DatabaseError{
		Operation: "CONNECT",
		Table:     "N/A",
		Message:   "connection refused",
		Code:      1001,
	}
}

// Separating Logging from Error Handling
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

// Retry Patterns
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

func simulateUnreliableOperation() error {
	// Simulate 70% failure rate
	if time.Now().UnixNano()%10 < 7 {
		return errors.New("operation failed")
	}
	return nil
}

// Smart retry with error type checking
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

// Fallback Patterns
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

func getDataFromPrimary() (string, error) {
	return "", errors.New("primary source unavailable")
}

func getDataFromSecondary() (string, error) {
	return "", errors.New("secondary source unavailable")
}

func getDataFromCache() (string, error) {
	return "cached data", nil  // Cache works!
}

// Section 7: Panics vs Errors - When to Use Each
func section7_PanicsVsErrors() {
	fmt.Println("\n⚠️  Section 7: Panics vs Errors")
	fmt.Println("--------------------------------")
	
	// What are Panics?
	fmt.Println("\n1. Understanding Panics:")
	
	// Panic example (commented out to avoid crashing)
	// result := divideByZero(10, 0)  // This would PANIC!
	fmt.Println("Panic example: divideByZero(10, 0) would crash the program")
	
	// When to Use Panics vs Errors
	fmt.Println("\n2. Panics vs Errors:")
	
	// Better approach: Use errors for expected failures
	result, err := divideSafely(10, 0)
	if err != nil {
		fmt.Printf("Division failed: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	
	// Recover for Handling Panics Safely (Rare Cases)
	fmt.Println("\n3. Using Recover (Advanced Topic):")
	
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}

// Panic example (don't use this in real code!)
func divideByZero(a, b int) int {
	if b == 0 {
		panic("division by zero")  // This will crash the program!
	}
	return a / b
}

// Better approach: Use errors for expected failures
func divideSafely(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

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