package main

import "fmt"

func main() {
	window := Window{
		Components: []Item{
			Glass{Cost: 50},
			Glass{Cost: 50},
			Frame{Cost: 30},
		},
	}

	fmt.Println(window.Display())
	fmt.Printf("Total price: $%.2f\n", window.Price())
}
