package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func getAnimeList() []string {
	animeData, err := ioutil.ReadFile("animes.txt")
	if err != nil {
		panic(err)
	}

	animes := strings.Split(string(animeData), "\n")
	for i := range animes {
		animes[i] = strings.TrimSpace(animes[i])
	}
	return animes
}

func findHrefsInLink(pageLink string) []string {
	resp, err := http.Get(pageLink)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	var hrefSlice []string
	for {
		currentTokenType := tokenizer.Next()

		switch currentTokenType {
		case html.ErrorToken:
			return hrefSlice
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						path, _ := url.PathUnescape(attr.Val)
						hrefSlice = append(hrefSlice, path)
					}
				}
			}
		}
	}
}

func getAnimePage(animeName string) string {
	fixedAnimeName := strings.ReplaceAll(animeName, " ", "%20")
	malSearchLink := "https://myanimelist.net/search/all?q=" + fixedAnimeName
	hrefs := findHrefsInLink(malSearchLink)
	var wantedHref string
	reg := regexp.MustCompile(`myanimelist.net/anime/[0-9]+/`)
	for _, href := range hrefs {
		if reg.FindString(href) != "" {
			wantedHref = href
			break
		}
	}
	return wantedHref
}

func findBroadcastDateInAnimePage(animePage string) string {
	resp, err := http.Get(animePage)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)

	dirtyBroadcastTimeRegex := regexp.MustCompile(`<span class="dark_text">Broadcast:</span>\s*.*\s*</div>`)
	cleanBroadcastTimeRegex := regexp.MustCompile(`.*at.*`)
	broadcastDate := dirtyBroadcastTimeRegex.FindString(string(html))
	cleanBroadcastDate := cleanBroadcastTimeRegex.FindString(broadcastDate)
	return cleanBroadcastDate
}

func printAnimeBroadcastDate(animeName string, date string) {
	fmt.Println(animeName + ": " + strings.TrimSpace(date))
}
