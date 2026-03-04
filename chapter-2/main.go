package main // packages are go's ways of organizing code similar to modules or libraries in other languages
// the main() function can only be in the main package

import "fmt"

func main() {
	fmt.Println("Hello world") //Println is the functino within the fmt package that writes to the std output
	// Println starts with capital P sicne any symbol starting with capital letters Uppercase → exported (accessible from other packages)
	// if its lowercase, its not accesible outside of the package
}

// // some notes for idiomacy
// if scope of variable is limited to two to three lines, one or two letter placeholders are fine
// consistency between functions
// go is camelCase
