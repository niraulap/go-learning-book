package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("🐹 Go Pointers - Chapter 9 🐹")
	fmt.Println("===============================")

	// ============================================================================
	// SECTION 1: Value vs Reference Types
	// ============================================================================
	section1_ValueVsReferenceTypes()

	// ============================================================================
	// SECTION 2: Basic Pointer Syntax
	// ============================================================================
	section2_BasicPointerSyntax()

	// ============================================================================
	// SECTION 3: Zero Values and Pointers
	// ============================================================================
	section3_ZeroValuesAndPointers()

	// ============================================================================
	// SECTION 4: Large Structs and Performance
	// ============================================================================
	section4_LargeStructsAndPerformance()

	// ============================================================================
	// SECTION 5: Method Receivers - Value vs Pointer
	// ============================================================================
	section5_MethodReceivers()

	// ============================================================================
	// SECTION 6: Common Pointer Patterns
	// ============================================================================
	section6_CommonPointerPatterns()

	// ============================================================================
	// SECTION 7: Copying vs Sharing State
	// ============================================================================
	section7_CopyingVsSharingState()

	// ============================================================================
	// SECTION 8: Pointer Safety Features
	// ============================================================================
	section8_PointerSafetyFeatures()

	fmt.Println("\n🎉 Chapter 9 Complete! You understand Go pointers!")
}

// ============================================================================
// SECTION 1: Value vs Reference Types
// ============================================================================
func section1_ValueVsReferenceTypes() {
	fmt.Println("\n📚 SECTION 1: Value vs Reference Types")
	fmt.Println("----------------------------------------")

	// Value types (copied when passed around)
	fmt.Println("Value Types (copied when passed around):")
	
	// int is a value type
	age1 := 25
	age2 := age1  // This creates a COPY
	fmt.Printf("age1: %d, age2: %d\n", age1, age2)
	
	age2 = 30     // Changing age2 doesn't affect age1
	fmt.Printf("After changing age2: age1: %d, age2: %d\n", age1, age2)
	
	// float64 is a value type
	price1 := 19.99
	price2 := price1  // Copy
	price2 = 29.99    // Changing price2 doesn't affect price1
	fmt.Printf("price1: %.2f, price2: %.2f\n", price1, price2)
	
	// string is a value type
	name1 := "Alice"
	name2 := name1  // Copy
	name2 = "Bob"   // Changing name2 doesn't affect name1
	fmt.Printf("name1: %s, name2: %s\n", name1, name2)

	// Reference types (shared when passed around)
	fmt.Println("\nReference Types (shared when passed around):")
	
	// Slice is a reference type
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := numbers1  // This shares the SAME underlying array
	fmt.Printf("numbers1: %v, numbers2: %v\n", numbers1, numbers2)
	
	numbers2[0] = 100     // Changing numbers2 affects numbers1!
	fmt.Printf("After changing numbers2[0]: numbers1: %v, numbers2: %v\n", numbers1, numbers2)
	
	// Map is a reference type
	scores1 := map[string]int{"Alice": 95, "Bob": 87}
	scores2 := scores1    // Shares the same map
	fmt.Printf("scores1: %v, scores2: %v\n", scores1, scores2)
	
	scores2["Charlie"] = 92  // Adding to scores2 affects scores1!
	fmt.Printf("After adding to scores2: scores1: %v, scores2: %v\n", scores1, scores2)

	// Struct is a value type (but can contain reference types)
	fmt.Println("\nStructs are value types (but can contain reference types):")
	
	person1 := Person{
		Name:     "Alice",
		Age:      25,
		Hobbies:  []string{"reading", "swimming"},
		Metadata: map[string]string{"city": "New York"},
	}
	
	person2 := person1  // This creates a COPY of the struct
	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)
	
	// But the slices and maps inside are shared!
	person2.Hobbies[0] = "coding"           // Changes person1 too!
	person2.Metadata["city"] = "Boston"     // Changes person1 too!
	
	fmt.Printf("After modifying person2: person1: %+v\n", person1)
	fmt.Printf("After modifying person2: person2: %+v\n", person2)
}

// ============================================================================
// SECTION 2: Basic Pointer Syntax
// ============================================================================
func section2_BasicPointerSyntax() {
	fmt.Println("\n📚 SECTION 2: Basic Pointer Syntax")
	fmt.Println("-----------------------------------")

	// & operator (address-of)
	fmt.Println("& operator (address-of):")
	
	age := 25
	fmt.Printf("age: %d\n", age)
	fmt.Printf("age address: %p\n", &age)
	
	name := "Alice"
	fmt.Printf("name: %s\n", name)
	fmt.Printf("name address: %p\n", &name)
	
	// * operator (dereference)
	fmt.Println("\n* operator (dereference):")
	
	agePointer := &age
	fmt.Printf("agePointer: %p\n", agePointer)
	fmt.Printf("Value at agePointer: %d\n", *agePointer)
	
	// Modifying value through pointer
	fmt.Println("\nModifying value through pointer:")
	fmt.Printf("Before: age = %d\n", age)
	*agePointer = 26
	fmt.Printf("After: age = %d\n", age)
	
	// Creating pointers to different types
	fmt.Println("\nCreating pointers to different types:")
	
	// Method 1: Address operator
	price := 19.99
	pricePointer := &price
	fmt.Printf("price: %.2f, pricePointer: %p, *pricePointer: %.2f\n", 
		price, pricePointer, *pricePointer)
	
	// Method 2: new() function
	scorePointer := new(int)
	*scorePointer = 95
	fmt.Printf("scorePointer: %p, *scorePointer: %d\n", scorePointer, *scorePointer)
	
	// Method 3: Pointer to struct literal
	personPointer := &Person{Name: "Bob", Age: 30}
	fmt.Printf("personPointer: %p, *personPointer: %+v\n", personPointer, *personPointer)
	
	// Pointer arithmetic (or lack thereof in Go)
	fmt.Println("\nPointer arithmetic (or lack thereof in Go):")
	fmt.Println("In Go, you cannot do:")
	fmt.Println("- pointer++ (increment pointer)")
	fmt.Println("- pointer + 4 (add to pointer)")
	fmt.Println("- pointer - 2 (subtract from pointer)")
	fmt.Println("This makes Go pointers much safer than C/C++!")
}

// ============================================================================
// SECTION 3: Zero Values and Pointers
// ============================================================================
func section3_ZeroValuesAndPointers() {
	fmt.Println("\n📚 SECTION 3: Zero Values and Pointers")
	fmt.Println("----------------------------------------")

	// Zero values for different types
	fmt.Println("Zero values for different types:")
	
	var intValue int
	var stringValue string
	var boolValue bool
	var sliceValue []int
	var mapValue map[string]int
	var structValue Person
	var pointerValue *Person
	
	fmt.Printf("int zero value: %d\n", intValue)
	fmt.Printf("string zero value: %q\n", stringValue)
	fmt.Printf("bool zero value: %t\n", boolValue)
	fmt.Printf("slice zero value: %v (nil: %t)\n", sliceValue, sliceValue == nil)
	fmt.Printf("map zero value: %v (nil: %t)\n", mapValue, mapValue == nil)
	fmt.Printf("struct zero value: %+v\n", structValue)
	fmt.Printf("pointer zero value: %v (nil: %t)\n", pointerValue, pointerValue == nil)

	// Nil pointers
	fmt.Println("\nNil pointers:")
	
	var nilPointer *Person
	fmt.Printf("nilPointer: %v\n", nilPointer)
	fmt.Printf("Is nilPointer nil? %t\n", nilPointer == nil)
	
	// This would cause a panic if we tried to access it
	// nilPointer.Name = "Test"  // PANIC!
	
	// Safe way to handle nil pointers
	if nilPointer != nil {
		nilPointer.Name = "Test"
	} else {
		fmt.Println("Pointer is nil, cannot access fields")
	}
	
	// Creating non-nil pointers
	fmt.Println("\nCreating non-nil pointers:")
	
	// Method 1: Address of existing value
	person := Person{Name: "Alice", Age: 25}
	personPointer := &person
	fmt.Printf("personPointer: %v (nil: %t)\n", personPointer, personPointer == nil)
	
	// Method 2: new() function
	newPersonPointer := new(Person)
	fmt.Printf("newPersonPointer: %v (nil: %t)\n", newPersonPointer, newPersonPointer == nil)
	
	// Method 3: Pointer to struct literal
	literalPointer := &Person{Name: "Bob", Age: 30}
	fmt.Printf("literalPointer: %v (nil: %t)\n", literalPointer, literalPointer == nil)
}

// ============================================================================
// SECTION 4: Large Structs and Performance
// ============================================================================
func section4_LargeStructsAndPerformance() {
	fmt.Println("\n📚 SECTION 4: Large Structs and Performance")
	fmt.Println("----------------------------------------------")

	// Large struct definition
	fmt.Println("Large struct definition:")
	
	largeStruct := LargeStruct{
		ID:           "user123",
		Name:         "Alice Johnson",
		Email:        "alice.johnson@example.com",
		Age:          25,
		Phone:        "+1-555-0123",
		Address:      "123 Main Street, New York, NY 10001",
		City:         "New York",
		State:        "NY",
		ZipCode:      "10001",
		Country:      "USA",
		DateOfBirth:  "1998-05-15",
		JoinDate:     "2020-01-15",
		LastLogin:    "2024-01-15T10:30:00Z",
		IsActive:     true,
		IsVerified:   true,
		Preferences:  map[string]string{"theme": "dark", "language": "en"},
		Tags:         []string{"premium", "verified", "early-adopter"},
		Metadata:     map[string]interface{}{"source": "web", "campaign": "winter2024"},
	}
	
	fmt.Printf("Large struct size: %d bytes\n", estimateStructSize(largeStruct))
	
	// Performance comparison: value vs pointer
	fmt.Println("\nPerformance comparison: value vs pointer:")
	
	// Test with value receiver
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		largeStruct.GetInfoValue()
	}
	valueTime := time.Since(start)
	fmt.Printf("Value receiver time: %v\n", valueTime)
	
	// Test with pointer receiver
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		largeStruct.GetInfoPointer()
	}
	pointerTime := time.Since(start)
	fmt.Printf("Pointer receiver time: %v\n", pointerTime)
	
	// Show the difference
	if valueTime > pointerTime {
		improvement := float64(valueTime-pointerTime) / float64(valueTime) * 100
		fmt.Printf("Pointer receiver is %.1f%% faster\n", improvement)
	} else {
		improvement := float64(pointerTime-valueTime) / float64(pointerTime) * 100
		fmt.Printf("Value receiver is %.1f%% faster\n", improvement)
	}
	
	// When to use pointers for performance
	fmt.Println("\nWhen to use pointers for performance:")
	fmt.Println("✅ Use pointers when:")
	fmt.Println("  - Struct is large (>100 bytes)")
	fmt.Println("  - You need to modify the original")
	fmt.Println("  - Passing to multiple functions")
	fmt.Println("  - Avoiding unnecessary copying")
	
	fmt.Println("\n❌ Don't use pointers when:")
	fmt.Println("  - Struct is small (<100 bytes)")
	fmt.Println("  - You want to ensure immutability")
	fmt.Println("  - Working with simple types (int, string, bool)")
}

// ============================================================================
// SECTION 5: Method Receivers - Value vs Pointer
// ============================================================================
func section5_MethodReceivers() {
	fmt.Println("\n📚 SECTION 5: Method Receivers - Value vs Pointer")
	fmt.Println("--------------------------------------------------")

	// Create a person to work with
	person := Person{Name: "Alice", Age: 25, Email: "alice@example.com", IsActive: true}
	
	fmt.Println("Original person:")
	fmt.Printf("  %+v\n", person)
	
	// Value receiver methods (don't modify original)
	fmt.Println("\nValue receiver methods (don't modify original):")
	
	person.Introduce()  // This method has a value receiver
	fmt.Printf("After Introduce(): %+v\n", person)  // Still same!
	
	person.GetAge()     // This method has a value receiver
	fmt.Printf("After GetAge(): %+v\n", person)     // Still same!
	
	// Pointer receiver methods (modify original)
	fmt.Println("\nPointer receiver methods (modify original):")
	
	person.HaveBirthday()  // This method has a pointer receiver
	fmt.Printf("After HaveBirthday(): %+v\n", person)  // Age changed!
	
	person.UpdateEmail("alice.new@example.com")  // This method has a pointer receiver
	fmt.Printf("After UpdateEmail(): %+v\n", person)   // Email changed!
	
	// Demonstrating the difference
	fmt.Println("\nDemonstrating the difference:")
	
	person2 := Person{Name: "Bob", Age: 30, Email: "bob@example.com", IsActive: true}
	fmt.Printf("Original person2: %+v\n", person2)
	
	// Value receiver - creates a copy
	person2.Introduce()  // Works on a copy
	fmt.Printf("After value receiver method: %+v\n", person2)  // Still same!
	
	// Pointer receiver - works on original
	person2.HaveBirthday()  // Works on original
	fmt.Printf("After pointer receiver method: %+v\n", person2)  // Age changed!
	
	// When to use each type
	fmt.Println("\nWhen to use each type of receiver:")
	fmt.Println("✅ Use value receivers when:")
	fmt.Println("  - Method doesn't modify the struct")
	fmt.Println("  - Struct is small")
	fmt.Println("  - You want to ensure immutability")
	fmt.Println("  - Method is read-only")
	
	fmt.Println("\n✅ Use pointer receivers when:")
	fmt.Println("  - Method modifies the struct")
	fmt.Println("  - Struct is large")
	fmt.Println("  - You need to modify the original")
	fmt.Println("  - Method changes state")
}

// ============================================================================
// SECTION 6: Common Pointer Patterns
// ============================================================================
func section6_CommonPointerPatterns() {
	fmt.Println("\n📚 SECTION 6: Common Pointer Patterns")
	fmt.Println("----------------------------------------")

	// Builder pattern with pointers
	fmt.Println("Builder pattern with pointers:")
	
	computer := NewComputerBuilder().
		SetCPU("Intel i9").
		SetRAM(32).
		SetStorage("1TB NVMe").
		SetGPU("RTX 4080").
		Build()
	
	fmt.Printf("Built computer: %+v\n", computer)
	
	// Factory pattern with pointers
	fmt.Println("\nFactory pattern with pointers:")
	
	// Create different types of shapes
	circle := NewShape("circle", 5.0)
	square := NewShape("square", 4.0)
	triangle := NewShape("triangle", 3.0)
	
	fmt.Printf("Circle: %+v, Area: %.2f\n", circle, circle.Area())
	fmt.Printf("Square: %+v, Area: %.2f\n", square, square.Area())
	fmt.Printf("Triangle: %+v, Area: %.2f\n", triangle, triangle.Area())
	
	// Singleton pattern (using pointers)
	fmt.Println("\nSingleton pattern (using pointers):")
	
	config1 := GetConfig()
	config2 := GetConfig()
	
	fmt.Printf("config1 address: %p\n", config1)
	fmt.Printf("config2 address: %p\n", config2)
	fmt.Printf("Are they the same instance? %t\n", config1 == config2)
	
	// Modify config through one reference
	config1.SetTheme("dark")
	fmt.Printf("config1 theme: %s\n", config1.GetTheme())
	fmt.Printf("config2 theme: %s\n", config2.GetTheme())  // Same instance!
	
	// Linked list pattern
	fmt.Println("\nLinked list pattern:")
	
	list := &LinkedList{Value: "first"}
	list.Next = &LinkedList{Value: "second"}
	list.Next.Next = &LinkedList{Value: "third"}
	
	// Print the list
	current := list
	for current != nil {
		fmt.Printf("Node: %s -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
	
	// Add a new node
	newNode := &LinkedList{Value: "fourth"}
	list.Next.Next.Next = newNode
	
	// Print again
	current = list
	for current != nil {
		fmt.Printf("Node: %s -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// ============================================================================
// SECTION 7: Copying vs Sharing State
// ============================================================================
func section7_CopyingVsSharingState() {
	fmt.Println("\n📚 SECTION 7: Copying vs Sharing State")
	fmt.Println("----------------------------------------")

	// Deep vs shallow copying
	fmt.Println("Deep vs shallow copying:")
	
	// Original person with slices and maps
	originalPerson := Person{
		Name:     "Alice",
		Age:      25,
		Hobbies:  []string{"reading", "swimming"},
		Metadata: map[string]string{"city": "New York", "country": "USA"},
	}
	
	fmt.Printf("Original person: %+v\n", originalPerson)
	
	// Shallow copy (shares slices and maps)
	shallowCopy := originalPerson
	fmt.Printf("Shallow copy: %+v\n", shallowCopy)
	
	// Modify the shallow copy
	shallowCopy.Hobbies[0] = "coding"
	shallowCopy.Metadata["city"] = "Boston"
	
	fmt.Printf("After modifying shallow copy:\n")
	fmt.Printf("  Original: %+v\n", originalPerson)   // Changed!
	fmt.Printf("  Shallow:  %+v\n", shallowCopy)
	
	// Deep copy (completely independent)
	fmt.Println("\nCreating deep copy:")
	deepCopy := originalPerson.DeepCopy()
	fmt.Printf("Deep copy: %+v\n", deepCopy)
	
	// Modify the deep copy
	deepCopy.Hobbies[0] = "painting"
	deepCopy.Metadata["city"] = "Los Angeles"
	
	fmt.Printf("After modifying deep copy:\n")
	fmt.Printf("  Original: %+v\n", originalPerson)   // Unchanged!
	fmt.Printf("  Deep:     %+v\n", deepCopy)
	
	// Pointer sharing
	fmt.Println("\nPointer sharing:")
	
	personPointer := &originalPerson
	sharedPointer := personPointer  // Both point to the same person
	
	fmt.Printf("personPointer: %p\n", personPointer)
	fmt.Printf("sharedPointer: %p\n", sharedPointer)
	fmt.Printf("Are they the same? %t\n", personPointer == sharedPointer)
	
	// Modify through one pointer
	personPointer.Age = 26
	fmt.Printf("Modified through personPointer: %+v\n", *personPointer)
	fmt.Printf("Accessed through sharedPointer: %+v\n", *sharedPointer)  // Same data!
}

// ============================================================================
// SECTION 8: Pointer Safety Features
// ============================================================================
func section8_PointerSafetyFeatures() {
	fmt.Println("\n📚 SECTION 8: Pointer Safety Features")
	fmt.Println("----------------------------------------")

	// No pointer arithmetic
	fmt.Println("No pointer arithmetic:")
	fmt.Println("In Go, you cannot do:")
	fmt.Println("- pointer++ (increment pointer)")
	fmt.Println("- pointer + 4 (add to pointer)")
	fmt.Println("- pointer - 2 (subtract from pointer)")
	fmt.Println("This makes Go pointers much safer than C/C++!")
	
	// Automatic dereferencing
	fmt.Println("\nAutomatic dereferencing:")
	
	person := &Person{Name: "Alice", Age: 25}
	
	// Go automatically dereferences when accessing fields
	fmt.Printf("person.Name: %s\n", person.Name)      // Same as (*person).Name
	fmt.Printf("person.Age: %d\n", person.Age)        // Same as (*person).Age
	
	// But you still need * for assignment
	person.Age = 26  // This works
	// person = Person{Name: "Bob", Age: 30}  // This would change the pointer
	
	// Nil pointer safety
	fmt.Println("\nNil pointer safety:")
	
	var nilPointer *Person
	fmt.Printf("nilPointer: %v\n", nilPointer)
	
	// Safe way to check before using
	if nilPointer != nil {
		fmt.Printf("Name: %s\n", nilPointer.Name)
	} else {
		fmt.Println("Pointer is nil, cannot access fields")
	}
	
	// Garbage collection
	fmt.Println("\nGarbage collection:")
	fmt.Println("Go automatically manages memory:")
	fmt.Println("- No manual memory allocation/deallocation")
	fmt.Println("- No memory leaks from forgotten pointers")
	fmt.Println("- Automatic cleanup when pointers go out of scope")
	
	// Example of automatic cleanup
	fmt.Println("\nExample of automatic cleanup:")
	{
		// This person will be automatically cleaned up when this block ends
		tempPerson := &Person{Name: "Temp", Age: 100}
		fmt.Printf("Created temporary person: %+v\n", *tempPerson)
	} // tempPerson is automatically cleaned up here
	
	fmt.Println("Block ended, temporary person cleaned up")
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
	Hobbies  []string
	Metadata map[string]string
}

// LargeStruct represents a large data structure
type LargeStruct struct {
	ID          string
	Name        string
	Email       string
	Age         int
	Phone       string
	Address     string
	City        string
	State       string
	ZipCode     string
	Country     string
	DateOfBirth string
	JoinDate    string
	LastLogin   string
	IsActive    bool
	IsVerified  bool
	Preferences map[string]string
	Tags        []string
	Metadata    map[string]interface{}
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

// Config represents application configuration
type Config struct {
	Theme     string
	Language  string
	Timezone  string
	DebugMode bool
}

// LinkedList represents a linked list node
type LinkedList struct {
	Value string
	Next  *LinkedList
}

// Person methods
func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

func (p Person) GetAge() int {
	return p.Age
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("Happy birthday! %s is now %d years old.\n", p.Name, p.Age)
}

func (p *Person) UpdateEmail(newEmail string) {
	p.Email = newEmail
	fmt.Printf("Email updated to: %s\n", p.Email)
}

func (p Person) DeepCopy() Person {
	// Deep copy the slices and maps
	hobbiesCopy := make([]string, len(p.Hobbies))
	copy(hobbiesCopy, p.Hobbies)
	
	metadataCopy := make(map[string]string)
	for k, v := range p.Metadata {
		metadataCopy[k] = v
	}
	
	return Person{
		Name:     p.Name,
		Age:      p.Age,
		Email:    p.Email,
		IsActive: p.IsActive,
		Hobbies:  hobbiesCopy,
		Metadata: metadataCopy,
	}
}

// LargeStruct methods
func (ls LargeStruct) GetInfoValue() string {
	return fmt.Sprintf("User %s (%s) is %d years old", ls.Name, ls.Email, ls.Age)
}

func (ls *LargeStruct) GetInfoPointer() string {
	return fmt.Sprintf("User %s (%s) is %d years old", ls.Name, ls.Email, ls.Age)
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

// Shape factory
func NewShape(shapeType string, size float64) *Shape {
	switch shapeType {
	case "circle":
		return &Shape{Type: "circle", Radius: size}
	case "square":
		return &Shape{Type: "square", Side: size}
	case "triangle":
		return &Shape{Type: "triangle", Base: size, Height: size}
	default:
		return &Shape{Type: "unknown"}
	}
}

// Config singleton
var configInstance *Config

func GetConfig() *Config {
	if configInstance == nil {
		configInstance = &Config{
			Theme:     "light",
			Language:  "en",
			Timezone:  "UTC",
			DebugMode: false,
		}
	}
	return configInstance
}

func (c *Config) SetTheme(theme string) {
	c.Theme = theme
}

func (c *Config) GetTheme() string {
	return c.Theme
}

// Helper functions
func estimateStructSize(s LargeStruct) int {
	// Rough estimation of struct size
	size := 0
	size += len(s.ID) + len(s.Name) + len(s.Email) + len(s.Phone)
	size += len(s.Address) + len(s.City) + len(s.State) + len(s.ZipCode)
	size += len(s.Country) + len(s.DateOfBirth) + len(s.JoinDate) + len(s.LastLogin)
	size += len(s.Preferences) * 20  // Rough estimate for map
	size += len(s.Tags) * 10         // Rough estimate for slice
	size += len(s.Metadata) * 30     // Rough estimate for interface{} map
	size += 50  // Fixed fields (int, bool, etc.)
	return size
} 