package main_test

import (
	"testing"
	"time"

	elo "github.com/distrobyte/multiplayer-elo"
)

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
