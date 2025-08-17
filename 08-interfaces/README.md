# Go Interfaces - Chapter 8 🐹

Welcome to Chapter 8! Now that you understand structs and methods, you're ready to learn about **interfaces** - one of Go's most powerful and elegant features. This chapter will teach you how to define behavior contracts and write flexible, reusable code.

## What You'll Learn

- What interfaces are and why they're powerful
- How to create and use interfaces effectively
- Interface composition and embedding
- The empty interface and type assertions
- Real-world interface examples
- How interfaces enable polymorphism in Go

## Why Interfaces Matter

### **Think of Interfaces Like This:**

Imagine you're hiring for a job:
- **Job Description** = The interface (defines what skills are required)
- **Candidates** = Types that implement the interface (have the required skills)
- **Hiring Process** = Code that works with any type implementing the interface

In Go:
- **Interfaces** define what methods a type must have
- **Any type** with the required methods automatically satisfies the interface
- **Functions** can work with any type that implements an interface

### **Go's Interface Philosophy:**

1. **Implicit implementation** - types don't declare which interfaces they implement
2. **Small interfaces** - prefer many small interfaces over few large ones
3. **Composition over inheritance** - combine interfaces to create new ones
4. **Duck typing** - "if it walks like a duck and quacks like a duck, it's a duck"

## Section 1: What Are Interfaces?

### **What is an Interface?**

An interface is a **contract** that defines what methods a type must have. It's like a blueprint that says "to be considered X, you must be able to do Y and Z."

```go
// Basic interface definition
type Shape interface {
    Area() float64
}

// This interface says: "To be a Shape, you must have an Area() method that returns float64"
```

### **The Magic of Implicit Implementation**

In Go, types **automatically implement interfaces** if they have the required methods. No explicit declaration needed!

```go
// Circle has an Area() method, so it automatically implements Shape
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Rectangle also has an Area() method, so it also implements Shape
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Both Circle and Rectangle can be used wherever a Shape is expected!
```

### **Why This is So Powerful**

1. **No coupling** - types don't need to know about interfaces
2. **Easy testing** - create mock types that implement interfaces
3. **Flexible code** - functions work with any type that has the right methods
4. **Clean architecture** - separate behavior definition from implementation

## Section 2: Creating and Using Interfaces

### **Basic Interface Definition**

```go
// Simple interface with one method
type Speaker interface {
    Speak() string
}

// Interface with multiple methods
type Animal interface {
    Speak() string
    Move() string
    GetName() string
}

// Interface with no methods (empty interface)
type Any interface{} // or just interface{}
```

### **Implementing Interfaces**

```go
// Dog implements the Animal interface
type Dog struct {
    Name  string
    Breed string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running on four legs"
}

func (d Dog) GetName() string {
    return d.Name
}

// Cat also implements the Animal interface
type Cat struct {
    Name  string
    Color string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (c Cat) Move() string {
    return "Walking gracefully"
}

func (c Cat) GetName() string {
    return c.Name
}
```

### **Using Interfaces in Functions**

```go
// Function that works with any Animal
func DescribeAnimal(animal Animal) {
    fmt.Printf("%s says: %s\n", animal.GetName(), animal.Speak())
    fmt.Printf("%s moves by: %s\n", animal.GetName(), animal.Move())
}

// Use it with different types
dog := Dog{Name: "Buddy", Breed: "Golden Retriever"}
cat := Cat{Name: "Whiskers", Color: "Orange"}

DescribeAnimal(dog)   // Works with Dog
DescribeAnimal(cat)   // Works with Cat
```

### **Collections of Interfaces**

```go
// Store different types that implement the same interface
var animals []Animal
animals = append(animals, dog)
animals = append(animals, cat)

// Process all animals the same way
for _, animal := range animals {
    DescribeAnimal(animal)
}
```

## Section 3: Interface Composition and Embedding

### **Combining Interfaces**

Interfaces can be combined to create new, more specific interfaces:

```go
// Basic interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Combined interface
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// This is equivalent to:
type ReadWriteCloser interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
}
```

### **Why Interface Composition is Powerful**

1. **Reusability** - combine existing interfaces
2. **Flexibility** - types can implement only what they need
3. **Clarity** - interfaces express exactly what's required
4. **Maintainability** - change one interface, affects all combined ones

### **Real Example: File Operations**

```go
// File implements ReadWriteCloser
type File struct {
    Name    string
    Content string
    IsOpen  bool
}

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

// Now File can be used anywhere a Reader, Writer, Closer, or ReadWriteCloser is expected
```

## Section 4: Empty Interface and Type Assertions

### **The Empty Interface**

The empty interface `interface{}` has no methods, so **every type implements it**. This makes it useful for storing values of unknown types.

```go
// Store any type in an empty interface
var anything interface{}
anything = "Hello"
anything = 42
anything = 3.14
anything = true
anything = Circle{Radius: 5.0}

// Empty interface slice
var items []interface{}
items = append(items, "string")
items = append(items, 123)
items = append(items, Circle{Radius: 3.0})
```

### **Type Assertions**

Type assertions let you extract the concrete type from an interface:

```go
// Basic type assertion
value, ok := interfaceValue.(Type)

// Examples
if str, ok := anything.(string); ok {
    fmt.Printf("It's a string: %s\n", str)
}

if num, ok := anything.(int); ok {
    fmt.Printf("It's an int: %d\n", num)
}

if circle, ok := anything.(Circle); ok {
    fmt.Printf("It's a circle with radius: %.2f\n", circle.Radius)
}
```

### **Type Switches**

Type switches are a cleaner way to check multiple types:

```go
switch v := value.(type) {
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
```

### **When to Use Empty Interfaces**

1. **Generic containers** - store different types
2. **JSON handling** - unmarshal into `interface{}`
3. **Reflection** - examine types at runtime
4. **Legacy code** - work with existing APIs

## Section 5: Real-World Interface Examples

### **Database Operations**

```go
// Database interface
type Database interface {
    Connect() error
    Query(query string) (string, error)
    Close()
    GetType() string
}

// MySQL implementation
type MySQLDatabase struct {
    ConnectionString string
    IsConnected     bool
}

func (m *MySQLDatabase) Connect() error {
    m.IsConnected = true
    return nil
}

func (m *MySQLDatabase) Query(query string) (string, error) {
    if !m.IsConnected {
        return "", fmt.Errorf("not connected")
    }
    return fmt.Sprintf("MySQL result: %s", query), nil
}

func (m *MySQLDatabase) Close() {
    m.IsConnected = false
}

func (m *MySQLDatabase) GetType() string {
    return "MySQL"
}

// PostgreSQL implementation
type PostgreSQLDatabase struct {
    ConnectionString string
    IsConnected     bool
}

// ... similar methods for PostgreSQL

// Function that works with any database
func ExecuteQuery(db Database, query string) (string, error) {
    if err := db.Connect(); err != nil {
        return "", fmt.Errorf("failed to connect: %w", err)
    }
    defer db.Close()
    
    return db.Query(query)
}

// Use with different databases
mysqlDB := &MySQLDatabase{ConnectionString: "mysql://localhost:3306/mydb"}
postgresDB := &PostgreSQLDatabase{ConnectionString: "postgres://localhost:5432/mydb"}

result1, _ := ExecuteQuery(mysqlDB, "SELECT * FROM users")
result2, _ := ExecuteQuery(postgresDB, "SELECT * FROM users")
```

### **HTTP Handlers**

```go
// HTTP handler interface
type HTTPHandler interface {
    Handle(method string, params map[string]string) string
    GetEndpoint() string
}

// User handler
type UserHandler struct {
    Endpoint string
}

func (u *UserHandler) Handle(method string, params map[string]string) string {
    return fmt.Sprintf("User handler: %s %s", method, u.Endpoint)
}

func (u *UserHandler) GetEndpoint() string {
    return u.Endpoint
}

// Product handler
type ProductHandler struct {
    Endpoint string
}

func (p *ProductHandler) Handle(method string, params map[string]string) string {
    return fmt.Sprintf("Product handler: %s %s", method, p.Endpoint)
}

func (p *ProductHandler) GetEndpoint() string {
    return p.Endpoint
}

// Router that works with any handler
func RouteRequest(handler HTTPHandler, method string, params map[string]string) string {
    return handler.Handle(method, params)
}
```

### **Sorting and Collections**

```go
// Sortable interface
type Sortable interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// IntSlice implements Sortable
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// StringSlice implements Sortable
type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Generic sort function
func Sort(sortable Sortable) {
    // Implementation of sorting algorithm
    // Works with any type implementing Sortable
}

// Use with different types
numbers := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
names := StringSlice{"Charlie", "Alice", "Bob", "David"}

Sort(numbers)
Sort(names)
```

## How to Run Your Program

1. Open your terminal
2. Go to the interfaces folder: `cd 08-interfaces`
3. Run the program: `go run main.go`

## Try It Yourself!

### Exercise 1: Create a Logger Interface
Create a logging interface with different implementations:

```go
type Logger interface {
    Log(level string, message string)
    SetLevel(level string)
}

type ConsoleLogger struct {
    Level string
}

type FileLogger struct {
    Filename string
    Level    string
}

// Implement the Logger interface for both types
// Then create a function that works with any Logger
```

### Exercise 2: Build a Payment Processor
Create a payment processing system:

```go
type PaymentMethod interface {
    Process(amount float64) error
    GetType() string
    IsAvailable() bool
}

type CreditCard struct {
    Number string
    Expiry string
    CVV    string
}

type PayPal struct {
    Email string
}

type BankTransfer struct {
    AccountNumber string
    RoutingNumber string
}

// Implement PaymentMethod for all three types
// Create a function that processes payments with any method
```

### Exercise 3: Create a Data Store
Build a flexible data storage system:

```go
type DataStore interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}) error
    Delete(key string) error
    List() []string
}

type MemoryStore struct {
    data map[string]interface{}
}

type FileStore struct {
    filename string
}

// Implement DataStore for both types
// Create a function that works with any DataStore
```

## Common Mistakes to Avoid

### 1. **Over-Engineering with Interfaces**
```go
// ❌ WRONG: Interface for everything
type StringProcessor interface {
    Process(s string) string
}

// ✅ CORRECT: Interface only when you need polymorphism
type StringProcessor interface {
    Process(s string) string
}

// Use it when you need to work with different types
func ProcessStrings(processor StringProcessor, strings []string) []string {
    // ... implementation
}
```

### 2. **Large Interfaces**
```go
// ❌ WRONG: Too many methods
type BigInterface interface {
    Method1()
    Method2()
    Method3()
    Method4()
    Method5()
    Method6()
    Method7()
    Method8()
}

// ✅ CORRECT: Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Combine when needed
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

### 3. **Ignoring Interface Satisfaction**
```go
// ❌ WRONG: Not checking if type implements interface
func ProcessData(data SomeInterface) {
    // ... implementation
}

// This will cause a compile error if the type doesn't implement SomeInterface
```

## Key Takeaways

1. **Interfaces define behavior** - specify what methods a type must have
2. **Implicit implementation** - types automatically implement interfaces
3. **Small interfaces are better** - prefer many small interfaces over few large ones
4. **Composition over inheritance** - combine interfaces to create new ones
5. **Empty interface for any type** - use `interface{}` when you need to store any type
6. **Type assertions extract types** - use `value.(Type)` to get concrete types
7. **Type switches are cleaner** - use `switch v := value.(type)` for multiple types
8. **Interfaces enable polymorphism** - write functions that work with any compatible type

## Next Steps

After this chapter, you'll be ready for:
- **Pointers** (understanding memory and references)
- **Error handling** (building robust applications with the error interface)
- **Testing** (using interfaces for mocking)
- **Packages** (organizing your code)
- **Concurrency** (goroutines and channels)

---

**Excellent work! You now understand Go interfaces! 🎉**

Interfaces are what make Go's type system so powerful and flexible. You've learned how to define behavior contracts, create polymorphic code, and build systems that can work with any type that implements the right methods.

Practice creating interfaces for real-world scenarios, combining them to create new interfaces, and using them to write flexible, testable code. This knowledge will help you write clean, maintainable, and extensible Go programs! 🚀 