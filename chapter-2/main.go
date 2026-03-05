package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"en",
		"The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
// type safety and it enhances readabiolity
type language string

// impementing a map <language,phrase>
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in the specified language
// the inpuyt language will be a string that represents a language
func greet(l language) string {
	greeting, ok := phrasebook[l] // returns the message associated with the key language l and a boolean that tells us if the key was found
	if !ok {                      // if ok is false
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
