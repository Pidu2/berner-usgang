package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeChessu(url string) ([]models.Event, error) {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
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
