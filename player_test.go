package elo_test

import (
	"testing"

	elo "github.com/distrobyte/multiplayer-elo"
)

func TestPlayer_AddPlayer(t *testing.T) {
	t.Run("AddPlayer", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("AddPlayerAlreadyExists", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player1")
		if err != elo.ErrPlayerAlreadyExists {
			t.Error(err)
		}
	})

	t.Run("AddMultiplePlayers", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player2")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player3")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("AddMultiplePlayersSameName", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player1")
		if err != elo.ErrPlayerAlreadyExists {
			t.Error(err)
		}
		err = l.AddPlayer("player1")
		if err != elo.ErrPlayerAlreadyExists {
			t.Error(err)
		}
	})
}

func TestPlayer_GetPlayer(t *testing.T) {
	t.Run("GetPlayer", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		p, err := l.GetPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		if p.Name != "player1" {
			t.Error("Player name does not match")
		}
	})

	t.Run("GetPlayerCaseInsensitive", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}

		p, err := l.GetPlayer("Player1")
		if err != nil {
			t.Error(err)
		}

		if p.Name != "player1" {
			t.Error("Player name does not match")
		}

		p, err = l.GetPlayer("pLaYeR1")
		if err != nil {
			t.Error(err)
		}

		if p.Name != "player1" {
			t.Error("Player name does not match")
		}
	})

	t.Run("GetPlayerNotFound", func(t *testing.T) {
		l := elo.NewLeague()
		_, err := l.GetPlayer("player1")
		if err != elo.ErrPlayerNotFound {
			t.Error(err)
		}
	})
}

func TestPlayer_RemovePlayer(t *testing.T) {
	t.Run("RemovePlayer", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.RemovePlayer("player1")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("RemovePlayerNotFound", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.RemovePlayer("player1")
		if err != elo.ErrPlayerNotFound {
			t.Error(err)
		}
	})
}

func TestPlayer_ResetPlayers(t *testing.T) {
	t.Run("ResetPlayers", func(t *testing.T) {
		l := elo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player2")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player3")
		if err != nil {
			t.Error(err)
		}
		l.ResetPlayers()
		for _, p := range l.Players {
			if p.ELO != elo.InitialElo {
				t.Error("Player ELO not reset")
			}
			if p.Stats.MatchesPlayed != 0 {
				t.Error("Player stats not reset")
			}
		}
	})

	t.Run("ResetPlayersEmpty", func(t *testing.T) {
		l := elo.NewLeague()
		l.ResetPlayers()
	})
}
