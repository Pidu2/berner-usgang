package main

import (
	"fmt"

	"github.com/Pidu2/berner-usgang/scrapers"
)

func main() {
	// TODO build REST endpoint
	// TODO ggf add rabe
	//dachEvs := scrapers.ScrapeDachstock("https://www.dachstock.ch/events")
	//chessEvs := scrapers.ScrapeChessu("https://gaskessel.ch/programm/")
	//iscEvs := scrapers.ScrapeISC("https://isc-club.ch/")        // het keni artists
	//hueEvs := scrapers.ScrapeHuebeli("https://bierhuebeli.ch/") // hie si artists mengisch o eifach subtitles
	//leEvs := scrapers.ScrapeLesAmis("https://www.lesamis.ch/wohnzimmer/")
	//deEvs := scrapers.ScrapeDeadEnd("https://dead-end.ch/programm/")
	//thEvs := scrapers.ScrapeTurnhalle("https://www.progr.ch/de/turnhalle/programm/")
	//kapEvs := scrapers.ScrapeKapitel("https://www.kapitel.ch/programm/")
	roEvs, err := scrapers.ScrapeRoessli("https://www.souslepont-roessli.ch/")
	if err != nil {
		return
	}
	for _, ev := range roEvs {
		fmt.Println("--------")
		fmt.Println("Date: " + ev.Date)
		fmt.Println("Title: " + ev.Title)
		fmt.Println("Artists: " + ev.Artists)
		fmt.Println("Genres: " + ev.Genre)
	}

}
