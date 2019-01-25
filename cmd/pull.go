package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/zianKazi/redit-bot/core"
)

//PullCmd Command for fetching content
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull Content from a path in Reddit",
	Long:  "Pull Content from a path in Reddit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetch time")

		var path = cmd.Flag("path").Value.String()
		var configLocation = cmd.Flag("agent-config").Value.String()
		var verbose = cmd.Flag("verbose").Value.String()
		var subscribe = cmd.Flag("subscribe").Value.String()

		fmt.Println("Reddit link: ", path)
		fmt.Println("Agent file location: ", configLocation)
		fmt.Println("Verbose: ", verbose)
		fmt.Println("Subscribe: ", subscribe)

		bot, err := core.CreateBot(configLocation)

		if err != nil {
			fmt.Println("Error when trying to create the bot", err)
			return
		}

		fmt.Println("The bot that was created is ", bot)

		//Pull information from Reddit
		status, err := bot.Pull(path)

		if err != nil {
			fmt.Println("Error when trying to pull Reddit listings for: ", path)
		} else {
			fmt.Println("Pull completed with status: ", status)
		}

		if subscribe == "true" {
			fmt.Println("Time to subscribe")
			errSub := bot.Subscribe(strings.TrimPrefix(path, "/r/"))

			if errSub != nil {
				fmt.Println("Error when trying to listen to Reddit listings for: ", path)
			}
		}

		return
	},
}
