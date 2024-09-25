# Notes
* Dächu
  * https://www.dachstock.ch/events Loop through `.event-list`, get Date from `.event-date`, get title from `.event-title`, get artists from `.artist-list .artist-name`, get genre from all `.tag`
* Chessu
  * https://gaskessel.ch/programm/ Loop through `.eventpreview`, get Date from `.eventdatum`, get title from `.eventname`, get artists from `.subtitle`, get genre from `.eventgenre`
* ISC
  * https://isc-club.ch/ Loop through `.event_preview`, get Date from `.event_title_date`, get title from `.event_title_title`, NO ARTIST, get genre from `.event_title_info_mobile`
* Cafete
  * https://cafete.ch/ Loop throgh `.event`, get Date from `.date`, get title from `.title`, get artists from `.acts`, get genre from `.style`
* Hübeli
  * https://bierhuebeli.ch/ `.datumlink` and `.stiltags` alternate, stiltags containing only genre. Loop through `.datumlink`, get Date from `.eventdatum`, get title from `.eventtitel`, get artists from `.byline`
* Les Amis
  * Bild https://www.lesamis.ch/wohnzimmer/
* Deadend
  * Bild https://dead-end.ch/programm/
* PROGR/Turnhalle
  * https://www.progr.ch/de/turnhalle/programm/ Loop through `.slideinner`, get Date from `.text <p>`, get title from `.text <h3`, get artists from `.text <h3>`, get genre from `.text <p>`
* Kapitel
  * https://www.kapitel.ch/programm/year-month Loop through `.event-list`, get Date from `.event-date-inner`, get title from `.size-medium`, get artists from `.event-artist-list`, get genre from `.event-tag`
* Edward Rabe Loop through `
  * Not yet available / seasonal
