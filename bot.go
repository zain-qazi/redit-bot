package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"
)

func main() {

	//Create Bot
	bot, err := reddit.NewBotFromAgentFile("C:\\Users\\qhasan\\go\\src\\reddit\\bot.agent", 0)

	//Handle errors, print and return
	if err != nil {
		fmt.Println("Failed to create bot handle", err)
		return
	}

	//Get Listing for the sub-reddit
	harvest, err := bot.Listing("/r/postrock", "")

	//Handle errors, print and return
	if err != nil {
		fmt.Println("Failed to fetch from /r/postrock", err)
		return
	}

	//Loop and print posts
	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}

}
