package main

import (
	"log"

	"tv-guide/cmd"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MainCommand()
}

func MainCommand() {
	rootCmd := &cobra.Command{Use: "tv-guide"}
	rootCmd.AddCommand(cmd.MovieCmd)
	rootCmd.Execute()
}
