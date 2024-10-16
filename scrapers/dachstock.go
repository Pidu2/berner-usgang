package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeDachstock(url string, limit int) ([]models.Event, error) {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	// Find the event list and iterate over each event item
	doc.Find(".teaser").Each(func(i int, eventItem *goquery.Selection) {
		if len(evList) == limit {
			return
		}
		// Extract event date
		eventDate := eventItem.Find(".event-date").Text()

		// Extract event title
		eventTitle := eventItem.Find(".event-title").Text()

		// Extract artist list (inside the .artist-list class)
		var artists string
		eventItem.Find(".artist-list .artist-name").Each(func(j int, artist *goquery.Selection) {
			artists = artists + ", " + artist.Text()
		})
		// remove first comma
		if artists != "" {
			artists = artists[2:]
		}

		// Extract all tags (inside the .tag class)
		var genre string
		eventItem.Find(".tag").Each(func(k int, tag *goquery.Selection) {
			genre = genre + ", " + tag.Text()
		})
		// remove first comma
		if genre != "" {
			genre = genre[2:]
		}

		// if there is no title set, then its probably a concert
		if eventTitle == "" {
			eventTitle = "Konzert"
		}
		evList = append(evList, models.Event{
			Title:   eventTitle,
			Date:    eventDate,
			Artists: artists,
			Genre:   genre,
			IsImage: false,
		})
	})
	return evList, nil
}
