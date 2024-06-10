package chapter_1

// You can only import whole libraries, not files or functions.
// The compiler sorts it all out for you.
import "fmt"

func HelloWorld() {
	// Sup bitches this is a hello world!
	fmt.Println("Hello, World!")
}

func DodgyString() {
	fmt.Printf("My special string is: %s")
}

// chapter_1/notes.go:13:2: fmt.Printf format %s reads arg #1, but call has 0 args
