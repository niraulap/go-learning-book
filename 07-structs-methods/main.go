package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("🐹 Go Structs and Methods - Chapter 7 🐹")
	fmt.Println("==========================================")

	// ============================================================================
	// SECTION 1: Creating and Using Structs
	// ============================================================================
	section1_CreatingStructs()

	// ============================================================================
	// SECTION 2: Methods - Functions Attached to Types
	// ============================================================================
	section2_Methods()

	// ============================================================================
	// SECTION 3: Advanced Struct Patterns
	// ============================================================================
	section3_AdvancedStructPatterns()

	// ============================================================================
	// SECTION 4: Structs with Collections
	// ============================================================================
	section4_StructsWithCollections()

	fmt.Println("\n🎉 Chapter 7 Complete! You understand Go structs and methods!")
}

// ============================================================================
// SECTION 1: Creating and Using Structs
// ============================================================================
func section1_CreatingStructs() {
	fmt.Println("\n📚 SECTION 1: Creating and Using Structs")
	fmt.Println("----------------------------------------")

	// Basic struct creation
	fmt.Println("Basic struct creation:")
	person1 := Person{
		Name:     "Alice",
		Age:      25,
		Email:    "alice@example.com",
		IsActive: true,
	}
	fmt.Printf("Person: %+v\n", person1)

	// Struct literals
	fmt.Println("\nStruct literals:")
	person2 := Person{"Bob", 30, "bob@example.com", true}
	fmt.Printf("Person: %+v\n", person2)

	// Using new() function
	fmt.Println("\nUsing new() function:")
	person3 := new(Person)
	person3.Name = "Charlie"
	person3.Age = 28
	person3.Email = "charlie@example.com"
	person3.IsActive = true
	fmt.Printf("Person: %+v\n", person3)

	// Nested structs
	fmt.Println("\nNested structs:")
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
	fmt.Printf("Employee: %+v\n", employee)

	// Structs with collections
	fmt.Println("\nStructs with collections:")
	team := Team{
		Members: []Person{person1, person2, person3},
		Info: map[string]string{
			"department": "Engineering",
			"location":   "New York",
		},
	}
	fmt.Printf("Team: %+v\n", team)
}

// ============================================================================
// SECTION 2: Methods - Functions Attached to Types
// ============================================================================
func section2_Methods() {
	fmt.Println("\n📚 SECTION 2: Methods - Functions Attached to Types")
	fmt.Println("--------------------------------------------------")

	person := Person{Name: "Alice", Age: 25, Email: "alice@example.com", IsActive: true}

	// Basic methods
	fmt.Println("Basic methods:")
	person.Introduce()
	person.HaveBirthday()
	fmt.Printf("After birthday: %+v\n", person)

	// Methods with parameters
	fmt.Println("\nMethods with parameters:")
	person.UpdateEmail("alice.new@example.com")
	fmt.Printf("After email update: %+v\n", person)

	// Methods that return values
	fmt.Println("\nMethods that return values:")
	isAdult := person.IsAdult()
	fmt.Printf("Is adult: %t\n", isAdult)

	// Methods on embedded structs
	fmt.Println("\nMethods on embedded structs:")
	employee := Employee{
		Person: person,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		Salary: 75000,
		Role:   "Software Engineer",
	}
	employee.DisplayInfo()

	// Method chaining
	fmt.Println("\nMethod chaining:")
	computer := NewComputerBuilder().
		SetCPU("Intel i9").
		SetRAM(32).
		SetStorage("1TB NVMe").
		SetGPU("RTX 4080").
		Build()
	fmt.Printf("Built computer: %+v\n", computer)
}

// ============================================================================
// SECTION 3: Advanced Struct Patterns
// ============================================================================
func section3_AdvancedStructPatterns() {
	fmt.Println("\n📚 SECTION 3: Advanced Struct Patterns")
	fmt.Println("----------------------------------------")

	// Builder pattern
	fmt.Println("Builder pattern:")
	car := NewCarBuilder().
		SetBrand("Tesla").
		SetModel("Model 3").
		SetYear(2024).
		SetColor("Red").
		Build()
	fmt.Printf("Built car: %+v\n", car)

	// Factory pattern
	fmt.Println("\nFactory pattern:")
	circle := NewShape("circle", 5.0)
	square := NewShape("square", 4.0)
	triangle := NewShape("triangle", 3.0)

	fmt.Printf("Circle area: %.2f\n", circle.Area())
	fmt.Printf("Square area: %.2f\n", square.Area())
	fmt.Printf("Triangle area: %.2f\n", triangle.Area())

	// Validation pattern
	fmt.Println("\nValidation pattern:")
	user := User{Name: "John", Age: 25, Email: "john@example.com"}
	if err := user.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("User is valid!")
	}

	// Invalid user
	invalidUser := User{Name: "", Age: 15, Email: "invalid-email"}
	if err := invalidUser.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("User is valid!")
	}
}

// ============================================================================
// SECTION 4: Structs with Collections
// ============================================================================
func section4_StructsWithCollections() {
	fmt.Println("\n📚 SECTION 4: Structs with Collections")
	fmt.Println("----------------------------------------")

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

	// Display team information
	fmt.Printf("Team: %s\n", team.Info["department"])
	fmt.Printf("Location: %s\n", team.Info["location"])
	fmt.Printf("Project: %s\n", team.Info["project"])
	fmt.Printf("Member count: %d\n", len(team.Members))

	// Filter active members
	fmt.Println("\nActive members:")
	for _, member := range team.Members {
		if member.IsActive {
			fmt.Printf("- %s (%s)\n", member.Name, member.Email)
		}
	}

	// Add a new member
	newMember := Person{Name: "Diana", Age: 27, Email: "diana@example.com", IsActive: true}
	team.AddMember(newMember)
	fmt.Printf("\nAfter adding member, total count: %d\n", len(team.Members))

	// Display all members
	fmt.Println("\nAll team members:")
	for i, member := range team.Members {
		status := "Active"
		if !member.IsActive {
			status = "Inactive"
		}
		fmt.Printf("%d. %s - %s (%s)\n", i+1, member.Name, member.Email, status)
	}
}

// ============================================================================
// HELPER FUNCTIONS AND STRUCTS
// ============================================================================

// Person represents a person with basic information
type Person struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
}

// Address represents a physical address
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// Employee represents an employee with person info and work details
type Employee struct {
	Person
	Address
	Salary int
	Role   string
}

// Team represents a team with members and info
type Team struct {
	Members []Person
	Info    map[string]string
}

// User represents a user for validation
type User struct {
	Name  string
	Age   int
	Email string
}

// Shape represents different geometric shapes
type Shape struct {
	Type   string
	Size   float64
	Radius float64
	Side   float64
	Base   float64
	Height float64
}

// Computer represents a computer system
type Computer struct {
	CPU     string
	RAM     int
	Storage string
	GPU     string
}

// Car represents a car
type Car struct {
	Brand string
	Model string
	Year  int
	Color string
}

// Person methods
func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

func (p Person) HaveBirthday() {
	p.Age++
	fmt.Printf("Happy birthday! %s is now %d years old.\n", p.Name, p.Age)
}

func (p Person) UpdateEmail(newEmail string) {
	p.Email = newEmail
	fmt.Printf("Email updated to: %s\n", p.Email)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Employee methods
func (e Employee) DisplayInfo() {
	fmt.Printf("Employee: %s\n", e.Person.Name)
	fmt.Printf("Role: %s\n", e.Role)
	fmt.Printf("Salary: $%d\n", e.Salary)
	fmt.Printf("Address: %s, %s, %s %s\n", e.Address.Street, e.Address.City, e.Address.State, e.Address.ZipCode)
}

// Team methods
func (t *Team) AddMember(person Person) {
	t.Members = append(t.Members, person)
	t.Info["count"] = fmt.Sprintf("%d", len(t.Members))
}

// User methods
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

// Shape methods
func (s Shape) Area() float64 {
	switch s.Type {
	case "circle":
		return math.Pi * s.Radius * s.Radius
	case "square":
		return s.Side * s.Side
	case "triangle":
		return 0.5 * s.Base * s.Height
	default:
		return 0
	}
}

// Computer builder
type ComputerBuilder struct {
	computer Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func (cb *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	cb.computer.CPU = cpu
	return cb
}

func (cb *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	cb.computer.RAM = ram
	return cb
}

func (cb *ComputerBuilder) SetStorage(storage string) *ComputerBuilder {
	cb.computer.Storage = storage
	return cb
}

func (cb *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	cb.computer.GPU = gpu
	return cb
}

func (cb *ComputerBuilder) Build() Computer {
	return cb.computer
}

// Car builder
type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
	cb.car.Brand = brand
	return cb
}

func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.car.Model = model
	return cb
}

func (cb *CarBuilder) SetYear(year int) *CarBuilder {
	cb.car.Year = year
	return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.car.Color = color
	return cb
}

func (cb *CarBuilder) Build() Car {
	return cb.car
}

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