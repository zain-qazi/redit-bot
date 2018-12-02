package core

import (
	"github.com/turnage/graw"
	"fmt"
	"github.com/turnage/graw/reddit"
)

//RedditBot Struct for the internal bot
type RedditBot struct {
	Bot reddit.Bot
}

//Post Handler of Posts
// func (e *Bot) Post(p *reddit.Post) error {
// 	//if strings.Contains(p.SelfText, "Java") {
// 	//	<-time.After(5 * time.Second)
// 	fmt.Printf("%s posted %s at %d", p.Author, p.Title, p.CreatedUTC)
// 	//}
// 	return nil
// }

//CreateBot Creation of bot
func CreateBot(configPath string) (*RedditBot, error) {
	bot, err := reddit.NewBotFromAgentFile(configPath, 0)

	if err != nil {
		fmt.Println("Failed to initialize the Reddit Bot")
		return nil, err
	}

	return &RedditBot{Bot: bot}, err
}

//Pull Pull content from Reddit
func (r *RedditBot) Pull(path string) (bool, error) {
		
		//Get Listing for the sub-reddit
		harvest, err := r.Bot.Listing(path, "")
	
		//Handle errors, print and return
		if err != nil {
			fmt.Println("Failed to fetch from", path, err)
			return false, err
		}
	
		//Loop and print posts
		for _, post := range harvest.Posts {
			fmt.Printf("%s posted %s\n", post.Author, post.Title)
		}
		return true, nil
}

//Post PostHandler that handles Post events
func (r *RedditBot) Post(p *reddit.Post) error {
	fmt.Println("Received new Post...")
	fmt.Println("Title: ", p.Title)
	fmt.Println("Text: ", p.SelfText)
	fmt.Println("Author: ", p.Author)
	return nil
}

//Subscribe subscribe to a path in Reddit
func (r *RedditBot) Subscribe(path string) error {
	cfg := graw.Config{Subreddits: []string{path}}
	fmt.Println("Bot is ", r.Bot)
	fmt.Println("Handler is ", r)
	fmt.Println("Config is ", cfg)
	if _ ,wait, err := graw.Run(r, r.Bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
		return err
	} else {
		fmt.Println("Graw run kicked off... ", wait())
		return nil
	}
	
}