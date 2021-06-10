package main

import (
	"fruits-cash-register/internal/app"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cash register",
		Short: "",
		Long: `Register cash coming in from fruits selling.`,
		Run: func(cmd *cobra.Command, args []string) {
			application := app.New()
			application.Run()
		},
	}
	rootCmd.Execute()
}
