// Closure functions
package main

import "fmt"

func main() {
	count, sum := createCounterAndSum()
	fmt.Println(count(), sum())
	fmt.Println(count(), sum())
	fmt.Println(count(), sum())
	fmt.Println(count(), sum())

}

// Function declarations

// func test() {
// 	// Function body
// }

// // Anonymous functions
// func() {
//     // Function body
// }()

// Function with closure
// func createCounter() func() int {
//     count := 0
//     return func() int {
//         count++
//         return count
//     }
// }

// Function with closure and multiple return values
func createCounterAndSum() (func() int, func() int) {
	count := 0
	sum := 0
	return func() int {
			count++
			sum += count
			return count
		}, func() int {
			return sum
		}
}
