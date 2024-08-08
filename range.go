package main

import "fmt"

// iterating over data structure
func main() {
	// nums := []int {6,7,8}

	// sum :=0

	// for indx,num := range nums{     //iterating over slices
	//sum = sum+num
	//fmt.Println(sum)
	// 	fmt.Println(indx,num)
	// }

	// iterating over map
	// m :=map[string]string{"fname":"jhon","lname":"doe"}

	// for k,v :=range m{
	// 	fmt.Println(k,v)
	// }

	// m :=map[string]string{"fname":"jhon","lname":"doe"}

	// for k :=range m{
	// 	fmt.Println(k)
	// }

	// String
	for i, c := range "golang" { // i is a starting byte of rune ||  below 255 -> 1byte
		// fmt.Println(i, c) // c is a unicode code point rune
		fmt.Println(i, string(c)) 

}
}
