package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/satori/go.uuid"
)

type Corner struct {
	Street    string
	Rating    float64 // 0.1 = ghetto, 0.9 = park avenue
	Dealers   []*Dealer
	Users     map[uuid.UUID]*User
	LocationX int
	LocationY int
	City      *City
}

type Dealer struct {
	Name  string
	Cash  int
	Drugs map[Drug]int
}

type Drug struct {
	Label string
	Type  int // 1 = marijuana, 2 = cocaine, 3 = lsd, 4 = heroin
}

type User struct {
	Id            uuid.UUID
	Name          string
	Sex           int // 1 = male, 2 = female
	CurrentAge    int
	Alive         bool
	UsingSince    int
	BornOn        int64
	DrugOfChoice  int     // 1 = marijuana, 2 = cocaine, 3 = lsd, 4 = heroin
	Dependency    float64 // static value: 0.0 = not dependent at all, 1.0 = extremely dependent, cannot survive on own
	Addiction     float64 // static value: 0.0 = no addictive personality, 1.0 = extremely addictive personality
	NumberOfUses  int
	CurrentHigh   float64 // 0.0 = clean and sober, 1.0 = dead
	CurrentDrug   int
	LastUsed      int64
	LastMoved     int64
	CurrentCorner *Corner
}

func NewUser(name string, sex int, age int, corner *Corner) *User {
	return &User{
		Id:            uuid.NewV4(),
		Name:          name,
		Sex:           sex,
		CurrentAge:    age,
		Alive:         true,
		BornOn:        time.Now().Unix(),
		UsingSince:    0,
		DrugOfChoice:  1, // e'rybody start out with the trees
		Dependency:    (float64(rand.Int()%10) + 1) * 0.1,
		Addiction:     (float64(rand.Int()%10) + 1) * 0.1,
		CurrentDrug:   0,
		CurrentHigh:   0,
		LastUsed:      0,
		NumberOfUses:  0,
		CurrentCorner: corner,
	}
}

func (user *User) Tick() {

	if user.CurrentHigh >= 0 {
		user.CurrentHigh = user.CurrentHigh - (0.08*float64(user.CurrentDrug)*(float64(rand.Int()%5)*0.01) + 0.10) // static linear decline of high

		if user.CurrentHigh < 0 {
			user.CurrentHigh = 0
		}
	} else {
		user.CurrentHigh = 0
	}

}

func (user *User) RandomMove() *Corner {
	x := -1
	y := -1
	maxX := len(user.CurrentCorner.City.Corners) - 1
	maxY := len(user.CurrentCorner.City.Corners[0]) - 1

	for x < 0 || y < 0 || x > maxX || y > maxY {
		switch rand.Int() % 4 {
		case 0:
			x = user.CurrentCorner.LocationX + 1
			y = user.CurrentCorner.LocationY
		case 1:
			x = user.CurrentCorner.LocationX
			y = user.CurrentCorner.LocationY + 1
		case 2:
			x = user.CurrentCorner.LocationX - 1
			y = user.CurrentCorner.LocationY
		case 3:
			x = user.CurrentCorner.LocationX
			y = user.CurrentCorner.LocationY - 1
		}
	}

	return user.MoveTo(x, y)
}

func (user *User) MoveTo(x, y int) *Corner {
	newCorner := user.CurrentCorner.City.Corner(x, y)
	currentCorner := user.CurrentCorner
	delete(currentCorner.Users, user.Id)
	newCorner.Users[user.Id] = user
	user.LastMoved = time.Now().Unix()

	return newCorner
}

func (user *User) LastFix() int64 {
	return time.Now().Unix() - user.LastUsed
}

func (user *User) LastMovedSeconds() int64 {
	return time.Now().Unix() - user.LastMoved
}

func (user *User) NeedsFix() bool {
	return user.CurrentHigh < user.Addiction
}

func (user *User) Use(drug int) {

	user.NumberOfUses = user.NumberOfUses + 1
	var docMultiplier float64

	if user.DrugOfChoice == drug {
		docMultiplier = 1
	} else {
		docMultiplier = 0.6
	}

	var high = docMultiplier * float64(drug)

	if user.NumberOfUses > 1 {
		high = high * math.Log10(float64(user.NumberOfUses+2))
	} else {
		high = high * 0.30
	}

	user.CurrentHigh = high
	user.LastUsed = time.Now().Unix()
	user.CurrentDrug = drug

	if user.CurrentHigh > 5 {
		fmt.Printf("USER OD'ED %s\n", user.Name)
		user.Alive = false
		return
	}

}

func (user *User) String() string {
	return fmt.Sprintf("[%s] %s (%d,%d) alive=%t [Dependency %f, Addiction: %f, CurrentHigh: %f, LastUsed: %d, Uses: %d]", user.Id, user.Name, user.CurrentCorner.LocationX, user.CurrentCorner.LocationY, user.Alive, user.Dependency, user.Addiction, user.CurrentHigh, user.LastUsed, user.NumberOfUses)
}

func (dealer *Dealer) String() string {
	return dealer.Name
}
