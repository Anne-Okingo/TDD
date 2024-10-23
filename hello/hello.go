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
	return name + s
}