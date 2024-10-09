package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeCafete(url string) []models.Event {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

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

		evList = append(evList, models.Event{
			Title:   eventTitle,
			Date:    eventDate,
			Artists: artists,
			Genre:   genre,
			IsImage: false,
		})
	})
	return evList
}
