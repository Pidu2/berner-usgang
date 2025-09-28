package globals

import (
	"os"

	"github.com/Pidu2/berner-usgang/scrapers"
)

type Scraper struct {
	Function    scrapers.ScraperFunc `json:"-"`
	DisplayName string               `json:"displayname"`
	URL         string               `json:"-"`
	Enabled     string               `json:"enabled"`
	Order       string               `json:"order"`
}

var (
	// weight for order
	DefaultOrder      = getEnv("DEFAULT_SCRAPER_ORDER", "99")
	AvailableScrapers = map[string]Scraper{
		"dachstock": {
			Function:    scrapers.ScrapeDachstock,
			DisplayName: "Dachstock",
			URL:         getEnv("URL_DACHSTOCK", "https://www.dachstock.ch/events"),
			Enabled:     getEnv("ENABLED_DACHSTOCK", "true"),
			Order:       getEnv("ORDER_DACHSTOCK", DefaultOrder),
		},
		"gaskessel": {
			Function:    scrapers.ScrapeChessu,
			DisplayName: "Gaskessel",
			URL:         getEnv("URL_GASKESSEL", "https://gaskessel.ch/programm/"),
			Enabled:     getEnv("ENABLED_GASKESSEL", "true"),
			Order:       getEnv("ORDER_GASKESSEL", DefaultOrder),
		},
		"isc": {
			Function:    scrapers.ScrapeISC,
			DisplayName: "ISC",
			URL:         getEnv("URL_ISC", "https://isc-club.ch/"),
			Enabled:     getEnv("ENABLED_ISC", "true"),
			Order:       getEnv("ORDER_ISC", DefaultOrder),
		},
		"bierhuebeli": {
			Function:    scrapers.ScrapeHuebeli,
			DisplayName: "Bierhübeli",
			URL:         getEnv("URL_HUEBELI", "https://bierhuebeli.ch/"),
			Enabled:     getEnv("ENABLED_HUEBELI", "true"),
			Order:       getEnv("ORDER_HUEBELI", DefaultOrder),
		},
		"lesamis": {
			Function:    scrapers.ScrapeLesAmis,
			DisplayName: "Les Amis",
			URL:         getEnv("URL_LESAMIS", "https://www.lesamis.ch/programm/"),
			Enabled:     getEnv("ENABLED_LESAMIS", "true"),
			Order:       getEnv("ORDER_LESAMIS", DefaultOrder),
		},
		"deadend": {
			Function:    scrapers.ScrapeDeadEnd,
			DisplayName: "Dead End",
			URL:         getEnv("URL_DEADEND", "https://dead-end.ch/programm/"),
			Enabled:     getEnv("ENABLED_DEADEND", "true"),
			Order:       getEnv("ORDER_DEADEND", DefaultOrder),
		},
		"turnhalle": {
			Function:    scrapers.ScrapeTurnhalle,
			DisplayName: "Turnhalle",
			URL:         getEnv("URL_TURNHALLE", "https://www.progr.ch/de/turnhalle/programm/"),
			Enabled:     getEnv("ENABLED_TURNHALLE", "true"),
			Order:       getEnv("ORDER_TURNHALLE", DefaultOrder),
		},
		"kapitel": {
			Function:    scrapers.ScrapeKapitel,
			DisplayName: "Kapitel",
			URL:         getEnv("URL_KAPITEL", "https://www.kapitel.ch/programm/"),
			Enabled:     getEnv("ENABLED_KAPITEL", "true"),
			Order:       getEnv("ORDER_KAPITEL", DefaultOrder),
		},
		"roessli": {
			Function:    scrapers.ScrapeRoessli,
			DisplayName: "Rössli",
			URL:         getEnv("URL_ROESSLI", "https://www.souslepont-roessli.ch/"),
			Enabled:     getEnv("ENABLED_ROESSLI", "true"),
			Order:       getEnv("ORDER_ROESSLI", DefaultOrder),
		},
		"cafete": {
			Function:    scrapers.ScrapeCafete,
			DisplayName: "Cafete",
			URL:         getEnv("URL_CAFETE", "https://cafete.ch/"),
			Enabled:     getEnv("ENABLED_CAFETE", "true"),
			Order:       getEnv("ORDER_CAFETE", DefaultOrder),
		},
		"stellwerk": {
			Function:    scrapers.ScrapeStellwerk,
			DisplayName: "Stellwerk",
			URL:         getEnv("URL_STELLWERK", "https://www.stellwerk.be/klub"),
			Enabled:     getEnv("ENABLED_STELLWERK", "true"),
			Order:       getEnv("ORDER_STELLWERK", DefaultOrder),
		},
	}
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
