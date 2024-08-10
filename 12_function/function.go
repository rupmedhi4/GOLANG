package main

import "fmt"

// func add(a , b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string) {
// 	return "golang", "js", "c"
// }

func processIt(fn func(a int) int ){
	fn(1)
}

func main() {
	// result := add(3, 5)
	//fmt.Println(result)

	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2 )

	fn:=func(a int) int{
		return 2
	}
	processIt(fn)
}
