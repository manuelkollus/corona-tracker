package main

import (
	"github.com/spf13/cobra"
	"os"
)

var baseCommand = &cobra.Command{
	Use:   "coronatracker",
	Short: "The coronatracker tracks the global data",
	Long:  ``,
}

func main() {
	if err := baseCommand.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	baseCommand.AddCommand(deathsCommand)

}
