package main

import (
	"os"
)

type Metadata struct {
	GameCount int
	PageCount int
}

type Spriter struct {
	games    *EntrySet
	consoles *EntrySet
	metadata *Metadata
}

var spriter *Spriter

func usage() {
	// info("\n+--------------------------- spriter v1.0.0 ---------------------------+\n\n")
	info("--- scanning ---\n")
	info("Scanning operations read specific pages of the spriter's resource to parse and cache certain information.\n")
	info("This is only relevant when using the search function, e.g. if you don't know the url for your target assets.\n")
	info("\n")
	info("spriter scan consoles\n")
	info(" - Scans available consoles \n")
	info("\n")
	info("spriter scan games [--meta | -m]\n")
	info(" - Fetches and caches game and page counts \n")
	info("\n")
	info("spriter scan games [--all | -a] \n")
	info(" - WARNING: This command may take longer than 10 minutes to finish\n")
	info(" - Scans all game pages and caches available game information \n")
	info("\n")
	info("spriter scan games --page <page-number>\n")
	info(" - Fetches and caches all games for a specific page number\n")
	info("\n")
	info("spriter scan games --pages <start> <end>\n")
	info(" - Fetches and caches all games for a specific range of pages (inclusive)\n")
	info("\n")
	os.Exit(0)
}

func main() {
	spriter = &Spriter{
		loadEntries("games.json"),
		loadEntries("consoles.json"),
		loadMeta(),
	}

	info("\n+--------------------------- spriter v1.0.0 ---------------------------+\n\n")

	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "scan":
		scan()
	default:
		usage()
	}
}
