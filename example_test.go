package u8p_test

import (
	"fmt"

	"github.com/catatsuy/u8p"
)

func ExampleFind() {
	// Example with Japanese characters
	a := "こんにちは世界"
	la := 10
	index, err := u8p.Find(a, la)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Japanese example - %s\n", a[:index])
	}

	// Example with an emoji
	b := "Hello, 🌍. Hi!"
	lb := 13
	index, err = u8p.Find(b, lb)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Emoji example - %s\n", b[:index])
	}

	// Output:
	// Japanese example - こんに
	// Emoji example - Hello, 🌍.
}
