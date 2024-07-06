package main

import "fmt"

type Student struct {
	Name       string
	Age        int
	RollNumber int
}

func main() {
	student1 := Student{Name: "John", Age: 13, RollNumber: 1234}
	student2 := Student{Name: "Tom", Age: 15, RollNumber: 12334}
	student_details(student1)
	student_details(student2)
}

func student_details(detail Student) {
	fmt.Println("Student Name: ", detail.Name)
	fmt.Println("Student Age: ", detail.Age)
	fmt.Println("Student Roll Number: ", detail.RollNumber)

}
