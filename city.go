package main

import (
	"fmt"
	"math/rand"

	"github.com/satori/go.uuid"
)

type City struct {
	Corners [][]*Corner
}

var names = []string{
	"John Smith",
	"Bob James",
	"Clarice Pantalones",
	"Shitty McShitterson",
	"Billy Buttfuck",
	"James Sloose",
	"Joe Ribz",
	"Tim Bellum",
	"Daniel Doubledown",
	"Pinky Ray",
	"Vuvu Zela",
	"Edwardoe Parsons",
	"Jack Endall",
	"Liesel Diesel",
	"Gotobed Joy",
	"Wheres Myphone",
	"Donald Drumpf",
	"Bernard Sand0rs",
	"Lollery Clinton",
	"Casey Pants",
	"Ducky",
}

func (city *City) Populate(sizeX, sizeY int) {
	city.Corners = make([][]*Corner, sizeX)

	for x, _ := range city.Corners {
		city.Corners[x] = make([]*Corner, sizeY)
		for y, _ := range city.Corners[x] {
			city.Corners[x][y] = &Corner{
				Street:    fmt.Sprintf("%d/%d Street", x, y),
				Rating:    float64(rand.Int()%10.0) * float64(0.1),
				LocationX: x,
				LocationY: y,
				City:      city,
			}
		}
	}

}

func (city *City) Corner(x, y int) *Corner {
	return city.Corners[x][y]
}

func (city *City) AllCorners() <-chan *Corner {
	corners := make(chan *Corner)
	go func() {
		for x, _ := range city.Corners {
			for y, _ := range city.Corners[x] {
				corners <- city.Corners[x][y]
			}
		}
		close(corners)
	}()
	return corners
}

func (city *City) AllJunkies() <-chan *User {
	users := make(chan *User)
	go func() {
		for x, _ := range city.Corners {
			for y, _ := range city.Corners[x] {
				for _, user := range city.Corners[x][y].Users {
					users <- user
				}
			}
		}
		close(users)
	}()
	return users
}

func (city *City) PopulateJunkies() {
	if city.Corners == nil {
		return
	}

	for corner := range city.AllCorners() {
		if corner.Users == nil {
			corner.Users = make(map[uuid.UUID]*User, 1)
			user := NewUser(names[rand.Int()%len(names)], 1, 18, corner)
			corner.Users[user.Id] = user
		}
	}
}

func (city *City) UpdateJunkies() {
	if city.Corners == nil {
		return
	}
	moved := 0
	for junkie := range city.AllJunkies() {
		if junkie.LastFix() > 1000*5 && junkie.LastMovedSeconds() > 10 {
			// 1 in 4 chance of moving
			if rand.Int()%4 == 0 {
				moved++
				//fmt.Printf("MOVE: %s (%d,%d)\n", junkie.Name, junkie.CurrentCorner.LocationX, junkie.CurrentCorner.LocationY)
				junkie.RandomMove()
			}
		}
		junkie.Tick()
	}
	//fmt.Printf("Moved: %d\n", moved)
}
