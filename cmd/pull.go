package cmd

import (
	"fmt"
	"github.com/zianKazi/redit-bot/core"

	"github.com/spf13/cobra"
)

//PullCmd Command for fetching content
var PullCmd = &cobra.Command {
	Use:   "pull",
	Short: "Pull Content from a path in Reddit",
	Long:  "Pull Content from a path in Reddit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetch time")

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

		status, err := bot.Pull(path)

		if err != nil {
			fmt.Println("Error when trying to pull Reddit listings for: ", path)
		} else {
			fmt.Println("Pull completed with status: ", status)
		}
		return
	},
}
