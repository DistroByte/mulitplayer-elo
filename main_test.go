package multielo_test

import (
	"testing"
	"time"

	elo "github.com/distrobyte/multielo"
	"github.com/stretchr/testify/assert"
)

func TestLeague_NewLeague(t *testing.T) {
	l := elo.NewLeague()
	assert.NotNil(t, l)

	assert.NotNil(t, l.Players)
	assert.NotNil(t, l.Matches)

	assert.Equal(t, 0, len(l.Players))
	assert.Equal(t, 0, len(l.Matches))

	assert.IsType(t, []*elo.Player{}, l.Players)
	assert.IsType(t, []elo.Match{}, l.Matches)

	assert.IsType(t, &elo.League{}, l)

	assert.IsType(t, []*elo.Player{}, l.Players)
	assert.IsType(t, []elo.Match{}, l.Matches)
}

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

func TestMatch_AddMatch(t *testing.T) {
	t.Run("AddMatch", func(t *testing.T) {
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
		err = l.AddPlayer("player4")
		if err != nil {
			t.Error(err)
		}

		// get the players
		player1, err := l.GetPlayer("player1")
		if err != nil {
			t.Error(err)
		}

		player2, err := l.GetPlayer("player2")
		if err != nil {
			t.Error(err)
		}

		player3, err := l.GetPlayer("player3")
		if err != nil {
			t.Error(err)
		}

		player4, err := l.GetPlayer("player4")
		if err != nil {
			t.Error(err)
		}

		// generate the match
		results := []*elo.MatchResult{
			{Player: player1, Position: 1},
			{Player: player2, Position: 2},
			{Player: player3, Position: 3},
			{Player: player4, Position: 4},
		}

		matchDiff, err := l.AddMatch(time.Now(), results...)
		if err != nil {
			t.Error(err)
		}

		// check the match
		match := l.Matches[0]

		if match.Date.IsZero() {
			t.Error("Match time is zero")
		}

		if len(match.Results) != 4 {
			t.Error("Match results not added correctly")
		}

		for i, result := range match.Results {
			if result.Player != results[i].Player {
				t.Error("Player not added correctly")
			}

			if result.Position != results[i].Position {
				t.Error("Position not added correctly")
			}
		}

		// check the match diff
		if len(matchDiff) != 4 {
			t.Error("Match diff not added correctly")
		}

		for i, diff := range matchDiff {
			if diff.Player != results[i].Player {
				t.Error("Player not added correctly")
			}

			if diff.Diff == 0 {
				t.Error("Diff not added correctly")
			}
		}

	})

	t.Run("AddMatchPlayerNotFound", func(t *testing.T) {
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

		// get the players
		player1, err := l.GetPlayer("player1")
		if err != nil {
			t.Error(err)
		}

		player2, err := l.GetPlayer("player2")
		if err != nil {
			t.Error(err)
		}

		// generate the match
		results := []elo.MatchResult{
			{Player: player1, Position: 1},
			{Player: player2, Position: 2},
			{Player: nil, Position: 3},
		}

		_, err = l.AddMatch(time.Now(), &results[0], &results[1], &results[2])
		if err.Error() != "nil player found. not recording match" {
			t.Error(err)
		}
	})

	t.Run("AddMatchLeagueNil", func(t *testing.T) {
		l := elo.NewLeague()

		_, err := l.AddMatch(time.Now())
		if err != elo.ErrNoPlayers {
			t.Error(err)
		}
	})
}
