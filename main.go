package main

import (
	"fmt"
	"os"
	"redit-bot/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println("An error occurred")
		os.Exit(1)
	}
}
