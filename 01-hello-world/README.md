# Go Hello World: Your First Program 🐹

Welcome to Go programming! This is your very first program.

## What You'll Learn

- How to write a simple Go program
- What `package main` means
- What `func main()` does
- How to run your program

## Your First Go Program

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go World!")
	fmt.Println("Welcome to Go programming!")
}
```

## What Each Part Does

### `package main`
- Every Go program starts with a package
- `main` means this is the main program (not a library)
- You need this line in every Go program

### `import "fmt"`
- `fmt` is a package that helps you print text
- `import` brings it into your program
- Without this, you can't use `fmt.Println()`

### `func main()`
- `func` means "function" (a block of code)
- `main` is the special name - Go looks for this function first
- `main()` is where your program starts running
- Everything inside the `{ }` runs when you start the program

### `fmt.Println()`
- Prints text on the screen
- Adds a new line after the text
- Put your message in quotes: `"Hello"`

## How to Run Your Program

1. Open your terminal/command prompt
2. Go to the hello-world folder: `cd 01-hello-world`
3. Run the program: `go run main.go`

**What you'll see:**
```
Hello, Go World!
Welcome to Go programming!
```

## Try It Yourself!

1. Change the text in the quotes
2. Add more `fmt.Println()` lines
3. Run it again with `go run main.go`

## Next Steps

After this, you'll learn about:
- Variables (storing information)
- Control structures (making decisions)
- Functions (reusable code)

---

**Congratulations! You've written your first Go program! 🎉** 