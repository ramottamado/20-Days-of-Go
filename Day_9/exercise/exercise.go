package main

import "fmt"

// Number type.
type Number int

// Add method.
func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

//Subtract method.
func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}
