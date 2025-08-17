# Go Structs and Methods - Chapter 7 🐹

Welcome to Chapter 7! Now that you understand functions, arrays, slices, and maps, you're ready to learn about **structs and methods** - the building blocks of Go's object-oriented programming approach. This chapter will teach you how to create custom data types and attach behavior to them.

## What You'll Learn

- Creating and using custom data types with structs
- Attaching functions (methods) to your custom types
- Struct embedding and composition
- Advanced patterns like builders and factories
- Working with structs and collections
- Building real-world data structures

## Why Structs and Methods Matter

### **Think of Structs Like This:**

Imagine you're building a house:
- **Struct** = The blueprint (defines what the house will have)
- **Fields** = The rooms, doors, windows (the data)
- **Methods** = The functions (turn on lights, open garage, etc.)

In Go:
- **Structs** define what data your type holds
- **Methods** define what your type can do
- **Composition** lets you build complex types from simple ones

### **Go's Approach to OOP:**

Unlike traditional object-oriented languages (Java, C#), Go uses:
- **Composition over inheritance** - build complex types by combining simple ones
- **Methods attached to types** - functions that belong to specific data types
- **Interface-based polymorphism** - we'll learn this in the next chapter!

## Section 1: Creating and Using Structs

### **What is a Struct?**

A struct is a **custom data type** that groups together related data. Think of it as a container that holds multiple pieces of information about one thing.

```go
// Basic struct definition
type Person struct {
    Name     string  // Person's name
    Age      int     // Person's age
    Email    string  // Person's email
    IsActive bool    // Whether person is active
}
```

### **Creating Structs (Three Ways)**

```go
// Method 1: Struct literal (create and initialize at once)
person1 := Person{
    Name:     "Alice",
    Age:      25,
    Email:    "alice@example.com",
    IsActive: true,
}

// Method 2: Create empty struct, then set fields
var person2 Person
person2.Name = "Bob"
person2.Age = 30
person2.Email = "bob@example.com"
person2.IsActive = false

// Method 3: Using new() function
person3 := new(Person)
person3.Name = "Charlie"
person3.Age = 28
person3.Email = "charlie@example.com"
person3.IsActive = true
```

### **Accessing Struct Fields**

```go
// Access individual fields
fmt.Printf("Name: %s\n", person1.Name)
fmt.Printf("Age: %d\n", person1.Age)
fmt.Printf("Email: %s\n", person1.Email)
fmt.Printf("Is Active: %t\n", person1.IsActive)

// Print entire struct
fmt.Printf("Person: %+v\n", person1)  // %+v shows field names
fmt.Printf("Person: %#v\n", person1)  // %#v shows Go syntax
```

### **Nested Structs (Structs Inside Structs)**

```go
// Address struct
type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

// Employee struct that contains Person and Address
type Employee struct {
    Person  Person  // Embedded struct
    Address Address // Embedded struct
    Salary  int
    Role    string
}

// Create an employee
address := Address{
    Street:  "123 Main St",
    City:    "New York",
    State:   "NY",
    ZipCode: "10001",
}

employee := Employee{
    Person:  person1,
    Address: address,
    Salary:  75000,
    Role:    "Software Engineer",
}

// Access nested fields
fmt.Printf("Employee name: %s\n", employee.Person.Name)
fmt.Printf("Employee city: %s\n", employee.Address.City)
fmt.Printf("Employee salary: $%d\n", employee.Salary)
```

### **Structs with Collections**

```go
// Team struct with slices and maps
type Team struct {
    Members []Person                    // Slice of people
    Info    map[string]string          // Map of team information
}

// Create a team
team := Team{
    Members: []Person{person1, person2, person3},
    Info: map[string]string{
        "department": "Engineering",
        "location":   "New York",
        "project":    "Go Learning Book",
    },
}

// Access team data
fmt.Printf("Team: %s\n", team.Info["department"])
fmt.Printf("Member count: %d\n", len(team.Members))

// Loop through team members
for i, member := range team.Members {
    fmt.Printf("Member %d: %s (%s)\n", i+1, member.Name, member.Email)
}
```

## Section 2: Methods - Functions Attached to Types

### **What are Methods?**

Methods are **functions that belong to a specific type**. They let you attach behavior to your data structures.

```go
// Method definition syntax:
// func (receiver Type) MethodName(parameters) returnType { ... }

// Method attached to Person type
func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

// Method with parameters
func (p Person) UpdateEmail(newEmail string) {
    p.Email = newEmail
    fmt.Printf("Email updated to: %s\n", p.Email)
}

// Method that returns a value
func (p Person) IsAdult() bool {
    return p.Age >= 18
}
```

### **Using Methods**

```go
// Create a person
person := Person{Name: "Alice", Age: 25, Email: "alice@example.com", IsActive: true}

// Call methods on the person
person.Introduce()                    // Output: Hi, I'm Alice and I'm 25 years old.
person.UpdateEmail("alice.new@example.com")
fmt.Printf("Is adult: %t\n", person.IsAdult())  // Output: Is adult: true
```

### **Methods on Embedded Structs**

When you embed one struct inside another, the outer struct can use methods from the inner struct:

```go
// Employee can use Person methods
func (e Employee) DisplayInfo() {
    fmt.Printf("Employee: %s\n", e.Person.Name)  // Use Person's Name field
    fmt.Printf("Role: %s\n", e.Role)
    fmt.Printf("Salary: $%d\n", e.Salary)
    fmt.Printf("Address: %s, %s, %s %s\n", 
        e.Address.Street, e.Address.City, e.Address.State, e.Address.ZipCode)
}

// Use it
employee.DisplayInfo()
```

### **Method Chaining**

Methods can return the receiver, allowing you to chain method calls:

```go
// Computer builder with method chaining
type ComputerBuilder struct {
    computer Computer
}

func (cb *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
    cb.computer.CPU = cpu
    return cb  // Return the builder for chaining
}

func (cb *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
    cb.computer.RAM = ram
    return cb
}

func (cb *ComputerBuilder) Build() Computer {
    return cb.computer
}

// Use method chaining
computer := NewComputerBuilder().
    SetCPU("Intel i9").
    SetRAM(32).
    SetStorage("1TB NVMe").
    SetGPU("RTX 4080").
    Build()
```

## Section 3: Advanced Struct Patterns

### **Builder Pattern**

The builder pattern creates complex objects step by step:

```go
// Car builder
type CarBuilder struct {
    car Car
}

func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
    cb.car.Brand = brand
    return cb
}

func (cb *CarBuilder) SetModel(model string) *CarBuilder {
    cb.car.Model = model
    return cb
}

func (cb *CarBuilder) Build() Car {
    return cb.car
}

// Use the builder
car := NewCarBuilder().
    SetBrand("Tesla").
    SetModel("Model 3").
    SetYear(2024).
    SetColor("Red").
    Build()
```

### **Factory Pattern**

The factory pattern creates objects without specifying their exact type:

```go
// Shape factory
func NewShape(shapeType string, size float64) Shape {
    switch shapeType {
    case "circle":
        return Shape{Type: "circle", Radius: size}
    case "square":
        return Shape{Type: "square", Side: size}
    case "triangle":
        return Shape{Type: "triangle", Base: size, Height: size}
    default:
        return Shape{Type: "unknown"}
    }
}

// Use the factory
circle := NewShape("circle", 5.0)
square := NewShape("square", 4.0)
triangle := NewShape("triangle", 3.0)
```

### **Validation Pattern**

Methods can validate data and return errors:

```go
func (u User) Validate() error {
    if u.Name == "" {
        return fmt.Errorf("name is required")
    }
    if u.Age < 18 {
        return fmt.Errorf("age must be at least 18, got %d", u.Age)
    }
    if !strings.Contains(u.Email, "@") {
        return fmt.Errorf("invalid email format: %s", u.Email)
    }
    return nil
}

// Use validation
user := User{Name: "John", Age: 25, Email: "john@example.com"}
if err := user.Validate(); err != nil {
    fmt.Printf("Validation error: %v\n", err)
} else {
    fmt.Println("User is valid!")
}
```

## Section 4: Structs with Collections

### **Working with Slices of Structs**

```go
// Create a team with members
team := Team{
    Members: []Person{
        {Name: "Alice", Age: 25, Email: "alice@example.com", IsActive: true},
        {Name: "Bob", Age: 30, Email: "bob@example.com", IsActive: true},
        {Name: "Charlie", Age: 28, Email: "charlie@example.com", IsActive: false},
    },
    Info: map[string]string{
        "department": "Engineering",
        "location":   "New York",
        "project":    "Go Learning Book",
    },
}

// Filter active members
fmt.Println("Active members:")
for _, member := range team.Members {
    if member.IsActive {
        fmt.Printf("- %s (%s)\n", member.Name, member.Email)
    }
}
```

### **Adding Methods to Collections**

```go
// Method to add members to a team
func (t *Team) AddMember(person Person) {
    t.Members = append(t.Members, person)
    t.Info["count"] = fmt.Sprintf("%d", len(t.Members))
}

// Use the method
newMember := Person{Name: "Diana", Age: 27, Email: "diana@example.com", IsActive: true}
team.AddMember(newMember)
fmt.Printf("After adding member, total count: %d\n", len(team.Members))
```

## How to Run Your Program

1. Open your terminal
2. Go to the structs-methods folder: `cd 07-structs-methods`
3. Run the program: `go run main.go`

## Try It Yourself!

### Exercise 1: Create a Book Struct
Create a struct for books with methods:

```go
type Book struct {
    Title       string
    Author      string
    Pages       int
    Published   int
    IsAvailable bool
}

func (b Book) IsLong() bool {
    return b.Pages > 300
}

func (b Book) GetAge() int {
    return time.Now().Year() - b.Published
}

func (b Book) Summary() string {
    return fmt.Sprintf("%s by %s (%d pages, published %d)", 
        b.Title, b.Author, b.Pages, b.Published)
}
```

### Exercise 2: Build a Shopping Cart
Create a shopping cart system:

```go
type CartItem struct {
    Product  string
    Quantity int
    Price    float64
}

type ShoppingCart struct {
    Items     []CartItem
    Customer  Person
    CreatedAt time.Time
}

func (sc *ShoppingCart) AddItem(item CartItem) {
    sc.Items = append(sc.Items, item)
}

func (sc ShoppingCart) Total() float64 {
    total := 0.0
    for _, item := range sc.Items {
        total += float64(item.Quantity) * item.Price
    }
    return total
}

func (sc ShoppingCart) ItemCount() int {
    count := 0
    for _, item := range sc.Items {
        count += item.Quantity
    }
    return count
}
```

## Common Mistakes to Avoid

### 1. **Forgetting to Initialize Structs**
```go
// ❌ WRONG: Accessing uninitialized fields
var person Person
fmt.Println(person.Name)  // Empty string (zero value)

// ✅ CORRECT: Initialize before use
person := Person{Name: "Alice", Age: 25}
fmt.Println(person.Name)  // Alice
```

### 2. **Not Using Method Receivers Properly**
```go
// ❌ WRONG: Method doesn't modify the struct
func (p Person) HaveBirthday() {
    p.Age++  // This modifies a copy, not the original!
}

// ✅ CORRECT: Use pointer receiver to modify original
func (p *Person) HaveBirthday() {
    p.Age++  // This modifies the original struct
}
```

### 3. **Ignoring Embedded Struct Behavior**
```go
// ❌ WRONG: Not understanding embedded struct access
employee := Employee{Person: person1, Salary: 75000}
fmt.Println(employee.Name)  // This works due to embedding

// ✅ CORRECT: Be explicit about embedded access
fmt.Println(employee.Person.Name)  // More explicit and clear
```

## Key Takeaways

1. **Structs define data structure** - group related data together
2. **Methods attach behavior** - functions that belong to specific types
3. **Embedding enables composition** - build complex types from simple ones
4. **Use meaningful field names** - make your code self-documenting
5. **Methods can return values** - calculate and return information
6. **Builders create complex objects** - step-by-step object construction
7. **Factories create objects** - without specifying exact types
8. **Validation ensures data quality** - check data before using it

## Next Steps

After this chapter, you'll be ready for:
- **Interfaces** (defining behavior contracts)
- **Pointers** (understanding memory and references)
- **Error handling** (building robust applications)
- **Testing** (ensuring your code works correctly)
- **Packages** (organizing your code)

---

**Excellent work! You now understand Go structs and methods! 🎉**

Structs and methods are the foundation of Go's approach to organizing data and behavior. You've learned how to create custom types, attach behavior to them, and build complex systems using composition.

Practice creating structs for real-world scenarios, implementing methods that make sense for your data, and using patterns like builders and factories. This knowledge will help you write clean, organized, and maintainable Go code! 🚀 