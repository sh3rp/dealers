package main

import (
	"fmt"
	"math/rand"
)

type Corner struct {
	Street  string
	Rating  float64 // 0.1 = ghetto, 0.9 = park avenue
	Dealers []*Dealer
	Users   []*User
}

type Dealer struct {
	Cash  int
	Drugs map[Drug]int
}

type Drug struct {
	Label string
	Type  int // 1 = marijuana, 2 = cocaine, 3 = lsd, 4 = heroin
}

type User struct {
	Name           string
	DrugOfChoice   int     // 1 = marijuana, 2 = cocaine, 3 = lsd, 4 = heroin
	Susceptibility float64 // 0.1 = casual user, 0.9 = full blown junkie
	CurrentHigh    float64 // 0.0 = clean and sober, 0.9 = high as fuck
	LastUsed       int64
}

type City struct {
	Corners [][]*Corner
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

func (city *City) PopulateJunkies() {
	if city.Corners == nil {
		return
	}

	for corner := range city.AllCorners() {
		if corner.Users == nil {
			corner.Users = make([]*User, 1)
			corner.Users[0] = &User{
				Name: "Beavis",
			}
		}
	}
}
