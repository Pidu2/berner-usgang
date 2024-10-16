package globals

import (
	"os"

	"github.com/Pidu2/berner-usgang/scrapers"
)

type Scraper struct {
	Key      string
	Function scrapers.ScraperFunc
	URL      string
	Enabled  string
}

var (
	AvailableScrapers = map[string]Scraper{
		"dachstock": {
			Function: scrapers.ScrapeDachstock,
			URL:      getEnv("URL_DACHSTOCK", "https://www.dachstock.ch/events"),
			Enabled:  getEnv("ENABLED_DACHSTOCK", "true"),
		},
		"gaskessel": {
			Function: scrapers.ScrapeChessu,
			URL:      getEnv("URL_GASKESSEL", "https://gaskessel.ch/programm/"),
			Enabled:  getEnv("ENABLED_GASKESSEL", "true"),
		},
		"isc": {
			Function: scrapers.ScrapeISC,
			URL:      getEnv("URL_ISC", "https://isc-club.ch/"),
			Enabled:  getEnv("ENABLED_ISC", "true"),
		},
		"bierhuebeli": {
			Function: scrapers.ScrapeHuebeli,
			URL:      getEnv("URL_HUEBELI", "https://bierhuebeli.ch/"),
			Enabled:  getEnv("ENABLED_HUEBELI", "true"),
		},
		"lesamis": {
			Function: scrapers.ScrapeLesAmis,
			URL:      getEnv("URL_LESAMIS", "https://www.lesamis.ch/wohnzimmer/"),
			Enabled:  getEnv("ENABLED_LESAMIS", "true"),
		},
		"deadend": {
			Function: scrapers.ScrapeDeadEnd,
			URL:      getEnv("URL_DEADEND", "https://dead-end.ch/programm/"),
			Enabled:  getEnv("ENABLED_DEADEND", "true"),
		},
		"turnhalle": {
			Function: scrapers.ScrapeTurnhalle,
			URL:      getEnv("URL_TURNHALLE", "https://www.progr.ch/de/turnhalle/programm/"),
			Enabled:  getEnv("ENABLED_TURNHALLE", "true"),
		},
		"kapitel": {
			Function: scrapers.ScrapeKapitel,
			URL:      getEnv("URL_KAPITEL", "https://www.kapitel.ch/programm/"),
			Enabled:  getEnv("ENABLED_KAPITEL", "true"),
		},
		"roessli": {
			Function: scrapers.ScrapeRoessli,
			URL:      getEnv("URL_ROESSLI", "https://www.souslepont-roessli.ch/"),
			Enabled:  getEnv("ENABLED_ROESSLI", "true"),
		},
		"cafete": {
			Function: scrapers.ScrapeCafete,
			URL:      getEnv("URL_CAFETE", "https://cafete.ch/"),
			Enabled:  getEnv("ENABLED_CAFETE", "true"),
		},
		"stellwerk": {
			Function: scrapers.ScrapeStellwerk,
			URL:      getEnv("URL_STELLWERK", "https://www.stellwerk.be/klub"),
			Enabled:  getEnv("ENABLED_STELLWERK", "false"),
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
	for scrapername, scraper := range AvailableScrapers {
		if scraper.Enabled == "true" {
			scraperList = append(scraperList, scrapername)
		}
	}
	return scraperList
}
