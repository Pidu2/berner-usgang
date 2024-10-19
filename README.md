# berner-usgang scraper
Scrapes Bidu's usual Usgang-Spots in Bern and provides the results as REST API.

## Usage

```bash
go run main.go # or use Containerfile
curl localhost:8080/scraper
curl localhost:8080/scraper/<scraper>?limit=10
```

Check `globals/globals.go` for ENV-Vars that can be passed.