package main

import (
	"fmt"
	"time"
)

func main() {
	animes := getAnimeList()
	for _, anime := range animes {
		if anime == "" {
			continue
		}
		animePage := getAnimePage(anime)
		if animePage == "" {
			fmt.Println(anime + ": Could not find anything.")
			continue
		}
		broadcastDate := findBroadcastDateInAnimePage(animePage)
		printAnimeBroadcastDate(anime, broadcastDate)
		time.Sleep(2)
	}
}
