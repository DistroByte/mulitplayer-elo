package main

import (
	"fmt"
	"math"
	"time"
)

func (l *League) AddMatch(date time.Time, results ...*MatchResult) ([]MatchDiff, error) {
	if len(l.Players) == 0 {
		return nil, ErrNoPlayers
	}

	populatedResults := make([]*MatchResult, 0, len(results))
	matchDiff := make([]MatchDiff, 0, len(results))

	// ensure all drivers are registered and flesh out the results with ELOs
	for _, result := range results {
		found := false
		for _, player := range l.Players {
			if player == nil {
				return []MatchDiff{}, fmt.Errorf("nil player found. not recording match")
			}

			if result.Player == nil {
				return []MatchDiff{}, fmt.Errorf("nil player found. not recording match")
			}

			if player.Name == result.Player.Name {
				result.Player.ELO = player.ELO
				populatedResults = append(populatedResults, result)
				found = true
				break
			}
		}

		if !found {
			return []MatchDiff{}, fmt.Errorf("player %q not found. not recording match", result.Player.Name)
		}
	}

	// calculate the ELO changes
	n := len(populatedResults)
	kValue := 32 / (n - 1)

	// loop over every result
	for player, result := range populatedResults {
		curELO := result.Player.ELO
		curPosition := result.Position

		// loop over every other result
		for opponentPlayer, opponentResult := range populatedResults {
			// skip comparing the driver to themselves
			if player == opponentPlayer {
				continue
			}

			opponentELO := opponentResult.Player.ELO
			opponentPosition := opponentResult.Position

			// calculate the expected score
			var S float64

			// if the driver finished higher than the other driver
			if curPosition < opponentPosition {
				S = 1.0
			} else {
				S = 0.0
			}

			// calculate the expected score
			E := 1.0 / (1.0 + math.Pow(10, float64(opponentELO-curELO)/400))

			// update the driver's ELO change
			result.Player.ELOChange += int(math.Round(float64(kValue) * (S - E)))
		}

		// update the driver's ELO
		result.Player.ELO += result.Player.ELOChange
		matchDiff = append(matchDiff, MatchDiff{
			Player: result.Player,
			Diff:   result.Player.ELOChange,
		})
	}

	// update the drivers' ELOs
	for _, result := range populatedResults {
		for _, player := range l.Players {
			if player.Name == result.Player.Name {
				player.ELO = result.Player.ELO

				player.Stats.MatchesPlayed += 1
				if result.Position == 1 {
					player.Stats.MatchesWon++
				}

				player.Stats.AllTimeAveragePlace += float64(result.Position)

				player.Stats.Last5Finish = append(player.Stats.Last5Finish, result.Position)
				if len(player.Stats.Last5Finish) > 5 {
					player.Stats.Last5Finish = player.Stats.Last5Finish[1:]
				}

				if player.ELO > player.Stats.PeakELO {
					player.Stats.PeakELO = player.ELO
				}
				break
			}
		}
	}

	// create the event
	err := l.createEvent(populatedResults)
	if err != nil {
		return []MatchDiff{}, err
	}

	return matchDiff, nil
}

func (l *League) createEvent(results []*MatchResult) error {
	event := Match{
		Results: results,
		Date:    time.Now(),
	}

	l.Matches = append(l.Matches, event)

	return nil
}
