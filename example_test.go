package u8p_test

import (
	"fmt"

	"github.com/catatsuy/u8p"
)

func ExampleFind() {
	// Example with Japanese characters
	a := "こんにちは世界"
	la := 10 // Consider the last 10 bytes of the string
	index, err := u8p.Find(a, la)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Japanese example - Index of UTF-8 lead byte: %d\n", index)
	}

	// Example with an emoji
	b := "Hello, 🌍"
	lb := 8 // Consider the last 8 bytes of the string
	index, err = u8p.Find(b, lb)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Emoji example - Index of UTF-8 lead byte: %d\n", index)
	}

	// Output:
	// Japanese example - Index of UTF-8 lead byte: 9
	// Emoji example - Index of UTF-8 lead byte: 7
}
