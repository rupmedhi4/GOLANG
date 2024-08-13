package main

import "fmt"

//by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("in changeNum", num)
// }


//by reference
func changeNum(num *int) {
	*num = 5 //* for pass refernces
	fmt.Println("in changeNum", num)
}

func main() {
	num :=1 

	changeNum(&num)  //& for memory addrerss

	fmt.Println("after changeNum in main",* num)
}