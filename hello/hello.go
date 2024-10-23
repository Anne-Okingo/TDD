package main

import(
	"fmt"
)

const name = "Hello, "
func main(){
	str := Hello("Hannah")

	fmt.Println(str)
}

func Hello(s string)string{
	if s == ""{
		s = "World"
	}
	return name + s
}