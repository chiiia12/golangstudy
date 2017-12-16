package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%p\n", s)     //0xc420016210
	fmt.Printf("%p\n", s[:2]) //0xc420016210
	fmt.Printf("%p\n", &s[1])
	fmt.Printf("%p\n", s[1:2])
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
	s = []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	//struct sample
	var dilbert Employee
	dilbert.Salary -= 5000
	fmt.Println("salary is ", dilbert.Salary)
	position := &dilbert.Position
	fmt.Println(position)

	type Point struct{ X, Y int }
	p := Point{1, 2}
	q := Point{1, 2}
	fmt.Println(p == q)   //true
	fmt.Println(&p == &q) //false
	fmt.Println(&p == &p) //true

}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] //移動してるだけ
	return slice[:len(slice)-1]
}
