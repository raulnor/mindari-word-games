package main

import (
	"fmt"
)

type GameStreak struct {
	Username string
	Game     string
	Current  int
	Longest  int
}

func GetStreakByUserAndGame(username string, game string) (*GameStreak, error) {
	db, err := getDatabase()
	if err != nil {
		return nil, err
	}
	sql := `
	WITH d AS (
		SELECT DISTINCT p.date AS day
		FROM scores s
		JOIN puzzles p ON s.game = p.game AND s.game_number = p.game_number
		WHERE s.game = ? AND s.username = ?
	),
	runs AS (
		SELECT MAX(day) AS run_end, COUNT(*) AS len
		FROM (
			SELECT day, julianday(day) - ROW_NUMBER() OVER (ORDER BY day) AS grp
			FROM d
		)
		GROUP BY grp
	)
	SELECT COALESCE(MAX(len), 0) AS longest,
		COALESCE(MAX(CASE WHEN run_end >= date(?, '-1 day') THEN len END), 0) AS current
	FROM runs`
	var streak GameStreak
	streak.Username = username
	streak.Game = game
	err = db.QueryRow(sql, game, username, getCurrentDate()).Scan(&streak.Longest, &streak.Current)
	if err != nil {
		return nil, fmt.Errorf("failed to get streaks by user and game: %v", err)
	}
	return &streak, nil
}
