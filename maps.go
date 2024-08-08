package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map
	// m := make(map[string]string)

	 //setting an elements
	// m["name"] = "golang"
	// m["area"] = "backend"

	 //get an element
	// fmt.Println(m["name"])

	 //IMP: if key  does not exists in the map than it returns zero value
	// fmt.Println((m["phone"]))

	// delete(m,"name")
	// clear(m)


	//m := map[string]int{"price":23,"phone":3}

	// fmt.Println(m)

	//to check elements exists or not in the map
//    v, ok :=m["price"]

//    fmt.Println(v)
//    if ok{
// 	fmt.Println("all ok")
//    }else{
// 	fmt.Println("not ok")
//    }


m1 :=map[string]int{"price":80,"phones":7}
m2 :=map[string]int{"price":80,"phones":7}

fmt.Println(maps.Equal(m1,m2))

}
 