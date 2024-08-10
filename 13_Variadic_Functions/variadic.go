package main

import "fmt"


func sum(nums ...int) int { // sum function accept only int value but if we use interface{} then it takes any type 
//func sum(nums ...interface{}) int {
	total :=0

	for _,num:=range nums{
		total =total+num
	}

	return total
}
func main() {
	result :=sum(3,4,5,6,7)
	fmt.Println((result))
}