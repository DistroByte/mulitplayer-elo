package elo

import (
	"errors"
)

var (
	ErrPlayerAlreadyExists = errors.New("player already exists")
	ErrPlayerNotFound      = errors.New("player not found")
	ErrMatchNotFound       = errors.New("match not found")
	ErrInvalidMatch        = errors.New("invalid match")
	ErrInvalidELOChange    = errors.New("invalid elo change")
	ErrInvalidPlayer       = errors.New("invalid player")
	ErrInvalidPlayerStats  = errors.New("invalid player stats")
	ErrInvalidLeague       = errors.New("invalid league")
	ErrNoPlayers           = errors.New("no players")
)
