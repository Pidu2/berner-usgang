package scrapers

import (
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

	doc.Find(".content-area").Find("img").Each(func(i int, eventItem *goquery.Selection) {
		if len(evList) == limit {
			return
		}
		programImg, _ := eventItem.Attr("src")

		evList = append(evList, models.Event{
			Title:   programImg,
			IsImage: true,
		})
	})

	return evList, nil
}
