package main

import "fmt"

// Liters type
type Liters float64

// Mililiters type
type Mililiters float64

// Gallons type
type Gallons float64

// ToMililiters method
func (l Liters) ToMililiters() Mililiters {
	return Mililiters(l * 1000)
}

// ToLiters method
func (m Mililiters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	l := Liters(3)
	fmt.Printf("%.1f liters is %.1f mililiters\n", l, l.ToMililiters())
	ml := Mililiters(500)
	fmt.Printf("%.1f mililiters is %.1f liters\n", ml, ml.ToLiters())
}
