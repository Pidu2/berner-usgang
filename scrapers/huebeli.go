package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeHuebeli(url string) ([]models.Event, error) {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	// Find the event list and iterate over each event item
	doc.Find(".datumlink").Each(func(i int, eventItem *goquery.Selection) {
		// Extract event date
		eventDate := eventItem.Find(".eventdatum").Text()

		// Extract event title
		eventTitle := eventItem.Find(".eventtitel").Text()

		// Extract artist
		artists := eventItem.Find(".byline").Text()

		evList = append(evList, models.Event{
			Title:   eventTitle,
			Date:    eventDate,
			Artists: artists,
			Genre:   "",
			IsImage: false,
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
		evList[i].Genre = genre
	})

	return evList, nil
}
