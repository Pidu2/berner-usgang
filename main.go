package main

import (
	"fmt"

	"github.com/Pidu2/berner-usgang/cache"
	"github.com/Pidu2/berner-usgang/globals"
)

func main() {
	// TODO build REST endpoint
	// TODO ggf add rabe

	roEvs, err := cache.ScrapeWithCache(globals.AvailableScrapers["roessli"], 5)
	if err != nil {
		return
	}
	for _, ev := range roEvs {
		fmt.Println("--------")
		fmt.Println("Date: " + ev.Date)
		fmt.Println("Title: " + ev.Title)
		fmt.Println("Artists: " + ev.Artists)
		fmt.Println("Genres: " + ev.Genre)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	roEvs2, err := cache.ScrapeWithCache(globals.AvailableScrapers["roessli"], 5)
	if err != nil {
		return
	}
	for _, ev := range roEvs2 {
		fmt.Println("--------")
		fmt.Println("Date: " + ev.Date)
		fmt.Println("Title: " + ev.Title)
		fmt.Println("Artists: " + ev.Artists)
		fmt.Println("Genres: " + ev.Genre)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	roEvs3, err := cache.ScrapeWithCache(globals.AvailableScrapers["roessli"], 6)
	if err != nil {
		return
	}
	for _, ev := range roEvs3 {
		fmt.Println("--------")
		fmt.Println("Date: " + ev.Date)
		fmt.Println("Title: " + ev.Title)
		fmt.Println("Artists: " + ev.Artists)
		fmt.Println("Genres: " + ev.Genre)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	roEvs4, err := cache.ScrapeWithCache(globals.AvailableScrapers["roessli"], 3)
	if err != nil {
		return
	}
	for _, ev := range roEvs4 {
		fmt.Println("--------")
		fmt.Println("Date: " + ev.Date)
		fmt.Println("Title: " + ev.Title)
		fmt.Println("Artists: " + ev.Artists)
		fmt.Println("Genres: " + ev.Genre)
	}

}
