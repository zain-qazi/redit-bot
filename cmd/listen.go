package cmd

import (
	"github.com/zianKazi/redit-bot/core"
	"fmt"
	"github.com/spf13/cobra"
)

//ListenCmd Command for listening to live content
var ListenCmd = &cobra.Command {
	Use: "listen",
	Short: "Listen for content from a path in Reddit",
	Long: "Listen for content from a path in Reddit",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: remove redundant code into a helper
		fmt.Println("Initializing Listeners")
		var path = cmd.Flag("path").Value.String()
		var configLocation = cmd.Flag("agent-config").Value.String()
		var verbose = cmd.Flag("verbose").Value.String()

		fmt.Println("Reddit link: ", path)
		fmt.Println("Agent file location: ", configLocation)
		fmt.Println("Verbose: ", verbose)

		bot, err := core.CreateBot(configLocation)

		if err != nil {
			fmt.Println("Error when trying to create the bot", err)
			return
		}

		fmt.Println("The bot that was created is ", bot)

		errSub := bot.Subscribe(path)

		if errSub != nil {
			fmt.Println("Error when trying to listen to Reddit listings for: ", path)
		}

	},
}
