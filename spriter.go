package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// func onHitConsole(e *colly.HTMLElement)

type console struct {
	name string
	path string
}

type game struct {
	name string
	path string
}

var consoles = make(map[string]console)
var games = make(map[string]game)

var c *colly.Collector = colly.NewCollector(
	colly.AllowedDomains("www.spriters-resource.com"),
)

func scanConsoles() {
	scanner := c.Clone()
	scanner.OnHTML("div#sidebar-left > a, div#other-systems > a", func(e *colly.HTMLElement) {
		dat := console{e.Text, e.Attr("href")}
		consoles[dat.name] = dat
	})
	scanner.Visit("https://www.spriters-resource.com/browse")
}

func scanGames() {
	scanner := c.Clone()
	scanner.OnHTML(".icondisplay > a.iconlink", func(e *colly.HTMLElement) {
		dat := game{strings.TrimSpace(e.Text), e.Attr("href")}
		games[dat.name] = dat
	})

	page := 1
	pageCount := 0
	pageNumberRegex := regexp.MustCompile(`\d+`)

	// Look for next page link
	scanner.OnHTML("table.pagination td:last-child > a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Next") {
			page++
		}
	})

	scanner.OnHTML("table.pagination td:nth-child(2)", func(e *colly.HTMLElement) {
		if pageCount != 0 {
			return
		}

		pageCount, _ = strconv.Atoi(pageNumberRegex.FindString(e.Text))
	})

	scanner.OnHTML("table.pagination td:last-child > span.disabledNav", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Next") {
			page = -1
		}
	})

	for 0 < page {
		scanner.Visit(fmt.Sprintf("https://www.spriters-resource.com/browse/games/page-%d", page))
		fmt.Printf("Scanned games page %d of %d\n", page-1, pageCount)
	}
}

func scan() {
	scanGames()
}

func main() {
	scan()
}
