package main

import (
	"tv-guide/cmd"

	"github.com/spf13/cobra"
)

func main() {
	MainCommand()
}

func MainCommand() {
	rootCmd := &cobra.Command{Use: "tv-guide"}
	rootCmd.AddCommand(cmd.MovieCmd)
	rootCmd.Execute()
}
