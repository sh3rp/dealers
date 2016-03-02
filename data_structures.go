package main

import "fmt"

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
	Name           string
	DrugOfChoice   int     // 1 = marijuana, 2 = cocaine, 3 = lsd, 4 = heroin
	Susceptibility float64 // 0.1 = casual user, 0.9 = full blown junkie
	CurrentHigh    float64 // 0.0 = clean and sober, 0.9 = high as fuck
	LastUsed       int64
}

func (user *User) String() string {
	return fmt.Sprintf("%s [Susceptibility: %f, CurrentHigh: %f, LastUsed: %d]", user.Name, user.Susceptibility, user.CurrentHigh, user.LastUsed)
}

func (dealer *Dealer) String() string {
	return dealer.Name
}
