package main

func NewLeague() *League {
	return &League{
		Players: []*Player{},
		Matches: []Match{},
	}
}
