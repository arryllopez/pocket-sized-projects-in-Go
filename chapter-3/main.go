package main

import (
	"fmt"
	"os"
)

func main() {
	// json formats store information in key:value
	// json values can be
	// decimal numbers
	// strings
	// arrays
	// booleans
	// objects

	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms : %s\n", err)
		os.Exit(1)
	}
	fmt.Println(bookworms)
}
