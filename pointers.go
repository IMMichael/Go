package main

import (
"fmt"
)

func main () {
	fmt.Println("Pointers")
	anInt := 42  //Value of integer anInt is 42.
	var p = &anInt //We declared that the var 'p' is a pointer to the memory address of anInt - that is it's pointing towards the memory address of anInt.
	fmt.Println("Value of p is: ", *p) // Here we asked for the contents of the memory address pointed out by p to be printed on screen. Ergo, what does this slot in the memory contains.

	aFloat1 := 42.21
	pointer1 := &aFloat1
	fmt.Println("The value of pointer1 is:", *pointer1)
	*pointer1 = *pointer1 / 42.21						// Here we have proven that changing the value which is being pointed also chnages the value of the initial variable (in this case aFloat1)
	fmt.Println("The value of pointer1 now is:", *pointer1)
	fmt.Println("The value of aFloat1 is:", aFloat1)
}