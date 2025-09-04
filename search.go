package main

import (
	"fmt"
	"os"
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

func search() {
	if len(os.Args) < 3 {
		usage()
	}

	var console string
	var game string

	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]

		if len(os.Args) <= i+1 {
			usage()
		}

		argv := os.Args[i+1]

		switch arg {
		case "--console":
			console = argv
			i++
		case "--game":
			game = argv
			i++
		default:
			usage()
		}
	}

	var entries []result

	if len(console) != 0 && len(game) != 0 {
		entries = runSearch(searchByGameAndConsole(game, console))
	} else if len(console) != 0 {
		entries = runSearch(searchByConsole(console))
	} else if len(game) != 0 {
		entries = runSearch(searchByGame(game))
	}

	info("%d Results (Showing 100 most relevant):\n", len(entries))

	i := 0
	var entry result

	for i, entry = range entries {
		if 100 <= i {
			break
		}

		info(" - %s\n", entry.name)
	}

	if 100 < len(entries) {
		info(" + %d less relevant results", len(entries)-i)
	}
}
