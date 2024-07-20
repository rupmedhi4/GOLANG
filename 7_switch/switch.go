package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("others")
	}


	//multiple condition switch

	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("its weekend")
		default :
			fmt.Println("its weekday")

	}


	//type switch
	whoAmI :=func(i interface{}){    //interface{} means any type of variable data received
		switch t := i.(type){
			case int:
				fmt.Println("i is int",t)
			case string:
				fmt.Println("i is string",t)
			case bool:
				fmt.Println("i is bool",t)
			default:
				fmt.Println("i is something else",t)
		}
	}

	whoAmI("bkj")
}