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
		//if junkie.NeedsFix() && junkie.Alive {
		//	junkie.Use(4)
		//}
		if junkie.LastFix() > 1000*5 && junkie.LastMovedSeconds() > 10 {
			junkie.RandomMove()
		}
		time.Sleep(1 * time.Second)
	}
}
