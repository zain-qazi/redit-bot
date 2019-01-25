package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

//RootCmd Root command for the application
var RootCmd = &cobra.Command{
	Use:   "rbot",
	Short: "Bot for Reddit data fetching",
	Long:  "Bot for Reddit data fetching",
	Run: func(cmd *cobra.Command, args []string) {
		//Do stuff here
		fmt.Println("Application starts!")
	},
}

func init() {

	//Flags for the Root command
	RootCmd.PersistentFlags().String("path", "", "Path of content to fetch from Reddit. Could be Sub-Reddit or a Post.")
	RootCmd.MarkPersistentFlagRequired("path")
	RootCmd.PersistentFlags().String("agent-config", "", "Location of the Agent file.")
	RootCmd.MarkPersistentFlagRequired("agent-config")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose execution")

	//Sub commands
	RootCmd.AddCommand(PullCmd)
	RootCmd.AddCommand(ListenCmd)

	//Flags for sub command
	PullCmd.Flags().Bool("subscribe", false, "Subscribe for new content")

}
