package main

import(
"fmt"
"os"
"strconv"
)


func main(){
	str := os.Args[1]
	num, _:= strconv.Atoi(os.Args[2])

	result := Repeat(str,num)
	fmt.Println(result)
}
func Repeat(str string, num int)string{
	result := ""
	for i := 0; i < num; i++{
		result += str
	}
	return result
}