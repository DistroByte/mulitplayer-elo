package main

import "time"

type Player struct {
	Name      string
	ELO       int
	ELOChange int
	Stats     *PlayerStats
}

type PlayerStats struct {
	MatchesPlayed       int
	MatchesWon          int
	AllTimeAveragePlace float64
	Last5Finish         []int
	PeakELO             int
}

type Match struct {
	Results []*MatchResult
	Date    time.Time
}

type MatchResult struct {
	Position int
	Player   *Player
}

type MatchDiff struct {
	Player *Player
	Diff   int
}

type League struct {
	Players []*Player
	Matches []Match
}
