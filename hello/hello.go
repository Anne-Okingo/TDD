package main

import(
	"fmt"
)

const Deutschprefix = "Hallo, "
const deutsch = "Deutsch"
func main(){
	str := Hello("Hannah", "Deutsch")
	
	fmt.Println(str)
}

func Hello(s, lang string)string{
	 name := "Hello, "
	if s == ""{
		s = "World"
	}
	if lang == deutsch{
		name = Deutschprefix
	}
	return name + s
}