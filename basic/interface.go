// When we don't know the type of variable then we can use the Interface
package main

import "fmt"

func main() {
	val1 := "Hello, world!"
	val2 := 123
	var test interface{}
	var test1 interface{}

	test = val1
	test1 = val2
	fmt.Println(test)
	fmt.Println(test1)
}
