package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gocolly/colly"
	"github.com/schollz/progressbar/v3"
)

func dispatch(folder string, entry *Entry, progress *progressbar.ProgressBar) {

	dir := path.Join("data/downloads", folder)
	pngPath := path.Join(dir, fmt.Sprintf("%s.png", entry.Name))
	zipPath := path.Join(dir, fmt.Sprintf("%s.zip", entry.Name))

	// Skip files that have already been downloaded
	if _, err := os.Stat(pngPath); err == nil {
		progress.ChangeMax(progress.GetMax() - 1)
		return
	} else if _, err := os.Stat(zipPath); err == nil {
		progress.ChangeMax(progress.GetMax() - 1)
		return
	}

	scanner := c.Clone()

	var extension string

	scanner.OnHTML("div#assetdisplay > img", func(e *colly.HTMLElement) {
		entry.Path = fmt.Sprintf("https://www.spriters-resource.com%s", e.Attr("src"))
		extension = ".png"
	})

	scanner.OnHTML("button#extract-zip", func(e *colly.HTMLElement) {
		entry.Path = fmt.Sprintf("https://www.spriters-resource.com%s", e.Attr("data-file"))
		extension = ".zip"
	})

	scanner.OnScraped(func(r *colly.Response) {
		progress.Add(1)
	})

	scanner.Visit(fmt.Sprintf("https://www.spriters-resource.com%s", entry.Path))

	scanner.Wait()

	if len(extension) == 0 {
		progress.ChangeMax(progress.GetMax() - 1)
		return
	}

	var res *http.Response
	var err error

	if res, err = http.Get(entry.Path); err != nil {
		fmt.Printf("Error retrieving asset from %s\n", entry.Path)
		return
		// panic(err)
	}

	defer res.Body.Close()

	if err = os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	var bytes []byte

	if bytes, err = io.ReadAll(res.Body); err != nil {
		panic(err)
	}

	switch extension {
	case ".png":
		if err := os.WriteFile(pngPath, bytes, 0644); err != nil {
			panic(err)
		}
	case ".zip":
		if err := os.WriteFile(zipPath, bytes, 0644); err != nil {
			panic(err)
		}
	}
}

func getSprites(page string) {
	info("Gathering sprite links... ")
	links := collectSpriteLinks(page)
	info("Discovered %d sprites.\n", len(links.Index))

	bar := progressbar.Default(int64(len(links.Index)), "Downloading sprites")

	for _, entry := range links.Index {
		dispatch(page, &entry, bar)
	}

	info("\nDownloaded %d sprites.\n", bar.GetMax())
}

func get() {
	if len(os.Args) < 3 {
		usage()
	}

	path := os.Args[2]

	info("Retrieving sprites from https://www.spriters-resource.com%s\n", path)

	getSprites(path)
}
