package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Englishprefix = "Hello, "
	Spanishprefix = "Hola, "
	Deutschprefix = "Hallo, "
	deutsch       = "deutsch"
	spanish       = "spanish"
	english       = "english"
)

func main(){

	if len( os.Args) <=1{
		fmt.Println("Hello, World\nPlease enter your name and langunage")
		return
	}

	name := os.Args[1]

	if len(os.Args) <= 2{
		s := Englishprefix + name

		fmt.Println(s)
		return
	}
	lang := os.Args[2]

	str := Hello(name, lang)
	fmt.Println(str)
}

func Hello(name, lang string) string {
	if name == ""{
		return "Hello, World"
	}
	lang = strings.ToLower(lang)

	// Use a switch statement for the greeting based on the language
	var prefix string
	switch lang {
	case english:
		prefix = Englishprefix
	case deutsch:
		prefix = Deutschprefix
	case spanish:
		prefix = Spanishprefix
	default:
		prefix = Englishprefix // Default to English if no match
	}

	return prefix + name
}