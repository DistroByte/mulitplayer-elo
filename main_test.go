package multielo_test

import (
	"fmt"
	"testing"

	"github.com/distrobyte/multielo"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/plot"
)

func TestLeague_NewLeague(t *testing.T) {
	l := multielo.NewLeague()
	assert.NotNil(t, l)

	assert.NotNil(t, l.Players)
	assert.NotNil(t, l.Matches)

	assert.Equal(t, 0, len(l.Players))
	assert.Equal(t, 0, len(l.Matches))

	assert.IsType(t, []*multielo.Player{}, l.Players)
	assert.IsType(t, []multielo.Match{}, l.Matches)

	assert.IsType(t, &multielo.League{}, l)

	assert.IsType(t, []*multielo.Player{}, l.Players)
	assert.IsType(t, []multielo.Match{}, l.Matches)
}

func TestPlayer_AddPlayer(t *testing.T) {
	t.Run("AddPlayer", func(t *testing.T) {
		l := multielo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("AddPlayerAlreadyExists", func(t *testing.T) {
		l := multielo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player1")
		if err != multielo.ErrPlayerAlreadyExists {
			t.Error(err)
		}
	})

	t.Run("AddMultiplePlayers", func(t *testing.T) {
		l := multielo.NewLeague()
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
		l := multielo.NewLeague()
		err := l.AddPlayer("player1")
		if err != nil {
			t.Error(err)
		}
		err = l.AddPlayer("player1")
		if err != multielo.ErrPlayerAlreadyExists {
			t.Error(err)
		}
		err = l.AddPlayer("player1")
		if err != multielo.ErrPlayerAlreadyExists {
			t.Error(err)
		}
	})
}

func TestPlayer_GetPlayer(t *testing.T) {
	t.Run("GetPlayer", func(t *testing.T) {
		l := multielo.NewLeague()
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
		l := multielo.NewLeague()
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
		l := multielo.NewLeague()
		_, err := l.GetPlayer("player1")
		if err != multielo.ErrPlayerNotFound {
			t.Error(err)
		}
	})
}

func TestPlayer_RemovePlayer(t *testing.T) {
	t.Run("RemovePlayer", func(t *testing.T) {
		l := multielo.NewLeague()
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
		l := multielo.NewLeague()
		err := l.RemovePlayer("player1")
		if err != multielo.ErrPlayerNotFound {
			t.Error(err)
		}
	})
}

func TestPlayer_ResetPlayers(t *testing.T) {
	t.Run("ResetPlayers", func(t *testing.T) {
		l := multielo.NewLeague()
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
			if p.ELO != multielo.InitialELO {
				t.Error("Player ELO not reset")
			}
			if p.Stats.MatchesPlayed != 0 {
				t.Error("Player stats not reset")
			}
		}
	})

	t.Run("ResetPlayersEmpty", func(t *testing.T) {
		l := multielo.NewLeague()
		l.ResetPlayers()
	})
}

func TestMatch_AddMatch(t *testing.T) {
	t.Run("AddMatch", func(t *testing.T) {
		l := multielo.NewLeague()
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
		results := []*multielo.MatchResult{
			{Player: player1, Position: 1},
			{Player: player2, Position: 2},
			{Player: player3, Position: 3},
			{Player: player4, Position: 4},
		}

		matchDiff, err := l.AddMatch(results)
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
		l := multielo.NewLeague()
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
		results := []*multielo.MatchResult{
			{Player: player1, Position: 1},
			{Player: player2, Position: 2},
			{Player: nil, Position: 3},
		}

		_, err = l.AddMatch(results)
		if err.Error() != "nil player found. not recording match" {
			t.Error(err)
		}
	})

	t.Run("AddMatchLeagueNil", func(t *testing.T) {
		l := multielo.NewLeague()

		_, err := l.AddMatch([]*multielo.MatchResult{})
		if err != multielo.ErrNoPlayers {
			t.Error(err)
		}
	})
}

func testTicker(t *testing.T, ticker plot.Ticker, start, end float64, expected int) {
	ticks := ticker.Ticks(start, end)
	if len(ticks) != expected {
		t.Errorf("expected %d ticks, got %v", expected, len(ticks))
	}
}

func TestTicks(t *testing.T) {
	tests := []struct {
		name     string
		ticker   plot.Ticker
		start    float64
		end      float64
		expected int
	}{
		{"raceTicker 0", multielo.RaceTicker{}, 0, 0, 0},
		{"eloTicker 0", multielo.ELOTicker{}, 0, 0, 0},
		{"raceTicker 100", multielo.RaceTicker{}, 0, 100, 33},
		{"eloTicker 100", multielo.ELOTicker{}, 0, 100, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTicker(t, tt.ticker, tt.start, tt.end, tt.expected)
		})
	}
}

func TestLeague_GetPlayers(t *testing.T) {
	l := multielo.NewLeague()
	err := l.AddPlayer("player1")
	assert.NoError(t, err)
	err = l.AddPlayer("player2")
	assert.NoError(t, err)

	players := l.GetPlayers()
	assert.Equal(t, 2, len(players))
	assert.Equal(t, "player1", players[0].Name)
	assert.Equal(t, "player2", players[1].Name)
}

func TestLeague_GetMatches(t *testing.T) {
	l := multielo.NewLeague()
	err := l.AddPlayer("player1")
	assert.NoError(t, err)
	err = l.AddPlayer("player2")
	assert.NoError(t, err)

	player1, err := l.GetPlayer("player1")
	assert.NoError(t, err)
	player2, err := l.GetPlayer("player2")
	assert.NoError(t, err)

	results := []*multielo.MatchResult{
		{Player: player1, Position: 1},
		{Player: player2, Position: 2},
	}

	_, err = l.AddMatch(results)
	assert.NoError(t, err)

	matches := l.GetMatches()
	assert.Equal(t, 1, len(matches))
	assert.Equal(t, 2, len(matches[0].Results))
}

func TestLeague_GetPlayerStats(t *testing.T) {
	l := multielo.NewLeague()
	err := l.AddPlayer("player1")
	assert.NoError(t, err)

	stats, err := l.GetPlayerStats("player1")
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 0, stats.MatchesPlayed)
	assert.Equal(t, 0, stats.MatchesWon)
	assert.Equal(t, 0.0, stats.AllTimeAveragePlace)
	assert.Equal(t, 0, len(stats.Last5Finish))
	assert.Equal(t, multielo.InitialELO, stats.PeakELO)
}

func TestLeague_GetPlayerELO(t *testing.T) {
	l := multielo.NewLeague()
	err := l.AddPlayer("player1")
	assert.NoError(t, err)

	elo, err := l.GetPlayerELO("player1")
	assert.NoError(t, err)
	assert.Equal(t, multielo.InitialELO, elo)
}

func TestLeague_GenerateGraph(t *testing.T) {
	l := multielo.NewLeague()
	err := l.AddPlayer("player1")
	assert.NoError(t, err)
	err = l.AddPlayer("player2")
	assert.NoError(t, err)

	_, err = l.GetPlayer("player1")
	assert.NoError(t, err)
	_, err = l.GetPlayer("player2")
	assert.NoError(t, err)

	var results1 []*multielo.MatchResult
	for i := 0; i < 2; i++ {
		results1 = append(results1, &multielo.MatchResult{
			Player:   &multielo.Player{Name: fmt.Sprintf("player%v", i+1)},
			Position: i + 1,
		})
	}

	_, err = l.AddMatch(results1)
	assert.NoError(t, err)

	assert.NoError(t, err)

	err = l.AddPlayer("player3")
	assert.NoError(t, err)

	_, err = l.GetPlayer("player3")
	assert.NoError(t, err)

	var results2 []*multielo.MatchResult
	for i := 0; i < 3; i++ {
		results2 = append(results2, &multielo.MatchResult{
			Player:   &multielo.Player{Name: fmt.Sprintf("player%v", i+1)},
			Position: i + 1,
		})
	}

	_, err = l.AddMatch(results2)
	assert.NoError(t, err)

	err = l.AddPlayer("player4")
	assert.NoError(t, err)

	_, err = l.GetPlayer("player4")
	assert.NoError(t, err)

	var results3 []*multielo.MatchResult
	for i := 0; i < 3; i++ {
		results3 = append(results3, &multielo.MatchResult{
			Player:   &multielo.Player{Name: fmt.Sprintf("player%v", i+1)},
			Position: i + 1,
		})
	}

	_, err = l.AddMatch(results3)
	assert.NoError(t, err)

	graphPath, err := l.GenerateGraph()
	assert.NoError(t, err)
	assert.Equal(t, "elo.png", graphPath)
}
