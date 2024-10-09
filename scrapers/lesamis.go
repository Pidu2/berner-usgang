package scrapers

import (
	"github.com/Pidu2/berner-usgang/models"
	"github.com/Pidu2/berner-usgang/utils"
)

func ScrapeLesAmis(url string) ([]models.Event, error) {
	evList := []models.Event{}
	doc, err := utils.ScrapePage(url)

	if err != nil {
		return nil, err
	}

	programImg, _ := doc.Find("img").Eq(1).Attr("src") //.First().Attr("src")
	evList = append(evList, models.Event{
		Title:   programImg,
		IsImage: true,
	})
	return evList, nil
}
