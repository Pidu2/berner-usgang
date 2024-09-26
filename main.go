package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

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

func scrapePage(url string) (*goquery.Document, []event) {
	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		return nil, errEvent(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errEvent(fmt.Sprintf("Failed to fetch page: %d %s", res.StatusCode, res.Status))
	}

	// Parse the page
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, errEvent(err.Error())
	}

	return doc, nil

}

func scrapeDachstock(url string) []event {
	evList := []event{}

	doc, err := scrapePage(url)

	if err != nil {
		return err
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
		// remove first comma
		if artists != "" {
			artists = artists[1:]
		}

		// Extract all tags (inside the .tag class)
		var genre string
		eventItem.Find(".tag").Each(func(k int, tag *goquery.Selection) {
			genre = genre + "," + tag.Text()
		})
		// remove first comma
		if genre != "" {
			genre = genre[1:]
		}

		// if there is no title set, then its probably a concert
		if eventTitle == "" {
			eventTitle = "Konzert"
		}
		evList = append(evList, event{
			title:   eventTitle,
			date:    eventDate,
			artists: artists,
			genre:   genre,
		})
	})
	return evList
}

func scrapeChessu(url string) []event {
	evList := []event{}

	doc, err := scrapePage(url)

	if err != nil {
		return err
	}

	// Find the event list and iterate over each event item
	doc.Find(".eventpreview").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := eventItem.Find(".eventdatum").Text()

		// Extract event title
		eventTitle := eventItem.Find(".eventname").Text()

		// Extract artist
		artists := eventItem.Find(".subtitle").Text()

		// Extract genre
		genre := eventItem.Find(".eventgenre").Text()

		evList = append(evList, event{
			title:   eventTitle,
			date:    eventDate,
			artists: artists,
			genre:   genre,
		})
	})
	return evList
}

func scrapeISC(url string) []event {
	evList := []event{}

	doc, err := scrapePage(url)

	if err != nil {
		return err
	}

	// Find the event list and iterate over each event item
	doc.Find(".event_preview").Each(func(i int, eventItem *goquery.Selection) {
		// ISC has lots of whitespaces, remove them
		reLeading := regexp.MustCompile(`^\s+`)
		reTrailing := regexp.MustCompile(`\s+$`)

		// Extract event date
		eventDate := eventItem.Find(".event_title_date").Text()

		// Extract event title
		eventTitle := eventItem.Find(".event_title_title").Text()

		// Extract genre
		genre := eventItem.Find(".event_title_info_mobile").Text()

		evList = append(evList, event{
			title:   reTrailing.ReplaceAllString(reLeading.ReplaceAllString(eventTitle, ""), ""),
			date:    reTrailing.ReplaceAllString(reLeading.ReplaceAllString(eventDate, ""), ""),
			artists: "", // ISC has no artists
			genre:   reTrailing.ReplaceAllString(reLeading.ReplaceAllString(genre, ""), ""),
		})
	})
	return evList
}

func scrapeCafete(url string) []event {
	evList := []event{}

	doc, err := scrapePage(url)

	if err != nil {
		return err
	}

	// Find the event list and iterate over each event item
	doc.Find(".event").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := eventItem.Find(".date").Text()

		// Extract event title
		eventTitle := eventItem.Find(".title").Text()

		// Extract artist
		artists := eventItem.Find(".acts").Text()

		// Extract genre
		genre := eventItem.Find(".style").Text()

		evList = append(evList, event{
			title:   eventTitle,
			date:    eventDate,
			artists: artists,
			genre:   genre,
		})
	})
	return evList
}

func scrapeHuebeli(url string) []event {
	evList := []event{}

	doc, err := scrapePage(url)

	if err != nil {
		return err
	}

	// Find the event list and iterate over each event item
	doc.Find(".datumlink").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := eventItem.Find(".eventdatum").Text()

		// Extract event title
		eventTitle := eventItem.Find(".eventtitel").Text()

		// Extract artist
		artists := eventItem.Find(".byline").Text()

		evList = append(evList, event{
			title:   eventTitle,
			date:    eventDate,
			artists: artists,
			genre:   "",
		})
	})

	doc.Find(".stiltags").Each(func(i int, eventItem *goquery.Selection) {
		genre := ""
		eventItem.Find("a").Each(func(j int, eventItemInner *goquery.Selection) {
			genre = genre + "," + eventItemInner.Text()
		})
		// remove first comma
		if genre != "" {
			genre = genre[1:]
		}
		evList[i].genre = genre
	})

	return evList
}

func main() {
	//dachEvs := scrapeDachstock("https://www.dachstock.ch/events")
	//chessEvs := scrapeChessu("https://gaskessel.ch/programm/")
	//iscEvs := scrapeISC("https://isc-club.ch/") // het keni artists
	hueEvs := scrapeHuebeli("https://bierhuebeli.ch/") // hie si artists mengisch o eifach subtitles
	for _, ev := range hueEvs {
		fmt.Println("--------")
		fmt.Println("Date: " + ev.date)
		fmt.Println("Title: " + ev.title)
		fmt.Println("Artists: " + ev.artists)
		fmt.Println("Genres: " + ev.genre)
	}

}
