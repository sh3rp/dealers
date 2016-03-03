package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Corner struct {
	Street  string
	Rating  float64 // 0.1 = ghetto, 0.9 = park avenue
	Dealers []*Dealer
	Users   []*User
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
	Name         string
	Sex          int // 1 = male, 2 = female
	CurrentAge   int
	Alive        bool
	UsingSince   int
	BornOn       int64
	DrugOfChoice int     // 1 = marijuana, 2 = cocaine, 3 = lsd, 4 = heroin
	Dependency   float64 // static value: 0.0 = not dependent at all, 1.0 = extremely dependent, cannot survive on own
	Addiction    float64 // static value: 0.0 = no addictive personality, 1.0 = extremely addictive personality
	NumberOfUses int
	CurrentHigh  float64 // 0.0 = clean and sober, 1.0 = dead
	CurrentDrug  int
	LastUsed     int64
}

func NewUser(name string, sex int, age int) *User {
	return &User{
		Name:         name,
		Sex:          sex,
		CurrentAge:   age,
		Alive:        true,
		BornOn:       time.Now().Unix(),
		UsingSince:   0,
		DrugOfChoice: 1, // e'rybody start out with the trees
		Dependency:   (float64(rand.Int()%10) + 1) * 0.1,
		Addiction:    (float64(rand.Int()%10) + 1) * 0.1,
		CurrentDrug:  0,
		CurrentHigh:  0,
		LastUsed:     0,
		NumberOfUses: 0,
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
	return fmt.Sprintf("%s alive=%t [Dependency %f, Addiction: %f, CurrentHigh: %f, LastUsed: %d, Uses: %d]", user.Name, user.Alive, user.Dependency, user.Addiction, user.CurrentHigh, user.LastUsed, user.NumberOfUses)
}

func (dealer *Dealer) String() string {
	return dealer.Name
}
