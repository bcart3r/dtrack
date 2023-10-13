package main

import (
	"dtrack/player"
	"fmt"
)

func main() {
	playerLink := player.GetPlayer("alec", "bryant", "Illinois Fighting Illini")
	fmt.Println(playerLink)
	gamelog := player.GetGameLog(playerLink, "2022")

	for _, n := range gamelog {
		fmt.Println(n.Text())
	}
}
