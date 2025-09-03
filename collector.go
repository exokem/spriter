package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var c *colly.Collector = colly.NewCollector(
	colly.AllowedDomains("www.spriters-resource.com"),
)

var numberRegex *regexp.Regexp = regexp.MustCompile(`\d+(,\d+)?`)

func getNumber(s string) int {
	s = strings.ReplaceAll(numberRegex.FindString(s), ",", "")

	n, _ := strconv.Atoi(s)
	return n
}

func collectConsoles() *EntrySet {
	set := NewEntrySet()
	scanner := c.Clone()

	scanner.OnHTML("div#sidebar-left > a, div#other-systems > a", func(e *colly.HTMLElement) {
		set.Store(e.Text, e.Attr("href"))
	})

	scanner.Visit("https://www.spriters-resource.com/browse")

	return set
}

func collectGamesInfo() (pageCount int, gameCount int) {
	scanner := c.Clone()
	pages := 0
	games := 0

	scanner.OnHTML("table.pagination td:nth-child(2)", func(e *colly.HTMLElement) {
		pages = getNumber(e.Text)
	})

	scanner.OnHTML("div#browse-sorting", func(e *colly.HTMLElement) {
		games = getNumber(e.Text)
	})

	scanner.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	scanner.Visit("https://www.spriters-resource.com/browse/games")

	// fmt.Printf("Discovered %d games pages\n", pages)

	return pages, games
}

// func scanGamesPageCount() int {
// 	scanner := c.Clone()
// 	pages := 0
// 	scanner.OnHTML("table.pagination td:nth-child(2)", func(e *colly.HTMLElement) {
// 		pages = GetNumber(e.Text)
// 	})

// 	scanner.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})

// 	scanner.Visit("https://www.spriters-resource.com/browse/games")

// 	fmt.Printf("Discovered %d games pages\n", pages)

// 	return pages
// }

func gameCollector(set *EntrySet) *colly.Collector {
	scanner := colly.NewCollector(
		colly.AllowedDomains("www.spriters-resource.com"),
		colly.MaxDepth(1),
		colly.Async(true),
	)

	scanner.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		RandomDelay: 5 * time.Second,
	})

	// scanner.WithTransport(&http.Transport{
	// 	Proxy: http.ProxyFromEnvironment,
	// 	DialContext: (&net.Dialer{
	// 		Timeout:   30 * time.Second,
	// 		KeepAlive: 30 * time.Second,
	// 		DualStack: true,
	// 	}).DialContext,
	// 	MaxIdleConns:          100,
	// 	IdleConnTimeout:       90 * time.Second,
	// 	TLSHandshakeTimeout:   10 * time.Second,
	// 	ExpectContinueTimeout: 1 * time.Second,
	// })

	scanner.OnHTML(".icondisplay > a.iconlink", func(e *colly.HTMLElement) {
		set.Store(strings.TrimSpace(e.Text), e.Attr("href"))
	})

	scanner.OnScraped(func(r *colly.Response) {
		fmt.Printf("Finished scanning '%d'\n", getNumber(r.Request.URL.String()))
	})

	scanner.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Failed to scan page %d\n", getNumber(r.Request.URL.String()))
	})

	return scanner
}

func collectAllGames() *EntrySet {
	pageCount, _ := collectGamesInfo()
	set := NewEntrySet()
	scanner := gameCollector(set)

	for page := 1; page < pageCount; page++ {
		scanner.Visit(fmt.Sprintf("https://www.spriters-resource.com/browse/games/page-%d", page))
	}

	scanner.Wait()

	return set
}

func collectGamesOnPage(url string) *EntrySet {
	set := NewEntrySet()
	scanner := gameCollector(set)

	scanner.Visit(url)
	scanner.Wait()

	return set
}

func collectGamesOnPageNumber(number int) *EntrySet {
	return collectGamesOnPage(fmt.Sprintf("https://www.spriters-resource.com/browse/games/page-%d", number))
}
