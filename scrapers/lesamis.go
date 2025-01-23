package scrapers

import (
	"strings"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeLesAmis(url string, limit int) ([]models.Event, error) {
	evList := []models.Event{}
	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	// Find the event list and iterate over each event item
	doc.Find(".siteorigin-widget-tinymce").Each(func(i int, monthItem *goquery.Selection) {

		monthItem.Find("p").Each(func(i int, eventItem *goquery.Selection) {
			if len(evList) == limit {
				return
			}
			eventTitle := eventItem.Find("strong").Text()
			genre := eventItem.Find("em").Text()
			eventDate := strings.Split(eventItem.Text(), "\n")[0]

			evList = append(evList, models.Event{
				Title:   eventTitle,
				Date:    eventDate,
				Genre:   genre,
				IsImage: false,
			})

		})
	})
	return evList, nil

}
