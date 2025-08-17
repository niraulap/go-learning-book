package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("🐹 Go Interfaces - Chapter 8 🐹")
	fmt.Println("=================================")

	// ============================================================================
	// SECTION 1: What Are Interfaces?
	// ============================================================================
	section1_WhatAreInterfaces()

	// ============================================================================
	// SECTION 2: Creating and Using Interfaces
	// ============================================================================
	section2_CreatingAndUsingInterfaces()

	// ============================================================================
	// SECTION 3: Interface Composition and Embedding
	// ============================================================================
	section3_InterfaceComposition()

	// ============================================================================
	// SECTION 4: The Empty Interface and Type Assertions
	// ============================================================================
	section4_EmptyInterfaceAndTypeAssertions()

	// ============================================================================
	// SECTION 5: Real-World Interface Examples
	// ============================================================================
	section5_RealWorldExamples()

	fmt.Println("\n🎉 Chapter 8 Complete! You understand Go interfaces!")
}

// ============================================================================
// SECTION 1: What Are Interfaces?
// ============================================================================
func section1_WhatAreInterfaces() {
	fmt.Println("\n📚 SECTION 1: What Are Interfaces?")
	fmt.Println("-----------------------------------")

	// Think of interfaces like contracts
	fmt.Println("Interfaces are like contracts that define what methods a type must have.")
	fmt.Println("Any type that implements all the required methods automatically satisfies the interface!")

	// The error interface (from Chapter 7)
	fmt.Println("\nThe error interface (from Chapter 7):")
	fmt.Println("type error interface { Error() string }")
	fmt.Println("Any type with an Error() method is automatically an error!")

	// Let's see this in action
	fmt.Println("\nDemonstrating the error interface:")
	
	// StringError implements the error interface
	stringErr := StringError("something went wrong")
	fmt.Printf("StringError: %v\n", stringErr)
	
	// CodeError implements the error interface
	codeErr := CodeError(404)
	fmt.Printf("CodeError: %v\n", codeErr)
	
	// PersonError implements the error interface
	personErr := PersonError{Name: "Alice", Message: "validation failed"}
	fmt.Printf("PersonError: %v\n", personErr)

	// All of these can be used wherever an error is expected
	fmt.Println("\nAll of these can be used wherever an error is expected!")
}

// ============================================================================
// SECTION 2: Creating and Using Interfaces
// ============================================================================
func section2_CreatingAndUsingInterfaces() {
	fmt.Println("\n📚 SECTION 2: Creating and Using Interfaces")
	fmt.Println("---------------------------------------------")

	// Basic interface definition
	fmt.Println("Basic interface definition:")
	fmt.Println("type Shape interface { Area() float64 }")
	
	// Any type with an Area() method is a Shape
	fmt.Println("\nAny type with an Area() method is a Shape:")
	
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}
	triangle := Triangle{Base: 3.0, Height: 4.0}
	
	// These all implement the Shape interface
	shapes := []Shape{circle, rectangle, triangle}
	
	fmt.Println("Calculating areas for different shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: Area = %.2f\n", i+1, shape.Area())
	}

	// Interface with multiple methods
	fmt.Println("\nInterface with multiple methods:")
	fmt.Println("type Animal interface { Speak() string; Move() string }")
	
	dog := Dog{Name: "Buddy", Breed: "Golden Retriever"}
	cat := Cat{Name: "Whiskers", Color: "Orange"}
	
	animals := []Animal{dog, cat}
	
	fmt.Println("Animals speaking and moving:")
	for _, animal := range animals {
		fmt.Printf("%s: %s, %s\n", 
			getAnimalName(animal), animal.Speak(), animal.Move())
	}

	// Function that works with any type implementing an interface
	fmt.Println("\nFunction that works with any type implementing an interface:")
	fmt.Printf("Total area of all shapes: %.2f\n", calculateTotalArea(shapes))
	fmt.Printf("All animals speaking: %s\n", makeAllAnimalsSpeak(animals))
}

// ============================================================================
// SECTION 3: Interface Composition and Embedding
// ============================================================================
func section3_InterfaceComposition() {
	fmt.Println("\n📚 SECTION 3: Interface Composition and Embedding")
	fmt.Println("------------------------------------------------")

	// Combining multiple interfaces
	fmt.Println("Combining multiple interfaces:")
	fmt.Println("type ReadWriteCloser interface { Reader; Writer; Closer }")
	
	// Simulate a file-like object
	file := &File{Name: "document.txt", Content: "Hello, Go!", IsOpen: true}
	
	// File implements ReadWriteCloser
	fmt.Printf("File: %s\n", file.Name)
	
	// Read from file
	data := make([]byte, 20)
	n, err := file.Read(data)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	}
	
	// Write to file
	writeData := []byte("New content!")
	n, err = file.Write(writeData)
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes\n", n)
	}
	
	// Close file
	err = file.Close()
	if err != nil {
		fmt.Printf("Close error: %v\n", err)
	} else {
		fmt.Println("File closed successfully")
	}

	// Interface embedding
	fmt.Println("\nInterface embedding:")
	fmt.Println("type AdvancedShape interface { Shape; Perimeter() float64 }")
	
	// Circle and Rectangle implement AdvancedShape
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}
	advancedShapes := []AdvancedShape{circle, rectangle}
	
	fmt.Println("Calculating areas and perimeters:")
	for i, shape := range advancedShapes {
		fmt.Printf("Shape %d: Area = %.2f, Perimeter = %.2f\n", 
			i+1, shape.Area(), shape.Perimeter())
	}
}

// ============================================================================
// SECTION 4: Empty Interface and Type Assertions
// ============================================================================
func section4_EmptyInterfaceAndTypeAssertions() {
	fmt.Println("\n📚 SECTION 4: Empty Interface and Type Assertions")
	fmt.Println("-------------------------------------------------")

	// The empty interface
	fmt.Println("The empty interface: interface{}")
	fmt.Println("Any type implements the empty interface (it has no methods)")
	
	// Store different types in a slice of empty interfaces
	var anything []interface{}
	anything = append(anything, "Hello")
	anything = append(anything, 42)
	anything = append(anything, 3.14)
	anything = append(anything, true)
	anything = append(anything, Circle{Radius: 5.0})
	
	fmt.Println("\nStoring different types in empty interface slice:")
	for i, item := range anything {
		fmt.Printf("Item %d: %v (type: %T)\n", i, item, item)
	}

	// Type assertions
	fmt.Println("\nType assertions:")
	fmt.Println("value, ok := interfaceValue.(Type)")
	
	for i, item := range anything {
		fmt.Printf("\nItem %d: ", i)
		
		// Try to assert as string
		if str, ok := item.(string); ok {
			fmt.Printf("String: %s\n", str)
			continue
		}
		
		// Try to assert as int
		if num, ok := item.(int); ok {
			fmt.Printf("Integer: %d\n", num)
			continue
		}
		
		// Try to assert as float64
		if f, ok := item.(float64); ok {
			fmt.Printf("Float: %.2f\n", f)
			continue
		}
		
		// Try to assert as bool
		if b, ok := item.(bool); ok {
			fmt.Printf("Boolean: %t\n", b)
			continue
		}
		
		// Try to assert as Circle
		if c, ok := item.(Circle); ok {
			fmt.Printf("Circle with radius: %.2f\n", c.Radius)
			continue
		}
		
		fmt.Printf("Unknown type: %T\n", item)
	}

	// Type switches
	fmt.Println("\nType switches:")
	fmt.Println("switch v := value.(type) { ... }")
	
	for i, item := range anything {
		fmt.Printf("Item %d: ", i)
		
		switch v := item.(type) {
		case string:
			fmt.Printf("String: %s\n", v)
		case int:
			fmt.Printf("Integer: %d\n", v)
		case float64:
			fmt.Printf("Float: %.2f\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		case Circle:
			fmt.Printf("Circle with radius: %.2f\n", v.Radius)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}
}

// ============================================================================
// SECTION 5: Real-World Interface Examples
// ============================================================================
func section5_RealWorldExamples() {
	fmt.Println("\n📚 SECTION 5: Real-World Interface Examples")
	fmt.Println("---------------------------------------------")

	// Database operations
	fmt.Println("Database operations:")
	
	// Simulate different database types
	mysqlDB := &MySQLDatabase{ConnectionString: "mysql://localhost:3306/mydb"}
	postgresDB := &PostgreSQLDatabase{ConnectionString: "postgres://localhost:5432/mydb"}
	
	databases := []Database{mysqlDB, postgresDB}
	
	for _, db := range databases {
		fmt.Printf("Connecting to %s...\n", db.GetType())
		if err := db.Connect(); err != nil {
			fmt.Printf("Connection failed: %v\n", err)
		} else {
			fmt.Printf("Connected to %s successfully!\n", db.GetType())
			
			// Execute a query
			result, err := db.Query("SELECT * FROM users")
			if err != nil {
				fmt.Printf("Query failed: %v\n", err)
			} else {
				fmt.Printf("Query result: %s\n", result)
			}
			
			// Close connection
			db.Close()
		}
	}

	// HTTP handlers
	fmt.Println("\nHTTP handlers:")
	
	// Different types of handlers
	userHandler := &UserHandler{Endpoint: "/users"}
	productHandler := &ProductHandler{Endpoint: "/products"}
	
	handlers := []HTTPHandler{userHandler, productHandler}
	
	for _, handler := range handlers {
		fmt.Printf("Handling request to %s...\n", handler.GetEndpoint())
		response := handler.Handle("GET", map[string]string{"id": "123"})
		fmt.Printf("Response: %s\n", response)
	}

	// Sortable collections
	fmt.Println("\nSortable collections:")
	
	// Different types that can be sorted
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	names := []string{"Charlie", "Alice", "Bob", "David"}
	
	fmt.Printf("Original numbers: %v\n", numbers)
	sortInts(numbers)
	fmt.Printf("Sorted numbers: %v\n", numbers)
	
	fmt.Printf("Original names: %v\n", names)
	sortStrings(names)
	fmt.Printf("Sorted names: %v\n", names)
}

// ============================================================================
// HELPER FUNCTIONS AND INTERFACES
// ============================================================================

// Basic interfaces
type Shape interface {
	Area() float64
}

type AdvancedShape interface {
	Shape
	Perimeter() float64
}

type Animal interface {
	Speak() string
	Move() string
}

// ReadWriteCloser interface (like io.ReadWriteCloser)
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// Database interface
type Database interface {
	Connect() error
	Query(query string) (string, error)
	Close()
	GetType() string
}

// HTTP handler interface
type HTTPHandler interface {
	Handle(method string, params map[string]string) string
	GetEndpoint() string
}

// Concrete types implementing interfaces
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name  string
	Color string
}

type File struct {
	Name    string
	Content string
	IsOpen  bool
}

type MySQLDatabase struct {
	ConnectionString string
	IsConnected     bool
}

type PostgreSQLDatabase struct {
	ConnectionString string
	IsConnected     bool
}

type UserHandler struct {
	Endpoint string
}

type ProductHandler struct {
	Endpoint string
}

// Error types (implementing the error interface)
type StringError string

func (e StringError) Error() string {
	return string(e)
}

type CodeError int

func (c CodeError) Error() string {
	return fmt.Sprintf("error code: %d", c)
}

type PersonError struct {
	Name    string
	Message string
}

func (e PersonError) Error() string {
	return fmt.Sprintf("person error for %s: %s", e.Name, e.Message)
}

// Shape methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Animal methods
func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running on four legs"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Walking gracefully"
}

// File methods (implementing ReadWriteCloser)
func (f *File) Read(p []byte) (n int, err error) {
	if !f.IsOpen {
		return 0, fmt.Errorf("file is not open")
	}
	
	copy(p, []byte(f.Content))
	return len(f.Content), nil
}

func (f *File) Write(p []byte) (n int, err error) {
	if !f.IsOpen {
		return 0, fmt.Errorf("file is not open")
	}
	
	f.Content = string(p)
	return len(p), nil
}

func (f *File) Close() error {
	if !f.IsOpen {
		return fmt.Errorf("file is already closed")
	}
	
	f.IsOpen = false
	return nil
}

// Database methods
func (m *MySQLDatabase) Connect() error {
	m.IsConnected = true
	return nil
}

func (m *MySQLDatabase) Query(query string) (string, error) {
	if !m.IsConnected {
		return "", fmt.Errorf("not connected to database")
	}
	return fmt.Sprintf("MySQL result for: %s", query), nil
}

func (m *MySQLDatabase) Close() {
	m.IsConnected = false
}

func (m *MySQLDatabase) GetType() string {
	return "MySQL"
}

func (p *PostgreSQLDatabase) Connect() error {
	p.IsConnected = true
	return nil
}

func (p *PostgreSQLDatabase) Query(query string) (string, error) {
	if !p.IsConnected {
		return "", fmt.Errorf("file is not open")
	}
	return fmt.Sprintf("PostgreSQL result for: %s", query), nil
}

func (p *PostgreSQLDatabase) Close() {
	p.IsConnected = false
}

func (p *PostgreSQLDatabase) GetType() string {
	return "PostgreSQL"
}

// HTTP handler methods
func (u *UserHandler) Handle(method string, params map[string]string) string {
	return fmt.Sprintf("User handler: %s %s with params %v", method, u.Endpoint, params)
}

func (u *UserHandler) GetEndpoint() string {
	return u.Endpoint
}

func (p *ProductHandler) Handle(method string, params map[string]string) string {
	return fmt.Sprintf("Product handler: %s %s with params %v", method, p.Endpoint, params)
}

func (p *ProductHandler) GetEndpoint() string {
	return p.Endpoint
}

// Helper functions
func getAnimalName(animal Animal) string {
	switch a := animal.(type) {
	case Dog:
		return a.Name
	case Cat:
		return a.Name
	default:
		return "Unknown"
	}
}

func calculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func makeAllAnimalsSpeak(animals []Animal) string {
	var responses []string
	for _, animal := range animals {
		responses = append(responses, animal.Speak())
	}
	return strings.Join(responses, ", ")
}

func sortInts(nums []int) {
	// Simple bubble sort for demonstration
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func sortStrings(strs []string) {
	// Simple bubble sort for demonstration
	for i := 0; i < len(strs)-1; i++ {
		for j := 0; j < len(strs)-i-1; j++ {
			if strs[j] > strs[j+1] {
				strs[j], strs[j+1] = strs[j+1], strs[j]
			}
		}
	}
} 