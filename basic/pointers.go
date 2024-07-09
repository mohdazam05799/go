// Pointer is being used to store the address of the variable in memory

package main

import "fmt"

func main() {

	a := 10
	var s *int = &a
	fmt.Println("Address of a: ", s)
}
