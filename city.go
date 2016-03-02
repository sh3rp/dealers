package main

import (
	"fmt"
	"math/rand"
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
}

func (city *City) Populate(sizeX, sizeY int) {
	city.Corners = make([][]*Corner, sizeX)

	for x, _ := range city.Corners {
		city.Corners[x] = make([]*Corner, sizeY)
		for y, _ := range city.Corners[x] {
			city.Corners[x][y] = &Corner{
				Street: fmt.Sprintf("%d/%d Street", x, y),
				Rating: float64(rand.Int()%10.0) * float64(0.1),
			}
		}
	}

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
			corner.Users = make([]*User, 1)
			corner.Users[0] = &User{
				Name:           names[rand.Int()%len(names)],
				Susceptibility: float64(rand.Int()%100) * float64(0.01),
				CurrentHigh:    float64(0),
			}
		}
	}
}

func (city *City) UpdateJunkies() {
	if city.Corners == nil {
		return
	}

	for junkie := range city.AllJunkies() {
		fmt.Println(junkie)
	}
}
