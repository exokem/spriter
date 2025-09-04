package main

import (
	"os"
)

func usage() {
	info("TBD - usage instructions")
	os.Exit(0)
}

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
