package main

import (
	"fmt"
)

// type console struct {
// 	name string
// 	path string
// }

// type game struct {
// 	name string
// 	path string
// }

// type Consoles struct {
// 	index map[string]console
// 	mu    sync.Mutex
// }

// type Games struct {
// 	index map[string]game
// 	mu    sync.Mutex
// }

// func (consoles *Consoles) Store(name string, path string) {
// 	consoles.mu.Lock()
// 	defer consoles.mu.Unlock()
// 	consoles.index[name] = console{name, path}
// }

// func (games *Games) Store(name string, path string) {
// 	games.mu.Lock()
// 	defer games.mu.Unlock()
// 	games.index[name] = game{name, path}
// }

// var consoles = Consoles{make(map[string]console), sync.Mutex{}}
// var games = Games{make(map[string]game), sync.Mutex{}}

// var c *colly.Collector = colly.NewCollector(
// 	colly.AllowedDomains("www.spriters-resource.com"),
// )

// var numberRegex *regexp.Regexp = regexp.MustCompile(`\d+(,\d+)?`)

// func GetNumber(s string) int {
// 	s = strings.ReplaceAll(numberRegex.FindString(s), ",", "")

// 	n, _ := strconv.Atoi(s)
// 	return n
// }

// func scanConsoles() {
// 	scanner := c.Clone()

// 	// scanner.OnHTML("div#sidebar-left > a, div#other-systems > a", func(e *colly.HTMLElement) {
// 	// 	consoles.Store(e.Text, e.Attr("href"))
// 	// })
// 	scanner.Visit("https://www.spriters-resource.com/browse")
// }

// func ScanGamesInfo() (availableGames int, pageCount int) {
// 	scanner := c.Clone()
// 	pages := 0
// 	games := 0

// 	scanner.OnHTML("table.pagination td:nth-child(2)", func(e *colly.HTMLElement) {
// 		pages = GetNumber(e.Text)
// 	})

// 	scanner.OnHTML("div#browse-sorting", func(e *colly.HTMLElement) {
// 		games = GetNumber(e.Text)
// 	})

// 	scanner.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})

// 	scanner.Visit("https://www.spriters-resource.com/browse/games")

// 	// fmt.Printf("Discovered %d games pages\n", pages)

// 	return pages, games
// }

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

// func GameCollector(set *EntrySet) *colly.Collector {
// 	scanner := colly.NewCollector(
// 		colly.AllowedDomains("www.spriters-resource.com"),
// 		colly.MaxDepth(1),
// 		colly.Async(true),
// 	)

// 	scanner.Limit(&colly.LimitRule{
// 		DomainGlob:  "*",
// 		Parallelism: 4,
// 		RandomDelay: 5 * time.Second,
// 	})

// 	// scanner.WithTransport(&http.Transport{
// 	// 	Proxy: http.ProxyFromEnvironment,
// 	// 	DialContext: (&net.Dialer{
// 	// 		Timeout:   30 * time.Second,
// 	// 		KeepAlive: 30 * time.Second,
// 	// 		DualStack: true,
// 	// 	}).DialContext,
// 	// 	MaxIdleConns:          100,
// 	// 	IdleConnTimeout:       90 * time.Second,
// 	// 	TLSHandshakeTimeout:   10 * time.Second,
// 	// 	ExpectContinueTimeout: 1 * time.Second,
// 	// })

// 	scanner.OnHTML(".icondisplay > a.iconlink", func(e *colly.HTMLElement) {
// 		set.Store(strings.TrimSpace(e.Text), e.Attr("href"))
// 	})

// 	scanner.OnScraped(func(r *colly.Response) {
// 		fmt.Printf("Finished scanning '%d'\n", GetNumber(r.Request.URL.String()))
// 	})

// 	scanner.OnError(func(r *colly.Response, err error) {
// 		fmt.Printf("Failed to scan page %d\n", GetNumber(r.Request.URL.String()))
// 	})

// 	return scanner
// }

// func ScanAllGames() *EntrySet {
// 	pageCount := scanGamesPageCount()
// 	set := NewEntrySet()
// 	scanner := GameCollector(set)

// 	for page := 1; page < pageCount; page++ {
// 		scanner.Visit(fmt.Sprintf("https://www.spriters-resource.com/browse/games/page-%d", page))
// 	}

// 	scanner.Wait()

// 	return set
// }

// func ScanGamesOnPage(url string) *EntrySet {
// 	set := NewEntrySet()
// 	scanner := GameCollector(set)

// 	scanner.Visit(url)
// 	scanner.Wait()

// 	return set
// }

// func ScanGamesOnPageNumber(number int) *EntrySet {
// 	return ScanGamesOnPage(fmt.Sprintf("https://www.spriters-resource.com/browse/games/page-%d", number))
// }

func scan() {
	fmt.Println(CollectGamesInfo())
	// ScanAllGames()
}

func main() {
	scan()
}
