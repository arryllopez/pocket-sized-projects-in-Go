package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	var bookworms []Bookworm

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil

}

func findCommonBooks(bookworms []Bookworm) []Book {
	return nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	// increment the count for each book in each bookworm's collection
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms) 
	// declaring a slice to hold the books that were found multiple times in teh bookworms' collections
	// slices in golang are dynamically sized, unlike arrays which have a fixed size . 
	// slices have three fields -- a pointer to the underlying array, the length of the slice, and the capacity of the slice.
	// length is the number of elements within the slice
	// and the capacity is how much elements the slice can store before dynamically resizing itself.
	// len and cap functions can be used to get the length and capacity of a slice, respectively, only when the slice is initialized with make
	var commonBooks []Book

	// 

	return nil
} 