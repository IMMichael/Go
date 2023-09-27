package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("The time now is:", n)
	//fmt.Printf("2nd line")

	t := time.Date(2009, time.April, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("The date is:", t)

	fmt.Println(t.Format(time.DateOnly))
	fmt.Println(t.Format(time.TimeOnly))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Print(t.Format("01/02/2006"))
}
