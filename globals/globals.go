package globals

import (
	"os"

	"github.com/Pidu2/berner-usgang/scrapers"
)

type Scraper struct {
	Key      string
	Function scrapers.ScraperFunc
	URL      string
}

var (
	AvailableScrapers = map[string]Scraper{
		"dachstock": {
			Function: scrapers.ScrapeDachstock,
			URL:      getEnv("URL_DACHSTOCK", "https://www.dachstock.ch/events"),
		},
		"gaskessel": {
			Function: scrapers.ScrapeChessu,
			URL:      getEnv("URL_GASKESSEL", "https://gaskessel.ch/programm/"),
		},
		"isc": {
			Function: scrapers.ScrapeISC,
			URL:      getEnv("URL_ISC", "https://isc-club.ch/"),
		},
		"bierhuebeli": {
			Function: scrapers.ScrapeHuebeli,
			URL:      getEnv("URL_HUEBELI", "https://bierhuebeli.ch/"),
		},
		"lesamis": {
			Function: scrapers.ScrapeLesAmis,
			URL:      getEnv("URL_LESAMIS", "https://www.lesamis.ch/wohnzimmer/"),
		},
		"deadend": {
			Function: scrapers.ScrapeDeadEnd,
			URL:      getEnv("URL_DEADEND", "https://dead-end.ch/programm/"),
		},
		"turnhalle": {
			Function: scrapers.ScrapeTurnhalle,
			URL:      getEnv("URL_TURNHALLE", "https://www.progr.ch/de/turnhalle/programm/"),
		},
		"kapitel": {
			Function: scrapers.ScrapeKapitel,
			URL:      getEnv("URL_KAPITEL", "https://www.kapitel.ch/programm/"),
		},
		"roessli": {
			Function: scrapers.ScrapeRoessli,
			URL:      getEnv("URL_ROESSLI", "https://www.souslepont-roessli.ch/"),
		},
		"cafete": {
			Function: scrapers.ScrapeCafete,
			URL:      getEnv("URL_CAFETE", "https://cafete.ch/"),
		},
		"stellwerk": {
			Function: scrapers.ScrapeStellwerk,
			URL:      getEnv("URL_STELLWERK", "https://www.stellwerk.be/klub"),
		},
	}
	ScraperList = getScraperList()
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getScraperList() []string {
	var scraperList []string
	for scraper := range AvailableScrapers {
		scraperList = append(scraperList, scraper)
	}
	return scraperList
}
