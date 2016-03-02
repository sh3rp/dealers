package main

import "fmt"

func main() {
	city := City{}
	city.Populate(10, 10)

	for corner := range city.AllCorners() {
		fmt.Println(corner)
	}
}
