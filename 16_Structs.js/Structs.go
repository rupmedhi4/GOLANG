package main

import (
	"fmt"
	"time"
)

// order structs
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
}

func newOrder(id string,amount float32,status string )*order{
	//initial setup goes here...
	myOrder :=order{
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string){
	o.status = status
}

func main(){


	myOrder :=newOrder("1",30.40,"received")
	fmt.Println(myOrder)
	
	// myOrder :=order{
	// 	id: "1234",
	// 	amount: 100.0,
	// 	status: "pending",		
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)

	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)
	// fmt.Println(" order structs" ,myOrder)

	// myOrder2 := order{
	// 	id:"2",
	// 	amount: 200.0,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }


	// fmt.Println(myOrder2)



} 
