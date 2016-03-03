package main

import (
	"fmt"
	"time"
)

func main() {
	city := City{}
	city.Populate(10, 10)
	city.PopulateJunkies()

	junkie := city.Corners[0][0].Users[0]

	for {
		city.UpdateJunkies()
		fmt.Println(junkie)
		if junkie.NeedsFix() && junkie.Alive {
			junkie.Use(4)
		}
		time.Sleep(1 * time.Second)
	}
}
