package main

import "fmt"

// Population type.
type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congratulations, Kevin and Anna! It's a girl!")
	population++
	fmt.Println("Sleep Creek County population:", population)
}
