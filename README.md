# Anime Release Schedule Finder
Tool that scrapes MyAnimeList to find the release schedule of specific anime.

## **Why?**
The goal for this tool is to integrate it in a notification system running on a Raspberry Pi. It's also just a nice excuse to try Golang out :)

## **How it works**
It reads anime names from the `animes.txt` file. For each name, it calls MyAnimeList's search endpoint, parses the returned HTML and picks the first search result.

## **Requirements**
* Golang
* `go get "golang.org/x/net/html"`

## **Usage**
1. Add the anime you want to know the release schedule for in `animes.txt`.
2. Type `go run .` in a terminal within this repo and read the printed output.
3. (Optional) Compile into a binary with `go build .`