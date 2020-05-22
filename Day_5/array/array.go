package main

import "fmt"

func main() {
	// init array
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for i := 0; i < len(notes); i++ { // unsafe looping over array
		fmt.Println(i, notes[i])
	}

	for index, value := range notes { // safe looping over array
		fmt.Println(index, value)
	}

	for _, value := range notes { // ignore index
		fmt.Println(value)
	}

	for index, _ := range notes { // ignore value
		fmt.Println(index)
	}

	for index := range notes { // ignore value, can do like this too
		fmt.Println(index)
	}
}
