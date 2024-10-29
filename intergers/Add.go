package main

import(
	"fmt"
)

func main(){
n1,n2 := 0,0
result := Sum(n1,n2)

fmt.Println(result)
}

func Sum(a, b int)int{
	return a + b
}