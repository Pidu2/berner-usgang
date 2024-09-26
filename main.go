package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type event struct {
	title   string
	date    string
	artists string
	genre   string
}

func errEvent(err string) []event {
	log.Fatal(err)
	return []event{event{
		title:   "error",
		date:    "error",
		artists: "error",
		genre:   "error",
	}}
}

func scrapeDachstock(url string) []event {
	evList := []event{}
	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		return errEvent(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errEvent(fmt.Sprintf("Failed to fetch page: %d %s", res.StatusCode, res.Status))
	}

	// Parse the page
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return errEvent(err.Error())
	}

	// Find the event list and iterate over each event item
	doc.Find(".teaser").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := eventItem.Find(".event-date").Text()

		// Extract event title
		eventTitle := eventItem.Find(".event-title").Text()

		// Extract artist list (inside the .artist-list class)
		var artists string
		eventItem.Find(".artist-list .artist-name").Each(func(j int, artist *goquery.Selection) {
			artists = artists + "," + artist.Text()
		})

		// Extract all tags (inside the .tag class)
		var genre string
		eventItem.Find(".tag").Each(func(k int, tag *goquery.Selection) {
			genre = genre + "," + tag.Text()
		})

		evList = append(evList, event{
			title:   eventTitle,
			date:    eventDate,
			artists: artists,
			genre:   genre,
		})
	})
	return evList
}

func main() {
	dachEvs := scrapeDachstock("https://www.dachstock.ch/events")
	fmt.Println("#### Dachstock ####")
	for _, ev := range dachEvs {
		fmt.Println("Date: " + ev.date)
		fmt.Println("Title: " + ev.title)
		fmt.Println("Artists: " + ev.artists)
		fmt.Println("Genres: " + ev.genre)
	}

}
