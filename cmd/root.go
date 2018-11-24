package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/turnage/graw/reddit"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "rbot",
	Short: "Bot for Reddit data fetching",
	Long:  "Bot for Reddit data fetching",
	Run: func(cmd *cobra.Command, args []string) {
		//Do stuff here
		fmt.Println("kicking off Graw bot")
		////Create Bot
		//if bot, err := reddit.NewBotFromAgentFile("C:\\Users\\qhasan\\go\\src\\reditbot-local\\bot.agent", 0);
		//	err != nil {
		//	fmt.Println("Failed to create bot handle: ", err)
		//} else {
		//	fmt.Println("at handle part")
		//	cfg := graw.Config{Subreddits: []string{"Music"}}
		//	handler := &eventBot{bot: bot}
		//	if _, wait, err := graw.Run(handler, bot, cfg);
		//		err != nil {
		//		fmt.Println("Failed to start graw run: ", err)
		//	} else {
		//		fmt.Println("graw listening")
		//		fmt.Println("graw running", wait())
		//	}
		//}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type eventBot struct {
	bot reddit.Bot
}

//Handler of Posts
func (e *eventBot) Post(p *reddit.Post) error {
	//if strings.Contains(p.SelfText, "Java") {
	//	<-time.After(5 * time.Second)
	fmt.Printf("%s posted %s at %d", p.Author, p.Title, p.CreatedUTC)
	//}
	return nil
}
