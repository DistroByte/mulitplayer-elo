package elo_test

import (
	"testing"

	elo "github.com/distrobyte/multiplayer-elo"
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
