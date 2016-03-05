package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chzyer/readline"
)

func main() {
	city := City{}
	city.Populate(10, 10)
	city.PopulateJunkies()

	//junkie := city.Corners[0][0].Users[0]

	go func() {
		for {
			city.UpdateJunkies()
			//if junkie.NeedsFix() && junkie.Alive {
			//	junkie.Use(4)
			//}
			time.Sleep(1 * time.Second)
		}
	}()

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	done := false
	for !done {
		line, err := rl.Readline()
		if err != nil {
			fmt.Printf("Error: %v", err)
			break
		}
		tokens := strings.Split(strings.ToLower(line), " ")
		switch tokens[0] {
		case "list":
			x, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("X value must be an integer.")
				break
			}
			y, err := strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println("Y value must be an integer.")
				break
			}
			corner := city.Corner(x, y)
			fmt.Printf("%s (%d,%d)\n", corner.Street, corner.LocationX, corner.LocationY)
			fmt.Println(" Users:")
			for _, user := range corner.Users {
				fmt.Printf("   %s (%s)\n", user.Name, user.Id)
			}
		case "exit":
			done = true
		}
	}
}
