package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// countdownCmd represents the countdown command
var countdownCmd = &cobra.Command{
	Use:   "countdown",
	Short: "Performs a countdown to a specific date",
	Long:  `Countdown to a specific date provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here, you would normally get the date from a flag or argument
		targetDate, _ := time.Parse("2006-01-02", "2024-09-26") // Example fixed date
		daysUntil := targetDate.Sub(time.Now()).Hours() / 24
		fmt.Printf("Countdown to September 26, 2024: %.0f days remaining!\n", daysUntil)
	},
}

func init() {
	rootCmd.AddCommand(countdownCmd)

	// Here you can define flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
