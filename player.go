package main

import "strings"

func (l *League) AddPlayer(name string) error {
	name = strings.ToLower(name)

	for _, p := range l.Players {
		if strings.ToLower(p.Name) == name {
			return ErrPlayerAlreadyExists
		}
	}

	l.Players = append(l.Players, &Player{
		Name: name,
		ELO:  InitialElo,
		Stats: &PlayerStats{
			Last5Finish:         []int{},
			MatchesPlayed:       0,
			MatchesWon:          0,
			AllTimeAveragePlace: 0,
			PeakELO:             InitialElo,
		},
	})

	return nil
}

func (l *League) GetPlayer(name string) (*Player, error) {
	name = strings.ToLower(name)

	for _, p := range l.Players {
		if strings.ToLower(p.Name) == name {
			return p, nil
		}
	}

	return nil, ErrPlayerNotFound
}

func (l *League) RemovePlayer(name string) error {
	name = strings.ToLower(name)

	for i, p := range l.Players {
		if strings.ToLower(p.Name) == name {
			l.Players = append(l.Players[:i], l.Players[i+1:]...)
			return nil
		}
	}

	return ErrPlayerNotFound
}

func (l *League) ResetPlayers() {
	for _, p := range l.Players {
		p.ELO = InitialElo
		p.Stats = &PlayerStats{}
	}
}
