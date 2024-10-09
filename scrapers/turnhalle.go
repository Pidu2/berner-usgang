package scrapers

import (
	"regexp"

	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeTurnhalle(url string) []models.Event {
	evList := []models.Event{}

	doc, err := utils.ScrapePage(url)

	if err != nil {
		return err
	}

	// Find the event list and iterate over each event item
	doc.Find(".slideinner").Each(func(i int, eventItem *goquery.Selection) {
		if i == 0 {
			return
		}

		textClassP, _ := eventItem.Find(".text").Find("p").Html()

		// Extract event date
		re := regexp.MustCompile("<br/?>")
		eventDate := re.Split(textClassP, -1)[1]

		// Extract event title
		eventTitle := ""
		eventItem.Find(".text").Find("span").Each(func(j int, titleItem *goquery.Selection) {
			eventTitle += titleItem.Text() + " "
		})

		// Extract artist
		genre := re.Split(textClassP, -1)[2]

		evList = append(evList, models.Event{
			Title:   eventTitle,
			Date:    eventDate,
			Artists: "",
			Genre:   genre,
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

	return evList
}
