package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"kollus.de/coronatracker/pkg/record"
)

var deathsCommand = &cobra.Command{
	Use:   "stats",
	Short: "Show the actual stats",
	Long:  `Show the actual stats`,
	Run:   RunStats,
}

func RunStats(
	command *cobra.Command,
	arguments []string) {
	record := record.Create()
	fmt.Println("Loaded following stats:")
	fmt.Println("Confirmed | Deaths | Recovered")
	fmt.Println(record)
}
