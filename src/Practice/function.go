package main 

import "fmt"

func plus(a int, b int) int {
	
	fmt.Println("Hello")
	return a+b
}


func plusPlus(a, b, c int) int {

	return a+b+c
}

func main() {

	res := plus(1,2)
	fmt.Println("Result of addition: ", res)

	res := plusPlus(1,2,3)
	fmt.Println("Result of addition: ", res)

}