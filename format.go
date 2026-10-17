package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func barMaxForScoresByGame(scores []Score, game string) int {
	switch {
	case strings.Contains(game, "Octordle"):
		return 100
	case game == "Wordle":
		return 7
	default:
		barMax := 1
		for _, s := range scores {
			f, err := strconv.ParseFloat(s.Score, 64)
			if err == nil {
				i := int(math.Ceil(f))
				if i > barMax {
					barMax = i
				}
			}
		}
		return barMax
	}
}

func formatGuessCount(v any) string {
	var guesses float64
	switch x := v.(type) {
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return x
		}
		guesses = f
	case float32:
		guesses = float64(x)
	case float64:
		guesses = x
	default:
		return fmt.Sprint(v)
	}

	if guesses != math.Trunc(guesses) {
		return fmt.Sprintf("%0.2f", guesses)
	} else {
		return strconv.FormatFloat(guesses, 'f', 0, 64)
	}
}

func formatTime(v any) string {
	var secs float64
	switch x := v.(type) {
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return x
		}
		secs = f
	case float32:
		secs = float64(x)
	case float64:
		secs = x
	default:
		return fmt.Sprint(v)
	}

	total := int(math.Floor(secs))
	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	} else {
		return fmt.Sprintf("%d:%02d", m, s)
	}
}

func formatScoreByGame(score any, game string) string {
	if isTimedGame(game) {
		return formatTime(score)
	} else {
		return formatGuessCount(score)
	}
}

func isTimedGame(game string) bool {
	switch game {
	case "Zip", "Mini Sudoku", "Crossclimb", "Queens", "Queens Ultimate - Daily Max", "Queens Ultimate - Daily Mini", "Tango":
		return true
	default:
		return false
	}
}
