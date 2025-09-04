package main

import (
	"os"
	"strconv"
)

func scanConsoles() {
	info("Collecting console information... ")
	spriter.consoles.Absorb(collectConsoles())
	info("Found %d consoles.\n", len(spriter.consoles.Index))
	info("Saving consoles... ")
	saveEntries("consoles.json", spriter.consoles)
	info("Done.\n")
}

func scanGameMeta() {
	info("Collecting game information... ")
	spriter.metadata.PageCount, spriter.metadata.GameCount = collectGamesInfo()
	info("Found %d games across %d pages.\n", spriter.metadata.GameCount, spriter.metadata.PageCount)

	info("Saving metadata... ")
	save("meta.json", spriter.metadata)
	info("Done.\n")
}

func saveGames() {
	info("Saving games... ")
	save("games.json", spriter.games.Index)
	info("Done.\n")
}

func scanAllGames() {
	scanGameMeta()

	spriter.games.Absorb(collectGamesInPageRange(1, spriter.metadata.PageCount))

	saveGames()
}

func scanSinglePage() {
	if len(os.Args) < 5 {
		usage()
	}

	var page int
	var err error

	if page, err = strconv.Atoi(os.Args[4]); err != nil {
		panic(err)
	}

	info("Scanning game page %d... ", page)
	spriter.games.Absorb(collectGamesOnPageNumber(page))

	saveGames()
}

func scanPageRange() {
	if len(os.Args) < 6 {
		usage()
	}

	var min, max int
	var err error

	if min, err = strconv.Atoi(os.Args[4]); err != nil {
		panic(err)
	}

	if max, err = strconv.Atoi(os.Args[5]); err != nil {
		panic(err)
	}

	scanGameMeta()

	info("Scanning game pages %d-%d.\n", min, max)
	spriter.games.Absorb(collectGamesInPageRange(min, max))

	saveGames()
}

func scanGames() {
	if len(os.Args) < 4 {
		usage()
	}

	switch os.Args[3] {
	case "-m", "--meta":
		scanGameMeta()
	case "-a", "--all":
		scanAllGames()
	case "--page":
		scanSinglePage()
	case "--pages":
		scanPageRange()
	default:
		usage()
	}
}

func scan() {
	if len(os.Args) < 3 {
		usage()
	}

	cmd := os.Args[2]

	switch cmd {
	case "consoles":
		scanConsoles()
	case "games":
		scanGames()
	default:
		usage()
	}
}
