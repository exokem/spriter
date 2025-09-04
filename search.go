package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func compare(query string, subject string) int {

	if strings.Contains(subject, query) || strings.Contains(subject, strings.ToUpper(query)) {
		return abs(len(query)-len(subject)) / 2
	}

	if len(subject) < len(query) {
		return compare(subject, query)
	}

	// base score is length difference - low score is more similar
	score := abs(len(query) - len(subject))

	matches := 0

	for iq := range len(query) {
		q := query[iq]
		has := false

		for is := iq; is < len(subject); is++ {
			s := subject[is]

			if q == s {
				has = true
				// add distance between matching characters
				score += abs(is - iq)
				matches++
				break
			} else if unicode.ToUpper(rune(q)) == unicode.ToUpper(rune(s)) {
				has = true
				// add distance between matching characters
				matches++
				score += abs(is-iq) + 1
				break
			}
		}

		if !has {
			score += len(subject)
		}
	}

	if matches == 0 {
		score *= 2
	}

	return score
}

type SearchMethod func(results map[string]int)

func searchByConsole(query string) SearchMethod {
	return func(results map[string]int) {
		for name, game := range spriter.games.Index {
			score := compare(query, game.Path)
			if score < 20 {
				results[fmt.Sprintf("%s (%s)", game.Path, name)] = score
			}
		}
	}
}

func searchByGame(query string) SearchMethod {
	return func(results map[string]int) {
		for name, game := range spriter.games.Index {
			score := compare(query, game.Path)
			if score < 20 {
				results[fmt.Sprintf("%s (%s)", game.Path, name)] = score
			}
		}
	}
}

func searchByGameAndConsole(game string, console string) SearchMethod {
	return func(results map[string]int) {
		for name, g := range spriter.games.Index {
			score := compare(game, g.Path) + compare(console, g.Path)
			if score < 30 {
				results[fmt.Sprintf("%s (%s)", g.Path, name)] = score
			}
		}
	}
}

type result struct {
	name  string
	score int
}

func runSearch(methods ...SearchMethod) []result {
	index := make(map[string]int)

	for _, method := range methods {
		method(index)
	}

	var pairs []result

	for name, score := range index {
		pairs = append(pairs, result{name, score})
	}

	slices.SortFunc(pairs, func(a, b result) int {
		return a.score - b.score
	})

	return pairs
}
